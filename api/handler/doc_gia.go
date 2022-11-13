package handler

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	docgia "daijoubuteam.xyz/se214-library-management/usecase/doc_gia"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type DocGiaDto struct {
	HoTen        string `json:"hoTen" binding:"required"`
	MaLoaiDocGia string `json:"maLoaiDocGia" binding:"required"`
	NgaySinh     string `json:"ngaySinh" binding:"required"`
	DiaChi       string `json:"diaChi" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NgayLapThe   string `json:"ngayLapThe" binding:"required"`
}

type DocGiaPresenter struct {
	MaDocGia   string               `json:"maDocGia" binding:"required"`
	HoTen      string               `json:"hoTen" binding:"required"`
	LoaiDocGia *LoaiDocGiaPresenter `json:"loaiDocGia" binding:"required"`
	NgaySinh   *time.Time           `json:"ngaySinh" binding:"required"`
	DiaChi     string               `json:"diaChi" binding:"required"`
	Email      string               `json:"email" binding:"required"`
	NgayLapThe *time.Time           `json:"ngayLapThe" binding:"required"`
	NgayHetHan *time.Time           `json:"ngayHetHan" binding:"required"`
	TongNo     uint                 `json:"tongNo" binding:"required"`
}

func NewDanhSachDocGiaPresenter(danhSachDocGia []*entity.DocGia) []*DocGiaPresenter {
	danhSachDocGiaPresenter := make([]*DocGiaPresenter, len(danhSachDocGia))
	for index, docGia := range danhSachDocGia {
		danhSachDocGiaPresenter[index] = NewDocGiaPresenter(docGia)
	}
	return danhSachDocGiaPresenter

}

func NewDocGiaPresenter(docGia *entity.DocGia) *DocGiaPresenter {
	return &DocGiaPresenter{
		MaDocGia:   docGia.MaDocGia.String(),
		HoTen:      docGia.HoTen,
		LoaiDocGia: NewLoaiDocGiaPresenter(docGia.LoaiDocGia),
		NgaySinh:   docGia.NgaySinh,
		DiaChi:     docGia.DiaChi,
		Email:      docGia.Email,
		NgayLapThe: docGia.NgayLapThe,
		NgayHetHan: docGia.NgayHetHan,
		TongNo:     docGia.TongNo,
	}
}

func GetDanhSachDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachDocGia, err := usecase.GetDanhSachDocGia()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDanhSachDocGiaPresenter(danhSachDocGia))
	}
}

func DeleteDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDocGia, err := entity.StringToID(context.Param("maDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		err = usecase.RemoveDocGia(maDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func GetDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDocGia, err := entity.StringToID(context.Param("maDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		docGia, err := usecase.GetDocGia(maDocGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDocGiaPresenter(docGia))
	}
}

func CreateDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var docGiaDto DocGiaDto
		err := context.ShouldBind(&docGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maLoaiDocGia, err := entity.StringToID(docGiaDto.MaLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		ngaySinh, err := time.Parse(utils.TimeLayout, docGiaDto.NgaySinh)
		ngayLapThe, err := time.Parse(utils.TimeLayout, docGiaDto.NgayLapThe)

		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		docGia, err := usecase.CreateDocGia(docGiaDto.HoTen, maLoaiDocGia, &ngaySinh, docGiaDto.DiaChi, docGiaDto.Email, &ngayLapThe)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, NewDocGiaPresenter(docGia))
	}
}

func UpdateDocGia(usecase docgia.DocGiaUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var docGiaDto DocGiaDto
		err := context.ShouldBind(&docGiaDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		maDocGia, err := entity.StringToID(context.Param("maDocGia"))
		if ErrorHandling(context, err) {
			return
		}
		maLoaiDocGia, err := entity.StringToID(docGiaDto.MaLoaiDocGia)
		if ErrorHandling(context, err) {
			return
		}
		ngaySinh, err := time.Parse(utils.TimeLayout, docGiaDto.NgaySinh)
		ngayLapThe, err := time.Parse(utils.TimeLayout, docGiaDto.NgayLapThe)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		docGia, err := usecase.UpdateDocGia(maDocGia, docGiaDto.HoTen, maLoaiDocGia, &ngaySinh, docGiaDto.DiaChi, docGiaDto.Email, &ngayLapThe)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, NewDocGiaPresenter(docGia))
	}
}

func MakeDocGiaHandler(r *gin.Engine, usecase docgia.DocGiaUsecase) {
	r.GET("/docgia", GetDanhSachDocGia(usecase))
	r.POST("/docgia", CreateDocGia(usecase))
	r.GET("/docgia/:maDocGia", GetDocGia(usecase))
	r.DELETE("/docgia/:maDocGia", DeleteDocGia(usecase))
	r.PUT("/docgia/:maDocGia", UpdateDocGia(usecase))
}
