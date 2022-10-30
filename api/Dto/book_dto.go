package dto

type CreateBookDto struct {
	NhaXuatBan string `json:"nhaxuatban"`
	SoLuong    uint   `json:"soluong"`
	TriGia     uint   `json:"trigia"`
	NamXuatBan uint   `json:"namxuatban"`
}
