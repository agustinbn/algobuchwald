package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
		c.ultimo = nuevoNodo
	}
}

func (c *colaEnlazada[T]) Desencolar() T {
	primero := c.VerPrimero()
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	}
	return primero
}
