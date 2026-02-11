package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"webmconv/converter"
	"webmconv/utils"
)

func main() {
	// Definisci i flag per la directory sorgente e quella di destinazione
	sourceDir := flag.String("source", "", "Directory contenente i file da convertire")
	destDir := flag.String("dest", "", "Directory dove salvare i file convertiti (opzionale, altrimenti usa la stessa directory)")
	flag.Parse()

	// Controlla se FFmpeg è disponibile
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		log.Fatal("FFmpeg non trovato nel sistema. Assicurati che sia installato e presente nel PATH.")
	}

	// Controlla se la directory sorgente è stata fornita
	if *sourceDir == "" {
		log.Fatal("La directory sorgente deve essere specificata con il flag -source")
	}

	// Controlla se la directory sorgente esiste
	if _, err := os.Stat(*sourceDir); os.IsNotExist(err) {
		log.Fatalf("La directory sorgente %s non esiste", *sourceDir)
	}

	// Se la directory di destinazione non è specificata, usa la directory sorgente
	if *destDir == "" {
		*destDir = *sourceDir
	}

	// Ottieni tutti i file supportati dalla directory sorgente
	files, err := utils.GetSupportedFiles(*sourceDir)
	if err != nil {
		log.Fatalf("Errore durante la scansione della directory: %v", err)
	}

	// Contatore per tenere traccia del progresso
	totalFiles := len(files)
	convertedCount := 0

	for _, file := range files {
		fmt.Printf("Conversione di %s in corso...", file)

		// Calcola il percorso di destinazione
		relPath, _ := filepath.Rel(*sourceDir, file)
		destPath := filepath.Join(*destDir, relPath)

		// Crea la directory di destinazione se non esiste
		destDirPath := filepath.Dir(destPath)
		if err := os.MkdirAll(destDirPath, 0755); err != nil {
			log.Printf("Impossibile creare la directory di destinazione %s: %v", destDirPath, err)
			continue
		}

		// Converti il file
		if err := converter.ConvertToWebM(file, destPath); err != nil {
			log.Printf(" Errore: %v\n", err)
		} else {
			fmt.Println(" Completato.")
			convertedCount++
		}
	}

	fmt.Printf("\nConversione terminata. %d/%d file convertiti con successo.\n", convertedCount, totalFiles)
}