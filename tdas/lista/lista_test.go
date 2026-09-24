package lista_test

import (
	TDALista "tdas/lista"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestListaRecienCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestInsertarPrimeroYUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	
	lista.InsertarUltimo(2)
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 2, lista.Largo())
}

func TestBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
}

func TestVerPrimeroYUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 2, lista.VerUltimo())
}

func TestLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.EqualValues(t, 0, lista.Largo())
	
	lista.InsertarPrimero(1)
	require.EqualValues(t, 1, lista.Largo())
	
	lista.InsertarUltimo(2)
	require.EqualValues(t, 2, lista.Largo())
	
	lista.BorrarPrimero()
	require.EqualValues(t, 1, lista.Largo())
	
	lista.BorrarPrimero()
	require.EqualValues(t, 0, lista.Largo())
}

func TestIteradorEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
}

func TestIteradorInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Insertar(1)

	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 1, iterador.VerActual())

	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.VerPrimero())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	iterador := lista.Iterador()
	iterador.Avanzar()
	iterador.Avanzar()
	require.False(t, iterador.HayAlgoMas())

	iterador.Insertar(3)

	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Avanzar()
	iterador.Insertar(2)

	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 2, iterador.VerActual())

	iterador.Avanzar()
	require.EqualValues(t, 3, iterador.VerActual())
	iterador.Avanzar()
	require.False(t, iterador.HayAlgoMas())

	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestIteradorBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.Borrar())

	require.EqualValues(t, 2, lista.VerPrimero())
	require.EqualValues(t, 3, lista.VerUltimo())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 2, iterador.VerActual())
}

func TestIteradorBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Avanzar()
	iterador.Avanzar()
	require.EqualValues(t, 3, iterador.Borrar())

	require.False(t, iterador.HayAlgoMas())
	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 2, lista.VerUltimo())
	require.EqualValues(t, 1, lista.VerPrimero())
}

func TestIteradorBorrarDelMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	iterador := lista.Iterador()
	iterador.Avanzar()
	require.EqualValues(t, 5, iterador.Borrar())

	require.EqualValues(t, 2, lista.Largo())
	require.EqualValues(t, 6, iterador.VerActual())
}

func TestIteradorBorrarUnicoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)

	iterador := lista.Iterador()
	require.EqualValues(t, 1, iterador.Borrar())

	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.False(t, iterador.HayAlgoMas())
}

func TestIteradorAvanzarYVerActualConPANIC(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)

	iterador := lista.Iterador()
	iterador.Avanzar()
	require.False(t, iterador.HayAlgoMas())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Avanzar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
}


func TestIterarSinCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	elementos := []int{}
	lista.Iterar(func(elemento int) bool {
		elementos = append(elementos, elemento)
		return true
	})
	require.EqualValues(t, []int{1, 2, 3}, elementos)
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	llamadas := 0
	lista.Iterar(func(elemento int) bool {
		llamadas++
		return true
	})
	require.EqualValues(t, 0, llamadas)
}

func TestIterarConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)

	elementos := []int{}
	lista.Iterar(func(elemento int) bool {
		elementos = append(elementos, elemento)
		return false
	})
	require.EqualValues(t, []int{1}, elementos)
}
