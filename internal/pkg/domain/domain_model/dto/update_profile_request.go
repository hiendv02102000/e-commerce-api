package dto

type UpdateProfileRequest struct {
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`
}
