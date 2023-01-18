package entity

type LoaiDocGia struct {
	MaLoaiDocGia        *ID
	TenLoaiDocGia       string
	SoSachToiDaDuocMuon int
	TienPhatTheoNgay    uint
	ThoiGianMuonToiDa   uint
}

func NewLoaiDocGia(tenLoaiDocGia string, soSachToiDaDuocMuon int, tienPhatTheoNgay uint, thoiGianMuonToiDa uint) *LoaiDocGia {
	newId := NewID()
	return &LoaiDocGia{
		MaLoaiDocGia:        &newId,
		TenLoaiDocGia:       tenLoaiDocGia,
		SoSachToiDaDuocMuon: soSachToiDaDuocMuon,
		TienPhatTheoNgay:    tienPhatTheoNgay,
		ThoiGianMuonToiDa:   thoiGianMuonToiDa,
	}
}
