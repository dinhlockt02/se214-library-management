package thuthu

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	"github.com/stretchr/testify/mock"
)

type MockThuThuUsecase struct {
	mock.Mock
}

func (mock *MockThuThuUsecase) GetDanhSachThuThu(email *string, phoneNumber *string) ([]*entity.ThuThu, error) {
	args := mock.Called(email, phoneNumber)
	return args.Get(0).([]*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuUsecase) GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error) {
	args := mock.Called(maThuThu)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuUsecase) GetThuThuByEmail(email string) (*entity.ThuThu, error) {
	args := mock.Called(email)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuUsecase) CreateThuThu(
	name string,
	ngaySinh *time.Time,
	email string,
	phoneNumber string,
	status bool,
	isAdminRole bool,
	password string,
) (*entity.ThuThu, error) {
	args := mock.Called(name, ngaySinh, email, phoneNumber, status, isAdminRole, password)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuUsecase) UpdateThuThu(maThuThu *entity.ID, name string, ngaySinh *time.Time, email string, phoneNumber string, status bool) (*entity.ThuThu, error) {
	args := mock.Called(maThuThu, name, ngaySinh, email, phoneNumber, status)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuUsecase) ChangePassword(maThuThu *entity.ID, newPassword string) (*entity.ThuThu, error) {
	args := mock.Called(maThuThu, newPassword)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}
