package test

import (
	"os/exec"
	"strings"
	"testing"
)

// TestTcCmd tests the tc command
func TestTcCmd(t *testing.T) {

	expectedOutput := "Searching for latest blogs on techcrunch...."
	cmd := exec.Command("scrapy", "tc")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the cli output
	got := strings.TrimSpace(string(output)[:44])
	if got != strings.Trim(expectedOutput, " ") {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// TestTcCmdWithCategoryFlag tests the tc command with the category flag

func TestTcCmdWithCategoryFlag(t *testing.T) {

	expectedOutput := "Searching for latest blogs on techcrunch...."
	cmd := exec.Command("scrapy", "tc", "--category", "startups")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the 1st line of the cli output
	got := strings.TrimSpace(string(output)[:44])
	if got != strings.Trim(expectedOutput, " ") {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

}

// // TestTcCmdWithInvalidCategoryFlag tests the tc command with the category flag
func TestTcCmdWithInvalidCategoryFlag(t *testing.T) {

	expectedOutput := "Searching for latest blogs on techcrunch...."
	stringsContains := "Invalid category. Showing homepage instead."
	cmd := exec.Command("scrapy", "tc", "--category", "football")

	// Capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("expected no error, but got: %v", err)
	}

	// Validate the 1st line of the cli output
	got := strings.TrimSpace(string(output)[:44])
	if got != strings.Trim(expectedOutput, " ") {
		t.Errorf("expected %v, but got: %v", expectedOutput, got)
	}

	// Validate the 2nd line of the cli output
	if strings.Contains(string(output), stringsContains) != true {
		t.Errorf("expected %v, but got: %v", stringsContains, string(output))
	}

}
