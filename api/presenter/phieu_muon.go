package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type PhieuMuonPresenter struct {
	MaPhieuMuon string           `json:"maPhieuMuon"`
	DocGia      *DocGiaPresenter `json:"docGia"`
	NgayMuon    *time.Time       `json:"ngayMuon"`
	Sach        *SachPresenter   `json:"sach"`
}

func NewPhieuMuonPresenter(phieuMuon *entity.PhieuMuon) PhieuMuonPresenter {
	return PhieuMuonPresenter{
		MaPhieuMuon: phieuMuon.MaPhieuMuon.String(),
		DocGia:      NewDocGiaPresenter(phieuMuon.DocGia),
		NgayMuon:    phieuMuon.NgayMuon,
		Sach:        NewSachPresenter(phieuMuon.Sach),
	}
}

func NewDanhSachPhieuMuonPresenter(danhSachPhieuMuon []*entity.PhieuMuon) []PhieuMuonPresenter {
	var danhSachPhieuMuonPresenter = make([]PhieuMuonPresenter, len(danhSachPhieuMuon))
	for i, _ := range danhSachPhieuMuon {
		danhSachPhieuMuonPresenter[i] = NewPhieuMuonPresenter(danhSachPhieuMuon[i])
	}
	return danhSachPhieuMuonPresenter
}
