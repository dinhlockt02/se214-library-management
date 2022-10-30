package handler

// import (
// 	"net/http"

// 	dto "daijoubuteam.xyz/se214-library-management/api/Dto"
// 	"daijoubuteam.xyz/se214-library-management/api/presenter"
// 	"daijoubuteam.xyz/se214-library-management/usecase/book"
// 	"github.com/gin-gonic/gin"
// )

// func GetBooks(service book.BookUsecase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		entityBooks, err := service.GetBooks()
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 		}
// 		var presenterBooks []*presenter.PresenterBook = make([]*presenter.PresenterBook, len(entityBooks))
// 		for index, entityBook := range entityBooks {
// 			presenterBooks[index] = presenter.FromEntityBook(entityBook)
// 		}
// 		c.JSON(http.StatusOK, presenterBooks)
// 	}
// }

// func PostBook(service book.BookUsecase) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var newBookDto dto.CreateBookDto
// 		err := c.ShouldBind(&newBookDto)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusBadRequest)
// 		}
// 		entityBook, err := service.CreateBook(newBookDto.Entity())
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusInternalServerError)
// 		}
// 		c.JSON(http.StatusCreated, entityBook)
// 	}
// }

// func MakeBookHandler(r *gin.Engine, bookService book.BookUsecase) {
// 	r.GET("/books", GetBooks(bookService))
// 	r.POST("/books", PostBook(bookService))
// }
