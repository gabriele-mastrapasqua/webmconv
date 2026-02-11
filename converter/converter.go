package converter

import (
	"fmt"
	"os/exec"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Quality represents the quality level for conversion
type Quality string

const (
	QualityMax   Quality = "max"
	QualityMedium Quality = "medium"
	QualityLow   Quality = "low"
)

// ConvertToWebM converts a video/gif file to WebM format using FFmpeg
func ConvertToWebM(inputPath string, outputPath string, quality Quality, rangeOption string) error {
	// Check if output path already contains .webm extension
	if filepath.Ext(outputPath) != ".webm" {
		outputPath = filepath.Join(filepath.Dir(outputPath), filepath.Base(inputPath)+".webm")
	}

	// Set quality parameters
	crfValue := "30" // Default value for medium quality
	switch quality {
	case QualityMax:
		crfValue = "15" // High quality (low CRF)
	case QualityLow:
		crfValue = "45" // Low quality (high CRF)
	}

	// Build the FFmpeg command with basic parameters
	args := []string{"-i", inputPath}

	// Add range parameters if specified
	if rangeOption != "" {
		parts := strings.Split(rangeOption, "-")
		if len(parts) == 2 {
			startTime, err := parseTimeString(parts[0])
			if err != nil {
				return fmt.Errorf("invalid start time format: %v", err)
			}

			endTime, err := parseTimeString(parts[1])
			if err != nil {
				return fmt.Errorf("invalid end time format: %v", err)
			}

			args = append(args, "-ss", fmt.Sprintf("%.6f", startTime), "-to", fmt.Sprintf("%.6f", endTime))
		}
	}

	// Add encoding parameters
	args = append(args, "-c:v", "libvpx-vp9", "-crf", crfValue, "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", outputPath)

	// Build the FFmpeg command
	cmd := exec.Command("ffmpeg", args...)

	// Execute the command and check for errors
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting %s: %v", inputPath, err)
	}

	return nil
}

// CheckFFmpeg checks if FFmpeg is available in the system
func CheckFFmpeg() error {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("ffmpeg not found in the system: %v", err)
	}
	return nil
}

// parseTimeString converts a time string in format "MM:SS" or "HH:MM:SS" to seconds
func parseTimeString(timeStr string) (float64, error) {
	timeStr = strings.TrimSpace(timeStr)
	
	// Handle format with seconds suffix (e.g., "10s")
	if strings.HasSuffix(timeStr, "s") {
		secondsStr := strings.TrimSuffix(timeStr, "s")
		return strconv.ParseFloat(secondsStr, 64)
	}
	
	parts := strings.Split(timeStr, ":")
	
	switch len(parts) {
	case 1: // Just seconds (e.g., "30")
		return strconv.ParseFloat(parts[0], 64)
	case 2: // MM:SS format (e.g., "1:30")
		minutes, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, err
		}
		seconds, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, err
		}
		return minutes*60 + seconds, nil
	case 3: // HH:MM:SS format (e.g., "1:10:30")
		hours, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, err
		}
		minutes, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return 0, err
		}
		seconds, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			return 0, err
		}
		return hours*3600 + minutes*60 + seconds, nil
	default:
		return 0, fmt.Errorf("invalid time format: %s", timeStr)
	}
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}