package dto

type DocGiaDto struct {
	MaDocGia     string `json:"maDocGia"`
	HoTen        string `json:"hoTen" binding:"required"`
	MaLoaiDocGia string `json:"maLoaiDocGia" binding:"required"`
	NgaySinh     string `json:"ngaySinh" binding:"required"`
	DiaChi       string `json:"diaChi" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NgayLapThe   string `json:"ngayLapThe" binding:"required"`
}
