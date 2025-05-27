package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Valid config file", func(t *testing.T) {
		tempDir := t.TempDir()
		configPath := filepath.Join(tempDir, "test_config.json")
		
		configContent := `[
			{
				"id": "test-log-1",
				"path": "/path/to/log1.log",
				"type": "nginx-access"
			},
			{
				"id": "test-log-2",
				"path": "/path/to/log2.log",
				"type": "custom-app"
			}
		]`
		
		err := os.WriteFile(configPath, []byte(configContent), 0644)
		if err != nil {
			t.Fatalf("Erreur lors de la création du fichier de test: %v", err)
		}
		
		config, err := LoadConfig(configPath)
		if err != nil {
			t.Fatalf("Erreur lors du chargement de la configuration: %v", err)
		}
		
		if len(config.Logs) != 2 {
			t.Errorf("Attendu 2 logs, obtenu %d", len(config.Logs))
		}
		
		if config.Logs[0].ID != "test-log-1" {
			t.Errorf("Attendu ID 'test-log-1', obtenu '%s'", config.Logs[0].ID)
		}
		
		if config.Logs[1].Type != "custom-app" {
			t.Errorf("Attendu type 'custom-app', obtenu '%s'", config.Logs[1].Type)
		}
	})
	
	t.Run("Non-existent config file", func(t *testing.T) {
		_, err := LoadConfig("/non/existent/config.json")
		if err == nil {
			t.Error("Attendu une erreur pour un fichier inexistant")
		}
	})
	
	t.Run("Invalid JSON", func(t *testing.T) {
		tempDir := t.TempDir()
		configPath := filepath.Join(tempDir, "invalid_config.json")
		
		invalidJSON := `{invalid json content`
		err := os.WriteFile(configPath, []byte(invalidJSON), 0644)
		if err != nil {
			t.Fatalf("Erreur lors de la création du fichier de test: %v", err)
		}
		
		_, err = LoadConfig(configPath)
		if err == nil {
			t.Error("Attendu une erreur pour un JSON invalide")
		}
	})
	
	t.Run("Duplicate IDs", func(t *testing.T) {
		tempDir := t.TempDir()
		configPath := filepath.Join(tempDir, "duplicate_config.json")
		
		duplicateConfig := `[
			{
				"id": "duplicate-id",
				"path": "/path/to/log1.log",
				"type": "nginx-access"
			},
			{
				"id": "duplicate-id",
				"path": "/path/to/log2.log",
				"type": "custom-app"
			}
		]`
		
		err := os.WriteFile(configPath, []byte(duplicateConfig), 0644)
		if err != nil {
			t.Fatalf("Erreur lors de la création du fichier de test: %v", err)
		}
		
		_, err = LoadConfig(configPath)
		if err == nil {
			t.Error("Attendu une erreur pour des IDs dupliqués")
		}
	})
	
	t.Run("Missing fields", func(t *testing.T) {
		tempDir := t.TempDir()
		configPath := filepath.Join(tempDir, "missing_fields_config.json")
		
		missingFieldsConfig := `[
			{
				"id": "",
				"path": "/path/to/log1.log",
				"type": "nginx-access"
			}
		]`
		
		err := os.WriteFile(configPath, []byte(missingFieldsConfig), 0644)
		if err != nil {
			t.Fatalf("Erreur lors de la création du fichier de test: %v", err)
		}
		
		_, err = LoadConfig(configPath)
		if err == nil {
			t.Error("Attendu une erreur pour un ID manquant")
		}
	})
}

func TestValidateConfig(t *testing.T) {
	t.Run("Empty config", func(t *testing.T) {
		err := ValidateConfig([]LogConfig{})
		if err == nil {
			t.Error("Attendu une erreur pour une configuration vide")
		}
	})
	
	t.Run("Valid config", func(t *testing.T) {
		logs := []LogConfig{
			{ID: "log1", Path: "/path/to/log1.log", Type: "nginx"},
			{ID: "log2", Path: "/path/to/log2.log", Type: "apache"},
		}
		
		err := ValidateConfig(logs)
		if err != nil {
			t.Errorf("Erreur inattendue pour une configuration valide: %v", err)
		}
	})
} 