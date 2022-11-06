package entity

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"fmt"
	"time"

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
		return false, coreerror.NewBadRequestError("ma doc gia is nil", nil)
	}
	if docGia.LoaiDocGia == nil {
		return false, coreerror.NewBadRequestError("loai doc gia is nil", nil)
	}

	if docGia.NgaySinh == nil {
		return false, coreerror.NewBadRequestError("ngay sinh is nil", nil)
	}

	tuoi := time.Now().Year() - docGia.NgaySinh.Year()
	if tuoi < 0 {
		return false, coreerror.NewBadRequestError(fmt.Sprintf("tuoi(%v) must be positive", tuoi), nil)
	}
	if uint(tuoi) > tuoiToiDa || uint(tuoi) < tuoiToiThieu {
		return false, coreerror.NewBadRequestError(fmt.Sprintf("tuoi(%v) is not between tuoi toi da(%v) and tuoi toi thieu(%v)", tuoi, tuoiToiThieu, tuoiToiDa), nil)
	}

	if diffMonths := utils.DiffMonths(*docGia.NgayHetHan, *docGia.NgayLapThe); diffMonths != int(thoiHanTheMonth) {
		return false, coreerror.NewBadRequestError("thoi han the is not match", nil)
	}
	return true, nil
}
