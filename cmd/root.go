package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Outil d'analyse de logs distribuée",
	Long: `LogAnalyzer est un outil en ligne de commande pour analyser des fichiers de logs
de manière distribuée et concurrente.

Il permet de:
- Analyser plusieurs logs en parallèle
- Gérer les erreurs de manière robuste
- Exporter les résultats au format JSON
- Afficher des statistiques détaillées

Exemple d'utilisation:
  loganalyzer analyze --config config.json --output report.json`,
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur lors de l'exécution de la commande: %v\n", err)
		os.Exit(1)
	}
}

func init() {


	rootCmd.Flags().BoolP("toggle", "t", false, "Aide pour toggle")
} 