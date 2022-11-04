package auth

import (
	"errors"
	"testing"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	mockThuThuUsecase := &thuthu.MockThuThuUsecase{}
	mockPasswordHasher := &coreservice.MockPasswordHasher{}
	mockJwtTokenService := &coreservice.MockJwtTokenService{}
	authService := NewAuthService(mockThuThuUsecase, mockPasswordHasher, mockJwtTokenService)
	t.Run("it should return a valid token if login info is correct", func(t *testing.T) {

		mockEmail := "mock@email.com"
		mockPassword := "mock-password"
		mockToken := "mock.token.mock"
		mockId := entity.NewID()
		mockThuThuUsecase.On("GetDanhSachThuThu", &mockEmail, (*string)(nil)).Return([]*entity.ThuThu{
			{
				MaThuThu: &mockId,
				Email:    mockEmail,
				Password: mockPassword,
			},
		}, nil).Once()
		mockPasswordHasher.On("VerifyPassword", mockPassword, mockPassword).Return(true, nil).Once()
		mockJwtTokenService.On("Encode", &mockId).Return(mockToken, nil).Once()
		rs, err := authService.Login(mockEmail, mockPassword)
		assert.Equal(t, *rs, mockToken)
		assert.Nil(t, err)
	})
	t.Run("it should return a error if find no thu thu", func(t *testing.T) {

		mockEmail := "mock@email.com"
		mockPassword := "mock-password"
		mockThuThuUsecase.On("GetDanhSachThuThu", &mockEmail, (*string)(nil)).Return([]*entity.ThuThu{}, nil).Once()
		rs, err := authService.Login(mockEmail, mockPassword)
		assert.Nil(t, rs)
		_, ok := err.(*coreerror.BusinessError)
		assert.Equal(t, true, ok)

	})
	t.Run("it should return a error if password not match", func(t *testing.T) {

		mockEmail := "mock@email.com"
		mockPassword := "mock-password"
		mockId := entity.NewID()
		mockThuThuUsecase.On("GetDanhSachThuThu", &mockEmail, (*string)(nil)).Return([]*entity.ThuThu{{
			MaThuThu: &mockId,
			Email:    mockEmail,
			Password: mockPassword,
		}}, nil).Once()
		mockPasswordHasher.On("VerifyPassword", mockPassword, mockPassword).Return(false, nil).Once()
		rs, err := authService.Login(mockEmail, mockPassword)
		assert.Nil(t, rs)
		_, ok := err.(*coreerror.BusinessError)
		assert.Equal(t, true, ok)

	})
	t.Run("it should return a error if password hasher throw error", func(t *testing.T) {

		mockEmail := "mock@email.com"
		mockPassword := "mock-password"
		mockId := entity.NewID()
		mockThuThuUsecase.On("GetDanhSachThuThu", &mockEmail, (*string)(nil)).Return([]*entity.ThuThu{{
			MaThuThu: &mockId,
			Email:    mockEmail,
			Password: mockPassword,
		}}, nil).Once()
		mockPasswordHasherErrorMessage := "mock-password-hasher-error-message"
		mockPasswordHasher.On("VerifyPassword", mockPassword, mockPassword).Return(false, errors.New(mockPasswordHasherErrorMessage)).Once()
		rs, err := authService.Login(mockEmail, mockPassword)
		assert.Nil(t, rs)
		assert.Error(t, err)
		assert.Equal(t, mockPasswordHasherErrorMessage, err.Error())
	})
	t.Run("it should return error if could not encode thuThu.MaThuThu", func(t *testing.T) {
		mockEmail := "mock@email.com"
		mockPassword := "mock-password"
		mockId := entity.NewID()
		mockEncodingErrorMessage := `mock-encoding-error-message`
		mockThuThuUsecase.On("GetDanhSachThuThu", &mockEmail, (*string)(nil)).Return([]*entity.ThuThu{
			{
				MaThuThu: &mockId,
				Email:    mockEmail,
				Password: mockPassword,
			},
		}, nil).Once()
		mockPasswordHasher.On("VerifyPassword", mockPassword, mockPassword).Return(true, nil).Once()
		mockJwtTokenService.On("Encode", &mockId).Return("", errors.New(mockEncodingErrorMessage)).Once()
		rs, err := authService.Login(mockEmail, mockPassword)
		assert.Nil(t, rs)
		assert.Error(t, err)
		assert.Equal(t, mockEncodingErrorMessage, err.Error())
	})
}

func TestVerifyToken(t *testing.T) {
	mockThuThuUsecase := &thuthu.MockThuThuUsecase{}
	mockPasswordHasher := &coreservice.MockPasswordHasher{}
	mockJwtTokenService := &coreservice.MockJwtTokenService{}
	authService := NewAuthService(mockThuThuUsecase, mockPasswordHasher, mockJwtTokenService)
	t.Run("it should return *entity.ThuThu if token is valid", func(t *testing.T) {
		mockToken := `mock.token.mock`
		mockId := entity.NewID()
		mockThuThu := &entity.ThuThu{
			MaThuThu: &mockId,
		}
		mockJwtTokenService.On("Decode", mockToken).Return(&mockId, nil).Once()
		mockThuThuUsecase.On("GetThuThu", &mockId).Return(mockThuThu, nil).Once()
		rs, err := authService.VerifyToken(mockToken)
		if assert.Nil(t, err) {
			assert.Equal(t, mockThuThu, rs)
		}
	})
	t.Run("it should return error if cant not decode token", func(t *testing.T) {
		mockToken := `mock.token.mock`
		mockDecodeErrorMessage := `mock-decode-error-message`
		mockJwtTokenService.On("Decode", mockToken).Return((*entity.ID)(nil), errors.New(mockDecodeErrorMessage)).Once()
		rs, err := authService.VerifyToken(mockToken)
		assert.Error(t, err)
		assert.Nil(t, rs)
		assert.Equal(t, mockDecodeErrorMessage, err.Error())
	})
	t.Run("it should return error if get thu thu throw error", func(t *testing.T) {
		mockToken := `mock.token.mock`
		mockId := entity.NewID()
		mockGetThuThuErrorMessage := `mock-get-thu-thu-error-message`
		mockJwtTokenService.On("Decode", mockToken).Return(&mockId, nil).Once()
		mockThuThuUsecase.On("GetThuThu", &mockId).Return((*entity.ThuThu)(nil), errors.New(mockGetThuThuErrorMessage)).Once()
		rs, err := authService.VerifyToken(mockToken)
		assert.Error(t, err)
		assert.Nil(t, rs)
		assert.Equal(t, mockGetThuThuErrorMessage, err.Error())
	})
}
