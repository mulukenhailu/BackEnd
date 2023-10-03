package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/mulukenhailu/FoodRecipe/auth"
	"github.com/mulukenhailu/FoodRecipe/handler"
)

func NewRedisDB() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return redisClient
}

func main() {
	r := gin.Default()
	redisClient := NewRedisDB()

	var rd = auth.NewAuth(redisClient)
	var tk = auth.NewToken()
	var service = handler.NewProfile(rd, tk)

	r.POST("/", handler.PublicUser)
	r.POST("/login", service.Login)
	r.Run(":3001")
}
