package v1dto

import "user-management-api/internal/utils"

// ErrorResponse documents the standard error response shape.
type ErrorResponse struct {
	Error  string `json:"error" example:"Invalid request"`
	Detail string `json:"detail,omitempty" example:"Validation failed"`
}

// MessageResponse documents success responses that contain only status/message.
type MessageResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Operation completed successfully"`
}

// LoginSuccessResponse documents a successful authentication response.
type LoginSuccessResponse struct {
	Status  string        `json:"status" example:"success"`
	Message string        `json:"message" example:"Login successfully"`
	Data    LoginResponse `json:"data"`
}

// UserSuccessResponse documents a response containing a single user.
type UserSuccessResponse struct {
	Status  string  `json:"status" example:"success"`
	Message string  `json:"message" example:"User returned successfully"`
	Data    UserDTO `json:"data"`
}

// UsersListData documents paginated user list data.
type UsersListData struct {
	Data       []UserDTO         `json:"data"`
	Pagination *utils.Pagination `json:"pagination"`
}

// UsersListSuccessResponse documents a paginated users response.
type UsersListSuccessResponse struct {
	Status  string        `json:"status" example:"success"`
	Message string        `json:"message" example:"User list successfully"`
	Data    UsersListData `json:"data"`
}
