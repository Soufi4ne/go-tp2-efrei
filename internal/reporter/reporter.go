package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
)

type Reporter struct{}

func NewReporter() *Reporter {
	return &Reporter{}
}

func (r *Reporter) ExportToJSON(results []analyzer.AnalysisResult, outputPath string) error {
	if err := r.ensureDirectoryExists(outputPath); err != nil {
		return fmt.Errorf("erreur lors de la création des répertoires: %w", err)
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la sérialisation JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, jsonData, 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier %s: %w", outputPath, err)
	}

	fmt.Printf("✅ Rapport exporté avec succès vers: %s\n", outputPath)
	return nil
}

func (r *Reporter) ensureDirectoryExists(filePath string) error {
	dir := filepath.Dir(filePath)
	if dir == "." {
		return nil
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("impossible de créer le répertoire %s: %w", dir, err)
	}

	return nil
}

func (r *Reporter) ValidateOutputPath(outputPath string) error {
	if filepath.Ext(outputPath) != ".json" {
		return fmt.Errorf("le fichier de sortie doit avoir l'extension .json")
	}

	dir := filepath.Dir(outputPath)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("impossible d'accéder ou créer le répertoire %s: %w", dir, err)
		}
	}

	return nil
} 