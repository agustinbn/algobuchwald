package cola_prioridad_test

import (
	TDAHeap "tdas/heap"

	"math/rand"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapSortVacio(t *testing.T) {
	arr := []int{}
	TDAHeap.HeapSort(arr, comparacionInt)
	require.Empty(t, arr)
}

func TestHeapSortUnElemento(t *testing.T) {
	arr := []int{7}
	TDAHeap.HeapSort(arr, comparacionInt)
	require.Equal(t, []int{7}, arr)
}

func TestHeapSortOrdenado(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	TDAHeap.HeapSort(arr, comparacionInt)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr)
}

func TestHeapSortInverso(t *testing.T) {
	arr := []int{6, 5, 4, 3, 2, 1}
	TDAHeap.HeapSort(arr, comparacionInt)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr)
}

func TestHeapSortRepetidos(t *testing.T) {
	arr := []int{3, 1, 3, 2, 1, 3, 2, 2}
	TDAHeap.HeapSort(arr, comparacionInt)
	require.Equal(t, []int{1, 1, 2, 2, 2, 3, 3, 3}, arr)
}

func TestHeapSortStrings(t *testing.T) {
	arr := []string{"pera", "manzana", "banana", "kiwi"}
	TDAHeap.HeapSort(arr, strings.Compare)
	require.Equal(t, []string{"banana", "kiwi", "manzana", "pera"}, arr)
}

func TestHeapSortVolumenGrande(t *testing.T) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = rand.Intn(10000)
	}
	esperado := make([]int, len(arr))
	copy(esperado, arr)
	sort.Ints(esperado)

	TDAHeap.HeapSort(arr, comparacionInt)
	require.Equal(t, esperado, arr)
}

func TestCrearHeapArrVacio(t *testing.T) {
	heap := TDAHeap.CrearHeapArr([]int{}, comparacionInt)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCrearHeapArrDesencolaOrdenado(t *testing.T) {
	heap := TDAHeap.CrearHeapArr([]int{4, 9, 1, 7, 3, 8, 2}, comparacionInt)
	require.EqualValues(t, 7, heap.Cantidad())
	require.EqualValues(t, 9, heap.VerMax())
	for _, esperado := range []int{9, 8, 7, 4, 3, 2, 1} {
		require.EqualValues(t, esperado, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestCrearHeapArrEncolarDespues(t *testing.T) {
	heap := TDAHeap.CrearHeapArr([]int{2, 5, 1}, comparacionInt)
	for i := 10; i < 50; i++ {
		heap.Encolar(i)
	}
	require.EqualValues(t, 43, heap.Cantidad())
	require.EqualValues(t, 49, heap.VerMax())
	require.EqualValues(t, 49, heap.Desencolar())
	require.EqualValues(t, 48, heap.VerMax())
}

func TestCrearHeapArrVacioYEncolar(t *testing.T) {
	heap := TDAHeap.CrearHeapArr([]int{}, comparacionInt)
	heap.Encolar(3)
	heap.Encolar(9)
	require.EqualValues(t, 9, heap.VerMax())
}

func TestCrearHeapArrNoModificaOriginal(t *testing.T) {
	arr := []int{4, 9, 1, 7, 3}
	heap := TDAHeap.CrearHeapArr(arr, comparacionInt)
	heap.Encolar(100)
	heap.Desencolar()
	heap.Desencolar()
	require.Equal(t, []int{4, 9, 1, 7, 3}, arr)
}
