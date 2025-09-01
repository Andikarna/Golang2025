package handler

import (
	"encoding/json"
	"net/http"

	"fundamental/internal/dto"
	"fundamental/internal/model"
	"fundamental/internal/repository"
)

var _ = model.User{}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// GetArticlesHandler godoc
// @Summary Get all Users
// @Description Get details of all available user
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} model.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUser()
	if err != nil {
		http.Error(w, "Error fetching articles", http.StatusInternalServerError)
		return
	}
	// json.NewEncoder(w).Encode(users)

	respondJSON(w, http.StatusOK, users)
}

// @Description Authentication for user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param loginRequest body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse "Login successful"
// @Router /api/login [post]
func Login(w http.ResponseWriter, r *http.Request) {

	var loginRequest dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, token, refreshToken, err := repository.Login(loginRequest.Email, loginRequest.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"user":         user,
		"token":        token,
		"refreshToken": refreshToken,
	}

	respondJSON(w, http.StatusOK, response)
}

// @Description Authentication for user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param refreshRequest body dto.RefreshRequest true "Refresh Param"
// @Success 200 {object} dto.RefreshResponse "200, Refresh Token Berhasil"
// @Router /api/refreshToken [post]
func RefreshToken(w http.ResponseWriter, r *http.Request) {

	var refreshRequest dto.RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&refreshRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newToken, newRefreshToken, err := repository.RefreshToken(refreshRequest.RefreshToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := map[string]interface{}{
		"newToken":        newToken,
		"newRefreshToken": newRefreshToken,
	}

	respondJSON(w, http.StatusOK, response)
}

// @Description Logout for user
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.LogoutResponse "200, Logout Berhasil"
// @Router /api/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"status":  "200",
		"message": "Berhasil melakukan logout",
	}

	respondJSON(w, http.StatusOK, response)
}

// @Description Register user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param loginRequest body dto.RegisterRequest true "Body Request"
// @Success 200 {object} dto.Response
// @Router /api/register [post]
func Register(w http.ResponseWriter, r *http.Request) {

	var registerRequest dto.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	status, message, err := repository.RegisterUser(registerRequest)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	response := map[string]interface{}{
		"status":  status,
		"message": message,
	}

	respondJSON(w, status, response)
}
