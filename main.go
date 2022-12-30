package main

import (
	"fmt"
	"urlshortener/api/redirect"
	"urlshortener/api/urls"

	_ "urlshortener/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

// @title Url Shortener
// @version 1.0
// @description Url Shortener API.
// @host localhost:8080
func main() {
	app := gin.Default()
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8080))
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
	app.Use(gin.Recovery())
	redirect.RouteUrls(app)
	urls.RouteUrls(app)

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
