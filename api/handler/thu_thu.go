package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	thuthu "daijoubuteam.xyz/se214-library-management/usecase/thu_thu"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MakeThuThuHandler(r *gin.Engine, usecase thuthu.ThuThuUsecase) {
	r.GET("/thuthu/:maThuThu", getThuThu(usecase))
	r.PUT("/thuthu/:maThuThu", updateThuThu(usecase))
}

func getThuThu(thuThuUsecase thuthu.ThuThuUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maThuThu, err := entity.StringToID(context.Param("maThuThu"))
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid id", err))
			return
		}
		var thuThu *entity.ThuThu
		thuThu, err = thuThuUsecase.GetThuThu(maThuThu)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewThuThuPresenter(thuThu))
	}
}
func updateThuThu(thuThuUsecase thuthu.ThuThuUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maThuThu, err := entity.StringToID(context.Param("maThuThu"))
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid id", err))
			return
		}
		var thuThuDto dto.ThuThuDto
		if err = context.ShouldBindJSON(&thuThuDto); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid data", err))
			return
		}
		var ngaySinh time.Time
		ngaySinh, err = time.Parse(utils.TimeLayout, thuThuDto.NgaySinh)
		if err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid time", err))
			return
		}
		var thuThu *entity.ThuThu
		thuThu, err = thuThuUsecase.UpdateThuThu(maThuThu, thuThuDto.HoTen, &ngaySinh, thuThuDto.Email, thuThuDto.SoDienThoai, true)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewThuThuPresenter(thuThu))
	}
}
