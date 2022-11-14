package handler

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	tacgia "daijoubuteam.xyz/se214-library-management/usecase/tac_gia"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TacGiaDto struct {
	TenTacGia string `json:"tenTacGia" binding:"required"`
}

type TacGiaPresenter struct {
	MaTacGia  string `json:"maTacGia" binding:"required"`
	TenTacGia string `json:"tenTacGia" binding:"required"`
}

func NewTacGiaPresenter(tacGia *entity.TacGia) *TacGiaPresenter {
	return &TacGiaPresenter{
		MaTacGia:  tacGia.MaTacGia.String(),
		TenTacGia: tacGia.TenTacGia,
	}
}

func NewDanhSachTacGiaPresenter(danhSachTacGia []*entity.TacGia) []*TacGiaPresenter {
	danhSachTacGiaPresenter := make([]*TacGiaPresenter, len(danhSachTacGia))
	for index, tacGia := range danhSachTacGia {
		danhSachTacGiaPresenter[index] = NewTacGiaPresenter(tacGia)
	}
	return danhSachTacGiaPresenter
}

func GetDanhSachTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachTacGia, err := usecase.GetDanhSachTacGia()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDanhSachTacGiaPresenter(danhSachTacGia))
	}
}

func GetTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTacGia, err := entity.StringToID(context.Param("maTacGia"))
		if ErrorHandling(context, err) {
			return
		}
		tacGia, err := usecase.GetTacGia(maTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewTacGiaPresenter(tacGia))
	}
}

func CreateTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var tacGiaDto TacGiaDto
		err := context.ShouldBind(&tacGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tacGia, err := usecase.CreateTacGia(tacGiaDto.TenTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, NewTacGiaPresenter(tacGia))
	}
}

func UpdateTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maTacGia, err := entity.StringToID(context.Param("maTacGia"))
		if ErrorHandling(context, err) {
			return
		}
		var tacGiaDto TacGiaDto
		err = context.ShouldBind(&tacGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		tacGia, err := usecase.UpdateTacGia(maTacGia, tacGiaDto.TenTacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewTacGiaPresenter(tacGia))
	}
}

func DeleteTacGia(usecase tacgia.TacGiaUsecase) gin.HandlerFunc {
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
	r.GET("/tacgia", GetDanhSachTacGia(usecase))
	r.POST("/tacgia", CreateTacGia(usecase))
	r.GET("/tacgia/:maTacGia", GetTacGia(usecase))
	r.PUT("/tacgia/:maTacGia", UpdateTacGia(usecase))
	r.DELETE("/tacgia/:maTacGia", DeleteTacGia(usecase))
}
