package entity

type LoaiDocGia struct {
	MaLoaiDocGia  ID
	TenLoaiDocGia string
}

func NewLoaiDocGia(tenLoaiDocGia string) *LoaiDocGia {
	return &LoaiDocGia{
		MaLoaiDocGia:  NewID(),
		TenLoaiDocGia: tenLoaiDocGia,
	}
}
