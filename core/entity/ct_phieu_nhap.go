package entity

type CtPhieuNhap struct {
	MaCTPN    *ID
	PhieuNhap *PhieuNhap
	Sach      *Sach
	SoLuong   uint
	DonGia    uint
	ThanhTien uint
}

func NewCtPhieuNhap(phieuNhap *PhieuNhap, sach *Sach, soLuong uint, donGia uint) *CtPhieuNhap {
	newId := NewID()

	return &CtPhieuNhap{
		MaCTPN:    &newId,
		PhieuNhap: phieuNhap,
		Sach:      sach,
		SoLuong:   soLuong,
		DonGia:    donGia,
		ThanhTien: donGia * soLuong,
	}
}

func (ctPhieuNhap *CtPhieuNhap) IsValid() bool {
	return true
}
