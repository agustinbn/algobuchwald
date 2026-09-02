package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaRecienCreada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestApilarDesapilarRespetaInvarianteLIFO(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)

	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())

	const cantidad = 10000
	for i := 0; i < cantidad; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for i := cantidad - 1; i >= 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaVaciadaSeComportaComoRecienCreada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Desapilar()
	pila.Desapilar()

	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(10)
	require.EqualValues(t, 10, pila.VerTope())
}

func TestPilaConDistintosTiposDeDatos(t *testing.T) {
	probarApilarDesapilar(t, []int{1, 2, 3})
	probarApilarDesapilar(t, []string{"pepe", "pedro", "juan"})
	probarApilarDesapilar(t, []bool{true, false, true})
}

func probarApilarDesapilar[T comparable](t *testing.T, valores []T) {
	pila := TDAPila.CrearPilaDinamica[T]()
	for _, v := range valores {
		pila.Apilar(v)
	}
	for i := len(valores) - 1; i >= 0; i-- {
		require.EqualValues(t, valores[i], pila.VerTope())
		require.EqualValues(t, valores[i], pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}
