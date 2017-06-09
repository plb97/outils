# outils

Le point de départ de cete exemple est le besoint de récupération de la liste des clés d'une 'map' quelque soit sont type.
Pour cela l'utilisation du package 'reflect' s'est imposée.

La fonction **Lister_cles**('map') retourne la liste des clés de la table passée en parametre.
Si c'est possible, la liste est triée par ordre croissant.


De là, l'idée d'utiliser ce même package pour introduire la notion d'ensembles comme dans le langage Pascal a émergé.
Cette notion peut s'appiquer au types simples mais aussi aux structures, tableaux et autres.
En particulier, cette notion peut s'appliquer à elle-même (des ensembles d'ensembles).
La fonction *Creer*([]'type') crée un ensemble d'éléments de type 'type'.

Les actions suivantes peuvent s'effecteur sur un 'Ensemble' :

**Methode** | Commentaire
----------- | -----------
**Ajouter**(le ...interface{}) Ensemble | ajouter les éléments passés en paramètre
**Retirer**(le ...interface{}) Ensemble | retirer les éléments passés en paramètre
**Lister**() interface{} | lister les élément dans un ordre déterministe
**Contient**(i interface{}) bool | vérifier si l'élément passé en parametre appartient à l'ensemble
**Nombre**() int | récupérer le nombre d'éléments de l'ensemble
**Vide**() bool | vérifier si l'ensemble est vide
**Egal**(x Ensemble) bool | comparer l'enemble passé en paratètre à l'ensemble
**Unir**(x Ensemble) Ensemble | unir l'ensemble passé en paramètre à l'ensemble
**Soustraire**(x Ensemble) Ensemble | soustraire l'ensemble passé en paramètre à l'ensemble
**Intersecter**(x Ensemble) Ensemble | croiser l'ensemble passé en paramètre avec l'ensemble
**Appeler**(i interface{}) interface{} | appeler la fonction passée en paramètre pour chaque élément de l'ensemble

