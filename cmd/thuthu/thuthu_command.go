package thuthucommand

import (
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func ThuThuCommand(db *sqlx.DB) *cobra.Command {
	command := &cobra.Command{
		Use:   `thuthu`,
		Short: `Modify thu thu using admin rights`,
	}
	command.AddCommand(LsThuThuCommand(db))
	command.AddCommand(GetThuThuCommand(db))
	command.AddCommand(LoginCommand(db))
	return command
}
