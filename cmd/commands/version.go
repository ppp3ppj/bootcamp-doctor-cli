package commands

import (
	"fmt"

	"github.com/ppp3ppj/bootcamp-doctor-cli/pkg/bootcamp/version"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
    return &cobra.Command{
        Use: "version",
        Short: fmt.Sprintf(
			"Shows the bootcamp doctor CLI version (saving time, it's: %s)",
			version.Print(),
		),
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println(version.Print())
        },
    }
}
