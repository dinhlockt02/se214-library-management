package entity

type Sach struct {
	MaSach     *ID
	DauSach    *DauSach
	NhaXuatBan string
	SoLuong    uint
	TriGia     uint
	NamXuatBan uint
}

func NewSach(dauSach *DauSach, nhaXuatBan string, soLuong uint, triGia uint, namXuatBan uint) *Sach {

	newId := NewID()

	return &Sach{
		MaSach:     &newId,
		DauSach:    dauSach,
		NhaXuatBan: nhaXuatBan,
		SoLuong:    soLuong,
		TriGia:     triGia,
		NamXuatBan: namXuatBan,
	}
}

func (sach *Sach) IsValid() bool {
	return true
}
