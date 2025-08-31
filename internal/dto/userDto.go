package dto

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type LoginResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type LogoutResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Berhasil melakukan logout"`
}

type RegisterRequest struct {
	FirstName       string `json:"firstName" example:"Jhon"`
	LastName        string `json:"lastName" example:"Doe"`
	Email           string `json:"email" example:"example@email.com"`
	Password        string `json:"password" example:"password123"`
	ConfirmPassword string `json:"confirmPassword" example:"password123"`
}

type Response struct {
	Status  string `json:"status" example:"200"`
	Message string `json:"message" example:"Sucess"`
}
