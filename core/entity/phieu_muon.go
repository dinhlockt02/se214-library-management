package entity

import "time"

type PhieuMuon struct {
	MaPhieuMuon *ID
	*DocGia
	NgayMuon *time.Time
	*Sach
	*PhieuTra
}

func NewPhieuMuon(docGia *DocGia, ngayMuon *time.Time, sach *Sach, phieuTra *PhieuTra) *PhieuMuon {
	id := NewID()
	return &PhieuMuon{
		MaPhieuMuon: &id,
		DocGia:      docGia,
		NgayMuon:    ngayMuon,
		Sach:        sach,
		PhieuTra:    phieuTra,
	}
}
