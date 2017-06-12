// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENSE.md (French) or LICENSE_EN.md (English) file.
package outils

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)


func TestListe(t *testing.T) {
	test := "TestListe"
	fmt.Println(test)

	var (
		mis = map[int]string{4:"quatre",2:"deux",3:"trois",1:"un",}
		li = []int{2,4,3,1,} // pas ordonnee (ordre alphabetique)
		msi = map[string]int{"deux":2,"quatre":4,"trois":3,"un":1,}
		ls = []string{"un","deux","trois","quatre"} // pas ordonnee (ordre numerique)
		mfs = map[float64]string{4:"quatre",2:"deux",3:"trois",1:"un",}
		lf = []float64{2,4,3,1,} // pas ordonnee (ordre alphabetique)
	)

	{ // liste 'int'
		attendu := make([]int,len(li))
		copy(attendu,li)
		obtenu := Lister_cles(mis).([]int) // ordonnee
		if reflect.DeepEqual(attendu, obtenu) { // attendu pas triee
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
		sort.Ints(attendu)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = Lister_cles_int(mis)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		for i, k := range obtenu {
			if mis[k] != ls[i] {
				t.Errorf(test+": attendu %v != obtenu %v\n", ls[i], mis[k])
			}
		}
	}

	{ // liste 'string'
		attendu := make([]string,len(ls))
		copy(attendu,ls)
		obtenu := Lister_cles(msi).([]string) // ordonnee
		if reflect.DeepEqual(attendu, obtenu) { // attendu pas triee
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
		sort.Strings(attendu)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = Lister_cles_string(msi)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		for i, k := range obtenu {
			if msi[k] != li[i] {
				t.Errorf(test+": attendu %v != obtenu %v\n", ls[i], msi[k])
			}
		}
	}

	{ // liste 'float64'
		attendu := make([]float64,len(lf))
		copy(attendu,lf)
		obtenu := Lister_cles(mfs).([]float64) // ordonnee
		if reflect.DeepEqual(attendu, obtenu) { // attendu pas triee
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
		sort.Float64s(attendu)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		obtenu = Lister_cles_float64(mfs)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		for i, k := range obtenu {
			if mfs[k] != ls[i] {
				t.Errorf(test+": attendu %v != obtenu %v\n", ls[i], mfs[k])
			}
		}
	}

}

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

// *ens_t

func TestEns(t *testing.T) {
	test := "TestEns"
	fmt.Println(test)
	var (
		ei12   = creer(li12)
		ei21   = creer(li21)
		ei23   = creer(li23)
		ei34   = creer(li34)
		ei123  = creer(li123)
		ei1234 = creer(li1234)
		ei1    = creer(li1)
		ei2    = creer(li2)
		ei3    = creer(li3)
		ei     = creer(li)
		esab   = creer(lsab)
	)

	{ // type
		attendu := reflect.Int
		obtenu := ei12.t.Kind()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // nombre
		attendu := len(li12)
		obtenu := ei12.Nombre()
		if attendu != obtenu {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei12.lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei21.lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := make([]int, len(li21))
		copy(attendu, li21)             // ([]{2,1,2})
		sort.Ints(attendu)              // ([]{1,2,2})
		obtenu := ei21.Lister().([]int) // ([]{1,2})
		if reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // egal
		attendu := ei21
		obtenu := ei12
		if attendu == obtenu { // pointeurs differents
			t.Errorf(test+": attendu %p == obtenu %p\n", ei12, obtenu)
		}
		if !attendu.egal(obtenu) { // ensembles egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", true, false)
		}
	}
	{ // copier + egal
		attendu := ei21
		obtenu := ei21.copier()
		if attendu == obtenu { // pointeurs differents
			t.Errorf(test+": attendu %p == obtenu %p\n", ei12, obtenu)
		}
		if !attendu.egal(obtenu) { // ensembles egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", true, false)
		}
	}
	{ // pas egal
		attendu := ei12
		obtenu := ei23
		if attendu.egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v == obtenu %v\n", false, true)
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
	//{ // soustraction
	//	var attendu *ens_t
	//	obtenu := soustraction(nil, ei12)
	//	if !attendu.egal(obtenu) { // ensembles pas egaux
	//		t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
	//	}
	//}
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

	{ // appeler
		mr := make(map[int]int)
		f := func(i int) {
			j := i+2
			mr[i] = j
		}
		ei23.appeler(f)
		for i, obtenu := range mr {
			attendu := i+2
			if attendu != obtenu {
				t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
			}
		}
	}
}

func TestEnsInd(t *testing.T) {
	test := "TestEnsInd"
	fmt.Println(test)
	var (
		ei12   = creer(li12)
		ei21   = creer(li21)
		ei23   = creer(li23)
		esab   = creer(lsab)

		eei     = creer([]*ens_t{})
		eei12   = creer([]*ens_t{ei12, ei21, ei12})
		eei23   = creer([]*ens_t{ei23})
		eei1223 = creer([]*ens_t{ei12, ei23})
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
		attendu := creer([]*ens_t{ei12,esab})
		obtenu := creer([]*ens_t{ei12}).ajouter(esab)
		if !obtenu.egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
}

func TestEnsPanic(t *testing.T) {
	test := "TestEnsPanic"
	fmt.Println(test)
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "creer"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		creer(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "ajouter"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := creer(li12)
		ei12.ajouter(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "unir"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := creer(li12)
		ei12.unir(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "egal"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := creer(li12)
		ei12.egal(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Intersection"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		intersection(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Union"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
				fmt.Println(fmt.Sprintf("Panic dans %#v", r))
			}
		}()
		union(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "soustraction"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		soustraction(nil)
	}()

}

// Ensemble

func TestEnsemble(t *testing.T) {
	test := "TestEnsemble"
	fmt.Println(test)
	var (
		ei12    = Creer(li12)
		ei21    = Creer(li21)
		ei23    = Creer(li23)
		ei34    = Creer(li34)
		ei123   = Creer(li123)
		ei1234  = Creer(li1234)
		ei1     = Creer(li1)
		ei2     = Creer(li2)
		ei3     = Creer(li3)
		ei      = Creer(li)
		esab    = Creer(lsab)
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
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := li12
		obtenu := ei21.Lister().([]int)
		if !reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
	{ // liste
		attendu := make([]int, len(li21))
		copy(attendu, li21)             // ([]{2,1,2})
		sort.Ints(attendu)              // ([]{1,2,2})
		obtenu := ei21.Lister().([]int) // ([]{1,2})
		if reflect.DeepEqual(attendu, obtenu) {
			t.Errorf(test+": attendu %v == obtenu %v\n", attendu, obtenu)
		}
	}
	{ // egal
		attendu := ei21
		obtenu := ei12
		if attendu == obtenu { // pointeurs differents
			t.Errorf(test+": attendu %p == obtenu %p\n", ei12, obtenu)
		}
		if !attendu.Egal(obtenu) { // ensembles egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", true, false)
		}
	}
	{ // Copier + egal
		attendu := ei21
		obtenu := ei21.Copier()
		if attendu == obtenu { // pointeurs differents
			t.Errorf(test+": attendu %p == obtenu %p\n", ei12, obtenu)
		}
		if !attendu.Egal(obtenu) { // ensembles egaux
			t.Errorf(test+": attendu %v != obtenu %v\n", true, false)
		}
	}
	{ // pas egal
		attendu := ei12
		obtenu := ei23
		if attendu.Egal(obtenu) { // ensembles pas egaux
			t.Errorf(test+": attendu %v == obtenu %v\n", false, true)
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

	{ // appeler
		mr := make(map[int]int)
		f := func(i int) int {
			j := i*2 + 1
			mr[i] = j
			return j
		}
		ei23.Appeler(f)
		for i, obtenu := range mr {
			attendu := f(i)
			if attendu != obtenu {
				t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
			}
		}
	}
}

func TestEnsembleInd(t *testing.T) {
	test := "TestEnsembleInd"
	fmt.Println(test)
	var (
		ei12    = Creer(li12)
		ei21    = Creer(li21)
		ei23    = Creer(li23)
		esab    = Creer(lsab)

		eei     = Creer([]Ensemble{})
		eei12   = Creer([]Ensemble{ei12, ei21, ei12})
		eei23   = Creer([]Ensemble{ei23})
		eei1223 = Creer([]Ensemble{ei12, ei23})
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
		attendu := Creer([]Ensemble{ei12,esab})
		obtenu := Creer([]Ensemble{ei12}).Ajouter(esab)
		if !obtenu.Egal(attendu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
		if !attendu.Egal(obtenu) {
			t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
		}
	}
}

func TestEnsemblePanic(t *testing.T) {
	test := "TestEnsemblePanic"
	fmt.Println(test)
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Creer"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		Creer(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Egal"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Egal(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Contient"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Contient(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Ajouter"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Ajouter(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Retirer"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Retirer(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Unir"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Unir(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Intersecter"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Intersecter(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Soustraire"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Soustraire(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Appeler"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		ei12    := Creer(li12)
		ei12.Appeler(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Intersection"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		Intersection(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Union"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		Union(nil)
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				attendu := "Soustraction"
				obtenu, ok := r.(string)
				if !ok {
					t.Errorf(test+": attendu %v != obtenu %v\n", !ok, ok)
				}
				if attendu != obtenu {
					t.Errorf(test+": attendu %v != obtenu %v\n", attendu, obtenu)
				}
			}
		}()
		Soustraction(nil)
	}()

}
