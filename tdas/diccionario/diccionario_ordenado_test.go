package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	TDADiccionario "tdas/diccionario"
)

func comparacionInt(a, b int) int {
	return a - b
}

func crearArbolInt(claves []int) TDADiccionario.DiccionarioOrdenado[int, int] {
	abb := TDADiccionario.CrearABB[int, int](comparacionInt)
	for _, clave := range claves {
		abb.Guardar(clave, clave*10)
	}
	return abb
}

func TestABBVacio(t *testing.T) {
	t.Log("Un ABB vacío no tiene claves y Obtener/Borrar entran en pánico")
	abb := TDADiccionario.CrearABB[int, int](comparacionInt)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(0) })
}

func TestABBGuardarEnteros(t *testing.T) {
	t.Log("Guarda enteros y verifica Pertenece, Obtener y Cantidad")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	require.EqualValues(t, 7, abb.Cantidad())
	for _, clave := range []int{50, 30, 70, 20, 40, 60, 80} {
		require.True(t, abb.Pertenece(clave))
		require.EqualValues(t, clave*10, abb.Obtener(clave))
	}
	require.False(t, abb.Pertenece(10))
	require.False(t, abb.Pertenece(45))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(45) })
}

func TestABBGuardarStrings(t *testing.T) {
	t.Log("Guarda strings usando strings.Compare")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar("Gato", "miau")
	abb.Guardar("Perro", "guau")
	abb.Guardar("Vaca", "moo")
	abb.Guardar("", "vacia")
	require.EqualValues(t, 4, abb.Cantidad())
	require.EqualValues(t, "miau", abb.Obtener("Gato"))
	require.EqualValues(t, "guau", abb.Obtener("Perro"))
	require.EqualValues(t, "moo", abb.Obtener("Vaca"))
	require.EqualValues(t, "vacia", abb.Obtener(""))
	require.False(t, abb.Pertenece("gato"))
}

func TestABBReemplazarDato(t *testing.T) {
	t.Log("Guardar una clave existente reemplaza el dato sin cambiar la cantidad")
	abb := crearArbolInt([]int{5, 3, 8})
	abb.Guardar(3, -1)
	abb.Guardar(5, -2)
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, -1, abb.Obtener(3))
	require.EqualValues(t, -2, abb.Obtener(5))
}

func TestABBBorrarHoja(t *testing.T) {
	t.Log("Borra un nodo sin hijos")
	abb := crearArbolInt([]int{50, 30, 70})
	require.EqualValues(t, 300, abb.Borrar(30))
	require.False(t, abb.Pertenece(30))
	require.True(t, abb.Pertenece(50))
	require.True(t, abb.Pertenece(70))
	require.EqualValues(t, 2, abb.Cantidad())
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(30) })
}

func TestABBBorrarUnHijo(t *testing.T) {
	t.Log("Borra nodos con un solo hijo, izquierdo y derecho")
	abb := crearArbolInt([]int{50, 30, 70, 20, 80})
	require.EqualValues(t, 300, abb.Borrar(30))
	require.EqualValues(t, 700, abb.Borrar(70))
	require.EqualValues(t, 3, abb.Cantidad())
	require.False(t, abb.Pertenece(30))
	require.False(t, abb.Pertenece(70))
	for _, clave := range []int{50, 20, 80} {
		require.EqualValues(t, clave*10, abb.Obtener(clave))
	}
}

func TestABBBorrarDosHijos(t *testing.T) {
	t.Log("Borra nodos con dos hijos, incluida la raíz, usando reemplazante")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80, 65})
	require.EqualValues(t, 300, abb.Borrar(30))
	require.EqualValues(t, 500, abb.Borrar(50))
	require.EqualValues(t, 700, abb.Borrar(70))
	require.EqualValues(t, 5, abb.Cantidad())
	for _, clave := range []int{30, 50, 70} {
		require.False(t, abb.Pertenece(clave))
	}
	for _, clave := range []int{20, 40, 60, 65, 80} {
		require.EqualValues(t, clave*10, abb.Obtener(clave))
	}
}

func TestABBBorrarTodoYReusar(t *testing.T) {
	t.Log("Borra todo (raíz incluida) y el árbol se puede seguir usando")
	abb := crearArbolInt([]int{2, 1, 3})
	abb.Borrar(2)
	abb.Borrar(1)
	abb.Borrar(3)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(2))
	abb.Guardar(7, 70)
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, 70, abb.Obtener(7))
}

func TestABBVolumen(t *testing.T) {
	t.Log("Guarda, obtiene y borra muchas claves en orden aleatorio")
	for _, n := range []int{1000, 10000, 50000} {
		t.Run(fmt.Sprintf("%d elementos", n), func(t *testing.T) {
			claves := rand.Perm(n)
			abb := crearArbolInt(claves)
			require.EqualValues(t, n, abb.Cantidad())

			for _, clave := range claves {
				require.EqualValues(t, clave*10, abb.Obtener(clave))
			}

			rand.Shuffle(n, func(i, j int) { claves[i], claves[j] = claves[j], claves[i] })
			for i, clave := range claves {
				require.EqualValues(t, clave*10, abb.Borrar(clave))
				require.False(t, abb.Pertenece(clave))
				require.EqualValues(t, n-i-1, abb.Cantidad())
			}
		})
	}
}
