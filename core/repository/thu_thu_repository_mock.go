package repository

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"github.com/stretchr/testify/mock"
)

type MockThuThuRepository struct {
	mock.Mock
}

func (mock *MockThuThuRepository) GetDanhSachThuThu(query *ThuThuSearchQuery) ([]*entity.ThuThu, error) {
	args := mock.Called(query)
	return args.Get(0).([]*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuRepository) GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error) {
	args := mock.Called(maThuThu)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuRepository) CreateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error) {
	args := mock.Called(thuThu)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}

func (mock *MockThuThuRepository) UpdateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error) {
	args := mock.Called(thuThu)
	return args.Get(0).(*entity.ThuThu), args.Error(1)
}
