package cola_prioridad

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, max(len(arreglo), _CAPACIDAD_INICIAL))
	copy(datos, arreglo)
	h := &heap[T]{datos: datos, cantidad: len(arreglo), cmp: cmp}
	heapify(h)
	return h
}

func HeapSort[T any](elementos []T, cmp func(T, T) int) {
	h := &heap[T]{datos: elementos, cantidad: len(elementos), cmp: cmp}
	heapify(h)
	for h.cantidad > 1 {
		h.cantidad--
		swap(&h.datos[0], &h.datos[h.cantidad])
		h.downHeap(0)
	}
}

func heapify[T any](h *heap[T]) {
	for i := h.cantidad/2 - 1; i >= 0; i-- {
		h.downHeap(i)
	}
}
