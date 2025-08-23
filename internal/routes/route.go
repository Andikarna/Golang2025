package routes

import (
    "net/http"

    "fundamental/internal/handler"
    "github.com/gorilla/mux"

    httpSwagger "github.com/swaggo/http-swagger"
  	_ "fundamental/docs"
)

func SetupRoutes() *mux.Router {
    app := mux.NewRouter().StrictSlash(true)

    app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Homepage Endpoint Hit"))
    }).Methods("GET")

    app.HandleFunc("/articles", handler.GetArticlesHandler).Methods("GET")
    app.HandleFunc("/articles", handler.CreateArticleHandler).Methods("POST")

    // Endpoint Users
    app.HandleFunc("/users", handler.GetUsers).Methods("GET")

    app.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
    return app
}
