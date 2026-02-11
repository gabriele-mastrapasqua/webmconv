package converter

import (
	"fmt"
	"os/exec"
	"os"
	"path/filepath"
	"strings"
)

// Quality rappresenta il livello di qualità per la conversione
type Quality string

const (
	QualityMax   Quality = "max"
	QualityMedium Quality = "medium"
	QualityLow   Quality = "low"
)

// ConvertToWebM converte un file video/gif in formato WebM usando FFmpeg
func ConvertToWebM(inputPath string, outputPath string, quality Quality, rangeOption string) error {
	// Controlla se il percorso di output contiene già l'estensione .webm
	if filepath.Ext(outputPath) != ".webm" {
		outputPath = filepath.Join(filepath.Dir(outputPath), filepath.Base(inputPath)+".webm")
	}

	// Imposta i parametri di qualità
	crfValue := "30" // Valore predefinito per qualità media
	switch quality {
	case QualityMax:
		crfValue = "15" // Qualità molto alta (CRF basso)
	case QualityLow:
		crfValue = "45" // Qualità bassa (CRF alto)
	}

	// Costruisci il comando FFmpeg con i parametri base
	args := []string{"-i", inputPath}

	// Aggiungi i parametri per il range se specificato
	if rangeOption != "" {
		parts := strings.Split(rangeOption, "-")
		if len(parts) == 2 {
			args = append(args, "-ss", parts[0], "-to", parts[1])
		}
	}

	// Aggiungi i parametri di codifica
	args = append(args, "-c:v", "libvpx-vp9", "-crf", crfValue, "-b:v", "0", "-b:a", "128k", "-c:a", "libopus", outputPath)

	// Costruisci il comando FFmpeg
	cmd := exec.Command("ffmpeg", args...)

	// Esegui il comando e controlla per eventuali errori
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("errore durante la conversione di %s: %v", inputPath, err)
	}

	return nil
}

// CheckFFmpeg verifica se FFmpeg è disponibile nel sistema
func CheckFFmpeg() error {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("ffmpeg non trovato nel sistema: %v", err)
	}
	return nil
}

// FileExists controlla se un file esiste
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}