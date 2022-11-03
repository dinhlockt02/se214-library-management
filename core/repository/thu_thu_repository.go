package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type ThuThuSearchQuery struct {
	Email       *string
	PhoneNumber *string
}

type ThuThuRepository interface {
	GetDanhSachThuThu(query *ThuThuSearchQuery) ([]*entity.ThuThu, error)
	GetThuThu(maThuThu *entity.ID) (*entity.ThuThu, error)
	CreateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error)
	UpdateThuThu(thuThu *entity.ThuThu) (*entity.ThuThu, error)
}
