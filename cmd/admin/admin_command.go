package admin

import (
	admincreate "daijoubuteam.xyz/se214-library-management/cmd/admin/create"
	"github.com/spf13/cobra"
)

func AdminCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   `admin`,
		Short: `Modify server using admin rights`,
	}
	command.AddCommand(admincreate.CreateAdminCommand())
	return command
}
