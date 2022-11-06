package service

import (
	"golang.org/x/crypto/openpgp/errors"
	"time"

	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/golang-jwt/jwt/v4"
)

type JwtTokenServiceImpl struct {
	jwtConfig config.JwtConfig
}

func NewJwtTokenServiceImpl(jwtConfig config.JwtConfig) *JwtTokenServiceImpl {
	return &JwtTokenServiceImpl{
		jwtConfig: jwtConfig,
	}
}

func (service *JwtTokenServiceImpl) Encode(sub *entity.ID) (string, error) {

	if sub == nil {
		return "", coreerror.NewInternalServerError("jwt token service: id not found", errors.InvalidArgumentError("sub argument must not be nil"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Audience: jwt.ClaimStrings{
			service.jwtConfig.Audience,
		},
		Issuer:  service.jwtConfig.Issuer,
		Subject: sub.String(),
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(service.jwtConfig.ExpDuration),
		},
		NotBefore: &jwt.NumericDate{
			Time: time.Now(),
		},
		IssuedAt: &jwt.NumericDate{
			Time: time.Now(),
		},
	})
	tokenString, err := token.SignedString(service.jwtConfig.Secret)
	if err != nil {
		return "", coreerror.NewInternalServerError("jwt token service: sign token failed", err)
	}
	return tokenString, nil
}
func (service *JwtTokenServiceImpl) Decode(token string) (*entity.ID, error) {

	decodedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return service.jwtConfig.Secret, nil
	})

	if err != nil {
		return nil, coreerror.NewBadRequestError("Invalid jwt token", err)
	}
	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, coreerror.NewBadRequestError("Invalid jwt token", errors.SignatureError("polluted token"))
	}
	sub := claims["sub"].(string)
	id, err := entity.StringToID(sub)
	if err != nil {
		return nil, coreerror.NewBadRequestError("Invalid jwt token", errors.SignatureError("invalid id"))
	}
	return id, nil
}
