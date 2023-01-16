package dto

type UpdateSachDto struct {
	NhaXuatBan string `json:"nhaXuatBan" binding:"required"`
	TriGia     uint   `json:"triGia" binding:"required"`
	NamXuatBan uint   `json:"namXuatBan" binding:"required"`
	TinhTrang  bool   `json:"tinhTrang"`
	GhiChu     string `json:"ghiChu" binding:"required"`
}
