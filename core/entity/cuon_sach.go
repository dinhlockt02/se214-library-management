package entity

type CuonSach struct {
	MaCuonSach *ID
	Sach       *Sach
	TinhTrang  bool
	CTPN       *CtPhieuNhap
}

func NewCuonSach(sach *Sach, ctpn *CtPhieuNhap, tinhTrang bool) *CuonSach {
	newId := NewID()

	return &CuonSach{
		MaCuonSach: &newId,
		Sach:       sach,
		CTPN:       ctpn,
		TinhTrang:  tinhTrang,
	}
}
