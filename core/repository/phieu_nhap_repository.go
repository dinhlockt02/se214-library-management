package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type PhieuNhapRepository interface {
	GetDanhSachPhieuNhap() ([]*entity.PhieuNhap, error)
	GetPhieuNhap(maPhieuNhap *entity.ID) (*entity.PhieuNhap, error)
	CreatePhieuNhap(phieuNhap *entity.PhieuNhap) (*entity.PhieuNhap, error)
	UpdatePhieuNhap(phieuNhap *entity.PhieuNhap) (*entity.PhieuNhap, error)
	RemovePhieuNhap(maPhieuNhap *entity.ID) error
}
