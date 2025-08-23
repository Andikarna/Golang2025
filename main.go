package main

import (
    "log"
    "net/http"

    "fundamental/config"
    "fundamental/internal/database"
    "fundamental/internal/model"
    "fundamental/internal/routes"
)

func main() {
    cfg := config.LoadConfig()

    database.Connect(cfg)

    database.DB.AutoMigrate(&model.Article{})
    database.DB.AutoMigrate(&model.User{})

    r := routes.SetupRoutes()

    log.Println("Server started at :8081")
    log.Fatal(http.ListenAndServe(":8081", r))
}