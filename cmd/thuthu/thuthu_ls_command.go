package thuthucommand

import (
	thuthucommandls "daijoubuteam.xyz/se214-library-management/cmd/thuthu/ls"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func LsThuThuCommand(db *sqlx.DB) *cobra.Command {

	var email string
	var phone string

	command := &cobra.Command{
		Use:   `ls`,
		Short: `List thu thu`,
		Run: func(cmd *cobra.Command, args []string) {
			thuthucommandls.ListThuThu(db, email, phone)
		},
	}
	command.PersistentFlags().StringVarP(&email, "email", "e", "", "Filter by email")
	command.PersistentFlags().StringVarP(&phone, "phone", "p", "", "Filter by phone")

	return command
}
