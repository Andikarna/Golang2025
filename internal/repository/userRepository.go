package repository

import (
	"fundamental/internal/database"
	"fundamental/internal/model"
)

var _ = model.User{};

func GetAllUser() ([]model.User, error) {
    var users []model.User
    result := database.DB.Find(&users)
    return users, result.Error
}

