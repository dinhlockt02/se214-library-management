package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	nhapsach "daijoubuteam.xyz/se214-library-management/usecase/nhap_sach"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getDanhSachPhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachPhieuNhap, err := nhapSachUsecase.GetDanhSachPhieuNhapSach()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachPhieuNhapPresenter(danhSachPhieuNhap))
	}
}

func getPhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maPhieuNhap, err := entity.StringToID(context.Param("maPhieuNhap"))
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		phieuNhap, err := nhapSachUsecase.GetPhieuNhapSach(maPhieuNhap)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewPhieuNhapPresenter(phieuNhap))
	}
}

func createPhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var phieuNhapDto dto.PhieuNhapDto
		err := context.ShouldBind(&phieuNhapDto)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
			return
		}
		ngayLap, err := time.Parse(utils.TimeLayout, phieuNhapDto.NgayLap)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
			return
		}
		phieuNhap, err := nhapSachUsecase.CreatePhieuNhapSach(&ngayLap, phieuNhapDto.CtPhieuNhap)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewPhieuNhapPresenter(phieuNhap))
	}
}
func removePhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {

		maPhieuNhap, err := entity.StringToID(context.Param("maPhieuNhap"))
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = nhapSachUsecase.RemovePhieuNhapSach(maPhieuNhap)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func updatePhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maPhieuNhap, err := entity.StringToID(context.Param("maPhieuNhap"))
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ma phieu nhap", err))
			return
		}
		var phieuNhapDto dto.UpdatePhieuNhapDto
		err = context.ShouldBind(&phieuNhapDto)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
			return
		}
		ngayLap, err := time.Parse(utils.TimeLayout, phieuNhapDto.NgayLap)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid time format", err))
			return
		}
		phieuNhap, err := nhapSachUsecase.UpdatePhieuNhapSach(maPhieuNhap, &ngayLap)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewPhieuNhapPresenter(phieuNhap))

	}
}

func MakePhieuNhapHandler(r *gin.Engine, nhapSachUsecase nhapsach.NhapSachUsecase) {
	r.GET("/phieunhap", getDanhSachPhieuNhap(nhapSachUsecase))
	r.POST("/phieunhap", createPhieuNhap(nhapSachUsecase))
	r.GET("/phieunhap/:maPhieuNhap", getPhieuNhap(nhapSachUsecase))
	r.PATCH("/phieunhap/:maPhieuNhap", updatePhieuNhap(nhapSachUsecase))
	r.DELETE("/phieunhap/:maPhieuNhap", removePhieuNhap(nhapSachUsecase))
}
