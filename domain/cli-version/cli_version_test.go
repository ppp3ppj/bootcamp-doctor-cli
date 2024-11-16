package domain

import "testing"

func TestCLIVersionCreation(t *testing.T) {
	version := NewCLIVersion("1.0.0", "build1234")

	// Test if the values are set correctly
	if version.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', but got '%s'", version.Version)
	}
	if version.Build != "build1234" {
		t.Errorf("Expected build 'build1234', but got '%s'", version.Build)
	}
}
