// @title Subscriptions API
// @version 1.0
// @description REST API для работы с подписками
// @host localhost:8080
// @BasePath /

package main

import (
	"log"
	"os"

	"subscriptions-service/internal/db"
	"subscriptions-service/internal/handler"
	"subscriptions-service/internal/repository"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "subscriptions-service/docs"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@127.0.0.1:5433/subscriptions?sslmode=disable"
	}

	pool := db.Connect(dbURL)
	defer pool.Close()

	repo := repository.NewSubscriprionRepo(pool)
	h := handler.NewSubscriprionHandler(repo)

	r := gin.Default()
	r.POST("/subscriptions", h.Create)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/subscriptions", h.GetAll)

	r.GET("/subscriptions/:id", h.GetByID)

	r.PUT("subscriptions/:id", h.Update)

	r.DELETE("subscriptions/:id", h.Delete)

	r.GET("subscriptions/total", h.GetTotalPrice)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
