package api

import (
	"daijoubuteam.xyz/se214-library-management/api/handler"
	"daijoubuteam.xyz/se214-library-management/api/middleware"
	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/wireimpl"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"math/rand"
	"net/http"
	"os"
	"time"
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
	rand.Seed(time.Now().UnixNano())

	authUsecase := wireimpl.InitAuthUsecase(db)
	loaiDocGiaUsecase := wireimpl.InitLoaiDocGiaUsecase(db)
	docGiaUsecase := wireimpl.InitDocGiaUsecase(db)
	theLoaiUsecase := wireimpl.InitTheLoaiUsecase(db)
	tacGiaUsecase := wireimpl.InitTacGiaUsecase(db)
	dauSachUsecase := wireimpl.InitDauSachUsecase(db)
	nhapSachUsecase := wireimpl.InitNhapSachUsecase(db)
	sachUsecase := wireimpl.InitSachUsecase(db)
	phieuMuonUsecase := wireimpl.InitPhieuMuonUsecase(db)
	phieuTraUsecase := wireimpl.InitPhieuTraUsecase(db)
	phieuThuTienUsecase := wireimpl.InitPhieuThuTienUsecase(db)
	reportUsecase := wireimpl.InitReportUsecase(db)

	r := gin.Default()

	if config.GetConfig().Mode == config.Release {
		gin.DisableConsoleColor()

		f, _ := os.Create("backend.log")
		ef, _ := os.Create("error.log")

		gin.DefaultWriter = f
		gin.DefaultErrorWriter = ef
	}

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
	handler.MakeSachHandler(r, sachUsecase)
	handler.MakePhieuMuonHandler(r, phieuMuonUsecase)
	handler.MakePhieuTraHandler(r, phieuTraUsecase)
	handler.MakePhieuThuTienHandler(r, phieuThuTienUsecase)
	handler.MakeReportHandler(r, reportUsecase)
	r.Run(":8080")
}
