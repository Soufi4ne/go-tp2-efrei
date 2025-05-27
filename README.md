# 🔍 LogAnalyzer - Outil d'Analyse de Logs Distribuée

[![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

LogAnalyzer est un outil en ligne de commande (CLI) développé en Go pour analyser des fichiers de logs de manière distribuée et concurrente. Il permet aux administrateurs système d'analyser plusieurs logs en parallèle et d'extraire des informations clés tout en gérant les erreurs de manière robuste.

## 🚀 Fonctionnalités

- ✅ **Analyse concurrente** : Traitement de plusieurs logs en parallèle avec des goroutines
- ✅ **Gestion d'erreurs robuste** : Erreurs personnalisées avec `errors.Is()` et `errors.As()`
- ✅ **Interface CLI intuitive** : Commandes et drapeaux avec Cobra
- ✅ **Export JSON** : Sauvegarde des résultats au format JSON
- ✅ **Statistiques détaillées** : Résumé des succès et échecs
- ✅ **Architecture modulaire** : Code organisé en packages logiques

## 📋 Prérequis

- Go 1.24 ou supérieur
- Système d'exploitation : Windows, macOS, Linux

## 🛠️ Installation

### Cloner le projet

```bash
git clone https://github.com/axellelanca/loganizer.git
cd loganizer
```

### Installer les dépendances

```bash
go mod tidy
```

### Compiler l'application

```bash
# Linux/macOS
go build -o loganalyzer .

# Windows
go build -o loganalyzer.exe .
```

## 📖 Utilisation

### Commande principale

```bash
loganalyzer --help
```

### Commande analyze

La commande `analyze` est le cœur de l'application. Elle lit un fichier de configuration JSON et analyse les logs spécifiés de manière concurrente.

#### Syntaxe

```bash
loganalyzer analyze --config <chemin_config> [--output <chemin_sortie>]
```

#### Options

- `-c, --config` : Chemin vers le fichier de configuration JSON (requis)
- `-o, --output` : Chemin vers le fichier de sortie JSON (optionnel)
- `-h, --help` : Afficher l'aide

#### Exemples

```bash
# Analyse avec affichage console uniquement
loganalyzer analyze --config config.json

# Analyse avec export JSON
loganalyzer analyze --config config.json --output report.json

# Utilisation des raccourcis
loganalyzer analyze -c config.json -o report.json
```

## 📁 Format du fichier de configuration

Le fichier de configuration doit être au format JSON et contenir un tableau d'objets représentant les logs à analyser :

```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  },
  {
    "id": "db-server-3",
    "path": "test_logs/mysql_error.log",
    "type": "mysql-error"
  }
]
```

### Champs requis

- `id` : Identifiant unique du log
- `path` : Chemin vers le fichier de log (absolu ou relatif)
- `type` : Type de log (pour classification)

## 📊 Format du rapport de sortie

Le rapport JSON généré contient les résultats de l'analyse :

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "fichier introuvable ou inaccessible: /non/existent/log.log"
  }
]
```

### Champs du rapport

- `log_id` : Identifiant du log analysé
- `file_path` : Chemin du fichier analysé
- `status` : Statut de l'analyse (`OK` ou `FAILED`)
- `message` : Message descriptif du résultat
- `error_details` : Détails de l'erreur (si applicable)

## 🏗️ Architecture du projet

```
loganizer/
├── cmd/                    # Commandes CLI
│   ├── root.go            # Commande racine
│   └── analyze.go         # Commande analyze
├── internal/              # Packages internes
│   ├── config/           # Gestion des configurations
│   │   └── config.go
│   ├── analyzer/         # Logique d'analyse
│   │   ├── analyzer.go
│   │   └── errors.go
│   └── reporter/         # Export des résultats
│       └── reporter.go
├── test_logs/            # Fichiers de test
├── config.json           # Configuration d'exemple
├── main.go              # Point d'entrée
├── go.mod               # Module Go
└── README.md            # Documentation
```

### Packages

#### `internal/config`
Gère la lecture et la validation des fichiers de configuration JSON.

#### `internal/analyzer`
Contient la logique d'analyse des logs, les erreurs personnalisées et l'affichage des résultats.

#### `internal/reporter`
Gère l'export des résultats au format JSON avec création automatique des répertoires.

#### `cmd`
Contient les commandes CLI utilisant le framework Cobra.

## 🔧 Concepts techniques implémentés

### Concurrence
- **Goroutines** : Chaque log est analysé dans une goroutine séparée
- **WaitGroups** : Synchronisation des goroutines
- **Channels** : Communication sécurisée entre goroutines

### Gestion d'erreurs
- **Erreurs personnalisées** : `FileNotFoundError` et `ParseError`
- **Wrapping d'erreurs** : Utilisation de `fmt.Errorf` avec `%w`
- **Inspection d'erreurs** : `errors.Is()` et `errors.As()`

### CLI avec Cobra
- **Commandes structurées** : Commande racine et sous-commandes
- **Drapeaux** : Options courtes et longues
- **Validation** : Drapeaux requis et validation des entrées

## 🧪 Tests et exemples

Le projet inclut des fichiers de test dans le répertoire `test_logs/` :

- `access.log` : Log d'accès web
- `errors.log` : Log d'erreurs d'application
- `mysql_error.log` : Log d'erreurs MySQL
- `empty.log` : Fichier vide
- `corrupted.log` : Fichier avec données corrompues

### Exécution des tests

```bash
# Test avec la configuration par défaut
./loganalyzer analyze -c config.json

# Test avec export JSON
./loganalyzer analyze -c config.json -o test_report.json
```

## 🎯 Fonctionnalités avancées

### Simulation d'erreurs
- **Erreurs aléatoires** : 10% de chance d'erreur de parsing
- **Délais variables** : Simulation réaliste avec délais de 50-200ms
- **Gestion des fichiers manquants** : Détection et rapport des fichiers inaccessibles

### Création automatique de répertoires
Si le chemin de sortie inclut des répertoires inexistants, ils sont créés automatiquement.

```bash
# Crée automatiquement le répertoire "rapports/2024/"
./loganalyzer analyze -c config.json -o rapports/2024/mon_rapport.json
```

## 👥 Équipe de développement

- **Soufiane** - Développeur principal
- **Axelle Lanca** - Contributeur original du repository

## 📝 Licence

Ce projet est sous licence MIT. Voir le fichier [LICENSE](LICENSE) pour plus de détails.

## 🤝 Contribution

Les contributions sont les bienvenues ! Pour contribuer :

1. Forkez le projet
2. Créez une branche pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. Committez vos changements (`git commit -m 'Add some AmazingFeature'`)
4. Poussez vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrez une Pull Request

## 📞 Support

Pour toute question ou problème, veuillez ouvrir une issue sur GitHub.

---

**LogAnalyzer** - Analyse de logs distribuée avec Go 🚀
