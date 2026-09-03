package pila

/* Definición del struct pila proporcionado por la cátedra. */

const CAPACIDAD_INICIAL = 10

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, CAPACIDAD_INICIAL),
		cantidad: 0,
	}
}

func (p *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if p.cantidad == len(p.datos) {
		p.redimensionar(len(p.datos) * 2)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	elemento := p.VerTope()
	p.cantidad--
	if len(p.datos) > CAPACIDAD_INICIAL && p.cantidad*4 <= len(p.datos) {
		nuevaCapacidad := len(p.datos) / 2
		if nuevaCapacidad < CAPACIDAD_INICIAL {
			nuevaCapacidad = CAPACIDAD_INICIAL
		}
		p.redimensionar(nuevaCapacidad)
	}
	return elemento
}
