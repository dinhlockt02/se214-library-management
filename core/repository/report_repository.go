package repository

import "daijoubuteam.xyz/se214-library-management/core/entity"

type ReportRepository interface {
	GetMonthReport(thang uint, nam uint) (*entity.MonthReportTheLoai, error)
	GetYearReport(nam uint) (*entity.YearReportTheLoai, error)
}
