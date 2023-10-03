package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type tokenservice struct{}

func NewToken() *tokenservice {
	return &tokenservice{}
}

type TokenInterface interface {
	CreateToken(userId string) (*TokenDetail, error)
	ExtractedMetaData(*http.Request) (*AccessDetail, error)
}

var _TokenInterface = &tokenservice{}

func (t *tokenservice) CreateToken(userId string) (*TokenDetail, error) {
	td := &TokenDetail{}

	td.AtExpire = time.Now().Add(time.Minute * 30).Unix()
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpire = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userId

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpire

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("Access_Secret")))
	if err != nil {
		return nil, err
	}

	td.RtExpire = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userId

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpire
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("Access_Secret")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, "")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("Access_Secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenVaild(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func extract(token *jwt.Token) (*AccessDetail, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if token.Valid && ok {
		access_uuid, ok := claims["access_uuid"].(string)
		user_id, userok := claims["user_id"].(string)

		if !ok || !userok {
			return nil, errors.New("UnAutherized")
		} else {
			return &AccessDetail{
				TokenUuid: access_uuid,
				UserId:    user_id,
			}, nil
		}
	}
	return nil, errors.New("something went Wrong")
}

func (t *tokenservice) ExtractedMetaData(r *http.Request) (*AccessDetail, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, err
	}

	accessDet, err := extract(token)
	if err != nil {
		return nil, err
	}
	return accessDet, nil
}
