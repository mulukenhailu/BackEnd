package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

type AuthInterface interface {
	CreateAuth(string, *TokenDetail) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetail) error
}

type services struct {
	client *redis.Client
}

var _AuthInterface = &services{}

func NewAuth(client *redis.Client) *services {
	return &services{client: client}
}

type AccessDetail struct {
	TokenUuid string
	UserId    string
}

type TokenDetail struct {
	AccessToken  string
	AtExpire     int64
	TokenUuid    string
	RefreshToken string
	RtExpire     int64
	RefreshUuid  string
}

func (tk *services) CreateAuth(userId string, td *TokenDetail) error {
	at := time.Unix(td.AtExpire, 0)
	rt := time.Unix(td.RtExpire, 0)
	now := time.Now()

	atCreated, err := tk.client.Set(td.TokenUuid, userId, at.Sub(now)).Result()
	if err != nil {
		return err
	}

	rtCreated, err := tk.client.Set(td.TokenUuid, userId, rt.Sub(now)).Result()
	if err != nil {
		return err
	}

	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}

	return nil
}

func (tk *services) FetchAuth(tokenUuid string) (string, error) {
	userId, err := tk.client.Get(tokenUuid).Result()
	if err != nil {
		return "", nil
	}
	return userId, nil
}

func (tk *services) DeleteTokens(authId *AccessDetail) error {
	refreshUuid := fmt.Sprintf("%s++%s", authId.TokenUuid, authId.UserId)

	deleteAt, err := tk.client.Del(authId.TokenUuid).Result()
	if err != nil {
		return err
	}

	deleteRt, err := tk.client.Del(refreshUuid).Result()
	if err != nil {
		return err
	}

	if deleteAt != 1 || deleteRt != 1 {
		return errors.New("something went wrong")
	}

	return nil

}

func (tk *services) DeleteRefresh(refreshUuid string) error {
	deleted, err := tk.client.Del(refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}

	return nil
}
