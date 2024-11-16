package commands

import (
	"bytes"
	"fmt"
	"testing"

	application "github.com/ppp3ppj/bootcamp-doctor-cli/application/cli-version"
	domain "github.com/ppp3ppj/bootcamp-doctor-cli/domain/cli-version"
	"github.com/spf13/cobra"
)

func NewCmdCLIVersion(useCase *application.CLIVersionUseCase) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Displays the current CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			cliVersion := useCase.GetCurrentVersion()
			fmt.Printf("CLI Version: %s\nBuild: %s\n", cliVersion.Version, cliVersion.Build)
		},
	}
}

func TestVersionCmd(t *testing.T) {
	version := domain.NewCLIVersion("1.0.0", "build1234")
	useCase := &application.CLIVersionUseCase{Version: version}
	command := NewCmdCLIVersion(useCase)

	var output bytes.Buffer
	command.SetOut(&output)

	err := command.Execute()

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if !contains(output.String(), "CLI Version: 1.0.0") {
		t.Errorf("Expected 'CLI Version: 1.0.0', but got %s", output.String())
	}

	if !contains(output.String(), "Build: build1234") {
		t.Errorf("Expected 'Build: build1234', but got %s", output.String())
	}
}

func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}
