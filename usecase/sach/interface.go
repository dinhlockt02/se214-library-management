package sach

import "daijoubuteam.xyz/se214-library-management/core/entity"

type SachUsecase interface {
	GetDanhSachSach() ([]*entity.Sach, error)
	GetSach(maSach *entity.ID) (*entity.Sach, error)
	UpdateSach(maSach *entity.ID, nhaXuatBan string, triGia uint, namXuatBan uint, tinhTrang bool, ghiChu string) (*entity.Sach, error)
}
