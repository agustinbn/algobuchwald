package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestEncolarDesencolarRespetaInvarianteFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)

	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestEncolarConVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cantidad := 200
	for i := 0; i < cantidad; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for i := 0; i < cantidad; i++ {
		require.EqualValues(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestColaVaciadaSeComportaComoRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Desencolar()
	cola.Desencolar()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(10)
	require.EqualValues(t, 10, cola.VerPrimero())
}

func TestColaConVariosDatos(t *testing.T) {
	probarEncolarDesencolar(t, []int{1, 2, 3})
	probarEncolarDesencolar(t, []string{"pepe", "pedro", "juan"})
	probarEncolarDesencolar(t, []bool{true, false, true})
}

func probarEncolarDesencolar[T comparable](t *testing.T, valores []T) {
	cola := TDACola.CrearColaEnlazada[T]()
	for _, v := range valores {
		cola.Encolar(v)
	}
	for _, v := range valores {
		require.EqualValues(t, v, cola.VerPrimero())
		require.EqualValues(t, v, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}
