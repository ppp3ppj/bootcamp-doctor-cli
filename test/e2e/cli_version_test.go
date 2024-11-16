package e2e

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func TestCLIVersionE2E(t *testing.T) {
	// Set the correct working directory
	cmd := exec.Command("go", "run", "cmd/main.go", "cliversion")
	cmd.Dir = "../.." // Adjust path if necessary to point to your project root

	// Capture output
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output

	// Run the command
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Error running command: %v\nOutput: %s", err, output.String())
	}

	// Validate output using regex to match version and build format
	versionPattern := `CLI Version: \d+\.\d+\.\d+` // Matches something like "CLI Version: 0.0.1"
	buildPattern := `Build: [a-zA-Z0-9]+`        // Matches something like "Build: dev"

	// Compile regex
	versionRegex := regexp.MustCompile(versionPattern)
	buildRegex := regexp.MustCompile(buildPattern)

	// Check if version matches expected format
	if !versionRegex.MatchString(output.String()) {
		t.Errorf("Unexpected version output: %s", output.String())
	}

	// Check if build matches expected format
	if !buildRegex.MatchString(output.String()) {
		t.Errorf("Unexpected build output: %s", output.String())
	}

	// Optionally check if version and build appear in the output together
	if !strings.Contains(output.String(), "CLI Version") || !strings.Contains(output.String(), "Build") {
		t.Errorf("Version and build info not found in output: %s", output.String())
	}
}

// Helper function to check if a string contains a substring
func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}
