package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type ThuThuRepository interface {
	GetDanhSachThuThu() ([]*entity.ThuThu, error)
	GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error)
	GetThuThuByEmail(email string) (*entity.ThuThu, error)
	CreateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error)
	UpdateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error)
}
