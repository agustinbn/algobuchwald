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