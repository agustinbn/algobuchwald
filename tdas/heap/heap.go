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
	h.upHeap(h.cantidad - 1)
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
	h.downHeap(0)
	return elemento
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func (h *heap[T]) downHeap(i int) {
	izq, der := 2*i+1, 2*i+2
	mayor := izq
	if der < h.cantidad {
		mayor = h.maxIndice(izq, der)
	}
	if izq < h.cantidad && h.cmp(h.datos[mayor], h.datos[i]) > 0 {
		swap(&h.datos[i], &h.datos[mayor])
		h.downHeap(mayor)
	}
}
func (h *heap[T]) maxIndice(a, b int) int {
	if h.cmp(h.datos[b], h.datos[a]) > 0 {
		return b
	}
	return a
}

func (h *heap[T]) upHeap(i int) {
	padre := (i - 1) / 2
	if i != 0 && h.cmp(h.datos[padre], h.datos[i]) < 0 {
		swap(&h.datos[padre], &h.datos[i])
		h.upHeap(padre)
	}
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}
