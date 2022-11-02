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
	AddChiTietPhieuNhapSach(maPhieuNhap *entity.ID, maSach *entity.ID, soLuong uint, donGia uint) (*entity.PhieuNhap, error)
	RemoveChiTietPhieuNhapSach(maChiTietPhieuNhap *entity.ID) (*entity.PhieuNhap, error)
	GetChiTietPhieuNhap(maChiTietPhieuNhap *entity.ID) (*entity.CtPhieuNhap, error)
}
