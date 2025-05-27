package analyzer

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/axellelanca/go_loganizer/internal/config"
)

func TestNewAnalyzer(t *testing.T) {
	analyzer := NewAnalyzer()
	if analyzer == nil {
		t.Error("NewAnalyzer() ne devrait pas retourner nil")
	}
	
	if analyzer.results == nil {
		t.Error("Le canal results ne devrait pas être nil")
	}
}

func TestAnalyzeLogs(t *testing.T) {
	tempDir := t.TempDir()
	
	validLogPath := filepath.Join(tempDir, "valid.log")
	err := os.WriteFile(validLogPath, []byte("test log content"), 0644)
	if err != nil {
		t.Fatalf("Erreur lors de la création du fichier de test: %v", err)
	}
	
	logs := []config.LogConfig{
		{
			ID:   "valid-log",
			Path: validLogPath,
			Type: "test",
		},
		{
			ID:   "invalid-log",
			Path: "/non/existent/path.log",
			Type: "test",
		},
	}
	
	analyzer := NewAnalyzer()
	results := analyzer.AnalyzeLogs(logs)
	
	if len(results) != 2 {
		t.Errorf("Attendu 2 résultats, obtenu %d", len(results))
	}
	
	hasSuccess := false
	hasFailure := false
	
	for _, result := range results {
		if result.Status == "OK" {
			hasSuccess = true
		}
		if result.Status == "FAILED" {
			hasFailure = true
		}
	}
	
	if !hasSuccess {
		t.Error("Attendu au moins un résultat avec statut OK")
	}
	
	if !hasFailure {
		t.Error("Attendu au moins un résultat avec statut FAILED")
	}
}

func TestCheckFileAccess(t *testing.T) {
	analyzer := NewAnalyzer()
	
	t.Run("Existing file", func(t *testing.T) {
		tempDir := t.TempDir()
		testFile := filepath.Join(tempDir, "test.log")
		
		err := os.WriteFile(testFile, []byte("test content"), 0644)
		if err != nil {
			t.Fatalf("Erreur lors de la création du fichier de test: %v", err)
		}
		
		err = analyzer.CheckFileAccess(testFile)
		if err != nil {
			t.Errorf("Erreur inattendue pour un fichier existant: %v", err)
		}
	})
	
	t.Run("Non-existent file", func(t *testing.T) {
		err := analyzer.CheckFileAccess("/non/existent/file.log")
		if err == nil {
			t.Error("Attendu une erreur pour un fichier inexistant")
		}
		
		var fileNotFoundErr *FileNotFoundError
		if !errors.As(err, &fileNotFoundErr) {
			t.Error("L'erreur devrait être de type FileNotFoundError")
		}
	})
	
	t.Run("Directory instead of file", func(t *testing.T) {
		tempDir := t.TempDir()
		
		err := analyzer.CheckFileAccess(tempDir)
		if err == nil {
			t.Error("Attendu une erreur pour un répertoire")
		}
		
		var fileNotFoundErr *FileNotFoundError
		if !errors.As(err, &fileNotFoundErr) {
			t.Error("L'erreur devrait être de type FileNotFoundError")
		}
	})
}

func TestAnalysisResult(t *testing.T) {
	result := AnalysisResult{
		LogID:        "test-log",
		FilePath:     "/path/to/test.log",
		Status:       "OK",
		Message:      "Test message",
		ErrorDetails: "",
	}
	
	if result.LogID != "test-log" {
		t.Errorf("Attendu LogID 'test-log', obtenu '%s'", result.LogID)
	}
	
	if result.Status != "OK" {
		t.Errorf("Attendu Status 'OK', obtenu '%s'", result.Status)
	}
}

func TestPrintResults(t *testing.T) {
	results := []AnalysisResult{
		{
			LogID:        "test-1",
			FilePath:     "/path/to/test1.log",
			Status:       "OK",
			Message:      "Success",
			ErrorDetails: "",
		},
		{
			LogID:        "test-2",
			FilePath:     "/path/to/test2.log",
			Status:       "FAILED",
			Message:      "Error",
			ErrorDetails: "File not found",
		},
	}
	
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PrintResults a paniqué: %v", r)
		}
	}()
	
	PrintResults(results)
} 