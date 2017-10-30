# Génération de scripts

Un script de generation de procédures cisco. Ce script permet de génèrer des fichier `.cisco` contenant les commandes requises pour configurer les machines décrites dans le fichier excel fournis.

### Utilisation

Vous devez vous assurer d'avoir installé *NodeJs* et *NPM* sur votre machine. Si ce n'est pas le cas, rendrez vous sur [le site officiel](https://nodejs.org) ou votre géstionnaire de paquets favoris pour l'installer. Ce script a été codé et testé pour la version *7.0.0* de Node.

Commencez par cloner ce dépot puis, apres vous être placé à la racine du projet installez les dépendances.

```
npm install
```

Configurez le programme en suivant les instructions décrites en dessous puis lancez la génération.

```
npm run generate
```

L'ensemble de vos scripts seront génerés dans le dossier `scripts`

### Configuration

La configuration s'éfféctue dans le fichier json `settings.json` et possède les paramètres suivants : 

* `srcfile` : Une *chaine de carractères* représentant le fichier xlsx utilé pour la génération
* `autosave` : Un *booleen* définissant si la commande `write mem` doit être éxécutée en fin de script
* `vtpclient` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *booleen* activant la configuration du client vtp sur toutes les machines
    - `domain` : Une *chaine de carractères* repréentant le nom de domain vtp
    - `password` : Une *chaine de carractères* représentant le mot de passe vtp
* `banner` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant la définition de la bannière de bienvenue
    - `text` : Une *chaine de carractères* définissant le contenu de la banière (**Ne doit pas contenir le signe "#"**)
* `domain` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant la définition du nom de domaine des machines
    - `domain` : Une *chaine de carractères* définissant le nom de domaine
* `enablesecret` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant le mot de passe lors du passage en mode privilègié
    - `password` : Une *chaine de carractères* définissant le mot de passe
* `consolesecret` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant le mot de passe lors de la connexion console
    - `password` : Une *chaine de carractères* définissant le mot de passe
* `telnet` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant le protocole de communication telnet
    - `password` : Une *chaine de carractères* définissant le mot de passe
* `ssh` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant le protocole de communication ssh en version 2
    - `keylength` : Un *nombre entier* définissant la taille de la clé a génerer (**Toute clé doit précédemment être supprimée sur la machine**)
    - `timeout` : Un *nombre entier* définissant le temps maximum de réponse ssh
    - `loginAttempts` : Un *nombre entier* définissant le nombre maximum d'echecs autorisés
* `admin` : Un *objet* contenant les paramètres suivants :
    - `enable` : Un *boolean* activant la connexion par utilisateur et mots de passe (**Obligatire avec l'utilisation de ssh**)
    - `username` : Une *chaine de carractères* définissant le nom d'utilisateur
    - `password` : Une *chaine de carractères* définissant le mot de passe
