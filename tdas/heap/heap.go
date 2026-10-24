package cola_prioridad

const (
	_CAPACIDAD_INICIAL   = 10
	_FACTOR_REDIMENSION  = 2
	_FACTOR_ACHICAMIENTO = 4
)

type cmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      cmpHeap[T]
}

func CrearHeap[T any](cmp cmpHeap[T]) ColaPrioridad[T] {
	return &heap[T]{
		datos:    make([]T, _CAPACIDAD_INICIAL),
		cantidad: 0,
		cmp:      cmp,
	}
}

func (h *heap[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, h.datos)
	h.datos = nuevosDatos
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elemento T) {
	if h.cantidad == len(h.datos) {
		h.redimensionar(len(h.datos) * _FACTOR_REDIMENSION)
	}
	h.datos[h.cantidad] = elemento
	h.cantidad++
	upHeap(h.datos[:h.cantidad], h.cantidad-1, h.cmp)
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	elemento := h.datos[0]
	swap(&h.datos[0], &h.datos[h.cantidad-1])
	h.cantidad--
	if len(h.datos) > _CAPACIDAD_INICIAL && h.cantidad*_FACTOR_ACHICAMIENTO <= len(h.datos) {
		nuevaCapacidad := max(len(h.datos)/_FACTOR_REDIMENSION, _CAPACIDAD_INICIAL)
		h.redimensionar(nuevaCapacidad)
	}
	downHeap(h.datos, h.cantidad, 0, h.cmp)
	return elemento
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func downHeap[T any](datos []T, cantidad, i int, cmp func(T, T) int) {
	izq, der := 2*i+1, 2*i+2
	if izq >= cantidad {
		return
	}
	mayor := izq
	if der < cantidad && cmp(datos[der], datos[izq]) > 0 {
		mayor = der
	}
	if cmp(datos[mayor], datos[i]) > 0 {
		swap(&datos[i], &datos[mayor])
		downHeap(datos, cantidad, mayor, cmp)
	}
}

func upHeap[T any](datos []T, i int, cmp func(T, T) int) {
	if i == 0 {
		return
	}
	padre := (i - 1) / 2
	if cmp(datos[padre], datos[i]) < 0 {
		swap(&datos[padre], &datos[i])
		upHeap(datos, padre, cmp)
	}
}

func heapify[T any](datos []T, cantidad int, cmp func(T, T) int) {
	for i := cantidad/2 - 1; i >= 0; i-- {
		downHeap(datos, cantidad, i, cmp)
	}
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}
