package commands

import (
	"fmt"

	application "github.com/ppp3ppj/bootcamp-doctor-cli/application/cli-version"
	"github.com/spf13/cobra"
)

func NewCmdCLIVersion(useCase *application.CLIVersionUseCase) *cobra.Command {
    return &cobra.Command{
        Use: "cliversion",
        Short: "Displays the current CLI version",
        Run: func(cmd *cobra.Command, args []string) {
            cliVersion := useCase.GetCurrentVersion()
			fmt.Printf("CLI Version: %s\nBuild: %s\n", cliVersion.Version, cliVersion.Build)
        },
    }
}
