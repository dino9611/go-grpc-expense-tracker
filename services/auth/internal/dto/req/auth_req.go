package req

type AuthReqDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthLoginReqDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
