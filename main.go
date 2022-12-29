package main

import (
	"fmt"
	"urlshortener/api/redirect"
	"urlshortener/api/urls"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}
func main() {
	app := gin.Default()
	app.Use(gin.Recovery())
	redirect.RouteUrls(app)
	urls.RouteUrls(app)

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
