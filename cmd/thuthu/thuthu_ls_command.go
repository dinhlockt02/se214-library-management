package thuthucommand

import (
	thuthucommandls "daijoubuteam.xyz/se214-library-management/cmd/thuthu/ls"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func LsThuThuCommand(db *sqlx.DB) *cobra.Command {
	
	command := &cobra.Command{
		Use:   `ls`,
		Short: `List thu thu`,
		Run: func(cmd *cobra.Command, args []string) {
			thuthucommandls.ListThuThu(db)
		},
	}

	return command
}
