package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/usecase/tra_sach"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MakePhieuTraHandler(r *gin.Engine, traSachUsecase tra_sach.Usecase) {
	r.GET("/phieu-tra", getDanhSachPhieuTra(traSachUsecase))
	r.POST("/phieu-tra", createPhieuTra(traSachUsecase))
}

func getDanhSachPhieuTra(traSachUsecase tra_sach.Usecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var err error

		var phieuTra []*entity.PhieuTra
		if phieuTra, err = traSachUsecase.GetPhieuTra(); err != nil {
			ErrorHandling(context, err)
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachPhieuTraPresenter(phieuTra))
		return
	}
}

func createPhieuTra(traSachUsecase tra_sach.Usecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var phieuTraDto dto.PhieuTraDto
		var err error
		if err = context.ShouldBindJSON(&phieuTraDto); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", nil))
			return
		}
		var maSach *entity.ID
		if maSach, err = entity.StringToID(phieuTraDto.MaSach); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid sach id", nil))
			return
		}
		var ngayTra time.Time
		if ngayTra, err = time.Parse(utils.TimeLayout, phieuTraDto.NgayTra); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ngay tra", nil))
			return
		}
		var phieuTra *entity.PhieuTra
		if phieuTra, err = traSachUsecase.CreatePhieuTra(maSach, phieuTraDto.GhiChu, &ngayTra); err != nil {
			ErrorHandling(context, err)
			return
		}
		context.JSON(http.StatusCreated, presenter.NewPhieuTraPresenter(phieuTra))
		return
	}
}
