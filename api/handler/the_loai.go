package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	theloai "daijoubuteam.xyz/se214-library-management/usecase/the_loai"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getDanhSachTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachTheLoai, err := usecase.GetDanhSachTheLoai()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachTheLoaiPresenter(danhSachTheLoai))
	}
}

func createTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var theLoaiDto dto.TheLoaiDto
		err := context.ShouldBind(&theLoaiDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		theLoai, err := usecase.CreateTheLoai(theLoaiDto.TenTheLoai)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewTheLoaiPresenter(theLoai))
	}
}

func getTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTheLoai, err := entity.StringToID(context.Param("maTheLoai"))
		if ErrorHandling(context, err) {
			return
		}
		theLoai, err := usecase.GetTheLoai(maTheLoai)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewTheLoaiPresenter(theLoai))
	}
}

func updateTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var theLoaiDto dto.TheLoaiDto
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
		context.JSON(http.StatusOK, presenter.NewTheLoaiPresenter(theLoai))
	}
}

func deleteTheLoai(usecase theloai.TheLoaiUsecase) gin.HandlerFunc {
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
	r.GET("/theloai", getDanhSachTheLoai(usecase))
	r.POST("/theloai", createTheLoai(usecase))
	r.GET("/theloai/:maTheLoai", getTheLoai(usecase))
	r.PUT("/theloai/:maTheLoai", updateTheLoai(usecase))
	r.DELETE("/theloai/:maTheLoai", deleteTheLoai(usecase))
}
