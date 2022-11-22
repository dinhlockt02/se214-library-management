package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/utils"
)

type PhieuNhapPresenter struct {
	MaPhieuNhap string                  `json:"maPhieuNhap" binding:"required"`
	NgayLap     *utils.JSONTime         `json:"ngayLap" binding:"required"`
	TongTien    uint                    `json:"tongTien" binding:"required"`
	CtPhieuNhap []*CtPhieuNhapPresenter `json:"ctPhieuNhap" binding:"required"`
}

func NewPhieuNhapPresenter(phieuNhap *entity.PhieuNhap) *PhieuNhapPresenter {
	return &PhieuNhapPresenter{
		MaPhieuNhap: phieuNhap.MaPhieuNhap.String(),
		NgayLap:     (*utils.JSONTime)(phieuNhap.NgayLap),
		TongTien:    phieuNhap.TongTien,
		CtPhieuNhap: NewDanhSachCtPhieuNhapPresenter(phieuNhap.CtPhieuNhap),
	}
}

func NewDanhSachPhieuNhapPresenter(danhSachPhieuNhap []*entity.PhieuNhap) []*PhieuNhapPresenter {
	danhSachPhieuNhapPresenter := make([]*PhieuNhapPresenter, len(danhSachPhieuNhap))
	for i, phieuNhap := range danhSachPhieuNhap {
		danhSachPhieuNhapPresenter[i] = NewPhieuNhapPresenter(phieuNhap)
	}
	return danhSachPhieuNhapPresenter
}
