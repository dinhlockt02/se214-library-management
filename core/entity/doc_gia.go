package entity

import (
	"fmt"
	"time"

	businessError "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/utils"
)

type DocGia struct {
	MaDocGia   *ID
	HoTen      string
	LoaiDocGia *LoaiDocGia
	NgaySinh   *time.Time
	DiaChi     string
	Email      string
	NgayLapThe *time.Time
	NgayHetHan *time.Time
	TongNo     uint
}

func NewDocGia(hoTen string, loaiDocGia *LoaiDocGia, ngaySinh *time.Time, diaChi string, email string, ngayLapThe *time.Time, ngayHetHan *time.Time) *DocGia {
	newId := NewID()
	return &DocGia{
		MaDocGia:   &newId,
		HoTen:      hoTen,
		LoaiDocGia: loaiDocGia,
		NgaySinh:   ngaySinh,
		DiaChi:     diaChi,
		Email:      email,
		NgayLapThe: ngayLapThe,
		NgayHetHan: ngayHetHan,
		TongNo:     0,
	}
}

func (docGia *DocGia) IsValid(tuoiToiDa uint, tuoiToiThieu uint, thoiHanTheMonth uint) (bool, error) {
	if docGia.MaDocGia == nil {
		return false, businessError.NewBusinessError("ma doc gia is nil")
	}
	if docGia.LoaiDocGia == nil {
		return false, businessError.NewBusinessError("loai doc gia is nil")
	}

	if docGia.NgaySinh == nil {
		return false, businessError.NewBusinessError("ngay sinh is nil")
	}

	tuoi := time.Now().Year() - docGia.NgaySinh.Year()
	if tuoi < 0 {
		return false, businessError.NewBusinessError(fmt.Sprintf("tuoi(%v) must be positive", tuoi))
	}
	if uint(tuoi) > tuoiToiDa || uint(tuoi) < tuoiToiThieu {
		return false, businessError.NewBusinessError(fmt.Sprintf("tuoi(%v) is not between tuoi toi da(%v) and tuoi toi thieu(%v)", tuoi, tuoiToiThieu, tuoiToiDa))
	}

	if diffMonths := utils.DiffMonths(*docGia.NgayHetHan, *docGia.NgayLapThe); diffMonths != int(thoiHanTheMonth) {
		return false, businessError.NewBusinessError("thoi han the is not match")
	}
	return true, nil
}
