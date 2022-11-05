package repository

import "github.com/stretchr/testify/mock"

type MockThamSoRepository struct {
	mock.Mock
}

func (mock *MockThamSoRepository) GetThoiHanThe() (uint, error) {
	args := mock.Called()
	return uint(args.Int(0)), args.Error(1)
}

func (mock *MockThamSoRepository) GetTuoiToiThieu() (uint, error) {
	args := mock.Called()
	return uint(args.Int(0)), args.Error(1)
}

func (mock *MockThamSoRepository) GetTuoiToiDa() (uint, error) {
	args := mock.Called()
	return uint(args.Int(0)), args.Error(1)
}

func (mock *MockThamSoRepository) GetDefaultPassword() (string, error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}
