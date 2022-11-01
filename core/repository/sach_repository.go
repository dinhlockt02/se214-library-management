package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type SachRepository interface {
	GetDanhSachSach() ([]*entity.Sach, error)
	GetSach(maSach *entity.ID) (*entity.Sach, error)
	CreateSach(sach *entity.Sach) (*entity.Sach, error)
	UpdateSach(sach *entity.Sach) (*entity.Sach, error)
	RemoveSach(maSach *entity.ID) error
}
