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

func (service *AuthService) Login(email string, password string) (*string, error) {

	dsThuThu, err := service.thuThuUsecase.GetDanhSachThuThu(&email, nil)
	if err != nil {
		return nil, err
	}
	if len(dsThuThu) == 0 || dsThuThu[0] == nil {
		return nil, businessError.NewBusinessError("invalid email or password")
	}

	thuThu := dsThuThu[0]

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
