# Navidrome Mix

Service de génération de playlists intelligentes pour Navidrome, similaire aux Daily Mix de Spotify.

## Vision du projet

Navidrome Mix est un système qui s'intègre avec [Navidrome](https://www.navidrome.org/), le serveur de musique open-source, pour créer des playlists personnalisées basées sur l'historique d'écoute des utilisateurs. Il vise à offrir une expérience similaire aux services de streaming commerciaux, mais entièrement sous votre contrôle et adaptée à votre bibliothèque musicale personnelle.

## Fonctionnalités principales

- Génération de playlists personnalisées basées sur l'historique d'écoute
- Création de mix cohérents avec des pistes musicalement similaires
- Découverte de nouvelles pistes pertinentes et non redondantes
- Interface utilisateur intégrée pour la gestion des playlists générées
- Enrichissement de métadonnées via des services externes (MusicBrainz, Last.fm)

## Architecture

Le projet est divisé en plusieurs composants :

1. **Navidrome Mix Service** : Service principal pour la génération de playlists
2. **Base de données de recommandation** : Stockage des données de recommandation et relations entre morceaux
3. **Navidrome Mix UI** : Interface utilisateur pour interagir avec le système
4. **Intégration Navidrome** : Communication avec le serveur Navidrome via l'API Subsonic

## Installation

```bash
# Cloner le repository
git clone https://gitlab.mserv.wtf/navidrome-mix.git
cd navidrome-mix

# Lancer les services avec Docker Compose
docker-compose up -d
```

## Configuration

Le fichier `config.yaml` permet de configurer les paramètres du service :

```yaml
navidrome:
  url: http://navidrome:4533
  username: admin
  password: password

recommendation:
  similarity_threshold: 0.7
  max_tracks_per_playlist: 25
  discovery_ratio: 0.3
```

## Développement

### Prérequis

- Go 1.20+
- Docker et Docker Compose
- Node.js 18+ (pour l'UI)

### Structure du projet

```
navidrome-mix/
├── api/                # API REST du service
├── cmd/                # Points d'entrée
├── config/             # Configuration
├── db/                 # Couche d'accès aux données
├── docker/             # Fichiers Docker
├── docs/               # Documentation
├── pkg/                # Bibliothèques partagées
├── recommender/        # Algorithmes de recommandation
├── scripts/            # Scripts utilitaires
├── ui/                 # Interface utilisateur
└── docker-compose.yml  # Configuration Docker Compose
```

## Licence

[MIT License](LICENSE)
