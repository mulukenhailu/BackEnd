package utilites

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/mulukenhailu/FoodRecipe/config"
)

func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

type UserByEmail struct {
	User []struct {
		Email     string `graphql:"email"`
		Password  string `graphql:"password"`
		User_name string `graphql:"user_name"`
		User_id   string `graphql:"user_id"`
	} `graphql:"User(where: {email: {_eq: $email}})"`
}

func GetUserByEmail(email string) (string, error) {
	client := config.GraphqlClient()

	var userdata UserByEmail

	variables := map[string]interface{}{
		"email": email,
	}

	err := client.Query(context.Background(), &userdata, variables)
	if err != nil {
		fmt.Println("Error accessing the database")
		return "", err
	}

	if len(userdata.User) == 1 {
		user := userdata.User[0]
		userJson, _ := json.Marshal(user)
		return string(userJson), nil
	} else {
		return "", errors.New("unexpected error occured")

	}
}
