package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type TacGiaPresenter struct {
	MaTacGia  string `json:"maTacGia" binding:"required"`
	TenTacGia string `json:"tenTacGia" binding:"required"`
}

func NewTacGiaPresenter(tacGia *entity.TacGia) *TacGiaPresenter {
	return &TacGiaPresenter{
		MaTacGia:  tacGia.MaTacGia.String(),
		TenTacGia: tacGia.TenTacGia,
	}
}

func NewDanhSachTacGiaPresenter(danhSachTacGia []*entity.TacGia) []*TacGiaPresenter {
	danhSachTacGiaPresenter := make([]*TacGiaPresenter, len(danhSachTacGia))
	for index, tacGia := range danhSachTacGia {
		danhSachTacGiaPresenter[index] = NewTacGiaPresenter(tacGia)
	}
	return danhSachTacGiaPresenter
}
