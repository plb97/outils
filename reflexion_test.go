// Copyright 2017 plb97. All rights reserved.
package outils

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

var (
	li12   = []int{1, 2}
	li21   = []int{2, 1, 2}
	li23   = []int{2, 3}
	li34   = []int{3, 4}
	li123  = []int{1, 2, 3}
	li1234 = []int{1, 2, 3, 4}
	li1    = []int{1}
	li2    = []int{2}
	li3    = []int{3}
	li     = []int{}

	lsab = []string{"a", "b"}
)

func TestEnsInt(t *testing.T) {
	test := "TestEnsInt"
	fmt.Println(test)
	var (
		ei12   = Creer(li12)
		ei21   = Creer(li21)
		ei23   = Creer(li23)
		ei34   = Creer(li34)
		ei123  = Creer(li123)
		ei1234 = Creer(li1234)
		ei1    = Creer(li1)
		ei2    = Creer(li2)
		ei3    = Creer(li3)
		ei     = Creer(li)
		esab   = Creer(lsab)
	)

	{ // type
		attendu := reflect.Int
		obtenu := ei12.t.Kind()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // nombre
		attendu := len(li12)
		obtenu := ei12.Nombre()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei12.lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei21.lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := make([]int, len(li21))
		copy(attendu, li21)             // ([]{2,1,2})
		sort.Ints(attendu)              // ([]{1,2,2})
		obtenu := ei21.Lister().([]int) // ([]{1,2})
		if reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // egal
		attendu := ei21
		obtenu := ei12
		if attendu == obtenu { // pointeurs differents
			t.Errorf(test+": attendu %p != obtenu %p\n", ei12, obtenu)
		}
		if !attendu.egal(obtenu) { // ensembles egaux
			t.Errorf(test+": attendu %v == obtenu %v\n", true, false)
		}
	}
	{ // pas egal
		attendu := ei12
		obtenu := ei23
		if attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", false, true)
		}
	}
	{ // union
		attendu := ei12
		obtenu := union(ei12)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", true, false)
		}
	}
	{ // union
		attendu := ei12
		obtenu := union(ei12, ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // union
		attendu := ei123
		obtenu := union(ei12, ei23)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // union
		attendu := ei1234
		obtenu := union(ei12, ei23, ei34)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei12
		obtenu := intersection(ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei12
		obtenu := intersection(ei12, ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei2
		obtenu := intersection(ei12, ei23)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei2
		obtenu := intersection(ei23, ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei
		obtenu := intersection(ei34, ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei
		obtenu := intersection(ei21, ei23, ei34)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei
		obtenu := soustraction(ei12, ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei12
		obtenu := soustraction(ei12, nil)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		var attendu *ens_t
		obtenu := soustraction(nil, ei12)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei12
		obtenu := soustraction(ei21, ei34)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei1
		obtenu := soustraction(ei21, ei23)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei3
		obtenu := soustraction(ei23, ei21)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei23
		obtenu := soustraction(ei23, esab)
		if !attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}

}

func TestEnsIntInd(t *testing.T) {
	test := "TestEnsIntInd"
	fmt.Println(test)
	var (
		ei12   = Creer(li12)
		ei21   = Creer(li21)
		ei23   = Creer(li23)
		esab   = Creer(lsab)

		eei     = Creer([]*ens_t{})
		eei12   = Creer([]*ens_t{ei12, ei21, ei12})
		eei23   = Creer([]*ens_t{ei23})
		eei1223 = Creer([]*ens_t{ei12, ei23})
	)
	{ // vide
		attendu := false
		obtenu := eei12.Vide()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // nombre
		attendu := 1
		obtenu := eei12.Nombre()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := []*ens_t{ei12}
		obtenu := eei12.lister().([]*ens_t)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // union + egal
		attendu := eei1223
		obtenu := union(eei12, eei23)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection + egal
		attendu := eei
		obtenu := intersection(eei12, eei23)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = intersection(eei23, eei12)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection + egal
		attendu := eei23
		obtenu := intersection(eei1223, eei23)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = intersection(eei23, eei1223)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction + egal
		attendu := eei12
		obtenu := soustraction(eei1223, eei23)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		attendu = eei
		obtenu = soustraction(eei23, eei1223)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // ajouter
		attendu := Creer([]*ens_t{ei12,esab})
		obtenu := Creer([]*ens_t{ei12}).ajouter(esab)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
}

func TestEnsembleInt(t *testing.T) {
	test := "TestEnsembleInt"
	fmt.Println(test)
	var (
		ei12 Ensemble   = Creer(li12)
		ei21 Ensemble   = Creer(li21)
		ei23 Ensemble   = Creer(li23)
		ei34 Ensemble   = Creer(li34)
		ei123 Ensemble  = Creer(li123)
		ei1234 Ensemble = Creer(li1234)
		ei1 Ensemble    = Creer(li1)
		ei2 Ensemble    = Creer(li2)
		ei3 Ensemble    = Creer(li3)
		ei Ensemble     = Creer(li)
		esab Ensemble   = Creer(lsab)
	)

	{ // nombre
		attendu := len(li12)
		obtenu := ei12.Nombre()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei12.Lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei21.Lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := make([]int, len(li21))
		copy(attendu, li21)             // ([]{2,1,2})
		sort.Ints(attendu)              // ([]{1,2,2})
		obtenu := ei21.Lister().([]int) // ([]{1,2})
		if reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // egal
		attendu := ei21
		obtenu := ei12
		if attendu == obtenu { // pointeurs differents
			t.Errorf(test+": attendu %p != obtenu %p\n", ei12, obtenu)
		}
		if !attendu.Egal(obtenu) { // ensembles egaux
			t.Errorf(test+": attendu %v == obtenu %v\n", true, false)
		}
	}
	{ // pas egal
		attendu := ei12
		obtenu := ei23
		if attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", false, true)
		}
	}
	{ // union
		attendu := ei12
		obtenu := Union(ei12)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", true, false)
		}
	}
	{ // union
		attendu := ei12
		obtenu := Union(ei12, ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // union
		attendu := ei123
		obtenu := Union(ei12, ei23)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // union
		attendu := ei1234
		obtenu := Union(ei12, ei23, ei34)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei12
		obtenu := Intersection(ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei12
		obtenu := Intersection(ei12, ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei2
		obtenu := Intersection(ei12, ei23)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei2
		obtenu := Intersection(ei23, ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei
		obtenu := Intersection(ei34, ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection
		attendu := ei
		obtenu := Intersection(ei21, ei23, ei34)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei
		obtenu := Soustraction(ei12, ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei12
		obtenu := Soustraction(ei12, nil)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei12
		obtenu := Soustraction(ei21, ei34)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei1
		obtenu := Soustraction(ei21, ei23)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei3
		obtenu := Soustraction(ei23, ei21)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction
		attendu := ei23
		obtenu := Soustraction(ei23, esab)
		if !attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}

}

func TestEnsembleIntInd(t *testing.T) {
	test := "TestEnsembleIntInd"
	fmt.Println(test)
	var (
		ei12 Ensemble   = Creer(li12)
		ei21 Ensemble   = Creer(li21)
		ei23 Ensemble   = Creer(li23)
		esab Ensemble   = Creer(lsab)

		eei Ensemble     = Creer([]Ensemble{})
		eei12 Ensemble   = Creer([]Ensemble{ei12, ei21, ei12})
		eei23 Ensemble   = Creer([]Ensemble{ei23})
		eei1223 Ensemble = Creer([]Ensemble{ei12, ei23})
	)
	{ // vide
		attendu := false
		obtenu := eei12.Vide()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // nombre
		attendu := 1
		obtenu := eei12.Nombre()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := []Ensemble{ei12}
		obtenu := eei12.Lister().([]Ensemble)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // union + egal
		attendu := eei1223
		obtenu := Union(eei12, eei23)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection + egal
		attendu := eei
		obtenu := Intersection(eei12, eei23)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = Intersection(eei23, eei12)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // intersection + egal
		attendu := eei23
		obtenu := Intersection(eei1223, eei23)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = Intersection(eei23, eei1223)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // soustraction + egal
		attendu := eei12
		obtenu := Soustraction(eei1223, eei23)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		attendu = eei
		obtenu = Soustraction(eei23, eei1223)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // ajouter
		var attendu Ensemble = Creer([]Ensemble{ei12,esab})
		var obtenu Ensemble = Creer([]Ensemble{ei12}).ajouter(esab)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
}

