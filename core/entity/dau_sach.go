package entity

type DauSach struct {
	MaDauSach  *ID
	TenDauSach string
	MaTheLoai  *ID
	MaTacGia   []*ID
}

func NewDauSach(maTheLoai *ID, tenDauSach string, maTacGia []*ID) *DauSach {
	newID := NewID()
	return &DauSach{
		MaDauSach:  &newID,
		TenDauSach: tenDauSach,
		MaTheLoai:  maTheLoai,
		MaTacGia:   maTacGia,
	}
}

func (dauSach *DauSach) IsValid() bool {
	return true
}
