package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	nhapsach "daijoubuteam.xyz/se214-library-management/usecase/nhap_sach"
	"daijoubuteam.xyz/se214-library-management/utils"
	"fmt"
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
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ngayLap, err := time.Parse(utils.TimeLayout, phieuNhapDto.NgayLap)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		phieuNhap, err := nhapSachUsecase.CreatePhieuNhapSach(&ngayLap)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewPhieuNhapPresenter(phieuNhap))
	}
}

func addCtPhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var ctPhieuNhapDto dto.CtPhieuNhapDto
		err := context.ShouldBind(&ctPhieuNhapDto)
		if err != nil {
			fmt.Println(err)
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maPhieuNhap, err := entity.StringToID(context.Param("maPhieuNhap"))
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maDauSach, err := entity.StringToID(ctPhieuNhapDto.MaDauSach)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ctPhieuNhap, err := nhapSachUsecase.AddChiTietPhieuNhapSach(
			maPhieuNhap,
			maDauSach,
			ctPhieuNhapDto.NhaXuatBan,
			ctPhieuNhapDto.TriGia,
			ctPhieuNhapDto.NamXuatBan,
			ctPhieuNhapDto.TinhTrang,
			ctPhieuNhapDto.DonGia,
		)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewCtPhieuNhapPresenter(ctPhieuNhap))
	}
}

func removeCtPhieuNhap(nhapSachUsecase nhapsach.NhapSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {

		maSach, err := entity.StringToID(context.Param("maSach"))
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}

		err = nhapSachUsecase.RemoveChiTietPhieuNhapSach(maSach)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
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

			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		var phieuNhapDto dto.PhieuNhapDto
		err = context.ShouldBind(&phieuNhapDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		ngayLap, err := time.Parse(utils.TimeLayout, phieuNhapDto.NgayLap)
		if err != nil {
			fmt.Println(err)
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		phieuNhap, err := nhapSachUsecase.UpdatePhieuNhapSach(maPhieuNhap, &ngayLap)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewPhieuNhapPresenter(phieuNhap))

	}
}

func MakePhieuNhapHandler(r *gin.Engine, nhapSachUsecase nhapsach.NhapSachUsecase) {
	r.GET("/phieunhap", getDanhSachPhieuNhap(nhapSachUsecase))
	r.POST("/phieunhap", createPhieuNhap(nhapSachUsecase))
	r.GET("/phieunhap/:maPhieuNhap", getPhieuNhap(nhapSachUsecase))
	r.PATCH("/phieunhap/:maPhieuNhap", updatePhieuNhap(nhapSachUsecase))
	r.DELETE("/phieunhap/:maPhieuNhap", removePhieuNhap(nhapSachUsecase))
	r.POST("/phieunhap/:maPhieuNhap", addCtPhieuNhap(nhapSachUsecase))
	r.DELETE("/phieunhap/:maPhieuNhap/:maSach", removeCtPhieuNhap(nhapSachUsecase))
}
