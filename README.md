![Logo Eagle](images/logo.jpg)
# Security information and event management
Eagle SIEM allows you to centralize, correlate and analyze logs in real time to detect and respond to security threats. It has agents for Linux and Windows, very customizable and adaptable to many situations. Its interface allows you to directly filter logs with SQL queries, save them and apply regex detection filters.

## Schéma architectural
![Schéma architectural v2](images/schéma%20de%20concept%20architectural%20v2.png)

## Architecture
**Agents / Linux** : Daemon codé en Go

**Agents / Windows** : Executable codé en Go lancé via le task scheduler

**Modules Agents** : Go/Python/Autre

**Indexer** : Node.js et MySQL (Docker)

**User Interface**: Python FastAPI + sqlite3 + Front-end en HTML/CSS/JS (Docker)

## Gestion des accès
La BDD de l'indexer possède une table Tokens qui contient les tokens générés via l'interface utilisateur. L'indexer n'enregistre que les logs contenant un token valide dans leur requête.

La gestion des accès à l'interface utilisateur est gérée via sqlite3 par JWT signé.

## Exemple d'utilisation
Un serveur Apache génère des logs dans des fichiers. L'agent vérifie toutes les 10 secondes ces fichiers et envoie à l'indexer les nouvelles lignes ainsi que le token agent.

L'indexer vérifie le token agent et récupère le nom de la machine correspondante. Il enregistre les logs dans la base de données MySQL.

Les logs peuvent être visionnés sur l'interface utilisateur.

Ayant détecté une tentative d'attaque sur le serveur, on peut envoyer l'ordre au serveur de bloquer cette IP (en lui envoyant une commande à exécuter).

## Fonctionnalités v2 du projet
- [ ] Agents Linux : code de base (Récupérer les nouveaux logs à intervalles réguliers dans une liste de fichiers définie)
- [ ] Agents Linux : code de base (Faire du programme un démon)
---
- [ ] Agents Windows : code de base (Récupérer les nouveaux logs à intervalles réguliers dans une liste de fichiers définie)
- [ ] Agents Windows : code de base (Script d'installation : compilation + planificateur de tâches)
- [ ] Agents Windows : code de base (Automatiser l'installation dans le planificateur de tâches)
---
- [ ] Indexer : code de base (Infrastructure Docker)
- [ ] Indexer : code de base (Connexion à la BDD MySQL)
- [ ] Indexer : code de base (Réception des logs)
- [ ] Indexer : code de base (Vérification du token)
- [ ] Indexer : code de base (Enregistrement des logs reçus dans la BDD si le token est valide)
---
- [ ] Interface utilisateur : code de base (Infrastructure Docker)
- [ ] Interface utilisateur : code de base (Connexion à la BDD SQLite3)
- [ ] Interface utilisateur : code de base (Connexion)
- [ ] Interface utilisateur : code de base (Création de compte)
- [ ] Interface utilisateur : code de base (Connexion à la BDD MySQL de l'indexer)
- [ ] Interface utilisateur : code de base (Création de token sur la BDD MySQL de l'indexer)
- [ ] Interface utilisateur : code de base (Récupération des logs sur la BDD MySQL de l'indexer)
- [ ] Interface utilisateur : code de base (UI minimale)
---
- [ ] Modules Agents : adapter le code des modules existants dans la v1

