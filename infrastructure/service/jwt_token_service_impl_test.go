package service

import (
	"regexp"
	"testing"
	"time"

	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {

	t.Run("it should return valid jwt token if valid id provided", func(t *testing.T) {
		newID := entity.NewID()
		jwtConfig := config.JwtConfig{
			Audience: "foo",
			Secret:   []byte("foo"),
		}
		jwtTokenService := NewJwtTokenServiceImpl(jwtConfig)
		encodedID, err := jwtTokenService.Encode(&newID)

		if assert.Nil(t, err, "TestEncode: err thrown") {
			r, _ := regexp.Compile(`.*\..*\..*`)
			assert.Equal(t, true, r.MatchString(encodedID), "TestEncode: invalid encode id")
		}
	})

	t.Run("it should return nil, err if id is nil", func(t *testing.T) {
		jwtConfig := config.JwtConfig{
			Audience: "foo",
			Secret:   []byte("foo"),
		}
		jwtTokenService := NewJwtTokenServiceImpl(jwtConfig)
		_, err := jwtTokenService.Encode(nil)

		if assert.NotNil(t, err) {
			_, ok := err.(*coreerror.InternalServerError)
			assert.Equal(t, ok, true)
		}
	})

}

func TestDecode(t *testing.T) {
	t.Run("it should return id token if a valid token provided", func(t *testing.T) {
		newId := entity.NewID()
		jwtConfig := config.JwtConfig{
			Audience:    "foo",
			Secret:      []byte("foo"),
			Issuer:      "bar",
			ExpDuration: time.Duration(1) * time.Minute,
		}
		jwtTokenService := NewJwtTokenServiceImpl(jwtConfig)
		token, err := jwtTokenService.Encode(&newId)
		if err != nil {
			assert.FailNow(t, "set up test case failed")
		}
		id, err := jwtTokenService.Decode(token)

		if assert.Nil(t, err, "err should be nil") {
			assert.Equal(t, newId.String(), id.String(), "id is not match")
		}
	})

	t.Run("it should return nil, err if invalid token provided", func(t *testing.T) {
		newId := entity.NewID()
		jwtConfig := config.JwtConfig{
			Audience: "foo",
			Secret:   []byte("foo"),
		}
		jwtTokenService := NewJwtTokenServiceImpl(jwtConfig)
		token, err := jwtTokenService.Encode(&newId)
		if err != nil {
			assert.FailNow(t, "set up test case failed")
		}
		_, err = jwtTokenService.Decode(token + `i`)

		if assert.NotNil(t, err, "err should not be nil") {
			_, ok := err.(*coreerror.BadRequestError)
			assert.Equal(t, ok, true)
		}
	})

}
