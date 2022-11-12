package handler

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	loaidocgia "daijoubuteam.xyz/se214-library-management/usecase/loai_doc_gia"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostLoaiDocGiaDto struct {
	TenLoaiDocGia string `json:"tenLoaiDocGia" binding:"required"`
}

type PutLoaiDocGiaDto struct {
	TenLoaiDocGia string `json:"tenLoaiDocGia" binding:"required"`
}

type LoaiDocGiaPresenter struct {
	MaLoaiDocGia  string `json:"maLoaiDocGia" binding:"required"`
	TenLoaiDocGia string `json:"tenLoaiDocGia" binding:"required"`
}

func NewLoaiDocGiaPresenter(loaiDocGia *entity.LoaiDocGia) *LoaiDocGiaPresenter {
	return &LoaiDocGiaPresenter{
		MaLoaiDocGia:  loaiDocGia.MaLoaiDocGia.String(),
		TenLoaiDocGia: loaiDocGia.TenLoaiDocGia,
	}
}

func NewDanhSachLoaiDocGiaPresenter(danhSachLoaiDocGia []*entity.LoaiDocGia) []*LoaiDocGiaPresenter {
	danhSachLoaiDocGiaPresenter := make([]*LoaiDocGiaPresenter, len(danhSachLoaiDocGia))
	for index, loaiDocGia := range danhSachLoaiDocGia {
		danhSachLoaiDocGiaPresenter[index] = NewLoaiDocGiaPresenter(loaiDocGia)
	}
	return danhSachLoaiDocGiaPresenter
}

func CreateLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var createLoaiDocGiaDto PostLoaiDocGiaDto
		err := context.ShouldBind(&createLoaiDocGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		loaiDocGia, err := usecase.CreateLoaiDocGia(createLoaiDocGiaDto.TenLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, NewLoaiDocGiaPresenter(loaiDocGia))
	}
}

func GetDanhSachLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachLoaiDocGia, err := usecase.GetDanhSachLoaiDocGia()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDanhSachLoaiDocGiaPresenter(danhSachLoaiDocGia))
	}
}

func GetLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maLoaiDocGia, err := entity.StringToID(context.Param("maLoaiDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		loaiDocGia, err := usecase.GetLoaiDocGia(maLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewLoaiDocGiaPresenter(loaiDocGia))
	}
}

func PutLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maLoaiDocGia, err := entity.StringToID(context.Param("maLoaiDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		var putLoaiDocGiaDto PutLoaiDocGiaDto
		err = context.ShouldBind(&putLoaiDocGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		loaiDocGia, err := usecase.UpdateLoaiDocGia(maLoaiDocGia, putLoaiDocGiaDto.TenLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewLoaiDocGiaPresenter(loaiDocGia))
	}
}

func DeleteLoaiDocGia(usecase loaidocgia.LoaiDocGiaUsecase) gin.HandlerFunc {
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
	r.POST("/loaidocgia", CreateLoaiDocGia(usecase))
	r.GET("/loaidocgia", GetDanhSachLoaiDocGia(usecase))
	r.GET("/loaidocgia/:maLoaiDocGia", GetLoaiDocGia(usecase))
	r.PUT("/loaidocgia/:maLoaiDocGia", PutLoaiDocGia(usecase))
	r.DELETE("/loaidocgia/:maLoaiDocGia", DeleteLoaiDocGia(usecase))
}
