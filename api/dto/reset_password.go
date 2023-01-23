package dto

type ResetPassword struct {
	Code     int    `json:"code"`
	Password string `json:"password"`
}
