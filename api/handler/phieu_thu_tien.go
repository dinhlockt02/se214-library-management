package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/usecase/thu_tien"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MakePhieuThuTienHandler(r *gin.Engine, thuTienUsecase thu_tien.Usecase) {
	r.GET("/phieu-thu-tien/:maDocGia", GetPhieuThuTienByMaDocGia(thuTienUsecase))
	r.POST("phieu-thu-tien", CreatePhieuThuTien(thuTienUsecase))
}

func CreatePhieuThuTien(thuTienUsecase thu_tien.Usecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var phieuThuTienDto dto.PhieuThuTienDto
		if err := context.ShouldBind(&phieuThuTienDto); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", nil))
			return
		}
		ngayThu, err := time.Parse(utils.TimeLayout, phieuThuTienDto.NgayThu)
		phieuThuTien, err := thuTienUsecase.CreatePhieuThuTien(phieuThuTienDto.MaDocGia, phieuThuTienDto.SoTienThu, &ngayThu)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewPhieuThuTienPresenter(phieuThuTien))
	}
}

func GetPhieuThuTienByMaDocGia(thuTienUsecase thu_tien.Usecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDocGia := context.Param("maDocGia")
		danhSachPhieuThuTien, err := thuTienUsecase.GetPhieuThuTienByMaDocGia(maDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachPhieuThuTienPresenter(danhSachPhieuThuTien))
	}
}
