package entity

type CtPhieuNhap struct {
	*Sach
	DonGia uint
}

func NewCtPhieuNhap(sach *Sach, donGia uint) *CtPhieuNhap {

	return &CtPhieuNhap{
		Sach:   sach,
		DonGia: donGia,
	}
}

func (ctPhieuNhap *CtPhieuNhap) IsValid() bool {
	return true && ctPhieuNhap.Sach.IsValid()
}
