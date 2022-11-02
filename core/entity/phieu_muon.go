package entity

import "time"

type PhieuMuon struct {
	MaPhieuMuon *ID
	DocGia      *DocGia
	NgayMuon    *time.Time
	HanTra      *time.Time
}
