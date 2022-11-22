package entity

type Sach struct {
	MaSach     *ID
	DauSach    *DauSach
	NhaXuatBan string
	TriGia     uint
	NamXuatBan uint
	TinhTrang  bool
}

func NewSach(dauSach *DauSach, nhaXuatBan string, triGia uint, namXuatBan uint, tinhTrang bool) *Sach {

	newId := NewID()

	return &Sach{
		MaSach:     &newId,
		DauSach:    dauSach,
		NhaXuatBan: nhaXuatBan,
		TriGia:     triGia,
		NamXuatBan: namXuatBan,
		TinhTrang:  tinhTrang,
	}
}

func (sach *Sach) IsValid() bool {
	return true
}
