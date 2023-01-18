package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type PhieuTraPresenter struct {
	TienPhat  uint               `json:"tienPhat"`
	NgayTra   *time.Time         `json:"ngayTra"`
	GhiChu    string             `json:"ghiChu"`
	PhieuMuon PhieuMuonPresenter `json:"phieuMuon"`
}

func NewPhieuTraPresenter(phieuTra *entity.PhieuTra) PhieuTraPresenter {
	return PhieuTraPresenter{
		TienPhat:  phieuTra.TienPhat,
		NgayTra:   phieuTra.NgayTra,
		GhiChu:    phieuTra.GhiChu,
		PhieuMuon: NewPhieuMuonPresenter(phieuTra.PhieuMuon),
	}
}

func NewDanhSachPhieuTraPresenter(danhSachPhieuTra []*entity.PhieuTra) []PhieuTraPresenter {
	var danhSachPhieuTraPresenter = make([]PhieuTraPresenter, len(danhSachPhieuTra))
	for i, _ := range danhSachPhieuTra {
		danhSachPhieuTraPresenter[i] = NewPhieuTraPresenter(danhSachPhieuTra[i])
	}
	return danhSachPhieuTraPresenter
}
