package analyzer

import (
	"errors"
	"fmt"
	"testing"
)

func TestFileNotFoundError(t *testing.T) {
	originalErr := fmt.Errorf("original error")
	filePath := "/path/to/missing/file.log"
	
	err := NewFileNotFoundError(filePath, originalErr)
	
	expectedMsg := "fichier introuvable ou inaccessible: /path/to/missing/file.log"
	if err.Error() != expectedMsg {
		t.Errorf("Attendu '%s', obtenu '%s'", expectedMsg, err.Error())
	}
	
	if !errors.Is(err, originalErr) {
		t.Error("L'erreur devrait wrapper l'erreur originale")
	}
	
	if err.FilePath != filePath {
		t.Errorf("Attendu FilePath '%s', obtenu '%s'", filePath, err.FilePath)
	}
	
	if err.Err != originalErr {
		t.Error("L'erreur interne ne correspond pas")
	}
}

func TestParseError(t *testing.T) {
	originalErr := fmt.Errorf("parsing failed")
	logID := "test-log-123"
	message := "données corrompues"
	
	err := NewParseError(logID, message, originalErr)
	
	expectedMsg := "erreur de parsing pour le log test-log-123: données corrompues"
	if err.Error() != expectedMsg {
		t.Errorf("Attendu '%s', obtenu '%s'", expectedMsg, err.Error())
	}
	
	if !errors.Is(err, originalErr) {
		t.Error("L'erreur devrait wrapper l'erreur originale")
	}
	
	if err.LogID != logID {
		t.Errorf("Attendu LogID '%s', obtenu '%s'", logID, err.LogID)
	}
	
	if err.Message != message {
		t.Errorf("Attendu Message '%s', obtenu '%s'", message, err.Message)
	}
	
	if err.Err != originalErr {
		t.Error("L'erreur interne ne correspond pas")
	}
}

func TestErrorsAs(t *testing.T) {
	t.Run("FileNotFoundError with errors.As", func(t *testing.T) {
		originalErr := fmt.Errorf("file not found")
		fileErr := NewFileNotFoundError("/test/path", originalErr)
		
		var targetErr *FileNotFoundError
		if !errors.As(fileErr, &targetErr) {
			t.Error("errors.As devrait fonctionner avec FileNotFoundError")
		}
		
		if targetErr.FilePath != "/test/path" {
			t.Error("Le FilePath ne correspond pas après errors.As")
		}
	})
	
	t.Run("ParseError with errors.As", func(t *testing.T) {
		originalErr := fmt.Errorf("parse failed")
		parseErr := NewParseError("log-id", "test message", originalErr)
		
		var targetErr *ParseError
		if !errors.As(parseErr, &targetErr) {
			t.Error("errors.As devrait fonctionner avec ParseError")
		}
		
		if targetErr.LogID != "log-id" {
			t.Error("Le LogID ne correspond pas après errors.As")
		}
	})
}

func TestErrorsIs(t *testing.T) {
	originalErr := fmt.Errorf("base error")
	
	fileErr := NewFileNotFoundError("/test/path", originalErr)
	if !errors.Is(fileErr, originalErr) {
		t.Error("errors.Is devrait fonctionner avec FileNotFoundError")
	}
	
	parseErr := NewParseError("log-id", "message", originalErr)
	if !errors.Is(parseErr, originalErr) {
		t.Error("errors.Is devrait fonctionner avec ParseError")
	}
	
	differentErr := fmt.Errorf("different error")
	if errors.Is(fileErr, differentErr) {
		t.Error("errors.Is ne devrait pas matcher une erreur différente")
	}
} 