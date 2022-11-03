package auth

import "daijoubuteam.xyz/se214-library-management/core/entity"

type AuthUsecase interface {
	Login(email string, password string) (*string, error)
	VerifyToken(token string) (*entity.ThuThu, error)
}
