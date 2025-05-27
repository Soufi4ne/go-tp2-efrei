package cmd

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/axellelanca/go_loganizer/internal/analyzer"
	"github.com/axellelanca/go_loganizer/internal/config"
	"github.com/axellelanca/go_loganizer/internal/reporter"
)

var (
	configPath string
	outputPath string
	statusFilter string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs spécifiés dans la configuration",
	Long: `La commande analyze lit un fichier de configuration JSON contenant
la liste des logs à analyser et les traite de manière concurrente.

Chaque log est analysé dans une goroutine séparée pour optimiser les performances.
Les résultats peuvent être exportés au format JSON.

Exemples:
  # Analyser avec affichage console uniquement
  loganalyzer analyze --config config.json

  # Analyser et exporter vers un fichier JSON
  loganalyzer analyze --config config.json --output report.json

  # Filtrer par statut
  loganalyzer analyze --config config.json --status FAILED

  # Utiliser les raccourcis
  loganalyzer analyze -c config.json -o report.json -s FAILED`,
	RunE: runAnalyze,
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin vers le fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin vers le fichier de sortie JSON (optionnel)")
	analyzeCmd.Flags().StringVarP(&statusFilter, "status", "s", "", "Filtrer les résultats par statut (ex: FAILED, OK)")

	analyzeCmd.MarkFlagRequired("config")
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	fmt.Printf("🔍 Démarrage de l'analyse des logs...\n")
	fmt.Printf("📁 Fichier de configuration: %s\n", configPath)

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("erreur lors du chargement de la configuration: %w", err)
	}

	fmt.Printf("📊 Nombre de logs à analyser: %d\n", len(cfg.Logs))

	var rep *reporter.Reporter
	if outputPath != "" {
		rep = reporter.NewReporter()
		// Add timestamp to output filename
		dir := filepath.Dir(outputPath)
		base := filepath.Base(outputPath)
		ext := filepath.Ext(base)
		name := strings.TrimSuffix(base, ext)
		timestamp := time.Now().Format("060102") // YYMMDD format
		outputPath = filepath.Join(dir, fmt.Sprintf("%s_%s%s", timestamp, name, ext))
		
		if err := rep.ValidateOutputPath(outputPath); err != nil {
			return fmt.Errorf("chemin de sortie invalide: %w", err)
		}
		fmt.Printf("📤 Fichier de sortie: %s\n", outputPath)
	}

	fmt.Println("\n🚀 Lancement de l'analyse concurrente...")

	logAnalyzer := analyzer.NewAnalyzer()
	startTime := time.Now()
	results := logAnalyzer.AnalyzeLogs(cfg.Logs)
	duration := time.Since(startTime)

	// Filter results if status filter is provided
	if statusFilter != "" {
		filteredResults := make([]analyzer.AnalysisResult, 0)
		for _, result := range results {
			if result.Status == statusFilter {
				filteredResults = append(filteredResults, result)
			}
		}
		results = filteredResults
	}

	analyzer.PrintResults(results)

	fmt.Printf("\n⏱️  Temps d'exécution total: %v\n", duration)

	if outputPath != "" && rep != nil {
		if err := rep.ExportToJSON(results, outputPath); err != nil {
			return fmt.Errorf("erreur lors de l'export JSON: %w", err)
		}
	}

	fmt.Println("\n✨ Analyse terminée avec succès!")
	return nil
}