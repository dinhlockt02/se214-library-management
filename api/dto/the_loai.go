package dto

type TheLoaiDto struct {
	TenTheLoai string `json:"tenTheLoai" binding:"required"`
}
