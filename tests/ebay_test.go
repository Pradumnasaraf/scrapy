package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestEbayCmd tests the ebay command
func TestEbayCmd(t *testing.T) {

	expectedOutput := "Searching for: iphone"
	cmd := exec.Command("scrapy", "ebay", "iphone")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:22])
	if got != strings.Trim(expectedOutput, " ") {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// TestEbayCmdPagesFlag tests the ebay command with the pages flag
func TestEbayCmdPagesFlag(t *testing.T) {

	expectedOutput := "Searching for: macbook"
	stringContains := "Scraping page 1 of 2"
	cmd := exec.Command("scrapy", "ebay", "macbook", "--pages", "2")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the 1st line of the cli output
	got := strings.TrimSpace(string(output)[:22])
	if got != strings.Trim(expectedOutput, " ") {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Validate the 2nd line of the cli output
	if strings.Contains(string(output), stringContains) != true {
		t.Errorf("expected %v, but got: %v", stringContains, string(output))
	}

}
