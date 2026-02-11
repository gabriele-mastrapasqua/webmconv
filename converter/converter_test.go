package converter

import (
	"path/filepath"
	"testing"
)

func TestConvertToWebM(t *testing.T) {
	// Verifica che FFmpeg sia disponibile
	if err := CheckFFmpeg(); err != nil {
		t.Skipf("ffmpeg non trovato: %v", err)
	}

	// Percorso del file di input di esempio
	inputPath := "./tiny_test_video.mp4" // File nella stessa directory del test

	// Controlla se il file di input esiste
	if !FileExists(inputPath) {
		t.Fatalf("File di input non trovato: %s", inputPath)
	}

	// Nome del file di output
	outputPath := filepath.Join(t.TempDir(), "output_test.webm")

	// Chiama la funzione di conversione
	err := ConvertToWebM(inputPath, outputPath, QualityMedium, "")

	// Verifica che non ci siano stati errori
	if err != nil {
		t.Errorf("Errore durante la conversione: %v", err)
	}

	// Controlla se il file di output è stato creato
	if !FileExists(outputPath) {
		t.Errorf("Il file di output non è stato creato: %s", outputPath)
	}

	// Verifica che il file di output abbia estensione .webm
	if filepath.Ext(outputPath) != ".webm" {
		t.Errorf("Il file di output non ha estensione .webm: %s", outputPath)
	}
}

func TestConvertToWebMWithRange(t *testing.T) {
	// Verifica che FFmpeg sia disponibile
	if err := CheckFFmpeg(); err != nil {
		t.Skipf("ffmpeg non trovato: %v", err)
	}

	// Percorso del file di input di esempio
	inputPath := "./tiny_test_video.mp4" // File nella stessa directory del test

	// Controlla se il file di input esiste
	if !FileExists(inputPath) {
		t.Skipf("File di input non trovato: %s, skipping test", inputPath)
	}

	// Nome del file di output
	outputPath := filepath.Join(t.TempDir(), "output_test_with_range.webm")

	// Chiama la funzione di conversione con range
	err := ConvertToWebM(inputPath, outputPath, QualityMedium, "0-5s") // Convert first 5 seconds

	// Verifica che non ci siano stati errori
	if err != nil {
		t.Errorf("Errore durante la conversione con range: %v", err)
	}

	// Controlla se il file di output è stato creato
	if !FileExists(outputPath) {
		t.Errorf("Il file di output con range non è stato creato: %s", outputPath)
	}

	// Verifica che il file di output abbia estensione .webm
	if filepath.Ext(outputPath) != ".webm" {
		t.Errorf("Il file di output con range non ha estensione .webm: %s", outputPath)
	}
}