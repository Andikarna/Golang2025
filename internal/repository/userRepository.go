package repository

import (
	"fmt"
	"fundamental/internal/database"
	"fundamental/internal/dto"
	"fundamental/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func Login(email, password string) (*model.User, string, string, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, "", "", fmt.Errorf("user not found")
	}

	if !CheckPassword(user.Password, password) {
		return nil, "Incorrect password", "", fmt.Errorf("Password Invalid")
	}

	token, err := GenerateJWT(strconv.FormatUint(uint64(user.ID), 10), 3*time.Minute)
	refreshToken, err := GenerateJWT(strconv.FormatUint(uint64(user.ID), 10), 5*time.Minute)

	SaveToken(user.ID, token, refreshToken)

	return user, token, refreshToken, nil
}

func RefreshToken(oldRefreshToken string) (string, string, error) {

	var userToken model.UserToken

	if err := database.DB.Where("refresh_token = ?", oldRefreshToken).First(&userToken).Error; err != nil {
		return "", "", fmt.Errorf("refresh token not found")
	}

	if time.Now().After(userToken.RefreshExpiredTime) {
		return "", "", fmt.Errorf("refresh token expired")
	}

	newAccessToken, err := GenerateJWT(strconv.FormatUint(uint64(userToken.UserID), 10), 3*time.Minute)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := GenerateJWT(strconv.FormatUint(uint64(userToken.UserID), 10), 5*time.Minute)
	if err != nil {
		return "", "", err
	}

	userToken.Token = newAccessToken
	userToken.RefreshToken = newRefreshToken
	userToken.ExpiredTime = time.Now().Add(3 * time.Minute)
	userToken.RefreshExpiredTime = time.Now().Add(5 * time.Minute)

	if err := database.DB.Save(&userToken).Error; err != nil {
		return "", "", fmt.Errorf("failed to update tokens: %v", err)
	}

	return newAccessToken, newRefreshToken, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return &user, nil
}

func SaveToken(userID uint, accessToken, refreshToken string) error {

	var exisToken model.UserToken

	if err := database.DB.Where("user_id = ?", userID).First(&exisToken).Error; err == nil {
		exisToken.Token = accessToken
		exisToken.RefreshToken = refreshToken
		exisToken.UpdateDate = time.Now()
		exisToken.ExpiredTime = time.Now().Add(3 * time.Minute)
		exisToken.RefreshExpiredTime = time.Now().Add(5 * time.Minute)

		if err := database.DB.Save(&exisToken).Error; err != nil {
			return fmt.Errorf("failed to update token: %v", err)
		}
	} else {
		userToken := model.UserToken{
			UserID:             userID,
			Token:              accessToken,
			RefreshToken:       refreshToken,
			CreatedDate:        time.Now(),
			ExpiredTime:        time.Now().Add(3 * time.Hour),
			RefreshExpiredTime: time.Now().Add(5 * time.Minute),
		}

		if err := database.DB.Create(&userToken).Error; err != nil {
			return fmt.Errorf("failed to save token: %v", err)
		}
	}

	return nil
}

var JwtKey = []byte("Key_DailyTask_React_Golang_2025")

func GenerateJWT(userID string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
