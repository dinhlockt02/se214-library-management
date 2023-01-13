package nhapsach

import (
	"time"

	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type NhapSachUsecase interface {
	GetDanhSachPhieuNhapSach() ([]*entity.PhieuNhap, error)
	GetPhieuNhapSach(maPhieuNhap *entity.ID) (*entity.PhieuNhap, error)
	CreatePhieuNhapSach(ngayLap *time.Time) (*entity.PhieuNhap, error)
	UpdatePhieuNhapSach(maPhieuNhap *entity.ID, ngayLap *time.Time) (*entity.PhieuNhap, error)
	RemovePhieuNhapSach(maPhieuNhap *entity.ID) error
	AddChiTietPhieuNhapSach(maPhieuNhap *entity.ID, maDauSach *entity.ID, nhaXuatBan string, triGia uint, namXuatBan uint, tinhTrang bool, donGia uint, ghiChu string) (*entity.CtPhieuNhap, error)
	RemoveChiTietPhieuNhapSach(maChiTietPhieuNhap *entity.ID) error
}
