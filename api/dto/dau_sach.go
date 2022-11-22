package dto

type DauSachDto struct {
	TenDauSach string   `json:"tenDauSach" binding:"required"`
	TheLoai    []string `json:"theLoai" binding:"required"`
	TacGia     []string `json:"tacGia" binding:"required"`
}
