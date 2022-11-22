package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
)

type DauSachPresenter struct {
	MaDauSach  string              `json:"maDauSach" binding:"required"`
	TenDauSach string              `json:"tenDauSach" binding:"required"`
	TheLoai    []*TheLoaiPresenter `json:"theLoai" binding:"required"`
	TacGia     []*TacGiaPresenter  `json:"tacGia" binding:"required"`
}

func NewDauSachPresenter(dauSach *entity.DauSach) *DauSachPresenter {
	return &DauSachPresenter{
		MaDauSach:  dauSach.MaDauSach.String(),
		TenDauSach: dauSach.TenDauSach,
		TheLoai:    NewDanhSachTheLoaiPresenter(dauSach.TheLoai),
		TacGia:     NewDanhSachTacGiaPresenter(dauSach.TacGia),
	}
}

func NewDanhSachDauSachPresenter(danhSachDauSach []*entity.DauSach) []*DauSachPresenter {
	danhSachDauSachPresenter := make([]*DauSachPresenter, 0, len(danhSachDauSach))
	for _, dauSach := range danhSachDauSach {
		danhSachDauSachPresenter = append(danhSachDauSachPresenter, NewDauSachPresenter(dauSach))
	}
	return danhSachDauSachPresenter
}
