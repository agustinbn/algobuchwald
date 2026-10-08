package diccionario

type hashCerrado[K comparable, V any] struct {
	tabla    []*celdaHash[K, V]
	tam      int
	cant     int
	borrados int
}

type interHash[K comparable, V any] struct {
	diccionario *hashCerrado[K, V]
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

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return false
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	var zero V
	return zero
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	var zero V
	return zero
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cant
}

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
	iter.pos = iter.diccionario.posOcupada(iter.pos + 1)
}

func (d *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(interHash[K, V])
	iter.diccionario = d
	iter.pos = d.posOcupada(0)
	return iter
}

func (d *hashCerrado[K, V]) Iterar(visitante func(clave K, dato V) bool) {
	iter := d.Iterador()
	for iter.HayAlgoMas() {
		clave, dato := iter.VerActual()
		if !visitante(clave, dato) {
			return
		}
		iter.Avanzar()
	}
}

func (d *hashCerrado[K, V]) posOcupada(pos int) int {
	for pos < d.tam && d.tabla[pos].estado != ocupada {
		pos++
	}
	return pos
}
