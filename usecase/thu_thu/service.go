package thuthu

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
	coreservice "daijoubuteam.xyz/se214-library-management/core/service"
)

type ThuThuService struct {
	passwordHasher coreservice.PasswordHasher
	thuThuRepo     repository.ThuThuRepository
	thamSoRepo     repository.ThamSoRepository
}

func NewThuThuService(
	passwordHasher coreservice.PasswordHasher,
	thuThuRepo repository.ThuThuRepository,
	thamSoRepo repository.ThamSoRepository,
) *ThuThuService {
	return &ThuThuService{
		passwordHasher: passwordHasher,
		thuThuRepo:     thuThuRepo,
		thamSoRepo:     thamSoRepo,
	}
}

func (service *ThuThuService) GetDanhSachThuThu() ([]*entity.ThuThu, error) {

	danhSachThuThu, err := service.thuThuRepo.GetDanhSachThuThu()

	if err != nil {
		return nil, err
	}

	return danhSachThuThu, nil
}

func (service *ThuThuService) GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error) {
	thuThu, err := service.thuThuRepo.GetThuThu(maThuThu)

	if err != nil {
		return nil, err
	}

	if thuThu == nil {
		return nil, coreerror.NewNotFoundError("thu thu not found", nil)
	}

	return thuThu, nil
}

func (service *ThuThuService) CreateThuThu(name string, ngaySinh *time.Time, email string, phoneNumber string, status bool, isAdminRole bool, password string) (*entity.ThuThu, error) {

	var err error

	if len(password) == 0 {
		password, err = service.thamSoRepo.GetDefaultPassword()
	}

	if err != nil {
		return nil, err
	}

	hashedPassword, err := service.passwordHasher.HashPassword(password)

	if err != nil {
		return nil, err
	}

	thuThu := entity.NewThuThu(name, ngaySinh, email, phoneNumber, hashedPassword, status, isAdminRole)
	thuThu, err = service.thuThuRepo.CreateThuThu(thuThu)

	if err != nil {
		return nil, err
	}

	return thuThu, nil
}

func (service *ThuThuService) UpdateThuThu(maThuThu *entity.ID, name string, ngaySinh *time.Time, email string, phoneNumber string, status bool) (*entity.ThuThu, error) {
	thuThu, err := service.GetThuThu(maThuThu)

	if err != nil {
		return nil, err
	}

	thuThu.Name = name
	thuThu.NgaySinh = ngaySinh
	thuThu.Email = email
	thuThu.PhoneNumber = phoneNumber
	thuThu.Status = status

	thuThu, err = service.thuThuRepo.UpdateThuThu(thuThu)

	if err != nil {
		return nil, err
	}

	return thuThu, nil
}

func (service *ThuThuService) ChangePassword(maThuThu *entity.ID, newPassword string) (*entity.ThuThu, error) {
	thuThu, err := service.GetThuThu(maThuThu)

	if err != nil {
		return nil, err
	}

	hashedPassword, err := service.passwordHasher.HashPassword(newPassword)

	if err != nil {
		return nil, err
	}

	thuThu.Password = hashedPassword

	thuThu, err = service.thuThuRepo.UpdateThuThu(thuThu)

	if err != nil {
		return nil, err
	}

	return thuThu, nil
}

func (service *ThuThuService) GetThuThuByEmail(email string) (*entity.ThuThu, error) {
	thuThu, err := service.thuThuRepo.GetThuThuByEmail(email)
	if err != nil {
		return nil, err
	}
	if thuThu == nil {
		return nil, coreerror.NewNotFoundError("Thu thu not found", nil)
	}
	return thuThu, nil
}
