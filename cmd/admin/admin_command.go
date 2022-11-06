package admin

import (
	admincreate "daijoubuteam.xyz/se214-library-management/cmd/admin/create"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func AdminCommand(db *sqlx.DB) *cobra.Command {
	command := &cobra.Command{
		Use:   `admin`,
		Short: `Modify server using admin rights`,
	}
	command.AddCommand(admincreate.CreateAdminCommand(db))
	return command
}
