package diccionario

type diccionarioHash[K comparable, V any] struct {
	tabla    []*celdaHash[K, V]
	tam      int
	cant     int
	borrados int
}

type interHash[K comparable, V any] struct {
	diccionario *diccionarioHash[K, V]
	pos         int
}

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado celdaEstado
}

type celdaEstado string

const (
	ocupada celdaEstado = "ocupada"
	vacia   celdaEstado = "vacia"
	borrada celdaEstado = "borrada"
)

func (iter *interHash[K, V]) HayAlgoMas() bool {
	return iter.pos < iter.diccionario.tam && iter.diccionario.tabla[iter.pos].estado == ocupada
}

func (iter *interHash[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	celda := iter.diccionario.tabla[iter.pos]
	return celda.clave, celda.dato
}

func (iter *interHash[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iter.pos = iter.diccionario.buscarOcupada(iter.pos + 1)
}

func (d *diccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(interHash[K, V])
	iter.diccionario = d
	iter.pos = d.buscarOcupada(0)
	return iter
}

func (d *diccionarioHash[K, V]) Iterar(visitante func(clave K, dato V) bool) {
	iter := d.Iterador()
	for iter.HayAlgoMas() {
		clave, dato := iter.VerActual()
		if !visitante(clave, dato) {
			return
		}
		iter.Avanzar()
	}
}

func (d *diccionarioHash[K, V]) buscarOcupada(pos int) int {
	for pos < d.tam && d.tabla[pos].estado != ocupada {
		pos++
	}
	return pos
}
