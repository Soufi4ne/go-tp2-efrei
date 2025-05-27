package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type Config struct {
	Logs []LogConfig `json:"-"`
}

func LoadConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("fichier de configuration introuvable: %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier de configuration: %w", err)
	}

	var logs []LogConfig
	if err := json.Unmarshal(data, &logs); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON: %w", err)
	}

	if err := ValidateConfig(logs); err != nil {
		return nil, fmt.Errorf("configuration invalide: %w", err)
	}

	return &Config{Logs: logs}, nil
}

func ValidateConfig(logs []LogConfig) error {
	if len(logs) == 0 {
		return fmt.Errorf("aucun log configuré")
	}

	ids := make(map[string]bool)
	for _, log := range logs {
		if log.ID == "" {
			return fmt.Errorf("ID manquant pour un log")
		}
		if log.Path == "" {
			return fmt.Errorf("chemin manquant pour le log %s", log.ID)
		}
		if log.Type == "" {
			return fmt.Errorf("type manquant pour le log %s", log.ID)
		}
		if ids[log.ID] {
			return fmt.Errorf("ID dupliqué: %s", log.ID)
		}
		ids[log.ID] = true
	}

	return nil
} 