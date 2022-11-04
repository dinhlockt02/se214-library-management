package coreservice

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"github.com/stretchr/testify/mock"
)

type MockJwtTokenService struct {
	mock.Mock
}

func (mock *MockJwtTokenService) Encode(sub *entity.ID) (string, error) {
	args := mock.Called(sub)
	return args.String(0), args.Error(1)
}

func (mock *MockJwtTokenService) Decode(token string) (*entity.ID, error) {
	args := mock.Called(token)
	return args.Get(0).(*entity.ID), args.Error(1)
}

type MockPasswordHasher struct {
	mock.Mock
}

func (mock *MockPasswordHasher) HashPassword(password string) (string, error) {
	args := mock.Called(password)
	return args.String(0), args.Error(1)
}

func (mock *MockPasswordHasher) VerifyPassword(rawPassword string, hashedPassword string) (bool, error) {
	args := mock.Called(rawPassword, hashedPassword)
	return args.Bool(0), args.Error(1)
}
