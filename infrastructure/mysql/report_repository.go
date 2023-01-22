package mysql

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"github.com/jmoiron/sqlx"
)

const (
	getmonthreportquery = `
		SELECT COUNT(MaPhieuMuon), TL.MaTheLoai MaTheLoai, TL.TenTheLoai TenTheLoai FROM PhieuMuon PM
		JOIN Sach S on PM.MaSach = S.MaSach
		JOIN DauSach DS on S.MaDauSach = DS.MaDauSach
		JOIN CT_TheLoai CTL on DS.MaDauSach = CTL.MaDauSach
		JOIN TheLoai TL on CTL.MaTheLoai = TL.MaTheLoai
		WHERE MONTH(PM.NgayMuon) = ?
		AND YEAR(PM.NgayMuon) = ?
		GROUP BY TL.MaTheLoai, TL.TenTheLoai`
	getyearreportquery = `
		SELECT COUNT(MaPhieuMuon), TL.MaTheLoai MaTheLoai, TL.TenTheLoai TenTheLoai FROM PhieuMuon PM
		JOIN Sach S on PM.MaSach = S.MaSach
		JOIN DauSach DS on S.MaDauSach = DS.MaDauSach
		JOIN CT_TheLoai CTL on DS.MaDauSach = CTL.MaDauSach
		JOIN TheLoai TL on CTL.MaTheLoai = TL.MaTheLoai
		WHERE YEAR(PM.NgayMuon) = ?
		GROUP BY TL.MaTheLoai, TL.TenTheLoai`
)

type ReportRepository struct {
	*sqlx.DB
}

func NewReportRepository(db *sqlx.DB) ReportRepository {
	return ReportRepository{db}
}

func (r ReportRepository) GetMonthReport(thang uint, nam uint) (*entity.MonthReportTheLoai, error) {
	var chiTietReportTheLoai []entity.ReportTheLoai
	rows, err := r.Queryx(getmonthreportquery, thang, nam)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create report", err)
	}
	for rows.Next() {
		var reportTheLoai = entity.ReportTheLoai{}
		if err = rows.Scan(&(reportTheLoai.SoLuotMuon), &(reportTheLoai.MaTheLoai), &(reportTheLoai.TenTheLoai)); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not create report", err)
		}
		chiTietReportTheLoai = append(chiTietReportTheLoai, reportTheLoai)
	}
	return &entity.MonthReportTheLoai{
		ChiTiet: chiTietReportTheLoai,
		Thang:   thang,
		Nam:     nam,
	}, nil
}

func (r ReportRepository) GetYearReport(nam uint) (*entity.YearReportTheLoai, error) {
	var chiTietReportTheLoai []entity.ReportTheLoai
	rows, err := r.Queryx(getyearreportquery, nam)
	if err != nil {
		return nil, coreerror.NewInternalServerError("database error: can't not create report", err)
	}
	for rows.Next() {
		var reportTheLoai = entity.ReportTheLoai{}
		if err = rows.Scan(&(reportTheLoai.SoLuotMuon), &(reportTheLoai.MaTheLoai), &(reportTheLoai.TenTheLoai)); err != nil {
			return nil, coreerror.NewInternalServerError("database error: can't not create report", err)
		}
		chiTietReportTheLoai = append(chiTietReportTheLoai, reportTheLoai)
	}
	return &entity.YearReportTheLoai{
		ChiTiet: chiTietReportTheLoai,
		Nam:     nam,
	}, nil
}
