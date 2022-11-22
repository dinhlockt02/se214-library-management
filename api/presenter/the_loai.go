package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type TheLoaiPresenter struct {
	MaTheLoai  string `json:"maTheLoai" binding:"required"`
	TenTheLoai string `json:"tenTheLoai" binding:"required"`
}

func NewTheLoaiPresenter(theLoai *entity.TheLoai) *TheLoaiPresenter {
	return &TheLoaiPresenter{
		MaTheLoai:  theLoai.MaTheLoai.String(),
		TenTheLoai: theLoai.TenTheLoai,
	}
}

func NewDanhSachTheLoaiPresenter(danhSachTheLoai []*entity.TheLoai) []*TheLoaiPresenter {
	danhSachTheLoaiPresenter := make([]*TheLoaiPresenter, len(danhSachTheLoai))
	for index, theLoai := range danhSachTheLoai {
		danhSachTheLoaiPresenter[index] = NewTheLoaiPresenter(theLoai)
	}
	return danhSachTheLoaiPresenter
}
