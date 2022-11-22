package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	"daijoubuteam.xyz/se214-library-management/core/entity"
	dausach "daijoubuteam.xyz/se214-library-management/usecase/dau_sach"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getDanhSachDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		danhSachDauSach, err := usecase.GetDanhSachDauSach()
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDanhSachDauSachPresenter(danhSachDauSach))
	}
}

func createDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var dauSachDto dto.DauSachDto
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
		context.JSON(http.StatusCreated, presenter.NewDauSachPresenter(dauSach))
	}
}

func updateDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var dauSachDto dto.DauSachDto
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
		context.JSON(http.StatusOK, presenter.NewDauSachPresenter(dauSach))
	}
}

func getDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		maDauSach, err := entity.StringToID(context.Param("maDauSach"))
		if ErrorHandling(context, err) {
			return
		}
		dauSach, err := usecase.GetDauSach(maDauSach)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, presenter.NewDauSachPresenter(dauSach))
	}
}

func deleteDauSach(usecase dausach.DauSachUsecase) gin.HandlerFunc {
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
	r.GET("/dausach", getDanhSachDauSach(usecase))
	r.POST("/dausach", createDauSach(usecase))
	r.GET("/dausach/:maDauSach", getDauSach(usecase))
	r.PUT("/dausach/:maDauSach", updateDauSach(usecase))
	r.DELETE("/dausach/:maDauSach", deleteDauSach(usecase))
}
