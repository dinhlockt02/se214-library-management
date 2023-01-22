package presenter

import "daijoubuteam.xyz/se214-library-management/core/entity"

type ReportTheLoaiPresenter struct {
	MaTheLoai  string `json:"maTheLoai"`
	TenTheLoai string `json:"tenTheLoai"`
	SoLuotMuon uint   `json:"soLuotMuon"`
}

type MonthReportTheLoaiPresenter struct {
	ChiTiet []ReportTheLoaiPresenter `json:"chiTiet"`
	Thang   uint                     `json:"thang"`
	Nam     uint                     `json:"nam"`
}

type YearReportTheLoaiPresenter struct {
	ChiTiet []ReportTheLoaiPresenter `json:"chiTiet"`
	Nam     uint                     `json:"nam"`
}

func NewMonthReportTheLoaiPresenter(tl *entity.MonthReportTheLoai) MonthReportTheLoaiPresenter {
	var chiTiet = make([]ReportTheLoaiPresenter, len(tl.ChiTiet))
	for i := 0; i < len(chiTiet); i++ {
		chiTiet[i] = ReportTheLoaiPresenter{
			MaTheLoai:  tl.ChiTiet[i].MaTheLoai.String(),
			TenTheLoai: tl.ChiTiet[i].TenTheLoai,
			SoLuotMuon: tl.ChiTiet[i].SoLuotMuon,
		}
	}
	return MonthReportTheLoaiPresenter{
		Thang:   tl.Thang,
		Nam:     tl.Nam,
		ChiTiet: chiTiet,
	}
}

func NewYearReportTheLoaiPresenter(tl *entity.YearReportTheLoai) YearReportTheLoaiPresenter {
	var chiTiet = make([]ReportTheLoaiPresenter, len(tl.ChiTiet))
	for i := 0; i < len(chiTiet); i++ {
		chiTiet[i] = ReportTheLoaiPresenter{
			MaTheLoai:  tl.ChiTiet[i].MaTheLoai.String(),
			TenTheLoai: tl.ChiTiet[i].TenTheLoai,
			SoLuotMuon: tl.ChiTiet[i].SoLuotMuon,
		}
	}
	return YearReportTheLoaiPresenter{
		Nam:     tl.Nam,
		ChiTiet: chiTiet,
	}
}
