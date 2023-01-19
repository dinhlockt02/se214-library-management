package entity

import "time"

type PhieuThuTien struct {
	MaPhieuThu *ID
	DocGia     *DocGia
	NgayThu    *time.Time
	TongNo     int
	SoTienThu  int
	ConLai     int
}

func NewPhieuThuTien(docGia *DocGia, ngayThu *time.Time, tongNo int, soTienThu int, conLai int) *PhieuThuTien {
	id := NewID()
	return &PhieuThuTien{
		MaPhieuThu: &id,
		DocGia:     docGia,
		NgayThu:    ngayThu,
		TongNo:     tongNo,
		SoTienThu:  soTienThu,
		ConLai:     conLai,
	}
}
