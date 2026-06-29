package dto

type SignupReq struct {
	Username        string `json:"username" binding:"required,min=1,max=64"`
	Email           string `json:"email" binding:"required,email,max=128"`
	Password        string `json:"password" binding:"required,min=6,max=255,password_complexity"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email,max=128"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}
