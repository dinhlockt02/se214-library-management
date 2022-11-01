package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type DauSachRepository interface {
	GetDanhSachDauSach() ([]*entity.DauSach, error)
	GetDauSach(maDauSach *entity.ID) (*entity.DauSach, error)
	CreateDauSach(dauSach *entity.DauSach) (*entity.DauSach, error)
	UpdateDauSach(dauSach *entity.DauSach) (*entity.DauSach, error)
	RemoveDauSach(maDauSach *entity.ID) error
}
