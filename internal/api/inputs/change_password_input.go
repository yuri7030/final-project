package inputs

type ChangePasswordInput struct {
	Email string `json:"email" binding:"required,email"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}
