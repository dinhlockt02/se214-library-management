package entity

import "time"

type BcTraTre struct {
	NgayThangNam time.Time
	MaCuonSach   ID
	MaPhieuMuon  ID
	SoNgayTraTre uint
}
