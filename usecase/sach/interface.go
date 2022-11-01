package sach

import "daijoubuteam.xyz/se214-library-management/core/entity"

type SachUsecase interface {
	GetDanhSachSach() ([]*entity.Sach, error)
	GetSach(maSach *entity.ID) (*entity.Sach, error)
	CreateSach(maDauSach *entity.ID, nhaXuatBan string, soLuong uint, triGia uint, namXuatBan uint) (*entity.Sach, error)
	UpdateSach(maSach *entity.ID, maDauSach *entity.ID, nhaXuatBan string, soLuong uint, triGia uint, namXuatBan uint) (*entity.Sach, error)
	RemoveSach(maSach *entity.ID) error
}
