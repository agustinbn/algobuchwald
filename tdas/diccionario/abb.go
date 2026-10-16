package diccionario

import TDAPila "tdas/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterAbb[K comparable, V any] struct {
	arbol *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:     nil,
		cantidad: 0,
		cmp:      funcion_cmp,
	}
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {
	if arbol.raiz == nil {
		arbol.raiz = crearNodoAbb(clave, dato)
		arbol.cantidad++
		return
	}

	nodo := arbol.raiz
	for {
		comparacion := arbol.cmp(clave, nodo.clave)
		if comparacion == 0 {
			nodo.dato = dato
			return
		} else if comparacion < 0 {
			if nodo.izquierdo == nil {
				nodo.izquierdo = crearNodoAbb(clave, dato)
				arbol.cantidad++
				return
			}
			nodo = nodo.izquierdo
		} else {
			if nodo.derecho == nil {
				nodo.derecho = crearNodoAbb(clave, dato)
				arbol.cantidad++
				return
			}
			nodo = nodo.derecho
		}
	}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	if arbol.raiz == nil {
		return false
	}
	nodo := arbol.raiz
	for nodo != nil {
		comparacion := arbol.cmp(clave, nodo.clave)
		if comparacion == 0 {
			return true
		} else if comparacion < 0 {
			nodo = nodo.izquierdo
		} else {
			nodo = nodo.derecho
		}
	}
	return false
}

func (arbol *abb[K, V]) Obtener(clave K) V {
	nodo := arbol.raiz
	for nodo != nil {
		comparacion := arbol.cmp(clave, nodo.clave)
		if comparacion == 0 {
			return nodo.dato
		} else if comparacion < 0 {
			nodo = nodo.izquierdo
		} else {
			nodo = nodo.derecho
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	dato := arbol.Obtener(clave)
	arbol.raiz = arbol.borrarNodo(arbol.raiz, clave)
	arbol.cantidad--
	return dato
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (iter *iterAbb[K, V]) HayAlgoMas() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iterAbb[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	iter.apilarIzq(nodo.derecho)
}

func (iter *iterAbb[K, V]) apilarIzq(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.arbol.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.derecho
		} else if iter.hasta != nil && iter.arbol.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izquierdo
		} else {
			iter.pila.Apilar(nodo)
			nodo = nodo.izquierdo
		}
	}
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	iter := arbol.IteradorRango(nil, nil)
	for iter.HayAlgoMas() {
		clave, dato := iter.VerActual()
		if !visitar(clave, dato) {
			return
		}
		iter.Avanzar()
	}
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iter := arbol.IteradorRango(desde, hasta)
	for iter.HayAlgoMas() {
		clave, dato := iter.VerActual()
		if !visitar(clave, dato) {
			return
		}
		iter.Avanzar()
	}
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		arbol: arbol,
		pila:  TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde: desde,
		hasta: hasta,
	}
	iter.apilarIzq(arbol.raiz)
	return iter
}
func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		clave:     clave,
		dato:      dato,
		izquierdo: nil,
		derecho:   nil,
	}
}

func (arbol *abb[K, V]) borrarNodo(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	comparacion := arbol.cmp(clave, nodo.clave)
	if comparacion < 0 {
		nodo.izquierdo = arbol.borrarNodo(nodo.izquierdo, clave)
		return nodo
	}
	if comparacion > 0 {
		nodo.derecho = arbol.borrarNodo(nodo.derecho, clave)
		return nodo
	}

	if nodo.izquierdo == nil {
		return nodo.derecho
	}
	if nodo.derecho == nil {
		return nodo.izquierdo
	}

	reemplazante := nodo.derecho
	for reemplazante.izquierdo != nil {
		reemplazante = reemplazante.izquierdo
	}
	nodo.clave, nodo.dato = reemplazante.clave, reemplazante.dato
	nodo.derecho = arbol.borrarNodo(nodo.derecho, reemplazante.clave)
	return nodo
}
