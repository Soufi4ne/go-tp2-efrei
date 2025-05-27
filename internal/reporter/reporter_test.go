package reporter

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
)

func TestNewReporter(t *testing.T) {
	reporter := NewReporter()
	if reporter == nil {
		t.Error("NewReporter() ne devrait pas retourner nil")
	}
}

func TestExportToJSON(t *testing.T) {
	reporter := NewReporter()
	
	results := []analyzer.AnalysisResult{
		{
			LogID:        "test-log-1",
			FilePath:     "/path/to/test1.log",
			Status:       "OK",
			Message:      "Analyse terminée avec succès.",
			ErrorDetails: "",
		},
		{
			LogID:        "test-log-2",
			FilePath:     "/path/to/test2.log",
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: "file not found error",
		},
	}
	
	t.Run("Simple file export", func(t *testing.T) {
		tempDir := t.TempDir()
		outputPath := filepath.Join(tempDir, "test_report.json")
		
		err := reporter.ExportToJSON(results, outputPath)
		if err != nil {
			t.Fatalf("Erreur lors de l'export JSON: %v", err)
		}
		
		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			t.Error("Le fichier de rapport n'a pas été créé")
		}
		
		data, err := os.ReadFile(outputPath)
		if err != nil {
			t.Fatalf("Erreur lors de la lecture du fichier: %v", err)
		}
		
		var loadedResults []analyzer.AnalysisResult
		err = json.Unmarshal(data, &loadedResults)
		if err != nil {
			t.Fatalf("Erreur lors du parsing JSON: %v", err)
		}
		
		if len(loadedResults) != len(results) {
			t.Errorf("Attendu %d résultats, obtenu %d", len(results), len(loadedResults))
		}
		
		if loadedResults[0].LogID != results[0].LogID {
			t.Errorf("LogID ne correspond pas: attendu %s, obtenu %s", results[0].LogID, loadedResults[0].LogID)
		}
	})
	
	t.Run("Export with directory creation", func(t *testing.T) {
		tempDir := t.TempDir()
		outputPath := filepath.Join(tempDir, "rapports", "2024", "test_report.json")
		
		err := reporter.ExportToJSON(results, outputPath)
		if err != nil {
			t.Fatalf("Erreur lors de l'export JSON avec création de répertoires: %v", err)
		}
		
		if _, err := os.Stat(outputPath); os.IsNotExist(err) {
			t.Error("Le fichier de rapport n'a pas été créé dans les sous-répertoires")
		}
		
		dirPath := filepath.Join(tempDir, "rapports", "2024")
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Error("Les répertoires n'ont pas été créés automatiquement")
		}
	})
}

func TestValidateOutputPath(t *testing.T) {
	reporter := NewReporter()
	
	t.Run("Valid JSON extension", func(t *testing.T) {
		err := reporter.ValidateOutputPath("report.json")
		if err != nil {
			t.Errorf("Erreur inattendue pour un chemin valide: %v", err)
		}
	})
	
	t.Run("Invalid extension", func(t *testing.T) {
		err := reporter.ValidateOutputPath("report.txt")
		if err == nil {
			t.Error("Attendu une erreur pour une extension invalide")
		}
	})
	
	t.Run("Complex valid path", func(t *testing.T) {
		tempDir := t.TempDir()
		complexPath := filepath.Join(tempDir, "rapports", "2024", "report.json")
		
		err := reporter.ValidateOutputPath(complexPath)
		if err != nil {
			t.Errorf("Erreur inattendue pour un chemin complexe valide: %v", err)
		}
	})
}

func TestEnsureDirectoryExists(t *testing.T) {
	reporter := NewReporter()
	
	t.Run("Current directory", func(t *testing.T) {
		err := reporter.ensureDirectoryExists("report.json")
		if err != nil {
			t.Errorf("Erreur inattendue pour le répertoire courant: %v", err)
		}
	})
	
	t.Run("Create directories", func(t *testing.T) {
		tempDir := t.TempDir()
		filePath := filepath.Join(tempDir, "new", "sub", "dir", "report.json")
		
		err := reporter.ensureDirectoryExists(filePath)
		if err != nil {
			t.Errorf("Erreur lors de la création des répertoires: %v", err)
		}
		
		dirPath := filepath.Join(tempDir, "new", "sub", "dir")
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Error("Les répertoires n'ont pas été créés")
		}
	})
}

func TestExportEmptyResults(t *testing.T) {
	reporter := NewReporter()
	tempDir := t.TempDir()
	outputPath := filepath.Join(tempDir, "empty_report.json")
	
	emptyResults := []analyzer.AnalysisResult{}
	
	err := reporter.ExportToJSON(emptyResults, outputPath)
	if err != nil {
		t.Fatalf("Erreur lors de l'export de résultats vides: %v", err)
	}
	
	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("Erreur lors de la lecture du fichier: %v", err)
	}
	
	var loadedResults []analyzer.AnalysisResult
	err = json.Unmarshal(data, &loadedResults)
	if err != nil {
		t.Fatalf("Erreur lors du parsing JSON: %v", err)
	}
	
	if len(loadedResults) != 0 {
		t.Errorf("Attendu 0 résultats, obtenu %d", len(loadedResults))
	}
} 