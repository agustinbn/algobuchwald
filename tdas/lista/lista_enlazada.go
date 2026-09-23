package lista

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	anterior *nodo[T]
	actual   *nodo[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	var cero T
	return cero
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
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return nil
}

func (iterador *iteradorLista[T]) VerActual() T {
	var cero T
	return cero
}

func (iterador *iteradorLista[T]) HayAlgoMas() bool {
	return false
}

func (iterador *iteradorLista[T]) Avanzar() {
}

func (iterador *iteradorLista[T]) Insertar(elemento T) {
}

func (iterador *iteradorLista[T]) Borrar() T {
	var cero T
	return cero
}
