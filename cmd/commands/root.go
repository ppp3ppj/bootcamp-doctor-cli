package commands

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ppp3ppj/bootcamp-doctor-cli/internal/tui"
	"github.com/spf13/cobra"
)

const (
    letters = `
▄▄▄▄·            ▄▄▄▄▄ ▄▄·  ▄▄▄· • ▌ ▄ ·.  ▄▄▄·    ·▄▄▄▄        ▄▄· ▄▄▄▄▄      ▄▄▄
▐█ ▀█▪▪     ▪    •██  ▐█ ▌▪▐█ ▀█ ·██ ▐███▪▐█ ▄█    ██▪ ██▪     ▐█ ▌▪•██  ▪     ▀▄ █·
▐█▀▀█▄ ▄█▀▄  ▄█▀▄ ▐█.▪██ ▄▄▄█▀▀█ ▐█ ▌▐▌▐█· ██▀·    ▐█· ▐█▌▄█▀▄ ██ ▄▄ ▐█.▪ ▄█▀▄ ▐▀▀▄
██▄▪▐█▐█▌.▐▌▐█▌.▐▌▐█▌·▐███▌▐█ ▪▐▌██ ██▌▐█▌▐█▪·•    ██. ██▐█▌.▐▌▐███▌ ▐█▌·▐█▌.▐▌▐█•█▌
·▀▀▀▀  ▀█▄▀▪ ▀█▄▀▪▀▀▀ ·▀▀▀  ▀  ▀ ▀▀  █▪▀▀▀.▀       ▀▀▀▀▀• ▀█▄▀▪·▀▀▀  ▀▀▀  ▀█▄▀▪.▀  ▀
`
)

func Execute() {
    rootCmd := &cobra.Command{
        Use: "bc",
        Long: letters,
        CompletionOptions: cobra.CompletionOptions{
            DisableDefaultCmd: true,
        },
        RunE: func(cmd *cobra.Command, args[] string) error {
            pg := tui.NewModel()
            if err := tea.NewProgram(pg).Start(); err != nil {
                return err
            }
            return nil
        },
    }

    // Register Top Level Commands
    //rootCmd.AddCommand()
    rootCmd.AddCommand(NewCmdVersion())
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


