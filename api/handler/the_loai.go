package handler

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	theloai "daijoubuteam.xyz/se214-library-management/usecase/the_loai"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TheLoaiDto struct {
	TenTheLoai string `json:"tenTheLoai" binding:"required"`
}

type TheLoaiPresenter struct {
	MaTheLoai  string `json:"maTheLoai" binding:"required"`
	TenTheLoai string `json:"tenTheLoai" binding:"required"`
}

func NewDanhSachTheLoaiPresenter(danhSachTheLoai []*entity.TheLoai) []*TheLoaiPresenter {
	danhSachTheLoaiPresenter := make([]*TheLoaiPresenter, len(danhSachTheLoai))
	for index, theLoai := range danhSachTheLoai {
		danhSachTheLoaiPresenter[index] = NewTheLoaiPresenter(theLoai)
	}
	return danhSachTheLoaiPresenter
}

func NewTheLoaiPresenter(theLoai *entity.TheLoai) *TheLoaiPresenter {
	return &TheLoaiPresenter{
		MaTheLoai:  theLoai.MaTheLoai.String(),
		TenTheLoai: theLoai.TenTheLoai,
	}
}

func GetDanhSachTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachTheLoai, err := usecase.GetDanhSachTheLoai()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDanhSachTheLoaiPresenter(danhSachTheLoai))
	}
}

func CreateTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var theLoaiDto TheLoaiDto
		err := context.ShouldBind(&theLoaiDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		theLoai, err := usecase.CreateTheLoai(theLoaiDto.TenTheLoai)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, NewTheLoaiPresenter(theLoai))
	}
}

func GetTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTheLoai, err := entity.StringToID(context.Param("maTheLoai"))
		if ErrorHandling(context, err) {
			return
		}
		theLoai, err := usecase.GetTheLoai(maTheLoai)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewTheLoaiPresenter(theLoai))
	}
}

func UpdateTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var theLoaiDto TheLoaiDto
		err := context.ShouldBind(&theLoaiDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maTheLoai, err := entity.StringToID(context.Param("maTheLoai"))
		if ErrorHandling(context, err) {
			return
		}
		theLoai, err := usecase.UpdateTheLoai(maTheLoai, theLoaiDto.TenTheLoai)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewTheLoaiPresenter(theLoai))
	}
}

func DeleteTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTheLoai, err := entity.StringToID(context.Param("maTheLoai"))
		if ErrorHandling(context, err) {
			return
		}
		err = usecase.RemoveTheLoai(maTheLoai)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func MakeTheLoaiHandler(r *gin.Engine, usecase theloai.TheLoaiUsecase) {
	r.GET("/theloai", GetDanhSachTheLoai(usecase))
	r.POST("/theloai", CreateTheLoai(usecase))
	r.GET("/theloai/:maTheLoai", GetTheLoai(usecase))
	r.PUT("/theloai/:maTheLoai", UpdateTheLoai(usecase))
	r.DELETE("/theloai/:maTheLoai", DeleteTheLoai(usecase))
}
