package presenter

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"time"
)

type PhieuThuTienPresenter struct {
	MaPhieuThu string           `json:"maPhieuThu"`
	DocGia     *DocGiaPresenter `json:"docGia"`
	NgayThu    *time.Time       `json:"ngayThu"`
	TongNo     int              `json:"tongNo"`
	SoTienThu  int              `json:"soTienThu"`
	ConLai     int              `json:"conLai"`
}

func NewPhieuThuTienPresenter(phieuThu *entity.PhieuThuTien) PhieuThuTienPresenter {
	return PhieuThuTienPresenter{
		MaPhieuThu: phieuThu.MaPhieuThu.String(),
		DocGia:     NewDocGiaPresenter(phieuThu.DocGia),
		NgayThu:    phieuThu.NgayThu,
		TongNo:     phieuThu.TongNo,
		SoTienThu:  phieuThu.SoTienThu,
		ConLai:     phieuThu.ConLai,
	}
}

func NewDanhSachPhieuThuTienPresenter(danhSachPhieuThuTien []*entity.PhieuThuTien) []PhieuThuTienPresenter {
	var danhSachPhieuThuTienPresenter = make([]PhieuThuTienPresenter, len(danhSachPhieuThuTien))
	for i, _ := range danhSachPhieuThuTien {
		danhSachPhieuThuTienPresenter[i] = NewPhieuThuTienPresenter(danhSachPhieuThuTien[i])
	}
	return danhSachPhieuThuTienPresenter
}
