package dausach

import "daijoubuteam.xyz/se214-library-management/core/entity"

type DauSachUsecase interface {
	GetDanhSachDauSach() ([]*entity.DauSach, error)
	GetDauSach(maDauSach *entity.ID) (*entity.DauSach, error)
	CreateDauSach(tenDauSach string, maTheLoai []*entity.ID, maTacGia []*entity.ID) (*entity.DauSach, error)
	UpdateDauSach(maDauSach *entity.ID, tenDauSach string, maTheLoai []*entity.ID, maTacGia []*entity.ID) (*entity.DauSach, error)
	RemoveDauSach(maDauSach *entity.ID) error
}
