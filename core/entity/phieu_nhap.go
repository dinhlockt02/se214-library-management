package entity

import "time"

type PhieuNhap struct {
	MaPhieuNhap *ID
	NgayLap     *time.Time
	TongTien    uint
	CtPhieuNhap []*CtPhieuNhap
}

func NewPhieuNhap(ngayNhap *time.Time) *PhieuNhap {
	newId := NewID()

	return &PhieuNhap{
		MaPhieuNhap: &newId,
		NgayLap:     ngayNhap,
		TongTien:    0,
		CtPhieuNhap: nil,
	}
}

func (phieuNhap *PhieuNhap) IsValid() bool {
	return true
}
