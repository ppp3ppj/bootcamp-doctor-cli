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
        Use: "bootcamp",
        Long: letters,
        CompletionOptions: cobra.CompletionOptions{
            DisableDefaultCmd: true,
        },
    }

    // Register Top Level Commands
    //rootCmd.AddCommand()
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


