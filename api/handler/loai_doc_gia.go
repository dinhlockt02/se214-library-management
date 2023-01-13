package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	loaidocgia "daijoubuteam.xyz/se214-library-management/usecase/loai_doc_gia"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var createLoaiDocGiaDto dto.PostLoaiDocGiaDto
		err := context.ShouldBind(&createLoaiDocGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		loaiDocGia, err := usecase.CreateLoaiDocGia(createLoaiDocGiaDto.TenLoaiDocGia, createLoaiDocGiaDto.SoSachToiDaDuocMuon)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewLoaiDocGiaPresenter(loaiDocGia))
	}
}

func getDanhSachLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachLoaiDocGia, err := usecase.GetDanhSachLoaiDocGia()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachLoaiDocGiaPresenter(danhSachLoaiDocGia))
	}
}

func getLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maLoaiDocGia, err := entity.StringToID(context.Param("maLoaiDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		loaiDocGia, err := usecase.GetLoaiDocGia(maLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewLoaiDocGiaPresenter(loaiDocGia))
	}
}

func putLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maLoaiDocGia, err := entity.StringToID(context.Param("maLoaiDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		var putLoaiDocGiaDto dto.PutLoaiDocGiaDto
		err = context.ShouldBind(&putLoaiDocGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		loaiDocGia, err := usecase.UpdateLoaiDocGia(maLoaiDocGia, putLoaiDocGiaDto.TenLoaiDocGia, putLoaiDocGiaDto.SoSachToiDaDuocMuon)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewLoaiDocGiaPresenter(loaiDocGia))
	}
}

func deleteLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maLoaiDocGia, err := entity.StringToID(context.Param("maLoaiDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		err = usecase.DeleteLoaiDocGia(maLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func MakeLoaiThuThuHandler(r *gin.Engine, usecase loaidocgia.LoaiDocGiaUsecase) {
	r.POST("/loaidocgia", createLoaiDocGia(usecase))
	r.GET("/loaidocgia", getDanhSachLoaiDocGia(usecase))
	r.GET("/loaidocgia/:maLoaiDocGia", getLoaiDocGia(usecase))
	r.PUT("/loaidocgia/:maLoaiDocGia", putLoaiDocGia(usecase))
	r.DELETE("/loaidocgia/:maLoaiDocGia", deleteLoaiDocGia(usecase))
}
