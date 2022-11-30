package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	docgia "daijoubuteam.xyz/se214-library-management/usecase/doc_gia"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getDanhSachDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachDocGia, err := usecase.GetDanhSachDocGia()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachDocGiaPresenter(danhSachDocGia))
	}
}

func deleteDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDocGia, err := entity.StringToID(context.Param("maDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		err = usecase.RemoveDocGia(maDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func getDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDocGia, err := entity.StringToID(context.Param("maDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		docGia, err := usecase.GetDocGia(maDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDocGiaPresenter(docGia))
	}
}

func createDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var docGiaDto dto.DocGiaDto
		err := context.ShouldBind(&docGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maLoaiDocGia, err := entity.StringToID(docGiaDto.MaLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		ngaySinh, err := time.Parse(utils.TimeLayout, docGiaDto.NgaySinh)
		ngayLapThe, err := time.Parse(utils.TimeLayout, docGiaDto.NgayLapThe)

		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		docGia, err := usecase.CreateDocGia(docGiaDto.HoTen, maLoaiDocGia, &ngaySinh, docGiaDto.DiaChi, docGiaDto.Email, &ngayLapThe)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewDocGiaPresenter(docGia))
	}
}

func updateDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var docGiaDto dto.DocGiaDto
		err := context.ShouldBind(&docGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maDocGia, err := entity.StringToID(context.Param("maDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		maLoaiDocGia, err := entity.StringToID(docGiaDto.MaLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		ngaySinh, err := time.Parse(utils.TimeLayout, docGiaDto.NgaySinh)
		ngayLapThe, err := time.Parse(utils.TimeLayout, docGiaDto.NgayLapThe)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		docGia, err := usecase.UpdateDocGia(maDocGia, docGiaDto.HoTen, maLoaiDocGia, &ngaySinh, docGiaDto.DiaChi, docGiaDto.Email, &ngayLapThe)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDocGiaPresenter(docGia))
		context.JSON(http.StatusOK, presenter.NewDocGiaPresenter(docGia))
	}
}

func MakeDocGiaHandler(r *gin.Engine, usecase docgia.DocGiaUsecase) {
	r.GET("/docgia", getDanhSachDocGia(usecase))
	r.POST("/docgia", createDocGia(usecase))
	r.GET("/docgia/:maDocGia", getDocGia(usecase))
	r.DELETE("/docgia/:maDocGia", deleteDocGia(usecase))
	r.PUT("/docgia/:maDocGia", updateDocGia(usecase))
}
