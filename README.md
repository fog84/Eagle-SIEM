![Logo Eagle](images/logo.jpg)
# Security information and event management
Eagle SIEM allows you to centralize, correlate and analyze logs in real time to detect and respond to security threats.

## Schéma architectural
![Schéma architectural v2](images/schéma%20de%20concept%20architectural%20v2.png)

## Architecture
**Agents / Linux** : Daemon codé en Go

**Modules Agents** : Go/Python/Autre

**Indexer** : Python FastAPI et MySQL (Docker)

**User Interface**: Python FastAPI + sqlite3 + Front-end en HTML/CSS/JS (Docker)

## Gestion des accès
L'Indexer autorise l'écriture et la lecture des logs sur la BDD via la vérifiaction des tokens correspondant (api_key, api_key_readlogs).

La gestion des accès à l'interface utilisateur est gérée via sqlite3 par JWT signé.

## Exemple d'utilisation
Un serveur Apache génère des logs dans des fichiers. L'agent vérifie toutes les 10 secondes ces fichiers et envoie à l'indexer les nouvelles lignes ainsi que le token agent.

L'indexer vérifie le token agent et récupère le nom de la machine correspondante. Il enregistre les logs dans la base de données MySQL.

Les logs peuvent être visionnés sur l'interface utilisateur.

## Installation
### Indexer
```
cd Indexer
bash generate_env.sh
docker compose up
```
### UserInterface
```
cd UserInterface
bash generate_env.sh
docker compose up
```

## Fonctionnalités v2.2 du projet
- [ ] Agents Linux :
---
- [X] Indexer : Faciliter l'installation via un script shell pour générer l'environement
---
- [X] Interface utilisateur : Faciliter l'installation via un script shell pour générer l'environement
- [X] Interface utilisateur : Style css login
- [X] Interface utilisateur : Ui pour la route save_indexer_api_key_readlogs_in_cookie
- [X] Interface utilisateur : CSS pour l'ui de la route save_indexer_api_key_readlogs_in_cookie
- [X] Interface utilisateur : Redirection / vers le login
- [X] Interface utilisateur : Ajout lien vers save_indexer_api_key_readlogs_in_cookie dans siem_ui
- [X] Interface utilisateur : Code de base (Création de compte)
---
- [ ] Modules Agents : Adapter le code des modules existants dans la v1
---
- [ ] Sécurité : Audit du code agent
- [ ] Sécurité : Audit du code Indexer
- [ ] Sécurité : Audit du code User Interface


## Fonctionnalités v2.1 du projet
- [X] Agents Linux : code de base (Récupérer les nouveaux logs à intervalles réguliers dans une liste de fichiers définie)
- [X] Agents Linux : code de base (Faire du programme un démon)
---
- [X] Indexer : code de base (Infrastructure Docker)
- [X] Indexer : code de base (Connexion à la BDD MySQL)
- [X] Indexer : code de base (Réception des logs)
- [X] Indexer : code de base (Vérification du token)
- [X] Indexer : code de base (Enregistrement des logs reçus dans la BDD si le token est valide)
---
- [X] Interface utilisateur : code de base (Infrastructure Docker)
- [X] Interface utilisateur : code de base (Connexion à la BDD SQLite3)
- [X] Interface utilisateur : code de base (Connexion)
- [X] Interface utilisateur : code de base (Récupération des logs de l'indexer)
- [X] Interface utilisateur : code de base (UI minimale)
---

