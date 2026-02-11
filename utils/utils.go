package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// GetSupportedFiles trova tutti i file video e gif in una directory in modo ricorsivo
func GetSupportedFiles(rootPath string) ([]string, error) {
	var files []string
	supportedExtensions := map[string]bool{
		".mp4":  true,
		".avi":  true,
		".mov":  true,
		".mkv":  true,
		".wmv":  true,
		".flv":  true,
		".gif":  true,
		".m4v":  true,
		".3gp":  true,
		".webm": true, // Potrebbe voler riconvertire anche webm
	}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if supportedExtensions[ext] {
				files = append(files, path)
			}
		}
		return nil
	})

	return files, err
}