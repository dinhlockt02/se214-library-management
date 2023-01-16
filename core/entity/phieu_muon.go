package entity

import "time"

type PhieuMuon struct {
	MaPhieuMuon *ID
	*DocGia
	NgayMuon *time.Time
	*Sach
	*PhieuTra
}
