package entity

import "time"

type PhieuThuTien struct {
	MaPhieuThu ID
	MaDocGia   ID
	NgayThu    time.Time
	SoTienThu  uint
	ConLai     uint
}
