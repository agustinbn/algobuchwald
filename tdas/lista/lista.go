package lista

type IteradorLista[T any] interface {
	VerActual() T
	HayAlgoMas() bool
	Avanzar()
	Insertar(T)
	Borrar() T
}

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)
	
	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista y lo devuelve. Si la lista está vacía, devuelve un panic de "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero devuelve el valor del primer elemento de la lista. Si la lista está vacía, devuelve un panic de "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el valor del último elemento de la lista. Si la lista está vacía, devuelve un panic de "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recorre la lista y aplica la función visitar a cada elemento. Si visitar devuelve false, se detiene la iteración.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador para la lista.
	Iterador() IteradorLista[T]
}