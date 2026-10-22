package cola_prioridad

import (
	"strings"
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func comparacionInt(a, b int) int {
	return a - b
}

func TestHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](comparacionInt)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar(5)
	heap.Encolar(3)
	heap.Encolar(8)
	heap.Encolar(1)

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 8, heap.VerMax())

	max := heap.Desencolar()
	require.EqualValues(t, 8, max)
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, 5, heap.VerMax())

	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](comparacionInt)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapRedimension(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](comparacionInt)

	for i := 0; i < 100; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	for i := 99; i >= 0; i-- {
		max := heap.Desencolar()
		require.EqualValues(t, i, max)
		require.EqualValues(t, i, heap.Cantidad())
	}
}

func TestHeapStrings(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](strings.Compare)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())

	heap.Encolar("Gato")
	heap.Encolar("Perro")
	heap.Encolar("Vaca")
	heap.Encolar("")

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, "Vaca", heap.VerMax())

	max := heap.Desencolar()
	require.EqualValues(t, "Vaca", max)
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, "Perro", heap.VerMax())

	heap.Desencolar()
	heap.Desencolar()
	heap.Desencolar()

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapEncolarDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](comparacionInt)

	for i := 0; i < 10; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	for i := 9; i >= 0; i-- {
		max := heap.Desencolar()
		require.EqualValues(t, i, max)
		require.EqualValues(t, i, heap.Cantidad())
	}
}

func TestHeapUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](comparacionInt)

	heap.Encolar(5)
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 5, heap.VerMax())

	max := heap.Desencolar()
	require.EqualValues(t, 5, max)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
}

func TestHeapVolumenGrande(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](comparacionInt)

	for i := 0; i < 10000; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	for i := 9999; i >= 0; i-- {
		max := heap.Desencolar()
		require.EqualValues(t, i, max)
		require.EqualValues(t, i, heap.Cantidad())
	}
}
