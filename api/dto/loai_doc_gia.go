package dto

type PostLoaiDocGiaDto struct {
	TenLoaiDocGia string `json:"tenLoaiDocGia" binding:"required"`
}

type PutLoaiDocGiaDto struct {
	TenLoaiDocGia string `json:"tenLoaiDocGia" binding:"required"`
}
