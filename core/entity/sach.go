package entity

type Sach struct {
	MaSach     *ID
	MaDauSach  *ID
	NhaXuatBan string
	SoLuong    uint
	TriGia     uint
	NamXuatBan uint
}

func NewSach(maDauSach *ID, nhaXuatBan string, soLuong uint, triGia uint, namXuatBan uint) *Sach {

	newId := NewID()

	return &Sach{
		MaSach:     &newId,
		MaDauSach:  maDauSach,
		NhaXuatBan: nhaXuatBan,
		SoLuong:    soLuong,
		TriGia:     triGia,
		NamXuatBan: namXuatBan,
	}
}

func (sach *Sach) IsValid() bool {
	return true
}
