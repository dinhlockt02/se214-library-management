package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/utils"
	"time"
)

type DocGiaPresenter struct {
	MaDocGia   string               `json:"maDocGia" binding:"required"`
	HoTen      string               `json:"hoTen" binding:"required"`
	LoaiDocGia *LoaiDocGiaPresenter `json:"loaiDocGia" binding:"required"`
	NgaySinh   *utils.JSONTime      `json:"ngaySinh" binding:"required"`
	DiaChi     string               `json:"diaChi" binding:"required"`
	Email      string               `json:"email" binding:"required"`
	NgayLapThe *time.Time           `json:"ngayLapThe" binding:"required"`
	NgayHetHan *time.Time           `json:"ngayHetHan" binding:"required"`
	TongNo     uint                 `json:"tongNo" binding:"required"`
}

func NewDanhSachDocGiaPresenter(danhSachDocGia []*entity.DocGia) []*DocGiaPresenter {
	danhSachDocGiaPresenter := make([]*DocGiaPresenter, len(danhSachDocGia))
	for index, docGia := range danhSachDocGia {
		danhSachDocGiaPresenter[index] = NewDocGiaPresenter(docGia)
	}
	return danhSachDocGiaPresenter

}

func NewDocGiaPresenter(docGia *entity.DocGia) *DocGiaPresenter {
	return &DocGiaPresenter{
		MaDocGia:   docGia.MaDocGia,
		HoTen:      docGia.HoTen,
		LoaiDocGia: NewLoaiDocGiaPresenter(docGia.LoaiDocGia),
		NgaySinh:   (*utils.JSONTime)(docGia.NgaySinh),
		DiaChi:     docGia.DiaChi,
		Email:      docGia.Email,
		NgayLapThe: docGia.NgayLapThe,
		NgayHetHan: docGia.NgayHetHan,
		TongNo:     docGia.TongNo,
	}
}
