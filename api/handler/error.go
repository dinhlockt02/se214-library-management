package handler

import (
	coreerror "daijoubuteam.xyz/se214-library-management/core/error"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandling(context *gin.Context, err error) bool {
	if err != nil {
		switch err.(type) {
		case *coreerror.UnauthorizedError:
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		case *coreerror.BadRequestError:
			context.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		case *coreerror.NotFoundError:
			context.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		case *coreerror.ForbiddenError:
			context.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		case *coreerror.ConflictError:
			context.AbortWithStatusJSON(http.StatusConflict, err.Error())
		case *coreerror.InternalServerError:
			context.AbortWithStatus(http.StatusInternalServerError)
			fmt.Println(err)
		default:
			context.AbortWithStatus(http.StatusInternalServerError)
			fmt.Println(err)
		}
		return true
	}
	return false
}
