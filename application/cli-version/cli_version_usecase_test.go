package application

import (
	"testing"

	domain "github.com/ppp3ppj/bootcamp-doctor-cli/domain/cli-version"
)

func TestVersionUseCase_GetCurrentVersion(t *testing.T) {
	version := domain.NewCLIVersion("1.0.0", "build1234")
	useCase := &CLIVersionUseCase{Version: version}

	result := useCase.GetCurrentVersion()

	// Test if the values are correct
	if result.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', but got '%s'", result.Version)
	}
	if result.Build != "build1234" {
		t.Errorf("Expected build 'build1234', but got '%s'", result.Build)
	}
}
