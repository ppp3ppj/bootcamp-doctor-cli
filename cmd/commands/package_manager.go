package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	tui "github.com/ppp3ppj/bootcamp-doctor-cli/internal/tui/package"
	"github.com/spf13/cobra"
)

func NewCmdManage() *cobra.Command {
	return &cobra.Command{
		Use:   "manage",
		Short: "Manage packages and their version commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			pg := tui.NewPackageManagerModel()
			if err := tea.NewProgram(pg).Start(); err != nil {
				return err
			}

			return nil
		},

	}
}
