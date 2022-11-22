package dto

type TacGiaDto struct {
	TenTacGia string `json:"tenTacGia" binding:"required"`
}
