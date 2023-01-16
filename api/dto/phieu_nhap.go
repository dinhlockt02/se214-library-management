package dto

type UpdatePhieuNhapDto struct {
	NgayLap string `json:"ngayLap" binding:"required"`
}

type PhieuNhapDto struct {
	NgayLap     string           `json:"ngayLap" binding:"required"`
	CtPhieuNhap []CtPhieuNhapDto `json:"ctPhieuNhap" binding:"required"`
}

type CtPhieuNhapDto struct {
	MaDauSach  string `json:"maDauSach" binding:"required"`
	NhaXuatBan string `json:"nhaXuatBan" binding:"required"`
	TriGia     uint   `json:"triGia" binding:"required"`
	NamXuatBan uint   `json:"namXuatBan"  binding:"required"`
	TinhTrang  bool   `json:"tinhTrang"  binding:"required"`
	DonGia     uint   `json:"donGia"  binding:"required"`
	GhiChu     string `json:"ghiChu" binding:"required"`
}
