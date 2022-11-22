package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	tacgia "daijoubuteam.xyz/se214-library-management/usecase/tac_gia"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getDanhSachTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachTacGia, err := usecase.GetDanhSachTacGia()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachTacGiaPresenter(danhSachTacGia))
	}
}

func getTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTacGia, err := entity.StringToID(context.Param("maTacGia"))
		if ErrorHandling(context, err) {
			return
		}
		tacGia, err := usecase.GetTacGia(maTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewTacGiaPresenter(tacGia))
	}
}

func createTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var tacGiaDto dto.TacGiaDto
		err := context.ShouldBind(&tacGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tacGia, err := usecase.CreateTacGia(tacGiaDto.TenTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewTacGiaPresenter(tacGia))
	}
}

func updateTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTacGia, err := entity.StringToID(context.Param("maTacGia"))
		if ErrorHandling(context, err) {
			return
		}
		var tacGiaDto dto.TacGiaDto
		err = context.ShouldBind(&tacGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tacGia, err := usecase.UpdateTacGia(maTacGia, tacGiaDto.TenTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewTacGiaPresenter(tacGia))
	}
}

func deleteTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTacGia, err := entity.StringToID(context.Param("maTacGia"))
		if ErrorHandling(context, err) {
			return
		}
		err = usecase.DeleteTacGia(maTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}
func MakeTacGiaHandler(r *gin.Engine, usecase tacgia.TacGiaUsecase) {
	r.GET("/tacgia", getDanhSachTacGia(usecase))
	r.POST("/tacgia", createTacGia(usecase))
	r.GET("/tacgia/:maTacGia", getTacGia(usecase))
	r.PUT("/tacgia/:maTacGia", updateTacGia(usecase))
	r.DELETE("/tacgia/:maTacGia", deleteTacGia(usecase))
}
