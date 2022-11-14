package entity

type DauSach struct {
	MaDauSach  *ID
	TenDauSach string
	TheLoai    []*TheLoai
	TacGia     []*TacGia
}

func NewDauSach(theLoai []*TheLoai, tenDauSach string, tacGia []*TacGia) *DauSach {
	newID := NewID()
	return &DauSach{
		MaDauSach:  &newID,
		TenDauSach: tenDauSach,
		TheLoai:    theLoai,
		TacGia:     tacGia,
	}
}

func (dauSach *DauSach) IsValid() bool {
	return true
}
