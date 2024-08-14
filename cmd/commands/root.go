package commands

import (
	"fmt"
	"os"
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
    }

    // Register Top Level Commands
    //rootCmd.AddCommand()
    rootCmd.AddCommand(NewCmdVersion())
    rootCmd.AddCommand(NewCmdDoctor())
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


