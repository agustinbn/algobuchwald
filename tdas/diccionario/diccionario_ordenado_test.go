package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	TDADiccionario "tdas/diccionario"

	"github.com/stretchr/testify/require"
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

func TestIterar(t *testing.T) {
	t.Log("Itera sobre todas las claves en orden")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	var claves []int
	abb.Iterar(func(clave int, dato int) bool {
		claves = append(claves, clave)
		return true
	})
	require.EqualValues(t, []int{20, 30, 40, 50, 60, 70, 80}, claves)
}

func TestIterarRango(t *testing.T) {
	t.Log("Itera sobre un rango de claves en orden")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	var claves []int
	desde := 35
	hasta := 75
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		claves = append(claves, clave)
		return true
	})
	require.EqualValues(t, []int{40, 50, 60, 70}, claves)
}

func TestIterador(t *testing.T) {
	t.Log("Itera sobre todas las claves en orden usando un iterador")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	iter := abb.Iterador()
	var claves []int
	for iter.HayAlgoMas() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Avanzar()
	}
	require.EqualValues(t, []int{20, 30, 40, 50, 60, 70, 80}, claves)
}

func TestIteradorRango(t *testing.T) {
	t.Log("Itera sobre un rango de claves en orden usando un iterador")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	desde := 35
	hasta := 75
	iter := abb.IteradorRango(&desde, &hasta)
	var claves []int
	for iter.HayAlgoMas() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Avanzar()
	}
	require.EqualValues(t, []int{40, 50, 60, 70}, claves)
}

func TestIterarSuma(t *testing.T) {
	t.Log("Itera sobre todas las claves en orden y suma los datos")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	suma := 0
	abb.Iterar(func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 3500, suma)
}

func TestIterarRangoSuma(t *testing.T) {
	t.Log("Itera sobre un rango de claves en orden y suma los datos")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	desde := 35
	hasta := 75
	suma := 0
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.EqualValues(t, 2200, suma)
}

func TestIterarCondicionCorte(t *testing.T) {
	t.Log("Itera sobre todas las claves en orden y corta la iteración con una condición")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	var claves []int
	abb.Iterar(func(clave int, dato int) bool {
		claves = append(claves, clave)
		return clave != 60
	})
	require.EqualValues(t, []int{20, 30, 40, 50, 60}, claves)
}

func TestIterarVacio(t *testing.T) {
	t.Log("Iterar sobre un ABB vacío no llama a la función")
	abb := TDADiccionario.CrearABB[int, int](comparacionInt)
	abb.Iterar(func(clave int, dato int) bool {
		require.Fail(t, "No debería llamarse a la función")
		return true
	})
}

func TestIterarRangoNilHasta(t *testing.T) {
	t.Log("Itera sobre un rango con hasta nil, incluyendo todas las claves mayores o iguales a desde")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	desde := 40
	var claves []int
	abb.IterarRango(&desde, nil, func(clave int, dato int) bool {
		claves = append(claves, clave)
		return true
	})
	require.EqualValues(t, []int{40, 50, 60, 70, 80}, claves)
}

func TestIterarRangoNilDesde(t *testing.T) {
	t.Log("Itera sobre un rango con desde nil, incluyendo todas las claves menores o iguales a hasta")
	abb := crearArbolInt([]int{50, 30, 70, 20, 40, 60, 80})
	hasta := 60
	var claves []int
	abb.IterarRango(nil, &hasta, func(clave int, dato int) bool {
		claves = append(claves, clave)
		return true
	})
	require.EqualValues(t, []int{20, 30, 40, 50, 60}, claves)
}
