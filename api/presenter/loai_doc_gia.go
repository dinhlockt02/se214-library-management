package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type LoaiDocGiaPresenter struct {
	MaLoaiDocGia        string `json:"maLoaiDocGia"`
	TenLoaiDocGia       string `json:"tenLoaiDocGia"`
	SoSachToiDaDuocMuon int    `json:"soSachToiDaDuocMuon"`
	TienPhatTheoNgay    uint   `json:"tienPhatTheoNgay"`
	ThoiGianMuonToiDa   uint   `json:"thoiGianMuonToiDa"`
}

func NewLoaiDocGiaPresenter(loaiDocGia *entity.LoaiDocGia) *LoaiDocGiaPresenter {
	return &LoaiDocGiaPresenter{
		MaLoaiDocGia:        loaiDocGia.MaLoaiDocGia.String(),
		TenLoaiDocGia:       loaiDocGia.TenLoaiDocGia,
		SoSachToiDaDuocMuon: loaiDocGia.SoSachToiDaDuocMuon,
		TienPhatTheoNgay:    loaiDocGia.TienPhatTheoNgay,
		ThoiGianMuonToiDa:   loaiDocGia.ThoiGianMuonToiDa,
	}
}

func NewDanhSachLoaiDocGiaPresenter(danhSachLoaiDocGia []*entity.LoaiDocGia) []*LoaiDocGiaPresenter {
	danhSachLoaiDocGiaPresenter := make([]*LoaiDocGiaPresenter, len(danhSachLoaiDocGia))
	for index, loaiDocGia := range danhSachLoaiDocGia {
		danhSachLoaiDocGiaPresenter[index] = NewLoaiDocGiaPresenter(loaiDocGia)
	}
	return danhSachLoaiDocGiaPresenter
}
