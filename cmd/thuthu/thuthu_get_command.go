package thuthucommand

import (
	thuthucommandget "daijoubuteam.xyz/se214-library-management/cmd/thuthu/get"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func GetThuThuCommand(db *sqlx.DB) *cobra.Command {

	command := &cobra.Command{
		Use:   `get`,
		Short: `Get thu thu`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Only 1 argument required")
				return
			}
			thuthucommandget.GetThuThu(db, args[0])
		},
	}

	return command
}
