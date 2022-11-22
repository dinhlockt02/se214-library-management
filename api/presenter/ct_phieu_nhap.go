package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type CtPhieuNhapPresenter struct {
	Sach   *SachPresenter `json:"sach" binding:"required"`
	DonGia uint           `json:"donGia" binding:"required"`
}

func NewDanhSachCtPhieuNhapPresenter(danhSachCtPhieuNhap []*entity.CtPhieuNhap) []*CtPhieuNhapPresenter {
	danhSachCtPhieuNhapPresenter := make([]*CtPhieuNhapPresenter, len(danhSachCtPhieuNhap))
	for i, ctPhieuNhap := range danhSachCtPhieuNhap {
		danhSachCtPhieuNhapPresenter[i] = NewCtPhieuNhapPresenter(ctPhieuNhap)
	}
	return danhSachCtPhieuNhapPresenter
}

func NewCtPhieuNhapPresenter(ctPhieuNhap *entity.CtPhieuNhap) *CtPhieuNhapPresenter {
	return &CtPhieuNhapPresenter{
		Sach:   NewSachPresenter(ctPhieuNhap.Sach),
		DonGia: ctPhieuNhap.DonGia,
	}
}
