package auth

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
)

type AuthService struct {
	thuThuUsecase   thuthu.ThuThuUsecase
	passwordHasher  coreservice.PasswordHasher
	jwtTokenService coreservice.JwtTokenService
}

func NewAuthService(
	thuThuUsecase thuthu.ThuThuUsecase,
	passwordHasher coreservice.PasswordHasher,
	jwtTokenService coreservice.JwtTokenService,
) *AuthService {
	return &AuthService{
		thuThuUsecase:   thuThuUsecase,
		passwordHasher:  passwordHasher,
		jwtTokenService: jwtTokenService,
	}
}

func (service *AuthService) Login(email string, password string) (*string, error) {

	thuThu, err := service.thuThuUsecase.GetThuThuByEmail(email)
	if err != nil {
		return nil, err
	}

	isPasswordMatch, err := service.passwordHasher.VerifyPassword(password, thuThu.Password)

	if err != nil {
		return nil, err
	}

	if !isPasswordMatch {
		return nil, businessError.NewBusinessError("invalid email or password")
	}

	token, err := service.jwtTokenService.Encode(thuThu.MaThuThu)

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (service *AuthService) VerifyToken(token string) (*entity.ThuThu, error) {
	maThuThu, err := service.jwtTokenService.Decode(token)
	if err != nil {
		return nil, err
	}

	thuThu, err := service.thuThuUsecase.GetThuThu(maThuThu)

	if err != nil {
		return nil, err
	}

	return thuThu, nil
}
