package thuthucommand

import "github.com/spf13/cobra"

func CreateGetThuThuCommand() *cobra.Command {

	command := &cobra.Command{
		Use:   `get`,
		Short: `Get thu thu`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return command
}
