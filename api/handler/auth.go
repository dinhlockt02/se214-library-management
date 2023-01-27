package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"daijoubuteam.xyz/se214-library-management/usecase/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(service auth.AuthUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginDto dto.LoginDto
		err := context.ShouldBind(&loginDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		token, maThuThu, err := service.Login(loginDto.Email, loginDto.Password)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, &presenter.LoginPresenter{
			Token:    *token,
			MaThuThu: *maThuThu,
		})
	}
}

func MakeAuthHandler(r *gin.Engine, authUsecase auth.AuthUsecase) {
	r.POST("/auth/login", login(authUsecase))
	r.POST("/auth/forget-password", forgetPassword(authUsecase))
	r.POST("/auth/reset-password", resetPassword(authUsecase))
}

func resetPassword(usecase auth.AuthUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var resetPasswordDto dto.ResetPassword
		if err := context.ShouldBindJSON(&resetPasswordDto); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
			return
		}
		if err := usecase.ResetPassword(resetPasswordDto.Code, resetPasswordDto.Password); ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}

func forgetPassword(usecase auth.AuthUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var forgetPasswordDto dto.ForgetPassword
		if err := context.ShouldBindJSON(&forgetPasswordDto); err != nil {
			ErrorHandling(context, coreerror.NewBadRequestError("Invalid json data", err))
			return
		}
		if err := usecase.ForgetPassword(forgetPasswordDto.Email); ErrorHandling(context, err) {
			return
		}
		context.Status(http.StatusOK)
	}
}
