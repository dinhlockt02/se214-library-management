package entity

import "time"

type PhieuMuon struct {
	MaPhieuMuon ID
	MaDocGia    ID
	NgayMuon    time.Time
	HanTra      time.Time
}
