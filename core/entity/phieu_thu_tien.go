package entity

import "time"

type PhieuThuTien struct {
	MaPhieuThu *ID
	DocGia     *DocGia
	NgayThu    *time.Time
	SoTienThu  uint
	ConLai     uint
}
