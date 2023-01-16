package entity

import "time"

type PhieuNhap struct {
	MaPhieuNhap *ID
	NgayLap     *time.Time
	TongTien    uint
	CtPhieuNhap []*CtPhieuNhap
}

func NewPhieuNhap(ngayNhap *time.Time, tongTien uint, ctPhieuNhap []*CtPhieuNhap) *PhieuNhap {
	newId := NewID()

	return &PhieuNhap{
		MaPhieuNhap: &newId,
		NgayLap:     ngayNhap,
		TongTien:    tongTien,
		CtPhieuNhap: ctPhieuNhap,
	}
}

func (phieuNhap *PhieuNhap) IsValid() bool {
	return true
}
