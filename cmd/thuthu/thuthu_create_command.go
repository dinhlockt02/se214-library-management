package thuthucommand

import (
	"daijoubuteam.xyz/se214-library-management/cmd/thuthu/create"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func CreateThuThuCommand(db *sqlx.DB) *cobra.Command {

	var name string
	var birth string
	var email string
	var PhoneNumber string
	var status bool
	var password string

	command := &cobra.Command{
		Use:   `create`,
		Short: `Create new thu thu`,
		Run: func(cmd *cobra.Command, args []string) {
			create.CreateThuThu(db, name, birth, email, PhoneNumber, status, password)
		},
	}

	command.PersistentFlags().StringVarP(&name, "name", "n", "", "Admin's name")
	command.PersistentFlags().StringVarP(&birth, "birth", "b", "2022-12-31", "Admin's birthday")
	command.PersistentFlags().StringVar(&PhoneNumber, "phone", "", "Admin's phone number")
	command.PersistentFlags().BoolVar(&status, "enable", false, "Admin's enable status")
	command.PersistentFlags().StringVarP(&email, "email", "e", "", "Admin's email")
	command.MarkPersistentFlagRequired("email")
	command.PersistentFlags().StringVarP(&password, "password", "p", "", "Admin's password")
	command.MarkPersistentFlagRequired("password")
	return command
}
