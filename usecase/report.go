package usecase

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	"daijoubuteam.xyz/se214-library-management/core/repository"
)

type ReportUsecase interface {
	GetMonthReport(thang uint, nam uint) (*entity.MonthReportTheLoai, error)
	GetYearReport(nam uint) (*entity.YearReportTheLoai, error)
}

type ReportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportUsecase {
	return ReportService{repo}
}

func (r ReportService) GetMonthReport(thang uint, nam uint) (*entity.MonthReportTheLoai, error) {
	return r.repo.GetMonthReport(thang, nam)
}

func (r ReportService) GetYearReport(nam uint) (*entity.YearReportTheLoai, error) {
	return r.repo.GetYearReport(nam)
}
