package thuthu

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type ThuThuUsecase interface {
	GetDanhSachThuThu(email *string, phoneNumber *string) ([]*entity.ThuThu, error)
	GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error)
	GetThuThuByEmail(email string) (*entity.ThuThu, error)
	CreateThuThu(name string, ngaySinh *time.Time, email string, phoneNumber string, status bool, isAdminRole bool, password string) (*entity.ThuThu, error)
	UpdateThuThu(maThuThu *entity.ID, name string, ngaySinh *time.Time, email string, phoneNumber string, status bool) (*entity.ThuThu, error)
	ChangePassword(maThuThu *entity.ID, newPassword string) (*entity.ThuThu, error)
}
