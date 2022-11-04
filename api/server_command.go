package api

import (
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

	r := gin.Default()

	r.Run(":8080")
}
