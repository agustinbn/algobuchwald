package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	_CAPACIDAD_INICIAL   = 10
	_FACTOR_REDIMENSION  = 2
	_FACTOR_ACHICAMIENTO = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, _CAPACIDAD_INICIAL),
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
		p.redimensionar(len(p.datos) * _FACTOR_REDIMENSION)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	elemento := p.VerTope()
	p.cantidad--
	if len(p.datos) > _CAPACIDAD_INICIAL && p.cantidad*_FACTOR_ACHICAMIENTO <= len(p.datos) {
		nuevaCapacidad := max(len(p.datos)/_FACTOR_REDIMENSION, _CAPACIDAD_INICIAL)
		p.redimensionar(nuevaCapacidad)
	}
	return elemento
}
