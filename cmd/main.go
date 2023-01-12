package main

import (
	"daijoubuteam.xyz/se214-library-management/api"
	"daijoubuteam.xyz/se214-library-management/cmd/admin"
	thuthucommand "daijoubuteam.xyz/se214-library-management/cmd/thuthu"
	"daijoubuteam.xyz/se214-library-management/config"
	_ "daijoubuteam.xyz/se214-library-management/config/env_config"
	"daijoubuteam.xyz/se214-library-management/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {

	// Connect to database
	db := utils.ConnectDB(config.GetConfig())

	rootCommand := CreateCommand()
	rootCommand.AddCommand(admin.AdminCommand(db))
	rootCommand.AddCommand(api.ServerCommand(db))
	rootCommand.AddCommand(thuthucommand.ThuThuCommand(db))
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func CreateCommand() *cobra.Command {

	command := &cobra.Command{
		Use:   `se214`,
		Short: `se214 is an console application`,
		Long:  `se214 is an console application for administate server`,
	}

	return command
}
