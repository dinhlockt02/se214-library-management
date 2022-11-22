package handler

import (
	"daijoubuteam.xyz/se214-library-management/api/dto"
	"daijoubuteam.xyz/se214-library-management/api/presenter"
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
		token, err := service.Login(loginDto.Email, loginDto.Password)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, &presenter.LoginPresenter{
			Token: *token,
		})
	}
}

func MakeAuthHandler(r *gin.Engine, authUsecase auth.AuthUsecase) {
	r.POST("/auth/login", login(authUsecase))
}
