package entity

import "time"

type PhieuNhap struct {
	MaPhieuNhap ID
	NgayLap     time.Time
	TongTien    uint
}
