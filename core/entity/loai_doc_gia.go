package entity

type LoaiDocGia struct {
	MaLoaiDocGia        *ID
	TenLoaiDocGia       string
	SoSachToiDaDuocMuon int
}

func NewLoaiDocGia(tenLoaiDocGia string, soSachToiDaDuocMuon int) *LoaiDocGia {
	newId := NewID()
	return &LoaiDocGia{
		MaLoaiDocGia:  &newId,
		TenLoaiDocGia: tenLoaiDocGia,
	}
}
