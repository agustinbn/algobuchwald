package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodoLista[T]
	actual   *nodoLista[T]
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	} else {
		nuevo.siguiente = lista.primero
	}
	lista.primero = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := crearNodo(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevo
	} else {
		lista.ultimo.siguiente = nuevo
	}
	lista.ultimo = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	primero := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return primero
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; actual = actual.siguiente {
		if !visitar(actual.dato) {
			break
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{
		lista:    lista,
		anterior: nil,
		actual:   lista.primero,
	}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HayAlgoMas() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	if iterador.anterior == nil {
		iterador.lista.InsertarPrimero(elemento)
		iterador.actual = iterador.lista.primero
	} else if !iterador.HayAlgoMas() {
		iterador.lista.InsertarUltimo(elemento)
		iterador.actual = iterador.lista.ultimo
	} else {
		nuevo := crearNodo(elemento)
		nuevo.siguiente = iterador.actual
		iterador.anterior.siguiente = nuevo
		iterador.actual = nuevo
		iterador.lista.largo++
	}
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	dato := iterador.actual.dato
	if iterador.anterior == nil {
		iterador.lista.BorrarPrimero()
		iterador.actual = iterador.lista.primero
	} else {
		if iterador.actual.siguiente == nil {
			iterador.anterior.siguiente = nil
			iterador.lista.ultimo = iterador.anterior
		} else {
			iterador.anterior.siguiente = iterador.actual.siguiente
		}
		iterador.actual = iterador.anterior.siguiente
		iterador.lista.largo--
	}
	return dato
}
