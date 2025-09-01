package repository

import (
	"fundamental/internal/database"
	"fundamental/internal/dto"
	"fundamental/internal/model"
)

var _ = model.User{}

func GetAttendance(userID string) ([]dto.AttendanceListResponse, error) {
	var attendance []dto.AttendanceListResponse

	query := database.DB.Model(&model.Attendance{})
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}

	result := query.Find(&attendance)

	return attendance, result.Error
}
