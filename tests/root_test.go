package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestRootCmd tests the root command (scrapy)
func TestRootCmd(t *testing.T) {

	expectedOutput := "Scrape the web from the command line"
	cmd := exec.Command("scrapy")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:36])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// TestRootCmdHelpFlag tests the root command (scrapy) with the help flag
func TestRootCmdHelpFlag(t *testing.T) {

	expectedOutput := "Scrape the web from the command line"
	cmd := exec.Command("scrapy", "--help")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:36])
	if got != expectedOutput {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}
