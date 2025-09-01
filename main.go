package main

import (
	"log"
	"net/http"

	"fundamental/config"
	"fundamental/internal/database"
	"fundamental/internal/model"
	"fundamental/internal/routes"

	"github.com/rs/cors"
)

// @title Fundamental API
// @version 1.0
// @description API documentation with JWT Auth

// @contact.name Andi Karna
// @contact.url http://localhost:8081
// @contact.email youremail@example.com

// @host localhost:8081
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	cfg := config.LoadConfig()

	database.Connect(cfg)

	database.DB.AutoMigrate(
		&model.Article{},
		&model.User{},
		&model.UserToken{},
		&model.Attendance{},
	)

	r := routes.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Domain FE yang diizinkan
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(r)

	log.Printf("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
