package entity

import "time"

type PhieuNhap struct {
	MaPhieuNhap *ID
	NgayLap     *time.Time
	TongTien    uint
}

func NewPhieuNhap(ngayNhap *time.Time) *PhieuNhap {
	newId := NewID()

	return &PhieuNhap{
		MaPhieuNhap: &newId,
		NgayLap:     ngayNhap,
		TongTien:    0,
	}
}

func (phieuNhap *PhieuNhap) IsValid() bool {
	return true
}
