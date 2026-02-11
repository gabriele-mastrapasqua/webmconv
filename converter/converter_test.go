package converter

import (
	"path/filepath"
	"testing"
)

func TestConvertToWebM(t *testing.T) {
	// Check that FFmpeg is available
	if err := CheckFFmpeg(); err != nil {
		t.Skipf("ffmpeg not found: %v", err)
	}

	// Path to sample input file
	inputPath := "./tiny_test_video.mp4" // File in the same directory as the test

	// Check if input file exists
	if !FileExists(inputPath) {
		t.Fatalf("Input file not found: %s", inputPath)
	}

	// Name of output file
	outputPath := filepath.Join(t.TempDir(), "output_test.webm")

	// Call the conversion function
	err := ConvertToWebM(inputPath, outputPath, QualityMedium, "")

	// Verify that there were no errors
	if err != nil {
		t.Errorf("Error during conversion: %v", err)
	}

	// Check if output file was created
	if !FileExists(outputPath) {
		t.Errorf("Output file was not created: %s", outputPath)
	}

	// Verify that output file has .webm extension
	if filepath.Ext(outputPath) != ".webm" {
		t.Errorf("Output file does not have .webm extension: %s", outputPath)
	}
}

func TestConvertToWebMWithRange(t *testing.T) {
	// Check that FFmpeg is available
	if err := CheckFFmpeg(); err != nil {
		t.Skipf("ffmpeg not found: %v", err)
	}

	// Path to sample input file
	inputPath := "./tiny_test_video.mp4" // File in the same directory as the test

	// Check if input file exists
	if !FileExists(inputPath) {
		t.Skipf("Input file not found: %s, skipping test", inputPath)
	}

	// Name of output file
	outputPath := filepath.Join(t.TempDir(), "output_test_with_range.webm")

	// Call the conversion function with range
	err := ConvertToWebM(inputPath, outputPath, QualityMedium, "0-5s") // Convert first 5 seconds

	// Verify that there were no errors
	if err != nil {
		t.Errorf("Error during conversion with range: %v", err)
	}

	// Check if output file was created
	if !FileExists(outputPath) {
		t.Errorf("Output file with range was not created: %s", outputPath)
	}

	// Verify that output file has .webm extension
	if filepath.Ext(outputPath) != ".webm" {
		t.Errorf("Output file with range does not have .webm extension: %s", outputPath)
	}
}

func TestParseTimeString(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"30", 30.0, false},
		{"30s", 30.0, false},
		{"1:30", 90.0, false},        // 1 minute 30 seconds = 90 seconds
		{"2:45", 165.0, false},       // 2 minutes 45 seconds = 165 seconds
		{"1:10:30", 4230.0, false},   // 1 hour 10 minutes 30 seconds = 4230 seconds
		{"0:30", 30.0, false},        // 30 seconds
		{"0:00:15", 15.0, false},     // 15 seconds
		{"invalid", 0.0, true},       // Invalid format
		{"1:2:3:4", 0.0, true},       // Too many colons
	}

	for _, tt := range tests {
		result, err := parseTimeString(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("parseTimeString(%q) expected error, got none", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("parseTimeString(%q) unexpected error: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("parseTimeString(%q) = %f, want %f", tt.input, result, tt.expected)
			}
		}
	}
}