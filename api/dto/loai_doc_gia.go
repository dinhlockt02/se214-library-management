package dto

type PostLoaiDocGiaDto struct {
	TenLoaiDocGia       string `json:"tenLoaiDocGia" binding:"required"`
	SoSachToiDaDuocMuon int    `json:"soSachToiDaDuocMuon" `
	TienPhatTheoNgay    uint   `json:"tienPhatTheoNgay" `
	ThoiGianMuonToiDa   uint   `json:"thoiGianMuonToiDa" `
}

type PutLoaiDocGiaDto struct {
	TenLoaiDocGia       string `json:"tenLoaiDocGia" binding:"required"`
	SoSachToiDaDuocMuon int    `json:"soSachToiDaDuocMuon" `
	TienPhatTheoNgay    uint   `json:"tienPhatTheoNgay" `
	ThoiGianMuonToiDa   uint   `json:"thoiGianMuonToiDa"`
}
