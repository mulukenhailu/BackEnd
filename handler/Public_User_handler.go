package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mulukenhailu/FoodRecipe/config"
)

type publicResponse struct {
	Recipes []struct {
		Title       string `graphal:"title"`
		Description string `graphal:"description"`
		User        struct {
			User_name string `graphql:"user_name"`
		} `graphql:"User"`
	} `graphql:"Recipes"`
}

func PublicUser(c *gin.Context) {
	client := config.GraphqlClient()
	var publicRecipe publicResponse

	err := client.Query(context.Background(), &publicRecipe, nil)
	if err != nil {
		fmt.Println("No Recipe to show", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"public Recipe": publicRecipe.Recipes})
}
