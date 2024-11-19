package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdManage() *cobra.Command {
	return &cobra.Command{
		Use:   "manage",
		Short: "Manage packages and their version commands",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", "hi it a manage")
		},
	}
}
