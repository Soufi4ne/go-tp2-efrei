package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/axellelanca/go_loganizer/internal/config"
)

var (
	addLogConfigPath string
	addLogID        string
	addLogPath      string
	addLogType      string
)

var addLogCmd = &cobra.Command{
	Use:   "add-log",
	Short: "Ajoute une nouvelle configuration de log au fichier config.json",
	Long: `La commande add-log permet d'ajouter manuellement une nouvelle configuration
de log au fichier config.json existant.

Exemples:
  # Ajouter une nouvelle configuration de log
  loganalyzer add-log --file config.json --id "new-log" --path "/var/log/new.log" --type "nginx"

  # Utiliser les raccourcis
  loganalyzer add-log -f config.json -i "new-log" -p "/var/log/new.log" -t "nginx"`,
	RunE: runAddLog,
}

func init() {
	rootCmd.AddCommand(addLogCmd)

	addLogCmd.Flags().StringVarP(&addLogConfigPath, "file", "f", "", "Chemin vers le fichier de configuration JSON (requis)")
	addLogCmd.Flags().StringVarP(&addLogID, "id", "i", "", "Identifiant unique du log (requis)")
	addLogCmd.Flags().StringVarP(&addLogPath, "path", "p", "", "Chemin vers le fichier de log (requis)")
	addLogCmd.Flags().StringVarP(&addLogType, "type", "t", "", "Type de log (requis)")

	addLogCmd.MarkFlagRequired("file")
	addLogCmd.MarkFlagRequired("id")
	addLogCmd.MarkFlagRequired("path")
	addLogCmd.MarkFlagRequired("type")
}

func runAddLog(cmd *cobra.Command, args []string) error {
	// Load existing config
	cfg, err := config.LoadConfig(addLogConfigPath)
	if err != nil {
		return fmt.Errorf("erreur lors du chargement de la configuration: %w", err)
	}

	// Check if log ID already exists
	for _, log := range cfg.Logs {
		if log.ID == addLogID {
			return fmt.Errorf("un log avec l'ID '%s' existe déjà", addLogID)
		}
	}

	// Create new log configuration
	newLog := config.LogConfig{
		ID:   addLogID,
		Path: addLogPath,
		Type: addLogType,
	}

	// Add new log to configuration
	cfg.Logs = append(cfg.Logs, newLog)

	// Write updated config back to file
	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de la sérialisation JSON: %w", err)
	}

	if err := os.WriteFile(addLogConfigPath, jsonData, 0644); err != nil {
		return fmt.Errorf("erreur lors de l'écriture du fichier %s: %w", addLogConfigPath, err)
	}

	fmt.Printf("✅ Nouvelle configuration de log ajoutée avec succès:\n")
	fmt.Printf("   ID: %s\n", newLog.ID)
	fmt.Printf("   Chemin: %s\n", newLog.Path)
	fmt.Printf("   Type: %s\n", newLog.Type)

	return nil
} 