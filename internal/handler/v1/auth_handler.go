package v1handler

import (
	"net/http"
	v1dto "user-management-api/internal/dto/v1"
	v1service "user-management-api/internal/service/v1"
	"user-management-api/internal/utils"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service v1service.AuthService
}

func NewAuthHandler(service v1service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// Login authenticates a user and returns access and refresh tokens.
// @Summary Login
// @Description Authenticate a user with email and password.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body v1dto.LoginInput true "Login payload"
// @Success 200 {object} v1dto.LoginSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /auth/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var input v1dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	accessToken, refreshToken, expiresIn, err := ah.service.Login(ctx, input.Email, input.Password)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	response := v1dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Login successfully", response)
}

// Logout revokes a refresh token.
// @Summary Logout
// @Description Revoke the current refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param payload body v1dto.RefreshTokenInput true "Refresh token payload"
// @Success 200 {object} v1dto.MessageResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /auth/logout [post]
func (ah *AuthHandler) Logout(ctx *gin.Context) {
	var input v1dto.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	if err := ah.service.Logout(ctx, input.RefreshToken); err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Logout successfully")
}

// RefreshToken refreshes an access token.
// @Summary Refresh token
// @Description Generate a new access token and refresh token from a valid refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body v1dto.RefreshTokenInput true "Refresh token payload"
// @Success 200 {object} v1dto.LoginSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /auth/refresh [post]
func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	var input v1dto.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	accessToken, refreshToken, expiresIn, err := ah.service.RefreshToken(ctx, input.RefreshToken)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	response := v1dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Refresh token generate successfully", response)
}

// RequestForgotPassword sends a password reset link.
// @Summary Request forgot password
// @Description Send a reset password link to the given email address.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body v1dto.RequestPasswordInput true "Forgot password payload"
// @Success 200 {object} v1dto.MessageResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /auth/forgot-password [post]
func (ah *AuthHandler) RequestForgotPassword(ctx *gin.Context) {
	var input v1dto.RequestPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	err := ah.service.RequestForgotPassword(ctx, input.Email)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Reset link sent to email")
}

// ResetPassword resets a password using a reset token.
// @Summary Reset password
// @Description Reset the user password using a reset token.
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param payload body v1dto.ResetPasswordInput true "Reset password payload"
// @Success 200 {object} v1dto.MessageResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /auth/reset-password [post]
func (ah *AuthHandler) ResetPassword(ctx *gin.Context) {
	var input v1dto.ResetPasswordInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	err := ah.service.ResetPassword(ctx, input.Token, input.NewPassword)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, "Password reset successfully")
}
