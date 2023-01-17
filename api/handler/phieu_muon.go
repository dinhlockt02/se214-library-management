package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/usecase/muon_sach"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getDanhSachPhieuMuon(muonSachUsecase muon_sach.Usecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		if danhSachPhieuMuon, err := muonSachUsecase.GetPhieuMuon(); ErrorHandling(context, err) {
		} else {
			context.JSONP(http.StatusOK, presenter.NewDanhSachPhieuMuonPresenter(danhSachPhieuMuon))
		}
		return
	}
}

func createPhieuMuon(usecase muon_sach.Usecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var phieuMuonDto dto.PhieuMuonDto
		var err error
		if err = context.ShouldBindJSON(&phieuMuonDto); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
		}
		var ngayMuon time.Time
		if ngayMuon, err = time.Parse(utils.TimeLayout, phieuMuonDto.NgayMuon); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ngay muon", err))
		}
		var maSach *entity.ID
		if maSach, err = entity.StringToID(phieuMuonDto.MaSach); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ma sach", err))
		}
		var maDocGia *entity.ID
		if maDocGia, err = entity.StringToID(phieuMuonDto.MaDocGia); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid ma doc gia", err))
		}
		var phieuMuon *entity.PhieuMuon
		if phieuMuon, err = usecase.CreatePhieuMuon(&ngayMuon, maSach, maDocGia); ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, presenter.NewPhieuMuonPresenter(phieuMuon))
		return
	}
}

func MakePhieuMuonHandler(r *gin.Engine, muonSachUsecase muon_sach.Usecase) {
	r.GET("/phieu-muon", getDanhSachPhieuMuon(muonSachUsecase))
	r.POST("/phieu-muon", createPhieuMuon(muonSachUsecase))
}
