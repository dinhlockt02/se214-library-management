package entity

type LoaiDocGia struct {
	MaLoaiDocGia  *ID
	TenLoaiDocGia string
}

func NewLoaiDocGia(tenLoaiDocGia string) *LoaiDocGia {
	newId := NewID()
	return &LoaiDocGia{
		MaLoaiDocGia:  &newId,
		TenLoaiDocGia: tenLoaiDocGia,
	}
}
