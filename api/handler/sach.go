package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/usecase/sach"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getDanhSachSach(sachUsecase sach.SachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachSach, err := sachUsecase.GetDanhSachSach()
		if err != nil {
			ErrorHandling(context, err)
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachSachPresenter(danhSachSach))
	}
}

func getSach(sachUsecase sach.SachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maSach, err := entity.StringToID(context.Param("maSach"))
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ma sach", nil))
		}
		sach, err := sachUsecase.GetSach(maSach)
		if err != nil {
			ErrorHandling(context, err)
			return
		}
		context.JSON(http.StatusOK, presenter.NewSachPresenter(sach))
	}
}
func updateSach(sachUsecase sach.SachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maSach, err := entity.StringToID(context.Param("maSach"))
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ma sach", err))
			return
		}
		var updateSachDto dto.UpdateSachDto
		err = context.ShouldBindJSON(&updateSachDto)
		if err != nil {
			fmt.Println(err)
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
			return
		}
		sach, err := sachUsecase.UpdateSach(maSach, updateSachDto.NhaXuatBan, updateSachDto.TriGia, updateSachDto.NamXuatBan, updateSachDto.TinhTrang, updateSachDto.GhiChu)
		if err != nil {
			ErrorHandling(context, err)
			return
		}
		context.JSON(http.StatusOK, presenter.NewSachPresenter(sach))
	}
}

func MakeSachHandler(r *gin.Engine, sachUsecase sach.SachUsecase) {
	r.GET("/sach", getDanhSachSach(sachUsecase))
	r.GET("/sach/:maSach", getSach(sachUsecase))
	r.PATCH("/sach/:maSach", updateSach(sachUsecase))
}
