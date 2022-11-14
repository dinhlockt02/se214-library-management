package handler

import (
	"daijoubuteam.xyz/se214-library-management/core/entity"
	dausach "daijoubuteam.xyz/se214-library-management/usecase/dau_sach"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DauSachDto struct {
	TenDauSach string   `json:"tenDauSach" binding:"required"`
	TheLoai    []string `json:"theLoai" binding:"required"`
	TacGia     []string `json:"tacGia" binding:"required"`
}

type DauSachPresenter struct {
	MaDauSach  string              `json:"maDauSach" binding:"required"`
	TenDauSach string              `json:"tenDauSach" binding:"required"`
	TheLoai    []*TheLoaiPresenter `json:"theLoai" binding:"required"`
	TacGia     []*TacGiaPresenter  `json:"tacGia" binding:"required"`
}

func NewDauSachPresenter(dauSach *entity.DauSach) *DauSachPresenter {
	return &DauSachPresenter{
		MaDauSach:  dauSach.MaDauSach.String(),
		TenDauSach: dauSach.TenDauSach,
		TheLoai:    NewDanhSachTheLoaiPresenter(dauSach.TheLoai),
		TacGia:     NewDanhSachTacGiaPresenter(dauSach.TacGia),
	}
}

func NewDanhSachDauSachPresenter(danhSachDauSach []*entity.DauSach) []*DauSachPresenter {
	danhSachDauSachPresenter := make([]*DauSachPresenter, 0, len(danhSachDauSach))
	for _, dauSach := range danhSachDauSach {
		danhSachDauSachPresenter = append(danhSachDauSachPresenter, NewDauSachPresenter(dauSach))
	}
	return danhSachDauSachPresenter
}

func GetDanhSachDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachDauSach, err := usecase.GetDanhSachDauSach()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDanhSachDauSachPresenter(danhSachDauSach))
	}
}

func CreateDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var dauSachDto DauSachDto
		err := context.ShouldBind(&dauSachDto)
		if err != nil || len(dauSachDto.TacGia) == 0 || len(dauSachDto.TheLoai) == 0 {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		theLoai := make([]*entity.ID, 0, len(dauSachDto.TheLoai))
		for _, tldto := range dauSachDto.TheLoai {
			tl, err := entity.StringToID(tldto)
			if ErrorHandling(context, err) {
				return
			}
			theLoai = append(theLoai, tl)
		}
		tacGia := make([]*entity.ID, 0, len(dauSachDto.TacGia))
		for _, tgdto := range dauSachDto.TacGia {
			tg, err := entity.StringToID(tgdto)
			if ErrorHandling(context, err) {
				return
			}
			tacGia = append(tacGia, tg)
		}
		dauSach, err := usecase.CreateDauSach(dauSachDto.TenDauSach, theLoai, tacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusCreated, NewDauSachPresenter(dauSach))
	}
}

func UpdateDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var dauSachDto DauSachDto
		maDauSach, err := entity.StringToID(context.Param("maDauSach"))
		if ErrorHandling(context, err) {
			return
		}
		err = context.ShouldBind(&dauSachDto)
		if err != nil || len(dauSachDto.TacGia) == 0 || len(dauSachDto.TheLoai) == 0 {
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		theLoai := make([]*entity.ID, 0, len(dauSachDto.TheLoai))
		for _, tldto := range dauSachDto.TheLoai {
			tl, err := entity.StringToID(tldto)
			if ErrorHandling(context, err) {
				return
			}
			theLoai = append(theLoai, tl)
		}
		tacGia := make([]*entity.ID, 0, len(dauSachDto.TacGia))
		for _, tgdto := range dauSachDto.TacGia {
			tg, err := entity.StringToID(tgdto)
			if ErrorHandling(context, err) {
				return
			}
			tacGia = append(tacGia, tg)
		}
		dauSach, err := usecase.UpdateDauSach(maDauSach, dauSachDto.TenDauSach, theLoai, tacGia)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDauSachPresenter(dauSach))
	}
}

func GetDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDauSach, err := entity.StringToID(context.Param("maDauSach"))
		if ErrorHandling(context, err) {
			return
		}
		dauSach, err := usecase.GetDauSach(maDauSach)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, NewDauSachPresenter(dauSach))
	}
}

func DeleteDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDauSach, err := entity.StringToID(context.Param("maDauSach"))
		if ErrorHandling(context, err) {
			return
		}
		err = usecase.RemoveDauSach(maDauSach)
		if ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func MakeDauSachHandler(r *gin.Engine, usecase dausach.DauSachUsecase) {
	r.GET("/dausach", GetDanhSachDauSach(usecase))
	r.POST("/dausach", CreateDauSach(usecase))
	r.GET("/dausach/:maDauSach", GetDauSach(usecase))
	r.PUT("/dausach/:maDauSach", UpdateDauSach(usecase))
	r.DELETE("/dausach/:maDauSach", DeleteDauSach(usecase))
}
