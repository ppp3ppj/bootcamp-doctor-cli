package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ppp3ppj/bootcamp-doctor-cli/internal/tui"
	"github.com/spf13/cobra"
)

func NewCmdDoctor() *cobra.Command {
    return &cobra.Command{
        Use: "doctor",
        Short: "doctor run check package version",
        RunE: func(cmd *cobra.Command, args []string) error {
            pg := tui.NewModel()
            if err := tea.NewProgram(pg).Start(); err != nil {
                return err
            }

            return nil
        },
    }
}
