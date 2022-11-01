package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type TheLoaiRepository interface {
	GetDanhSachTheLoai() ([]*entity.TheLoai, error)
	GetTheLoai(maTheLoai *entity.ID) (*entity.TheLoai, error)
	CreateTheLoai(theLoai *entity.TheLoai) (*entity.TheLoai, error)
	UpdateTheLoai(theLoai *entity.TheLoai) (*entity.TheLoai, error)
	RemoveTheLoai(maTheLoai *entity.ID) error
}
