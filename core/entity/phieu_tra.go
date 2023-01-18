package entity

import "time"

type PhieuTra struct {
	TienPhat uint
	NgayTra  *time.Time
	GhiChu   string
	*PhieuMuon
}
