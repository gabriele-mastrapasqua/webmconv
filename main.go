package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"sync/atomic"

	"webmconv/converter"
	"webmconv/utils"
)

func main() {
	// Define flags for source and destination directories
	sourceDir := flag.String("source", "", "Directory containing the files to convert")
	destDir := flag.String("dest", "", "Directory to save the converted files (optional, otherwise uses the same directory)")
	qualityStr := flag.String("quality", "medium", "Quality level for conversion: max, medium, low")
	rangeOpt := flag.String("range", "", "Time range for conversion in format start-end (e.g., 0-100s, 10-50s)")
	help := flag.Bool("help", false, "Show this help message")
	flag.Parse()

	// Check if help was requested
	if *help {
		showHelp()
		os.Exit(0)
	}

	// Parse quality level
	var quality converter.Quality
	switch *qualityStr {
	case "max":
		quality = converter.QualityMax
	case "low":
		quality = converter.QualityLow
	default:
		quality = converter.QualityMedium
	}

	// Check if FFmpeg is available
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal("FFmpeg not found in the system. Please ensure it is installed and in your PATH.")
	}

	// Check if source directory was provided
	if *sourceDir == "" {
		log.Fatal("Source directory must be specified with the -source flag")
	}

	// Check if source directory exists
	if _, err := os.Stat(*sourceDir); os.IsNotExist(err) {
		log.Fatalf("Source directory %s does not exist", *sourceDir)
	}

	// If destination directory is not specified, use the source directory
	if *destDir == "" {
		*destDir = *sourceDir
	}

	// Get all supported files from the source directory
	files, err := utils.GetSupportedFiles(*sourceDir)
	if err != nil {
		log.Fatalf("Error scanning directory: %v", err)
	}

	// Counter to track progress
	totalFiles := len(files)
	convertedCount := int64(0) // Used with atomic for thread safety

	// Number of goroutines for concurrent conversion
	numWorkers := 4 // Can be made configurable via flag
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, numWorkers) // Limits the number of concurrent conversions

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			
			// Acquire a slot in the semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release when done

			fmt.Printf("Converting %s...", f)

			// Calculate destination path
			relPath, _ := filepath.Rel(*sourceDir, f)
			destPath := filepath.Join(*destDir, relPath)

			// Create destination directory if it doesn't exist
			destDirPath := filepath.Dir(destPath)
			if err := os.MkdirAll(destDirPath, 0755); err != nil {
				log.Printf("Could not create destination directory %s: %v", destDirPath, err)
				return
			}

			// Convert the file
			if err := converter.ConvertToWebM(f, destPath, quality, *rangeOpt); err != nil {
				log.Printf(" Error: %v\n", err)
			} else {
				fmt.Println(" Completed.")
				// Thread-safe increment of the counter
				atomic.AddInt64(&convertedCount, 1)
			}
		}(file)
	}

	wg.Wait()

	fmt.Printf("\nConversion finished. %d/%d files converted successfully.\n", convertedCount, totalFiles)
}

// showHelp displays a help message with usage instructions
func showHelp() {
	fmt.Println("webmconv - A tool to convert video and GIF files to WebM format")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  webmconv -source <source_directory> [-dest <destination_directory>] [-quality <quality_level>] [-range <time_range>]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -source    Directory containing the files to convert (required)")
	fmt.Println("  -dest      Directory to save the converted files (optional, otherwise uses the same directory)")
	fmt.Println("  -quality   Quality level for conversion: max, medium, low (default: medium)")
	fmt.Println("  -range     Time range for conversion in format start-end (e.g., 0-100s, 10-50s)")
	fmt.Println("  -help      Show this help message")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  webmconv -source /path/to/videos -dest /path/to/output -quality max")
	fmt.Println("  webmconv -source /path/to/videos -quality low -range 0-30s")
	fmt.Println("")
	fmt.Println("Note: Ensure FFmpeg is installed and in your system PATH.")
}