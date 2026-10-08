package diccionario

import (
	TDALista "tdas/lista"
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{}
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	return false
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	var zero V
	return zero
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	var zero V
	return zero
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return nil
}
