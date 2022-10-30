package entity

import "time"

type DocGia struct {
	MaDocGia     ID
	HoTen        string
	MaLoaiDocGia ID
	NgaySinh     time.Time
	DiaChi       string
	Email        string
	NgayLapThe   time.Time
	NgayHetHan   time.Time
	TongNo       uint
}

func NewDocGia(hoTen string, loaiDocGia ID, ngaySinh time.Time, diaChi string, email string, ngayLapThe time.Time, ngayHetHan time.Time) *DocGia {
	return &DocGia{
		MaDocGia:     NewID(),
		HoTen:        hoTen,
		MaLoaiDocGia: loaiDocGia,
		NgaySinh:     ngaySinh,
		DiaChi:       diaChi,
		Email:        email,
		NgayLapThe:   ngayLapThe,
		NgayHetHan:   ngayHetHan,
		TongNo:       0,
	}
}
