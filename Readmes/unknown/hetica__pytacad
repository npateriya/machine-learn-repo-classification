# Pytacad

Ce script à été créé pour retrouver **rapidement** des informations sur un ou plusieurs étudiants du site netacad.com.  
Cisco ne fournit pas de champs de recherche, si bien qu'il est plus que fastidieux de retrouver un stagiaire dans la base.  
Et il est encore plus compliqué de retrouver les classes sur lesquelles il est inscrit.

##Deux parties 

### pytacad-server  

Son rôle est de parser toutes les classes d'une académie Cisco sur le site netacad.com, et d'en extraire des informations :

* chaque classe est stockée dans un fichier texte
* un fichier contient l'ensemble des stagiaires, et les classes auxquels ils appartiennent  

On pourra exécuter le serveur une ou plusieurs fois par jour à l'aide de la crontab

### pytacad-client  

Permet d'afficher des informations extraire par pytacad-server :  

* chercher un ou des stagiaires ou classes  
* mettre à jour la base de données (exécute pytacad-server)  
* afficher des infos générales (nombre de classes/stagiaires)  

##webeni

Le projet webeni est complémentaire de pytacad, il fournit une application web, plus conviviale pour chercher
des informations stockées par pytacad.
