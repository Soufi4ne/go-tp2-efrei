# 🎬 Démonstration LogAnalyzer

Ce fichier contient des exemples d'utilisation de LogAnalyzer pour démontrer ses capacités.

## 🚀 Exemples d'utilisation

### 1. Analyse basique avec affichage console

```bash
./loganalyzer analyze -c config.json
```

**Résultat attendu :**
- Analyse de 6 logs en parallèle
- Affichage des résultats sur la console
- Statistiques de succès/échecs
- Temps d'exécution affiché

### 2. Analyse avec export JSON

```bash
./loganalyzer analyze -c config.json -o report.json
```

**Résultat attendu :**
- Même analyse que précédemment
- Export des résultats vers `report.json`
- Création automatique du fichier

### 3. Analyse avec création de répertoires

```bash
./loganalyzer analyze -c config.json -o rapports/2024/monthly_report.json
```

**Résultat attendu :**
- Création automatique du répertoire `rapports/2024/`
- Export du rapport dans le nouveau répertoire

### 4. Affichage de l'aide

```bash
./loganalyzer --help
./loganalyzer analyze --help
```

## 🧪 Tests de fonctionnalités

### Concurrence
L'application traite tous les logs en parallèle. Vous pouvez observer :
- Les résultats apparaissent dans un ordre différent à chaque exécution
- Le temps total est optimisé grâce au traitement concurrent

### Gestion d'erreurs
Le fichier `config.json` inclut volontairement :
- Un fichier inexistant (`/non/existent/log.log`)
- Des fichiers valides pour tester les succès
- Une simulation d'erreurs de parsing (10% de chance)

### Validation
L'application valide :
- L'existence du fichier de configuration
- La structure JSON du fichier de configuration
- L'unicité des IDs de logs
- L'extension `.json` pour les fichiers de sortie

## 📊 Exemple de sortie

```
🔍 Démarrage de l'analyse des logs...
📁 Fichier de configuration: config.json
📊 Nombre de logs à analyser: 6
📤 Fichier de sortie: report.json

🚀 Lancement de l'analyse concurrente...

=== Résultats de l'analyse ===
Total des logs analysés: 6

ID: web-server-1
Chemin: test_logs/access.log
Statut: OK
Message: Analyse terminée avec succès.
---

ID: invalid-path
Chemin: /non/existent/log.log
Statut: FAILED
Message: Fichier introuvable.
Détails de l'erreur: fichier introuvable ou inaccessible: /non/existent/log.log
---

Statistiques:
✅ Succès: 5
❌ Échecs: 1

⏱️  Temps d'exécution total: 150ms
✅ Rapport exporté avec succès vers: report.json

✨ Analyse terminée avec succès!
```

## 🔧 Personnalisation

### Modifier la configuration

Éditez `config.json` pour ajouter vos propres logs :

```json
[
  {
    "id": "mon-log-custom",
    "path": "chemin/vers/mon/log.log",
    "type": "custom-type"
  }
]
```

### Tester avec vos fichiers

1. Créez vos fichiers de logs
2. Modifiez `config.json`
3. Lancez l'analyse

## 🎯 Points techniques démontrés

- ✅ **Goroutines** : Traitement concurrent des logs
- ✅ **Channels** : Communication sécurisée entre goroutines
- ✅ **WaitGroups** : Synchronisation des goroutines
- ✅ **Erreurs personnalisées** : `FileNotFoundError` et `ParseError`
- ✅ **CLI avec Cobra** : Interface utilisateur intuitive
- ✅ **JSON** : Import/export de données structurées
- ✅ **Architecture modulaire** : Packages organisés logiquement 