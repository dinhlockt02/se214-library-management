package thuthu

import (
	"errors"
	"testing"
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDanhSachPhuThu(t *testing.T) {

	passwordHasher := &coreservice.MockPasswordHasher{}
	thuThuRepo := &repository.MockThuThuRepository{}
	thamSoRepo := &repository.MockThamSoRepository{}

	thuThuService := NewThuThuService(
		passwordHasher,
		thuThuRepo,
		thamSoRepo,
	)

	t.Run("it should return list of thu thu if no error occur", func(t *testing.T) {
		mockDanhSachThuThu := []*entity.ThuThu{}
		thuThuRepo.On("GetDanhSachThuThu", mock.Anything).Return(mockDanhSachThuThu, nil).Once()
		rs, err := thuThuService.GetDanhSachThuThu()
		assert.Nil(t, err)
		assert.Equal(t, mockDanhSachThuThu, rs)
		thuThuRepo.AssertCalled(t, "GetDanhSachThuThu", mock.MatchedBy(func(i interface{}) bool {
			return true
		}))
	})
	t.Run("it should yield an error query from repository failed", func(t *testing.T) {
		errorMessage := `error-message`
		thuThuRepo.On("GetDanhSachThuThu", mock.Anything).Return([]*entity.ThuThu(nil), errors.New(errorMessage)).Once()
		rs, err := thuThuService.GetDanhSachThuThu()
		assert.Nil(t, rs)
		assert.Error(t, err)
		assert.Equal(t, errorMessage, err.Error())
	})
}

func TestGetThuThu(t *testing.T) {
	passwordHasher := &coreservice.MockPasswordHasher{}
	thuThuRepo := &repository.MockThuThuRepository{}
	thamSoRepo := &repository.MockThamSoRepository{}

	thuThuService := NewThuThuService(
		passwordHasher,
		thuThuRepo,
		thamSoRepo,
	)

	t.Run("it should return thu thu if thu thu is found", func(t *testing.T) {
		mockMaThuThu := entity.NewID()
		mockThuThu := &entity.ThuThu{}
		thuThuRepo.On("GetThuThu", &mockMaThuThu).Return(mockThuThu, nil)
		rs, err := thuThuService.GetThuThu(&mockMaThuThu)

		assert.Nil(t, err)
		assert.Equal(t, mockThuThu, rs)
	})

	t.Run("it should yield an error if thu thu is nil", func(t *testing.T) {
		mockMaThuThu := entity.NewID()

		thuThuRepo.On("GetThuThu", &mockMaThuThu).Return((*entity.ThuThu)(nil), nil)
		rs, err := thuThuService.GetThuThu(&mockMaThuThu)

		_, ok := err.(*coreerror.BusinessError)

		assert.True(t, ok)
		assert.Nil(t, rs)
	})

	t.Run("it should yield an error if thu thu repository yield error", func(t *testing.T) {
		mockMaThuThu := entity.NewID()
		mockErrorMessage := `mock-error-message`
		thuThuRepo.On("GetThuThu", &mockMaThuThu).Return((*entity.ThuThu)(nil), errors.New(mockErrorMessage))
		rs, err := thuThuService.GetThuThu(&mockMaThuThu)

		assert.Error(t, err)
		assert.Equal(t, mockErrorMessage, err.Error())
		assert.Nil(t, rs)
	})
}

func TestCreateThuThu(t *testing.T) {

	passwordHasher := &coreservice.MockPasswordHasher{}
	thuThuRepo := &repository.MockThuThuRepository{}
	thamSoRepo := &repository.MockThamSoRepository{}

	thuThuService := NewThuThuService(
		passwordHasher,
		thuThuRepo,
		thamSoRepo,
	)

	mockDefaultPassword := `mock-default-password`

	t.Run("it should create a thu thu if valid data provided", func(t *testing.T) {

		mockHashedPassword := `mock-hashed-password`

		mockName := `mock-name`
		mockNgaySinh := utils.Ptr(time.Date(2000, 2, 2, 2, 2, 2, 2, time.UTC))
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockIsAdmin := true
		mockStatus := false

		mockThuThu := entity.NewThuThu(mockName, mockNgaySinh, mockEmail, mockPhoneNumber, mockHashedPassword, mockStatus, mockIsAdmin)

		thamSoRepo.On("GetDefaultPassword").Return(mockDefaultPassword, nil).Once()
		passwordHasher.On("HashPassword", mockDefaultPassword).Return(mockHashedPassword, nil).Once()
		thuThuRepo.On("CreateThuThu", mock.Anything).Return(mockThuThu, nil).Once()

		rs, err := thuThuService.CreateThuThu(mockName, mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus, mockIsAdmin, "")

		assert.Nil(t, err)
		assert.Equal(t, mockThuThu, rs)
	})

	t.Run("it should yield an error if cant get default password", func(t *testing.T) {

		mockName := `mock-name`
		mockNgaySinh := utils.Ptr(time.Date(2000, 2, 2, 2, 2, 2, 2, time.UTC))
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockIsAdmin := true
		mockStatus := false

		mockThamSoRepoErrorMessage := `mock-thamso-repo-message`

		thamSoRepo.On("GetDefaultPassword").Return("", errors.New(mockThamSoRepoErrorMessage)).Once()

		rs, err := thuThuService.CreateThuThu(mockName, mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus, mockIsAdmin, "")

		assert.Nil(t, rs)
		assert.Error(t, err)
		assert.Equal(t, mockThamSoRepoErrorMessage, err.Error())
	})

	t.Run("it should yield an error if hash password failed", func(t *testing.T) {

		mockName := `mock-name`
		mockNgaySinh := utils.Ptr(time.Date(2000, 2, 2, 2, 2, 2, 2, time.UTC))
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockIsAdmin := true
		mockStatus := false

		hashPasswordErrorMesage := `hash-password-error-message`

		thamSoRepo.On("GetDefaultPassword").Return(mockDefaultPassword, nil).Once()
		passwordHasher.On("HashPassword", mockDefaultPassword).Return("", errors.New(hashPasswordErrorMesage)).Once()

		rs, err := thuThuService.CreateThuThu(mockName, mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus, mockIsAdmin, "")

		assert.Empty(t, rs)
		assert.Error(t, err)
		assert.Equal(t, hashPasswordErrorMesage, err.Error())
	})

	t.Run("it should yield an error if repository cant create thu thu", func(t *testing.T) {

		mockHashedPassword := `mock-hashed-password`

		mockName := `mock-name`
		mockNgaySinh := utils.Ptr(time.Date(2000, 2, 2, 2, 2, 2, 2, time.UTC))
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockIsAdmin := true
		mockStatus := false

		mockRepositoryError := `mock-repository-error`

		thamSoRepo.On("GetDefaultPassword").Return(mockDefaultPassword, nil).Once()
		passwordHasher.On("HashPassword", mockDefaultPassword).Return(mockHashedPassword, nil).Once()
		thuThuRepo.On("CreateThuThu", mock.Anything).Return((*entity.ThuThu)(nil), errors.New(mockRepositoryError)).Once()

		rs, err := thuThuService.CreateThuThu(mockName, mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus, mockIsAdmin, "")

		assert.Nil(t, rs)
		assert.Error(t, err)
		assert.Equal(t, mockRepositoryError, err.Error())
	})
}

func TestUpdateThuThu(t *testing.T) {
	passwordHasher := &coreservice.MockPasswordHasher{}
	thuThuRepo := &repository.MockThuThuRepository{}
	thamSoRepo := &repository.MockThamSoRepository{}

	thuThuService := NewThuThuService(
		passwordHasher,
		thuThuRepo,
		thamSoRepo,
	)

	t.Run("it should update thu thu correctly", func(t *testing.T) {

		mockMaThuThu := entity.NewID()
		mockName := `mock-name`
		mockNgaySinh := time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockStatus := false

		mockThuThu := &entity.ThuThu{
			MaThuThu: &(mockMaThuThu),
		}

		thuThuRepo.On("GetThuThu", &mockMaThuThu).Return(mockThuThu, nil).Once()

		thuThuRepo.On("UpdateThuThu", mockThuThu).Return(mockThuThu, nil).Once()

		rs, err := thuThuService.UpdateThuThu(&(mockMaThuThu), mockName, &mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus)

		assert.Nil(t, err)
		assert.Equal(t, mockThuThu, rs)
	})

	t.Run("it should yield error if thu thu cant not be found", func(t *testing.T) {

		mockMaThuThu := entity.NewID()
		mockName := `mock-name`
		mockNgaySinh := time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockStatus := false

		mockThuThu := &entity.ThuThu{
			MaThuThu: &(mockMaThuThu),
		}

		thuThuRepo.On("GetThuThu", &mockMaThuThu).Return((*entity.ThuThu)(nil), nil).Once()

		thuThuRepo.On("UpdateThuThu", mockThuThu).Return(mockThuThu, nil).Once()

		rs, err := thuThuService.UpdateThuThu(&(mockMaThuThu), mockName, &mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus)

		assert.Error(t, err)
		assert.Nil(t, rs)
	})

	t.Run("it should yield an error if cant not update thu thu", func(t *testing.T) {

		mockMaThuThu := entity.NewID()
		mockName := `mock-name`
		mockNgaySinh := time.Date(2000, 1, 1, 1, 1, 1, 1, time.UTC)
		mockEmail := `mock@email.com`
		mockPhoneNumber := `0123456789`
		mockStatus := false

		mockThuThu := &entity.ThuThu{
			MaThuThu: &(mockMaThuThu),
		}

		mockUpdateThuThuErr := `mock-update-thu-thu-error`

		thuThuRepo.On("GetThuThu", &mockMaThuThu).Return(mockThuThu, nil).Once()

		thuThuRepo.On("UpdateThuThu", mockThuThu).Return((*entity.ThuThu)(nil), errors.New(mockUpdateThuThuErr)).Once()

		rs, err := thuThuService.UpdateThuThu(&(mockMaThuThu), mockName, &mockNgaySinh, mockEmail, mockPhoneNumber, mockStatus)

		assert.Error(t, err)
		assert.Equal(t, mockUpdateThuThuErr, err.Error())
		assert.Nil(t, rs)
	})
}

func TestChangePassword(t *testing.T) {
	passwordHasher := &coreservice.MockPasswordHasher{}
	thuThuRepo := &repository.MockThuThuRepository{}
	thamSoRepo := &repository.MockThamSoRepository{}

	thuThuService := NewThuThuService(
		passwordHasher,
		thuThuRepo,
		thamSoRepo,
	)
	t.Run("it should update hashed password correctly", func(t *testing.T) {
		mockMaThuThu := utils.Ptr(entity.NewID())
		mockThuThu := &entity.ThuThu{
			MaThuThu: mockMaThuThu,
		}
		mockPassword := `mock-password`
		mockHashedPassword := `mock-hashed-password`
		thuThuRepo.On("GetThuThu", mockMaThuThu).Return(mockThuThu, nil).Once()
		passwordHasher.On("HashPassword", mockPassword).Return(mockHashedPassword, nil).Once()
		thuThuRepo.On("UpdateThuThu", mockThuThu).Return(mockThuThu, nil).Once()

		rs, err := thuThuService.ChangePassword(mockMaThuThu, mockPassword)

		thuThuRepo.AssertCalled(t, "UpdateThuThu", mock.MatchedBy(func(thuThu *entity.ThuThu) bool {
			return thuThu.Password == mockHashedPassword
		}))
		thuThuRepo.AssertNumberOfCalls(t, "UpdateThuThu", 1)

		assert.Nil(t, err)
		assert.Equal(t, mockThuThu, rs)
	})
	t.Run("it should yield error if thu thu not found", func(t *testing.T) {
		mockMaThuThu := utils.Ptr(entity.NewID())
		mockThuThu := &entity.ThuThu{
			MaThuThu: mockMaThuThu,
		}
		mockPassword := `mock-password`
		// mockHashedPassword := `mock-hashed-password`
		thuThuRepo.On("GetThuThu", mockMaThuThu).Return((*entity.ThuThu)(nil), nil).Once()
		// passwordHasher.On("HashPassword", mockPassword).Return(mockHashedPassword, nil).Once()
		thuThuRepo.On("UpdateThuThu", mockThuThu).Return(mockThuThu, nil).Once()

		rs, err := thuThuService.ChangePassword(mockMaThuThu, mockPassword)

		assert.Error(t, err)
		assert.Nil(t, rs)
	})

	t.Run("it should yield error if hash password failed", func(t *testing.T) {
		mockMaThuThu := utils.Ptr(entity.NewID())
		mockThuThu := &entity.ThuThu{
			MaThuThu: mockMaThuThu,
		}
		mockPassword := `mock-password`
		mockHashedPasswordErr := `mock-hash-password-failed`

		thuThuRepo.On("GetThuThu", mockMaThuThu).Return(mockThuThu, nil).Once()
		passwordHasher.On("HashPassword", mockPassword).Return("", errors.New(mockHashedPasswordErr)).Once()
		thuThuRepo.On("UpdateThuThu", mockThuThu).Return(mockThuThu, nil).Once()

		rs, err := thuThuService.ChangePassword(mockMaThuThu, mockPassword)

		assert.Error(t, err)
		assert.Nil(t, rs)
		assert.Equal(t, mockHashedPasswordErr, err.Error())
	})

	t.Run("it should yield error if update failed", func(t *testing.T) {
		mockMaThuThu := utils.Ptr(entity.NewID())
		mockThuThu := &entity.ThuThu{
			MaThuThu: mockMaThuThu,
		}
		mockPassword := `mock-password`
		mockHashedPassword := `mock-hashed-password`
		mockUpdateErr := `mock-update-failed`

		thuThuRepo.On("GetThuThu", mockMaThuThu).Return(mockThuThu, nil).Once()
		passwordHasher.On("HashPassword", mockPassword).Return(mockHashedPassword, nil).Once()
		thuThuRepo.On("UpdateThuThu", mockThuThu).Return((*entity.ThuThu)(nil), errors.New(mockUpdateErr)).Once()

		rs, err := thuThuService.ChangePassword(mockMaThuThu, mockPassword)

		assert.Error(t, err)
		assert.Nil(t, rs)
		assert.Equal(t, mockUpdateErr, err.Error())
	})
}
