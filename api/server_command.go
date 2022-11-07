package api

import (
	"daijoubuteam.xyz/se214-library-management/api/handler"
	"daijoubuteam.xyz/se214-library-management/wireimpl"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
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

	r := gin.Default()

	r.Static("/api-docs", "./docs/dist")

	handler.MakeAuthHandler(r, authUsecase)

	r.Run(":80")
}
