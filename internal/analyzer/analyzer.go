package analyzer

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/axellelanca/go_loganizer/internal/config"
)

type AnalysisResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

type Analyzer struct {
	results chan AnalysisResult
	wg      sync.WaitGroup
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		results: make(chan AnalysisResult, 100),
	}
}

func (a *Analyzer) AnalyzeLogs(logs []config.LogConfig) []AnalysisResult {
	for _, log := range logs {
		a.wg.Add(1)
		go a.analyzeLog(log)
	}

	go func() {
		a.wg.Wait()
		close(a.results)
	}()

	var allResults []AnalysisResult
	for result := range a.results {
		allResults = append(allResults, result)
	}

	return allResults
}

func (a *Analyzer) analyzeLog(log config.LogConfig) {
	defer a.wg.Done()

	result := AnalysisResult{
		LogID:    log.ID,
		FilePath: log.Path,
	}

	if err := a.CheckFileAccess(log.Path); err != nil {
		var fileNotFoundErr *FileNotFoundError
		if errors.As(err, &fileNotFoundErr) {
			result.Status = "FAILED"
			result.Message = "Fichier introuvable."
			result.ErrorDetails = err.Error()
		} else {
			result.Status = "FAILED"
			result.Message = "Erreur d'accès au fichier."
			result.ErrorDetails = err.Error()
		}
		a.results <- result
		return
	}

	analysisTime := time.Duration(50+rand.Intn(151)) * time.Millisecond
	time.Sleep(analysisTime)

	if rand.Float32() < 0.1 {
		parseErr := NewParseError(log.ID, "données corrompues détectées", fmt.Errorf("format de log invalide"))
		result.Status = "FAILED"
		result.Message = "Erreur de parsing."
		result.ErrorDetails = parseErr.Error()
		a.results <- result
		return
	}

	result.Status = "OK"
	result.Message = "Analyse terminée avec succès."
	result.ErrorDetails = ""
	a.results <- result
}

func (a *Analyzer) CheckFileAccess(filePath string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return NewFileNotFoundError(filePath, err)
		}
		return fmt.Errorf("erreur d'accès au fichier %s: %w", filePath, err)
	}

	if info.IsDir() {
		return NewFileNotFoundError(filePath, fmt.Errorf("le chemin pointe vers un répertoire"))
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("impossible d'ouvrir le fichier %s: %w", filePath, err)
	}
	file.Close()

	return nil
}

func PrintResults(results []AnalysisResult) {
	fmt.Println("\n=== Résultats de l'analyse ===")
	fmt.Printf("Total des logs analysés: %d\n\n", len(results))

	for _, result := range results {
		fmt.Printf("ID: %s\n", result.LogID)
		fmt.Printf("Chemin: %s\n", result.FilePath)
		fmt.Printf("Statut: %s\n", result.Status)
		fmt.Printf("Message: %s\n", result.Message)
		if result.ErrorDetails != "" {
			fmt.Printf("Détails de l'erreur: %s\n", result.ErrorDetails)
		}
		fmt.Println("---")
	}

	successCount := 0
	failureCount := 0
	for _, result := range results {
		if result.Status == "OK" {
			successCount++
		} else {
			failureCount++
		}
	}

	fmt.Printf("\nStatistiques:\n")
	fmt.Printf("✅ Succès: %d\n", successCount)
	fmt.Printf("❌ Échecs: %d\n", failureCount)
} 