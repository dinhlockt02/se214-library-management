package auth

import "daijoubuteam.xyz/se214-library-management/core/entity"

type AuthUsecase interface {
	Login(email string, password string) (*string, *string, error)
	VerifyToken(token string) (*entity.ThuThu, error)
	ForgetPassword(email string) error
	ResetPassword(code int, password string) error
}
