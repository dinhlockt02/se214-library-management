package auth

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"math/rand"
)

var resetPasswordCode = make(map[int]string)

type AuthService struct {
	thuThuUsecase    thuthu.ThuThuUsecase
	passwordHasher   coreservice.PasswordHasher
	jwtTokenService  coreservice.JwtTokenService
	sendEmailService coreservice.SendEmailService
}

func NewAuthService(
	thuThuUsecase thuthu.ThuThuUsecase,
	passwordHasher coreservice.PasswordHasher,
	jwtTokenService coreservice.JwtTokenService,
	sendEmailService coreservice.SendEmailService,
) *AuthService {
	return &AuthService{
		thuThuUsecase:    thuThuUsecase,
		passwordHasher:   passwordHasher,
		jwtTokenService:  jwtTokenService,
		sendEmailService: sendEmailService,
	}
}

func (service *AuthService) Login(email string, password string) (*string, *string, error) {

	thuThu, err := service.thuThuUsecase.GetThuThuByEmail(email)
	if err != nil {
		return nil, nil, err
	}

	if isPasswordMatch, err := service.passwordHasher.VerifyPassword(password, thuThu.Password); !isPasswordMatch {
		return nil, nil, coreerror.NewBadRequestError("invalid email or password", err)
	}

	token, err := service.jwtTokenService.Encode(thuThu.MaThuThu)

	if err != nil {
		return nil, nil, err
	}
	maThuThu := thuThu.MaThuThu.String()
	return &token, &maThuThu, nil
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

func generateRandomCode() int {
	return rand.Intn(9999-1000) + 1000
}

func (service *AuthService) ForgetPassword(email string) error {
	t, err := service.thuThuUsecase.GetThuThuByEmail(email)
	if err != nil {
		return err
	}
	var code int
	for {
		code = generateRandomCode()
		if _, ok := resetPasswordCode[code]; !ok {
			resetPasswordCode[code] = email
			break
		}
	}
	return service.sendEmailService.SendResetPasswordMail(t.Name, email, code)
}

func (service *AuthService) ResetPassword(code int, password string) error {
	resetPasswordCode[1234] = "dinhlockt02@gmail.com"
	var email string
	if e, ok := resetPasswordCode[code]; ok {
		email = e
	} else {
		return coreerror.NewBadRequestError("Invalid code", nil)
	}
	t, err := service.thuThuUsecase.GetThuThuByEmail(email)
	if err != nil {
		return err
	}
	_, err = service.thuThuUsecase.ChangePassword(t.MaThuThu, password)
	return err
}
