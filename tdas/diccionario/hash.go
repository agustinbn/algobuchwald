package diccionario

import (
	"fmt"
	"hash/fnv"
)

const (
	_CAPACIDAD_INICIAL        = 13
	_FACTOR_CARGA_DENOMINADOR = 10
	_FACTOR_CARGA_NUMERADOR   = 7
	_FACTOR_EXPANSION         = 2
	_FACTOR_REDUCCION         = 4
	_POSICION_INVALIDA        = -1
)

type celdaEstado string

const (
	_OCUPADA celdaEstado = "OCUPADA"
	_VACIA   celdaEstado = "VACIA"
	_BORRADA celdaEstado = "BORRADA"
)

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

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{tabla: crearTabla[K, V](_CAPACIDAD_INICIAL), tam: _CAPACIDAD_INICIAL}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if pos := hash.buscar(clave); pos != _POSICION_INVALIDA {
		hash.tabla[pos].dato = dato
		return
	}

	if (hash.cant+hash.borrados+1)*_FACTOR_CARGA_DENOMINADOR >= hash.tam*_FACTOR_CARGA_NUMERADOR {
		hash.redimensionar(hash.tam * _FACTOR_EXPANSION)
	}

	primeroBorrado := _POSICION_INVALIDA
	pos := hash.funcionHash(clave)

	for intentos := 0; intentos < hash.tam; intentos++ {
		if hash.tabla[pos].estado == _BORRADA {
			if primeroBorrado == _POSICION_INVALIDA {
				primeroBorrado = pos
			}
			pos = (pos + 1) % hash.tam
			continue
		}

		if hash.tabla[pos].estado == _VACIA {
			if primeroBorrado != _POSICION_INVALIDA {
				pos = primeroBorrado
				hash.borrados--
			}

			hash.tabla[pos] = crearCelda(clave, dato)
			hash.cant++
			return
		}
		pos = (pos + 1) % hash.tam
	}
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return hash.buscar(clave) != _POSICION_INVALIDA
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos := hash.buscar(clave)
	if pos == _POSICION_INVALIDA {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[pos].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos := hash.buscar(clave)
	if pos == _POSICION_INVALIDA {
		panic("La clave no pertenece al diccionario")
	}

	dato := hash.tabla[pos].dato
	hash.tabla[pos].estado = _BORRADA
	hash.cant--
	hash.borrados++
	hash.achicarSiCorresponde()
	return dato
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cant
}

func (iter *interHash[K, V]) HayAlgoMas() bool {
	return iter.pos < iter.diccionario.tam && iter.diccionario.tabla[iter.pos].estado == _OCUPADA
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

func (hash *hashCerrado[K, V]) Iterar(visitante func(clave K, dato V) bool) {
	iter := hash.Iterador()
	for iter.HayAlgoMas() {
		clave, dato := iter.VerActual()
		if !visitante(clave, dato) {
			return
		}
		iter.Avanzar()
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(interHash[K, V])
	iter.diccionario = hash
	iter.pos = hash.posOcupada(0)
	return iter
}

func (hash *hashCerrado[K, V]) posOcupada(pos int) int {
	for pos < hash.tam && hash.tabla[pos].estado != _OCUPADA {
		pos++
	}
	return pos
}

func crearTabla[K comparable, V any](tam int) []*celdaHash[K, V] {
	tabla := make([]*celdaHash[K, V], tam)
	for i := range tabla {
		tabla[i] = &celdaHash[K, V]{estado: _VACIA}
	}
	return tabla
}

func crearCelda[K comparable, V any](clave K, dato V) *celdaHash[K, V] {
	return &celdaHash[K, V]{clave: clave, dato: dato, estado: _OCUPADA}
}

func (hash *hashCerrado[K, V]) funcionHash(clave K) int {
	h := fnv.New64a()
	_, _ = fmt.Fprint(h, clave)
	return int(h.Sum64() % uint64(hash.tam))
}

func (hash *hashCerrado[K, V]) buscar(clave K) int {
	pos := hash.funcionHash(clave)

	for intentos := 0; intentos < hash.tam; intentos++ {
		celda := hash.tabla[pos]
		if celda.estado == _VACIA {
			return _POSICION_INVALIDA
		}
		if celda.estado == _OCUPADA && celda.clave == clave {
			return pos
		}
		pos = (pos + 1) % hash.tam
	}

	return _POSICION_INVALIDA
}

func (hash *hashCerrado[K, V]) redimensionar(nuevoTam int) {
	viejaTabla := hash.tabla
	hash.tabla = crearTabla[K, V](nuevoTam)
	hash.tam = nuevoTam
	hash.cant = 0
	hash.borrados = 0
	for _, celda := range viejaTabla {
		if celda.estado == _OCUPADA {
			hash.insertarSinRedimensionar(celda.clave, celda.dato)
		}
	}
}

func (hash *hashCerrado[K, V]) insertarSinRedimensionar(clave K, dato V) {
	pos := hash.funcionHash(clave)

	for intentos := 0; intentos < hash.tam; intentos++ {
		if hash.tabla[pos].estado != _OCUPADA || hash.tabla[pos].estado == _BORRADA {
			hash.tabla[pos] = crearCelda(clave, dato)
			hash.cant++
			return
		}
		pos = (pos + 1) % hash.tam
	}
}

func (hash *hashCerrado[K, V]) achicarSiCorresponde() {
	if hash.tam > _CAPACIDAD_INICIAL && hash.cant*_FACTOR_REDUCCION <= hash.tam {
		nuevoTam := hash.tam / _FACTOR_EXPANSION
		nuevoTam = max(nuevoTam, _CAPACIDAD_INICIAL)
		hash.redimensionar(nuevoTam)
	}
}
