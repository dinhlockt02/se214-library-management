package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type SachPresenter struct {
	MaSach     string `json:"maSach" binding:"required"`
	DauSach    *DauSachPresenter
	NhaXuatBan string `json:"nhaXuatBan" binding:"required"`
	TriGia     uint   `json:"triGia" binding:"required"`
	NamXuatBan uint   `json:"namXuatBan" binding:"required"`
	TinhTrang  bool   `json:"tinhTrang" binding:"required"`
}

func NewSachPresenter(sach *entity.Sach) *SachPresenter {
	return &SachPresenter{
		MaSach:     sach.MaSach.String(),
		DauSach:    NewDauSachPresenter(sach.DauSach),
		NhaXuatBan: sach.NhaXuatBan,
		TriGia:     sach.TriGia,
		NamXuatBan: sach.NamXuatBan,
		TinhTrang:  sach.TinhTrang,
	}
}

func NewDanhSachSachPresenter(danhSachSach []*entity.Sach) []*SachPresenter {
	danhSachSachPresenter := make([]*SachPresenter, len(danhSachSach))
	for i, _ := range danhSachSach {
		danhSachSachPresenter[i] = NewSachPresenter(danhSachSach[i])
	}
	return danhSachSachPresenter
}
