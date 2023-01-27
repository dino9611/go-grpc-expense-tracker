package dto

type AuthReq struct {
	Username string `json:"username" validate:"required,alphanumunicode"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanumunicode"`
}

type AuthRes struct {
	Id       int64  `json:"id" `
	Email    string `json:"email" `
	Username string `json:"password"`
}
