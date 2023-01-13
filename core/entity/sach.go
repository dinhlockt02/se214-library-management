package entity

type Sach struct {
	MaSach     *ID
	DauSach    *DauSach
	NhaXuatBan string
	TriGia     uint
	NamXuatBan uint
	TinhTrang  bool
	GhiChu     string
}

func NewSach(dauSach *DauSach, nhaXuatBan string, triGia uint, namXuatBan uint, tinhTrang bool, ghiChu string) *Sach {

	newId := NewID()

	return &Sach{
		MaSach:     &newId,
		DauSach:    dauSach,
		NhaXuatBan: nhaXuatBan,
		TriGia:     triGia,
		NamXuatBan: namXuatBan,
		TinhTrang:  tinhTrang,
		GhiChu:     ghiChu,
	}
}

func (sach *Sach) IsValid() bool {
	return true
}
