// Copyright 2017 plb97. All rights reserved.
package outils

import (
	"fmt"
	"reflect"
	"sort"
)

// la fonction 'lister_cles' relourne la liste des cles d'une 'map' quelconque
// si cela est possible les cles sont triees par ordre croissant
func lister_cles(i interface{}) interface{} {
	// controles
	if nil == i {
		return nil
	}
	ti := reflect.TypeOf(i) // recuperer le 'type' du parametre
	if reflect.Map != ti.Kind() {
		panic("lister_cles : le parametre n'est pas de type 'map'")
	}
	tik := ti.Key() // recuperer le 'type' des cles de la table ('map')

	// collecte des cles
	vi := reflect.ValueOf(i)                                   // recuperer la 'value' du parametre
	lk := reflect.MakeSlice(reflect.SliceOf(tik), 0, vi.Len()) // creer la liste des cles a retourner (Value)
	lmk := vi.MapKeys()                                        // liste des cles ([]Value)
	for _, k := range lmk {                                    // transformer 'lmk' de type '[]Value' en 'lk' de type 'Value'
		lk = reflect.Append(lk, k)
	}

	// tri des cles (si possible)
	// fonctions de comparaison dependant du 'type' des cles
	compString := func(ki, kj interface{}) bool { return ki.(string) < kj.(string) }
	compInt := func(ki, kj interface{}) bool { return ki.(int) < kj.(int) }
	compInt8 := func(ki, kj interface{}) bool { return ki.(int8) < kj.(int8) }
	compInt16 := func(ki, kj interface{}) bool { return ki.(int16) < kj.(int16) }
	compInt32 := func(ki, kj interface{}) bool { return ki.(int32) < kj.(int32) }
	compInt64 := func(ki, kj interface{}) bool { return ki.(int64) < kj.(int64) }
	compUint := func(ki, kj interface{}) bool { return ki.(uint) < kj.(uint) }
	compUint8 := func(ki, kj interface{}) bool { return ki.(uint8) < kj.(uint8) }
	compUint16 := func(ki, kj interface{}) bool { return ki.(uint16) < kj.(uint16) }
	compUint32 := func(ki, kj interface{}) bool { return ki.(uint32) < kj.(uint32) }
	compUint64 := func(ki, kj interface{}) bool { return ki.(uint64) < kj.(uint64) }
	compFloat32 := func(ki, kj interface{}) bool { return ki.(float32) < kj.(float32) }
	compFloat64 := func(ki, kj interface{}) bool { return ki.(float64) < kj.(float64) }
	// 'closure' retournant la fonction de comparaison utilisee lors de l'appel a 'sort.SliceStable'
	less := func(comp func(ki, kj interface{}) bool) func(i, j int) bool {
		return func(i, j int) bool {
			ki := lk.Index(i).Interface() // cle 'i'
			kj := lk.Index(j).Interface() // cle 'j'
			return comp(ki, kj)           // retourne la comparaison
		}
	}

	lki := lk.Interface() // 'interface' representant la liste des cles
	switch reflect.TypeOf(lki).Elem().Kind() {
	//// cas standards prevus dans le package 'sort'
	//case reflect.String:  sort.Strings(lki.([]string));
	//case reflect.Int:     sort.Ints(lki.([]int))
	//case reflect.Float64: sort.Float64s(lki.([]float64))
	// les cles doivent etre de 'type' "ordonnable"
	// c-a-d que 'x' < 'y' est autorise par le langage
	// par exemple : '1' < '0' est autorise mais 'false' < 'true' ne l'est pas
	case reflect.String:
		sort.SliceStable(lki, less(compString))
	case reflect.Int:
		sort.SliceStable(lki, less(compInt))
	case reflect.Int8:
		sort.SliceStable(lki, less(compInt8))
	case reflect.Int16:
		sort.SliceStable(lki, less(compInt16))
	case reflect.Int32:
		sort.SliceStable(lki, less(compInt32))
	case reflect.Int64:
		sort.SliceStable(lki, less(compInt64))
	case reflect.Uint:
		sort.SliceStable(lki, less(compUint))
	case reflect.Uint8:
		sort.SliceStable(lki, less(compUint8))
	case reflect.Uint16:
		sort.SliceStable(lki, less(compUint16))
	case reflect.Uint32:
		sort.SliceStable(lki, less(compUint32))
	case reflect.Uint64:
		sort.SliceStable(lki, less(compUint64))
	case reflect.Float32:
		sort.SliceStable(lki, less(compFloat32))
	case reflect.Float64:
		sort.SliceStable(lki, less(compFloat64))
	}

	return lki // retourner la liste des cles sous forme d' 'interface' "castable" par lki.([]'type')
	// par exemple : lki.([]int) si les cles sont des 'int' ou lki.([]string) si les cles sont des 'string'
}

var Lister_cles = lister_cles

func Lister_cles_string(i interface{}) []string {
	if lk := lister_cles(i); nil == lk {
		return nil
	} else {
		return lk.([]string)
	}
}
func Lister_cles_int(i interface{}) []int {
	if lk := lister_cles(i); nil == lk {
		return nil
	} else {
		return lk.([]int)
	}
}
func Lister_cles_float64(i interface{}) []float64 {
	if lk := lister_cles(i); nil == lk {
		return nil
	} else {
		return lk.([]float64)
	}
}

// REMARQUE : 'Renverser' n'est pas equivalente a sort.Reverse
// sauf losque la liste 'i' en entree est triee
func Renverser(i interface{}) {
	if nil == i {
		return
	}
	iv := reflect.ValueOf(i)
	if reflect.Slice != iv.Kind() {
		panic("Renverser : le parametre n'est pas de type 'slice'")
	}
	l := iv.Len()
	swap := reflect.Swapper(i)
	m, n := l/2, l-1
	for k := 0; k < m; k++ {
		swap(k, n-k)
	}
}

var nul = &ens_t{t: reflect.TypeOf(nil)}
var nulle = reflect.Value{}

// la fonction 'ident' retourne une chaine identifiant l'interface
// REMARQUE : peut etre amelioree...
func ident(i interface{}) string {
	return fmt.Sprintf("%T:%v", i, i)
}

// la fonction 'conv' convertit le parametre 'x' en pointeur *ens_t
func conv(x Ensemble) *ens_t {
	if nil == x {
		return nul
	}
	px, ok := x.(*ens_t)
	if !ok {
		panic("conv")
	}
	return px
}

// 'Ensemble' represente un ensemble d'elements de meme type
// sur lesquels (les ensembles) peuvent s'effectuer des operations
type Ensemble interface {
	Ajouter(le ...interface{}) Ensemble
	Retirer(le ...interface{}) Ensemble
	Lister() interface{}
	Contient(i interface{}) bool
	Nombre() int
	Vide() bool
	Egal(x Ensemble) bool
	Unir(x Ensemble) Ensemble       // comparable a Ajouter
	Soustraire(x Ensemble) Ensemble // comparable a Retirer
	Intersecter(x Ensemble) Ensemble
	Appeler(i interface{}) interface{}

	String() string
}

// 'ens_t' est une structure concrete pour les ensembles
type ens_t struct {
	t   reflect.Type   // type des elements
	ind bool           // indicateur d'indirection
	m   reflect.Value  // table ('map') representant l'ensemble
	msi map[string]int // table de correspondance id -> index (utilisee lorsqu'il y a indirection)
}

// la fonction new_ens_t' retourne un pointeur '*ens_t' sur un ensemble vide
// dont les elements doivent etre de type 't'
// REMARQUE : les ensembles vides sont types contrairement a l'ensemble vide des mathematiques
func new_ens_t(ind bool, t reflect.Type) *ens_t {
	var m reflect.Value
	var msi map[string]int
	if ind { // en cas d'indirection
		msi = make(map[string]int)
		m = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(0), t))
	} else { // pas d'indirection
		m = reflect.MakeMap(reflect.MapOf(t, reflect.TypeOf(true)))
	}
	e := ens_t{t: t, ind: ind, m: m, msi: msi}
	return &e
}
// la fonction 'ajouter' permet d'ajouter un element a l'ensemble
func (pe *ens_t) ajouter(i interface{}) *ens_t {
	if nil == i {
		panic("ajouter")
	}
	v := reflect.ValueOf(i)
	if !v.Type().ConvertibleTo(pe.t) {
		panic(fmt.Sprintf("ajouter : pe.t=%v v=%v",pe.t.Kind(), v.Kind()))
	}
	v = v.Convert(pe.t)
	if pe.ind {
		k := ident(i)
		if _, ok := pe.msi[k]; !ok {
			ki := pe.m.Len() // nouvel index
			pe.m.SetMapIndex(reflect.ValueOf(ki), v)
			pe.msi[k] = ki
		}
	} else {
		pe.m.SetMapIndex(v, reflect.ValueOf(true))
	}
	return pe
}

// la fonction 'ajouter_liste' permet d'ajouter tous les elements d'une liste a l'ensemble
// REMARQUE : la fonction 'ajouter' pourrait etre rendue "elliptique" mais cela ne serait pas equivalent
//            car une liste d'interfaces n'est pas une interface liste
func (pe *ens_t) ajouter_liste(vl reflect.Value) *ens_t {
	for i := 0; i < vl.Len(); i++ {
		pe.ajouter(vl.Index(i).Interface())
	}
	return pe
}

// la fonction 'copier' permet de "cloner" un ensemble
// REMARQUE : les deux ensembles sont egaux mais ils ne sont pas le meme (les pointeurs *ens_t sont differents)
func (pe *ens_t) copier() *ens_t {
	px := new_ens_t(pe.ind,pe.t)
	vl := reflect.ValueOf(pe.lister())
	for i := 0; i < vl.Len(); i++ {
		elmt := vl.Index(i).Interface()
		px.ajouter(elmt)
	}
	return px
}

// la fonction 'retirer' permet de retirer un element de l'ensemble
func (pe *ens_t) retirer(v reflect.Value) *ens_t {
	if pe.t.Kind() != v.Kind() {
		return pe
	}
	if pe.ind {
		k := ident(v.Interface())
		if ki, ok := pe.msi[k]; ok {
			pe.m.SetMapIndex(reflect.ValueOf(ki), nulle)
			delete(pe.msi, k)
		}
	} else {
		pe.m.SetMapIndex(v, nulle)
	}
	return pe
}

// la fonction 'lister' permet de lister les elements de l'ensemble (dans un ordre deterministe)
func (pe *ens_t) lister() interface{} {
	if nil == pe {
		return nil
	}
	if pe.ind {
		li := lister_cles(pe.m.Interface()).([]int)
		tik := pe.m.Type().Elem()
		lk := reflect.MakeSlice(reflect.SliceOf(tik), 0, pe.m.Len())
		for _, i := range li {
			k := pe.m.MapIndex(reflect.ValueOf(i))
			lk = reflect.Append(lk, k)
		}
		return lk.Interface()
	} else {
		return lister_cles(pe.m.Interface())
	}
}

// la fonction 'contient' permet de verifier si l'element 'v' appartient a l'ensemble
func (pe *ens_t) contient(v reflect.Value) bool {
	if nil == pe {
		return false
	}
	if pe.t.Kind() != v.Kind() {
		return false
	}
	if pe.ind {
		k := ident(v.Interface())
		_, ok := pe.msi[k]
		return ok // REMARQUE : l'egalite des identifiants implique l'egalite des elements eux-memes
	} else {
		ok := reflect.Invalid != pe.m.MapIndex(v).Kind()
		return ok
	}
}

// la fonction 'nombre' retourne le nombre d'elements (cardinal) de l'ensemble
func (pe *ens_t) nombre() int {
	if nil == pe {
		return 0
	}
	return pe.m.Len()
}

// la fonction 'vide' retourne 'true' si l'ensemble est vide ou 'nil' et 'false' sinon
func (pe *ens_t) vide() bool {
	if nil == pe {
		return true
	}
	return 0 == pe.m.Len()
}
// la fonction 'egal' permet de verifier l'egalite de l'ensemble 'px' et de l'ensemble 'pe'
// REMARQUE : reflect.DeepEqual(pe,px) ne donne pas le resultat voulu
func (pe *ens_t) egal(px *ens_t) bool {
	if nil == px {
		return pe == px // vrai si pe == nil faux sinon
	}
	if nil == pe {
		return false
	}
	ok := pe.t.Kind() == px.t.Kind() && pe.m.Len() == px.m.Len()
	// REMARQUE : deux ensembles sont egaux s'il ont le meme type d'elements y compris les ensembles vides
	if ok {
		if pe.ind {
			// REMARQUE : tout repose sur l'identifiant associe a l'element
			if len(pe.msi) != len(px.msi) {
				panic("egal") // TODO ne doit pas arriver
			}
			if ok {
				for k := range pe.msi {
					if _, ok := px.msi[k]; !ok {
						ok = false
						break
					}
				}
			}
		} else {
			le := reflect.ValueOf(pe.Lister())
			lx := reflect.ValueOf(px.Lister())
			if le.Len() != lx.Len() {
				panic("egal") // TODO ne doit pas arriver
			}
			if ok {
				n := le.Len()
				if n != lx.Len() {
					return false
				}
				for i := 0; i < n; i++ {
					if le.Index(i).Interface() != lx.Index(i).Interface() {
						ok = false
						break
					}
				}
			}
		}
	}
	return ok
}

// la fonction 'unir' permet d'ajouter les elements de l'ensemble 'px' a l'ensemble 'pe'
func (pe *ens_t) unir(px *ens_t) *ens_t {
	if nil == pe {
		panic("unir")
	}
	if nil == px {
		return pe
	}
	if pe.t.Kind() != px.t.Kind() {
		panic("unir")
	}
	for _, elmt := range px.m.MapKeys() {
		if px.ind {
			elmt = px.m.MapIndex(elmt)
		}
		if !pe.contient(elmt) {
			pe.ajouter(elmt.Interface())
		}
	}
	return pe
}

// la fonction 'intersecter' permet de retirer les elements de l'ensemble 'pe' qui n'appartienent pas a l'ensemble 'px'
func (pe *ens_t) intersecter(px *ens_t) *ens_t {
	if nil == pe {
		panic("intersecter")
	}
	if nil == px {
		return pe
	}
	if pe.t.Kind() != px.t.Kind() {
		return pe
	}
	for _, elmt := range pe.m.MapKeys() {
		if pe.ind {
			elmt = pe.m.MapIndex(elmt)
		}
		if !px.contient(elmt) {
			pe.retirer(elmt)
		}
		if 0 == pe.m.Len() {
			break // inutile de continuer
		}
	}
	return pe
}

// la fonction 'soustraire' permet de retirer les elements de l'ensemble 'pe' qui appartienent aussi a l'ensemble 'px'
func (pe *ens_t) soustraire(px *ens_t) *ens_t {
	if nil == pe || nil == pe.t {
		panic("soustraire")
	}
	if nil == px || nil == px.t {
		return pe
	}
	if pe.t.Kind() != px.t.Kind() {
		return pe
	}
	if pe.m.Len() < px.m.Len() {
		for _, elmt := range pe.m.MapKeys() {
			if pe.ind {
				elmt = pe.m.MapIndex(elmt)
			}
			if px.contient(elmt) {
				pe.retirer(elmt)
			}
			if 0 == pe.m.Len() {
				break // inutile de continuer
			}
		}
	} else {
		for _, elmt := range px.m.MapKeys() {
			if px.ind {
				elmt = px.m.MapIndex(elmt)
			}
			if pe.contient(elmt) {
				pe.retirer(elmt)
			}
			if 0 == pe.m.Len() {
				break // inutile de continuer
			}
		}
	}
	return pe
}

// la fonction 'Creer' permet de creer un ensemble dont le type correspond a la liste (eventuellemnt vide)
// representee par l'interface passee en parametre
func Creer(i interface{}) *ens_t {
	vi := reflect.ValueOf(i)
	if reflect.Slice != vi.Kind() {
		panic("Creer")
	}
	t := vi.Type().Elem()
	ind := !t.Comparable() ||
		reflect.Ptr == t.Kind() ||
		reflect.Interface == t.Kind() // indicateur d'indirection
	e := new_ens_t(ind,t).ajouter_liste(vi)
	return e
}

// la fonction 'Ajouter' permet d'ajouter des elements a l'ensemble
func (pe *ens_t) Ajouter(li ...interface{}) Ensemble {
	if nil == pe {
		panic("Ajouter")
	}
	for _, i := range li {
		pe.ajouter(reflect.ValueOf(i).Interface())
	}
	return pe
}

// la fonction 'Retirer' permet de retirer des elements de l'ensemble
func (pe *ens_t) Retirer(li ...interface{}) Ensemble {
	if nil == pe {
		panic("Retirer")
	}
	for _, i := range li {
		pe.retirer(reflect.ValueOf(i))
		if 0 == pe.m.Len() {
			break // iuntile de continuer
		}
	}
	return pe
}

// la fonction 'Lister' permet de lister les elements de l'ensemble dans un ordre deterministe
func (pe *ens_t) Lister() interface{} {
	return pe.lister()
}

// la fonction 'Contient' permet de verifier si l'element 'i' appartient a l'ensemble
func (pe *ens_t) Contient(i interface{}) bool {
	return pe.contient(reflect.ValueOf(i))
}

// la fonction 'Nombre' retourne le nombre d'elements (cardinal) de l'enmbme l'ensemble
func (pe *ens_t) Nombre() int {
	return pe.nombre()
}

// la fonction 'Vide' retourne 'true' si l'ensemble est vide ou 'nil' et 'false' sinon
func (pe *ens_t) Vide() bool {
	return pe.vide()
}

// la fonction 'Egal' permet de verifier l'egalite de l'Ensemble 'x' et de l'ensemble 'pe'
func (pe *ens_t) Egal(x Ensemble) bool {
	px := conv(x)
	return pe.egal(px)
}

// la fonction 'Unir' permet d'ajouter les elements de l'ensemble 'x' a l'ensemble 'pe'
func (pe *ens_t) Unir(x Ensemble) Ensemble {
	px := conv(x)
	return pe.unir(px)
}

// la fonction 'Intersecter' permet de retirer les elements de l'ensemble 'pe' qui n'appartienent pas a l'ensemble 'x'
func (pe *ens_t) Intersecter(x Ensemble) Ensemble {
	px := conv(x)
	return pe.intersecter(px)
}

// la fonction 'Soustraire' permet de retirer les elements de l'ensemble 'pe' qui appartienent aussi a l'ensemble 'x'
func (pe *ens_t) Soustraire(x Ensemble) Ensemble {
	px := conv(x)
	return pe.soustraire(px)
}

// la fonction 'Appeler' permet pour chaque element de l'ensemble d'appeler la fonction representee par 'i'
func (pe *ens_t) Appeler(i interface{}) interface{} {
	if nil == pe || nil == i {
		panic("Appliquer")
	}
	vf := reflect.ValueOf(i)
	tf := vf.Type()
	if reflect.Func != tf.Kind() {
		panic("Appeler")
	}
	ne := tf.NumIn()
	lx := reflect.ValueOf(pe.Lister())
	if 1 != ne {
		panic("Appliquer")
	}
	ns := tf.NumOut()
	var ts reflect.Type
	switch ns {
	case 0: // rien
	case 1:
		ts = tf.Out(0)
	default:
		lst := make([]reflect.StructField, 0)
		for i := 0; i < ns; i++ {
			lst = append(lst, reflect.StructField{Name: fmt.Sprintf("R%v", i), Type: tf.Out(i)})
		}
		ts = reflect.StructOf(lst)
	}

	switch ns {
	case 0:
		for i := 0; i < lx.Len(); i++ {
			in := reflect.ValueOf(lx.Index(i).Interface())
			vf.Call([]reflect.Value{in})
		}
		return nil
	case 1:
		ls := reflect.MakeSlice(reflect.SliceOf(ts), 0, lx.Len())
		for i := 0; i < lx.Len(); i++ {
			in := reflect.ValueOf(lx.Index(i).Interface())
			out := vf.Call([]reflect.Value{in})[0]
			ls = reflect.Append(ls, out)
		}
		return ls.Interface()
	default:
		ls := reflect.MakeSlice(reflect.SliceOf(ts), 0, lx.Len())
		for i := 0; i < lx.Len(); i++ {
			in := reflect.ValueOf(lx.Index(i).Interface())
			out := vf.Call([]reflect.Value{in})
			vst := reflect.Indirect(reflect.New(ts))
			for j := range out {
				vst.Field(j).Set(out[j])
			}
			ls = reflect.Append(ls, vst)
		}
		return ls.Interface()
	}
}

// la fonction 'String' retourne une chaine representant l'ensemble
func (pe *ens_t) String() string {
	if nil == pe {
		return "<nil>"
	}
	lk := pe.Lister()
	str := fmt.Sprintf("%v", lk)
	return str
}

// la fonction 'texte' retourne une chaine representant l'ensemble y compris 't' et 'ind'
func (pe *ens_t) texte() string {
	if nil == pe {
		return "<nil>"
	}
	str := "["
	str += fmt.Sprintf("t=%v ind=%v ", pe.t, pe.ind)
	lk := pe.Lister()
	str += fmt.Sprintf("%v", lk)
	str += "]"
	return str
}

// la fonction 'intersection' retourne l'ensemble des elements communs a tous les ensembles de la liste de parametres
func intersection(lpx ...*ens_t) *ens_t {
	if 0 == len(lpx) {
		//return nil
		panic("Intersection")
	}
	// tri par ordre croissant du nombre d'elements
	for i := 0; i < len(lpx); i++ {
		for j := i + 1; j < len(lpx); j++ {
			if lpx[i].Nombre() > lpx[j].Nombre() {
				lpx[i], lpx[j] = lpx[j], lpx[i]
			}
		}
	}
	var pe *ens_t
	if nil != lpx[0] {
		pe = lpx[0].copier()
		lpx = lpx[1:]
		for _, px := range lpx {
			pe.intersecter(px)
		}
	}
	return pe
}

// la fonction 'union' retourne l'ensemble des elements de tous les ensembles de la liste de parametres
func union(lpx ...*ens_t) *ens_t {
	if 0 == len(lpx) {
		//return nil
		panic("Union")
	}
	var pe *ens_t
	for i := 0; i < len(lpx); i++ {
		if nil != lpx[i] {
			pe = lpx[i].copier()
			lpx = lpx[1:]
			break
		}
		lpx = lpx[1:]
	}
	for _, px := range lpx {
		pe.unir(px)
	}
	return pe
}

// la fonction 'soustraction' retourne l'ensemble des elements de py qui n'appartiennent a aucun des ensembles de la liste 'lpx'
func soustraction(py *ens_t, lpx ...*ens_t) *ens_t {
	if nil == py {
		return nil
	}
	pe := py.copier()
	for _, px := range lpx {
		pe.soustraire(px)
	}
	return pe
}

// la fonction 'Intersection' retourne l'ensemble des elements communs a tous les ensembles de la liste de parametres
func Intersection(lpe ...Ensemble) Ensemble {
	lpx := make([]*ens_t, len(lpe))
	for i, x := range lpe {
		px := conv(x)
		lpx[i] = px
	}
	return intersection(lpx...)
}

// la fonction 'Union' retourne l'ensemble des elements de tous les ensembles de la liste de parametres
func Union(lpe ...Ensemble) Ensemble {
	lpx := make([]*ens_t, len(lpe))
	for i, x := range lpe {
		px := conv(x)
		lpx[i] = px
	}
	return union(lpx...)
}

// la fonction 'Soustraction' retourne l'ensemble des elements de py qui n'appartiennent a aucun des ensembles de la liste 'lpx'
func Soustraction(y Ensemble, lpe ...Ensemble) Ensemble {
	py := conv(y)
	lpx := make([]*ens_t, len(lpe))
	for i, x := range lpe {
		px := conv(x)
		lpx[i] = px
	}
	return soustraction(py, lpx...)
}
