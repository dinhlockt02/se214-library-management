package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func MakeReportHandler(r *gin.Engine, usecase usecase.ReportUsecase) {
	r.GET("/report/:year/:month", getMonthReport(usecase))
	r.GET("/report/:year", getYearReport(usecase))
}

func getYearReport(reportUsecase usecase.ReportUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		year := context.Param("year")
		rs, err := strconv.Atoi(year)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid year", err))
		}
		result, err := reportUsecase.GetYearReport(uint(rs))
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewYearReportTheLoaiPresenter(result))
	}
}

func getMonthReport(reportUsecase usecase.ReportUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		year := context.Param("year")
		rsyear, err := strconv.Atoi(year)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid year", err))
		}
		month := context.Param("month")
		rsmonth, err := strconv.Atoi(month)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid month", err))
		}
		result, err := reportUsecase.GetMonthReport(uint(rsmonth), uint(rsyear))
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewMonthReportTheLoaiPresenter(result))
	}
}
