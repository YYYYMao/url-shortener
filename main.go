package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urlshortener/api/redirect"
	"urlshortener/api/urls"
	urlsRepo "urlshortener/repositories/urls"
	"urlshortener/utils/pg"
	"urlshortener/utils/redis"

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
	app := setupRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}

func setupRouter() *gin.Engine {
	app := gin.Default()
	if mode := gin.Mode(); mode == gin.DebugMode {
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8080))
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
	app.Use(gin.Recovery())

	urlRepo := urlsRepo.NewUrlRepository(pg.Db)
	urlService := urls.NewUrlService(urlRepo)
	cacheRepo := redis.NewRedisRepository(redis.Client)
	redirectService := redirect.NewRedirectService(urlRepo, cacheRepo)

	redirectController := &redirect.RedirectController{
		RedirectService: redirectService,
	}

	urlController := &urls.UrlController{
		UrlService: urlService,
	}

	redirect.RouteUrls(app, redirectController)
	urls.RouteUrls(app, urlController)

	return app
}
