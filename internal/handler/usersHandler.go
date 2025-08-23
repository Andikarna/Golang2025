package handler

import (
	"encoding/json"
	"net/http"

	"fundamental/internal/model"
  "fundamental/internal/repository"
)

var _ = model.User{};

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

		respondJSON(w, http.StatusOK,users);
}
