package repository

import (
	"fmt"
	"fundamental/internal/database"
	"fundamental/internal/dto"
	"fundamental/internal/model"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var _ = model.User{}

func GetAllUser() ([]model.User, error) {
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func RegisterUser(request dto.RegisterRequest) (int, string, error) {

	if _, err := GetUserByEmail(request.Email); err == nil {
		return http.StatusBadRequest, "Email sudah terdaftar", nil
	}

	if request.Password != request.ConfirmPassword {
		return http.StatusBadRequest, "Password dan Confirm Password tidak sesuai!", nil
	}

	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		return http.StatusBadRequest, "Password yang dimasukan tidak bsa dihash!", nil
	}

	user := model.User{
		Username:    request.FirstName + " " + request.LastName,
		Email:       request.Email,
		Password:    hashedPassword,
		CreatedDate: time.Now(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return http.StatusConflict, "Failed to register user", fmt.Errorf("failed to register user: %v", err)
	}

	return 200, "User berhasil didaftarkan!", nil
}

func Login(email, password string) (*model.User, string, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, "", fmt.Errorf("user not found")
	}

	if !CheckPassword(user.Password, password) {
		return nil, "Incorrect password", fmt.Errorf("Password Invalid")
	}

	token := generateToken(int(user.ID))
	refreshToken := generateToken(int(user.ID))

	SaveToken(user.ID, token, refreshToken)

	return user, token, nil
}

func RefreshToken(oldRefreshToken string) (string, error) {
	var userToken model.UserToken

	if err := database.DB.Where("refresh_token = ?", oldRefreshToken).First(&userToken).Error; err != nil {
		return "", fmt.Errorf("refresh token not found")
	}

	if time.Now().After(userToken.RefreshExpiredTime) {
		return "", fmt.Errorf("refresh token expired")
	}

	newAccessToken := generateToken(int(userToken.UserID))
	userToken.Token = newAccessToken
	userToken.ExpiredTime = time.Now().Add(1 * time.Hour)

	// update token di database
	if err := database.DB.Save(&userToken).Error; err != nil {
		return "", fmt.Errorf("failed to update access token: %v", err)
	}

	return newAccessToken, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &user, nil
}

func SaveToken(userID uint, accessToken, refreshToken string) error {
	userToken := model.UserToken{
		UserID:             userID,
		Token:              accessToken,
		RefreshToken:       refreshToken,
		CreatedDate:        time.Now(),
		ExpiredTime:        time.Now().Add(1 * time.Hour),      // access token 1 jam
		RefreshExpiredTime: time.Now().Add(7 * 24 * time.Hour), // refresh token 7 hari
	}

	if err := database.DB.Create(&userToken).Error; err != nil {
		return fmt.Errorf("failed to save token: %v", err)
	}

	return nil
}

func generateToken(userID int) string {
	return fmt.Sprintf("token-%d-%d", userID, time.Now().Unix())
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
