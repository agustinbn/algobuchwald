package cola_prioridad

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, max(len(arreglo), _CAPACIDAD_INICIAL))
	copy(datos, arreglo)
	heapify(datos, len(arreglo), cmp)
	return &heap[T]{datos: datos, cantidad: len(arreglo), cmp: cmp}
}

func HeapSort[T any](elementos []T, cmp func(T, T) int) {
	heapify(elementos, len(elementos), cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		swap(&elementos[0], &elementos[i])
		downHeap(elementos, i, 0, cmp)
	}
}
