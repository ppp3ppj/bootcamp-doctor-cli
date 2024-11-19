package commands

import (
	"fmt"
	"os"

	application "github.com/ppp3ppj/bootcamp-doctor-cli/application/cli-version"
	domain "github.com/ppp3ppj/bootcamp-doctor-cli/domain/cli-version"
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
		Use:  "bc",
		Long: letters,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cliVersion := domain.NewCLIVersion("0.0.1", "dev")
	cliVersionUseCase := &application.CLIVersionUseCase{cliVersion}



	// Register Top Level Commands
	//rootCmd.AddCommand()
	rootCmd.AddCommand(NewCmdDoctor())
	rootCmd.AddCommand(NewCmdCLIVersion(cliVersionUseCase))
    rootCmd.AddCommand(NewCmdManage())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
