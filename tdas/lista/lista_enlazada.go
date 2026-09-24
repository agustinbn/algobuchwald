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

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := &nodoLista[T]{dato: elemento}
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	} else {
		nuevo.siguiente = lista.primero
	}
	lista.primero = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := &nodoLista[T]{dato: elemento}
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
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	// puse esto para que buildee
	return &iterListaEnlazada[T]{
		lista:    lista,
		anterior: nil,
		actual:   lista.primero,
	}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	var t T
	return t
}

func (iterador *iterListaEnlazada[T]) HayAlgoMas() bool {
	return false
}

func (iterador *iterListaEnlazada[T]) Avanzar() {
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {

}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	var t T
	return t
}
