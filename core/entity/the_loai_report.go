package entity

type ReportTheLoai struct {
	MaTheLoai  *ID
	TenTheLoai string
	SoLuotMuon uint
}

type MonthReportTheLoai struct {
	ChiTiet []ReportTheLoai
	Thang   uint
	Nam     uint
}

type YearReportTheLoai struct {
	ChiTiet []ReportTheLoai
	Nam     uint
}
