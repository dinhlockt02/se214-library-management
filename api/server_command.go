package api

import (
	"daijoubuteam.xyz/se214-library-management/api/handler"
	"daijoubuteam.xyz/se214-library-management/api/middleware"
	"daijoubuteam.xyz/se214-library-management/wireimpl"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"net/http"
)

func ServerCommand(db *sqlx.DB) *cobra.Command {
	command := &cobra.Command{
		Use:   `server`,
		Short: `Server administrate`,
	}
	command.AddCommand(StartServerCommand(db))
	return command
}

func StartServerCommand(db *sqlx.DB) *cobra.Command {
	command := &cobra.Command{
		Use:   `start`,
		Short: `Start server`,
		Run: func(cmd *cobra.Command, args []string) {
			StartServer(db)
		},
	}
	return command
}

func StartServer(db *sqlx.DB) {

	authUsecase := wireimpl.InitAuthUsecase(db)
	loaiDocGiaUsecase := wireimpl.InitLoaiDocGiaUsecase(db)
	docGiaUsecase := wireimpl.InitDocGiaUsecase(db)
	theLoaiUsecase := wireimpl.InitTheLoaiUsecase(db)
	tacGiaUsecase := wireimpl.InitTacGiaUsecase(db)
	dauSachUsecase := wireimpl.InitDauSachUsecase(db)
	nhapSachUsecase := wireimpl.InitNhapSachUsecase(db)

	r := gin.Default()

	r.Use(middleware.Cors())

	r.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "OK",
		})
	})

	handler.MakeAuthHandler(r, authUsecase)
	handler.MakeLoaiThuThuHandler(r, loaiDocGiaUsecase)
	handler.MakeDocGiaHandler(r, docGiaUsecase)
	handler.MakeTheLoaiHandler(r, theLoaiUsecase)
	handler.MakeTacGiaHandler(r, tacGiaUsecase)
	handler.MakeDauSachHandler(r, dauSachUsecase)
	handler.MakePhieuNhapHandler(r, nhapSachUsecase)
	r.Run(":8080")
}
