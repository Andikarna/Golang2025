package routes

import (
	"net/http"

	"fundamental/internal/handler"
	"fundamental/internal/middleware"

	"github.com/gorilla/mux"

	_ "fundamental/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes() *mux.Router {
	app := mux.NewRouter().StrictSlash(true)

	app.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Homepage Endpoint Hit"))
	}).Methods("GET")

	// Endpoint Users
	app.HandleFunc("/users", handler.GetUsers).Methods("GET")

	//Auth
	app.HandleFunc("/api/login", handler.Login).Methods("POST")
	app.HandleFunc("/api/logout", handler.Logout).Methods("POST")
	app.HandleFunc("/api/refreshToken", handler.RefreshToken).Methods("POST")
	app.HandleFunc("/api/register", handler.Register).Methods("POST")

	app.Handle("/api/attendance", middleware.JWTAuth(http.HandlerFunc(handler.GetAttendance))).Methods("GET")

	app.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return app
}
