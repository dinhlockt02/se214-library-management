package entity

import "time"

type BcTraTre struct {
	MaBaoCaoTraTre *ID
	NgayThangNam   *time.Time
	CuonSach       *CuonSach
	PhieuMuon      *PhieuMuon
	SoNgayTraTre   uint
}
