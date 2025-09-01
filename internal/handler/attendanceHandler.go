package handler

import (
	"fundamental/internal/components"
	"fundamental/internal/model"
	"fundamental/internal/repository"
	"net/http"
)

var _ = model.User{}

// GetAttendance godoc
// @Summary Get attendance
// @Description Get history attendance list
// @Tags Attendance
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Attendance
// @Failure 401 {object} map[string]string "Unauthorized"
// @Security BearerAuth
// @Router /api/attendance [get]
func GetAttendance(w http.ResponseWriter, r *http.Request) {
	// Ambil user_id langsung dari context yang sudah diisi middleware JWT
	userID, ok := r.Context().Value("user_id").(string)
	if !ok || userID == "" {
		components.RespondError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	attendance, err := repository.GetAttendance(userID)
	if err != nil {
		components.RespondError(w, http.StatusInternalServerError, "Error fetching attendance")
		return
	}

	components.RespondJSON(w, http.StatusOK, attendance)
}
