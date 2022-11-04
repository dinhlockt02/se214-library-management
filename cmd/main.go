package main

import (
	"daijoubuteam.xyz/se214-library-management/api"
	"daijoubuteam.xyz/se214-library-management/cmd/admin"
	"daijoubuteam.xyz/se214-library-management/config"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/spf13/cobra"
)

func main() {

	// Connect to database
	db := utils.ConnectDB(config.DevConfig)

	rootCommand := CreateCommand()
	rootCommand.AddCommand(admin.AdminCommand())
	rootCommand.AddCommand(api.ServerCommand(db))
	rootCommand.Execute()
}

func CreateCommand() *cobra.Command {

	command := &cobra.Command{
		Use:   `se214`,
		Short: `se214 is an console application`,
		Long:  `se214 is an console application for administate server`,
	}

	return command
}
