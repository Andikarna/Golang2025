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

//go:generate swag init -g main.go -o ./docs

func main() {

	cfg := config.LoadConfig()

	database.Connect(cfg)

	database.DB.AutoMigrate(&model.Article{})
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.UserToken{})

	r := routes.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Domain FE yang diizinkan
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With", "Accept", "Origin"},
		AllowCredentials: true,
		Debug:            true, // Set ke false di production
	})

	handler := c.Handler(r)

	// log.Println("Server started at :8081")
	// log.Fatal(http.ListenAndServe(":8081", r))

	log.Printf("Server started at :8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
