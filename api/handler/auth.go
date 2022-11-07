package handler

import (
	"daijoubuteam.xyz/se214-library-management/usecase/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginPresenter struct {
	Token string `json:"token"`
}

func Login(service auth.AuthUsecase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginDto LoginDto
		err := context.ShouldBind(&loginDto)
		if err != nil {
			context.AbortWithStatus(http.StatusBadRequest)
		}
		token, err := service.Login(loginDto.Email, loginDto.Password)
		if ErrorHandling(context, err) {
			return
		}
		context.JSON(http.StatusOK, &LoginPresenter{
			Token: *token,
		})
	}
}

func MakeAuthHandler(r *gin.Engine, authUsecase auth.AuthUsecase) {
	r.POST("/auth/login", Login(authUsecase))
}
