package v1handler

import (
	"net/http"
	v1dto "user-management-api/internal/dto/v1"
	v1service "user-management-api/internal/service/v1"
	"user-management-api/internal/utils"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service v1service.UserService
}

func NewUserHandler(service v1service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetAllUsers returns paginated active users.
// @Summary List users
// @Description Get a paginated list of active users.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param search query string false "Search keyword" minlength(3) maxlength(50)
// @Param page query int false "Page number" minimum(1)
// @Param limit query int false "Items per page" minimum(1) maximum(500)
// @Param order_by query string false "Order field" Enums(user_id,user_created_at)
// @Param sort query string false "Sort direction" Enums(asc,desc)
// @Success 200 {object} v1dto.UsersListSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users [get]
func (uh *UserHandler) GetAllUsers(ctx *gin.Context) {
	var params v1dto.GetUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	users, total, err := uh.service.GetAllUsers(ctx, params.Search, params.Order, params.Sort, params.Page, params.Limit, false)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	usersDTO := v1dto.MapUsersToDTO(users)

	paginationResp := utils.NewPagiantionResponse(usersDTO, params.Page, params.Limit, total)

	utils.ResponseSuccess(ctx, http.StatusOK, "User list successfully", paginationResp)
}

// CreateUser creates a new user.
// @Summary Create user
// @Description Create a user account.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param payload body v1dto.CreateUserInput true "Create user payload"
// @Success 201 {object} v1dto.UserSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 409 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users [post]
func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input v1dto.CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	user := input.MapCreateInputToModel()

	createdUser, err := uh.service.CreateUser(ctx, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(createdUser)

	utils.ResponseSuccess(ctx, http.StatusCreated, "User created successfully", userDTO)
}

// GetUserByUUID returns a user by UUID.
// @Summary Get user by UUID
// @Description Get one active user by UUID.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param uuid path string true "User UUID" format(uuid)
// @Success 200 {object} v1dto.UserSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 404 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users/{uuid} [get]
func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	var params v1dto.GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	user, err := uh.service.GetUserByUuid(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(user)

	utils.ResponseSuccess(ctx, http.StatusOK, "Get user successfully", userDTO)
}

// GetUserSoftDeleted returns paginated soft-deleted users.
// @Summary List soft-deleted users
// @Description Get a paginated list of soft-deleted users.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param search query string false "Search keyword" minlength(3) maxlength(50)
// @Param page query int false "Page number" minimum(1)
// @Param limit query int false "Items per page" minimum(1) maximum(500)
// @Param order_by query string false "Order field" Enums(user_id,user_created_at)
// @Param sort query string false "Sort direction" Enums(asc,desc)
// @Success 200 {object} v1dto.UsersListSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users/soft-deleted [get]
func (uh *UserHandler) GetUserSoftDeleted(ctx *gin.Context) {
	var params v1dto.GetUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	users, total, err := uh.service.GetAllUsers(ctx, params.Search, params.Order, params.Sort, params.Page, params.Limit, true)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	usersDTO := v1dto.MapUsersToDTO(users)

	paginationResp := utils.NewPagiantionResponse(usersDTO, params.Page, params.Limit, total)

	utils.ResponseSuccess(ctx, http.StatusOK, "List user soft deleted successfully", paginationResp)
}

// UpdateUser updates a user by UUID.
// @Summary Update user
// @Description Update an existing user by UUID.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param uuid path string true "User UUID" format(uuid)
// @Param payload body v1dto.UpdateUserInput true "Update user payload"
// @Success 200 {object} v1dto.UserSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 404 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users/{uuid} [put]
func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var params v1dto.GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	var input v1dto.UpdateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	user := input.MapUpdateInputToModel(userUuid)

	updatedUser, err := uh.service.UpdateUser(ctx, user)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(updatedUser)

	utils.ResponseSuccess(ctx, http.StatusOK, "User updated successfully", userDTO)
}

// SoftDeleteUser soft deletes a user by UUID.
// @Summary Soft delete user
// @Description Soft delete an active user by UUID.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param uuid path string true "User UUID" format(uuid)
// @Success 200 {object} v1dto.UserSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 404 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users/{uuid} [delete]
func (uh *UserHandler) SoftDeleteUser(ctx *gin.Context) {
	var params v1dto.GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	softDeleteUser, err := uh.service.SoftDeleteUser(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(softDeleteUser)

	utils.ResponseSuccess(ctx, http.StatusOK, "User deleted successfully", userDTO)
}

// RestoreUser restores a soft-deleted user by UUID.
// @Summary Restore user
// @Description Restore a soft-deleted user by UUID.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param uuid path string true "User UUID" format(uuid)
// @Success 200 {object} v1dto.UserSuccessResponse
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 404 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users/{uuid}/restore [put]
func (uh *UserHandler) RestoreUser(ctx *gin.Context) {
	var params v1dto.GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	softDeleteUser, err := uh.service.RestoreUser(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	userDTO := v1dto.MapUserToDTO(softDeleteUser)

	utils.ResponseSuccess(ctx, http.StatusOK, "Restore user successfully", userDTO)
}

// DeleteUser permanently deletes a user by UUID.
// @Summary Permanently delete user
// @Description Permanently delete a user from trash by UUID.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param uuid path string true "User UUID" format(uuid)
// @Success 204 "No Content"
// @Failure 400 {object} v1dto.ErrorResponse
// @Failure 401 {object} v1dto.ErrorResponse
// @Failure 404 {object} v1dto.ErrorResponse
// @Failure 500 {object} v1dto.ErrorResponse
// @Router /users/{uuid}/trash [delete]
func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var params v1dto.GetUserByUuidParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandleValidationErrors(err))
		return
	}

	userUuid, err := uuid.Parse(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	err = uh.service.DeleteUser(ctx, userUuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseStatusCode(ctx, http.StatusNoContent)
}
