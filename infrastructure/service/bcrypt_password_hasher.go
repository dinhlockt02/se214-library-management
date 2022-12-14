package service

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct {
}

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (passwordHasher *BcryptPasswordHasher) HashPassword(password string) (string, error) {
	bytesPassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", coreerror.NewInternalServerError("password hasher: hash password failed", err)
	}
	return string(hashedPassword), nil
}
func (passwordHasher *BcryptPasswordHasher) VerifyPassword(rawPassword string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return false, coreerror.NewBadRequestError("password not match", err)
	} else {
		return true, nil
	}
}
