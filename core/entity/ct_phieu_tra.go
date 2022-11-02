package entity

type CtPhieuTra struct {
	MaChiTietPhieuTra *ID
	PhieuTra          *PhieuTra
	CuonSach          *CuonSach
	PhieuMuon         *PhieuMuon
	SoNgayMuon        uint
	TienPhat          uint
}
