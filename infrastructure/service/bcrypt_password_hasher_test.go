package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("it should return hashed password", func(t *testing.T) {
		passwordHasher := NewBcryptPasswordHasher()
		mockPassword := `mock-password`
		rs, err := passwordHasher.HashPassword(mockPassword)
		if assert.Nil(t, err) {
			assert.NotEqual(t, mockPassword, rs)
		}
	})
}

func TestVerifyPassword(t *testing.T) {
	t.Run("it should verify password correctly", func(t *testing.T) {
		passwordHasher := NewBcryptPasswordHasher()
		mockPassword := `mock-password`
		hashedPassword, err := passwordHasher.HashPassword(mockPassword)
		if !assert.Nil(t, err) {
			return
		}
		isMatch, err := passwordHasher.VerifyPassword(mockPassword, hashedPassword)
		if assert.Nil(t, err) {
			assert.Equal(t, isMatch, true)
		}
	})
	t.Run("it should verify password fail if password not match", func(t *testing.T) {
		passwordHasher := NewBcryptPasswordHasher()
		mockPassword := `mock-password`
		hashedPassword, err := passwordHasher.HashPassword(mockPassword)
		if !assert.Nil(t, err) {
			return
		}
		isMatch, err := passwordHasher.VerifyPassword(mockPassword+`failed`, hashedPassword)
		if assert.NotNil(t, err) {
			assert.Equal(t, isMatch, false)
		}
	})
}
