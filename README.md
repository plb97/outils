# Exemples d'utilisation du langage **Go** (version 1.8 minimum) de *Google*

Le point de départ de ces **exemples** est le besoin de récupération de la liste des clés d'une *'map'* quel que soit sont type.

Pour cela l'utilisation du package __*'reflect'*__ s'est imposée.
Les conclusions des nombreux tâtonnements autour de l'utilisation de ce package sont données ici.

## Liste des clés d'une 'map'
La fonction **Lister_cles**('map') retourne sous forme d'__*interface*__ la liste des clés de la table passée en paramètre.
Si c'est possible, la liste est triée par ordre croissant.
Le résultat obtenu doit être "casté" par exemple : liste.([]int) pour une liste d'entiers ou liste.([]float64) pour une liste de réels.

En complément, trois fonctions spécialisées :

**Fonction** | Commentaire
------------ | -----------
**Lister_cles_string**(i interface{}) []string | liste des clés *'string'*
**Lister_cles_int**(i interface{}) []int | liste des clés *'int'*
**Lister_cles_float64**(i interface{}) []float64 | liste des clés *'float64'*


## Ensemble d'éléments de même 'type'
De là, l'idée d'utiliser ce même package pour introduire la notion d'ensembles - comme dans le langage Pascal - a émergé en réponse à un besoin et s'est avérée intéressante.

Cette notion peut s'appliquer au types simples mais aussi aux structures, tableaux et autres.
En particulier, elle peut s'appliquer à elle-même (des ensembles d'ensembles).

La fonction **Creer**([]'type') crée un ensemble d'éléments de type 'type'.

Les actions suivantes peuvent s'effecteur sur un *'Ensemble'* :

**Methode** | Commentaire
----------- | -----------
**Ajouter**(le ...interface{}) Ensemble | ajouter les éléments passés en paramètre à l'ensemble
**Retirer**(le ...interface{}) Ensemble | retirer les éléments passés en paramètre à l'ensemble
**Lister**() interface{} | lister les éléments de l'ensemble dans un ordre déterministe (croissant si possible)
**Contient**(i interface{}) bool | vérifier si l'élément passé en paramètre appartient à l'ensemble
**Nombre**() int | récupérer le nombre d'éléments de l'ensemble
**Vide**() bool | vérifier si l'ensemble est vide
**Egal**(x Ensemble) bool | comparer l'ensemble passé en paramètre à l'ensemble
**Unir**(x Ensemble) Ensemble | unir l'ensemble passé en paramètre à l'ensemble
**Soustraire**(x Ensemble) Ensemble | soustraire l'ensemble passé en paramètre de l'ensemble
**Intersecter**(x Ensemble) Ensemble | croiser l'ensemble passé en paramètre avec l'ensemble
**Appeler**(i interface{}) interface{} | appeler la fonction passée en paramètre pour chaque élément de l'ensemble