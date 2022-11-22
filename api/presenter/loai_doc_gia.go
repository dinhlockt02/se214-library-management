package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type LoaiDocGiaPresenter struct {
	MaLoaiDocGia  string `json:"maLoaiDocGia" binding:"required"`
	TenLoaiDocGia string `json:"tenLoaiDocGia" binding:"required"`
}

func NewLoaiDocGiaPresenter(loaiDocGia *entity.LoaiDocGia) *LoaiDocGiaPresenter {
	return &LoaiDocGiaPresenter{
		MaLoaiDocGia:  loaiDocGia.MaLoaiDocGia.String(),
		TenLoaiDocGia: loaiDocGia.TenLoaiDocGia,
	}
}

func NewDanhSachLoaiDocGiaPresenter(danhSachLoaiDocGia []*entity.LoaiDocGia) []*LoaiDocGiaPresenter {
	danhSachLoaiDocGiaPresenter := make([]*LoaiDocGiaPresenter, len(danhSachLoaiDocGia))
	for index, loaiDocGia := range danhSachLoaiDocGia {
		danhSachLoaiDocGiaPresenter[index] = NewLoaiDocGiaPresenter(loaiDocGia)
	}
	return danhSachLoaiDocGiaPresenter
}
