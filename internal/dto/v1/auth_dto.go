package v1dto

type LoginInput struct {
	Email    string `json:"email" binding:"required,email,email_advanced" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"Password123!"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"refresh-token-value"`
}

type RequestPasswordInput struct {
	Email string `json:"email" binding:"required,email,email_advanced" example:"user@example.com"`
}

type ResetPasswordInput struct {
	Token       string `json:"token" binding:"required" example:"reset-token-value"`
	NewPassword string `json:"new_password" binding:"required,min=8" example:"NewPassword123!"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"access-token-value"`
	RefreshToken string `json:"refresh_token" example:"refresh-token-value"`
	ExpiresIn    int    `json:"expires_in" example:"900"`
}
