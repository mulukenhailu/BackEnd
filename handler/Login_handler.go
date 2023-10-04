package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/FoodRecipe/auth"
)

type profileHandler struct {
	rd auth.AuthInterface
	tk auth.TokenInterface
}

func NewProfile(rd auth.AuthInterface, tk auth.TokenInterface) *profileHandler {
	return &profileHandler{
		rd: rd,
		tk: tk,
	}
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ph *profileHandler) Login(c *gin.Context) {
	// var u User
	// if err := c.ShouldBindJSON(&u); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, "Invalid Json Provided")
	// 	return
	// }

	// if u.Email == "" || u.Password == "" {
	// 	c.JSON(http.StatusBadRequest, "Please Provide the Necessary Credential")
	// 	return
	// }

	// valid := utilites.IsEmailValid(u.Email)
	// if !valid {
	// 	c.JSON(http.StatusBadRequest, "Invalid Email or Password")
	// 	return
	// }

	// responseJson, err := utilites.GetUserByEmail(u.Email)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }

	// type userDetail struct {
	// 	User_name string
	// 	Password  string
	// 	User_id   string
	// 	Email     string
	// }

	// var UD userDetail
	// err = json.Unmarshal([]byte(responseJson), &UD)
	// if err != nil {
	// 	c.JSON(http.StatusNoContent, gin.H{"unable to parse": err.Error()})
	// 	return
	// }

	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{"err:": err.Error()})
	// 	return
	// }

	// if !Security.CheckPasswordHash(u.Password, UD.Password) {
	// 	c.JSON(http.StatusUnauthorized, "Please provide valid login details")
	// 	return
	// }

	// td, err := ph.tk.CreateToken(UD.User_id)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, "Something went wrong!")
	// 	return
	// }

	// saveErr := ph.rd.CreateAuth(UD.User_id, td)
	// if saveErr != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	// 	return
	// }

	// tokens := map[string]string{
	// 	"access_token": td.AccessToken,
	// }

	// c.SetCookie("access_token", td.AccessToken, 1800, "/", "localhost", false, true)
	// c.SetCookie("referesh_token", td.RefreshToken, 3600*24*7, "/", "localhost", false, true)

	// c.JSON(http.StatusOK, gin.H{"result": tokens})
	c.JSON(http.StatusOK, "login")

}
