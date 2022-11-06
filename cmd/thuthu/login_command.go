package thuthucommand

import (
	thuthucommandlogin "daijoubuteam.xyz/se214-library-management/cmd/thuthu/login"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func LoginCommand(db *sqlx.DB) *cobra.Command {
	var email string
	var password string
	command := &cobra.Command{
		Use:   `login`,
		Short: `Login and get token`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(email) > 0 && len(password) > 0 {
				thuthucommandlogin.Login(db, email, password)
			}
		},
	}
	command.PersistentFlags().StringVarP(&email, "email", "e", "", "Thu thu's email")
	command.PersistentFlags().StringVarP(&password, "password", "p", "", "Thu thu's password")
	command.MarkFlagsRequiredTogether("password", "email")
	return command
}
