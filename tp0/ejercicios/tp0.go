package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	posicion_maximo := -1
	for k, v := range vector {
		if posicion_maximo == -1 || v > vector[posicion_maximo] {
			posicion_maximo = k
		}
	}
	return posicion_maximo
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	for k := range vector1 {
		if k == len(vector2) {
			return 1
		}
		if vector1[k] < vector2[k] {
			return -1
		}
		if vector1[k] > vector2[k] {
			return 1
		}
	}
	if len(vector1) < len(vector2) {
		return -1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := 0; i < len(vector)-1; i++ {
		posicion_minimo := i
		for j := i + 1; j < len(vector); j++ {
			if vector[j] < vector[posicion_minimo] {
				posicion_minimo = j
			}
		}
		vector[i], vector[posicion_minimo] = vector[posicion_minimo], vector[i]
	}
}

func sumarValores(vector []int, pos int, acumulado int) int {
	if pos == len(vector) {
		return acumulado
	}
	return sumarValores(vector, pos+1, acumulado+vector[pos])
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	return sumarValores(vector, 0, 0)
}

func esCapicua(cadena string, cadena_inversa string, pos int) bool {
	r := []rune(cadena)

	if pos == len(r) {
		return cadena == cadena_inversa
	}

	nueva_cadena_inversa := cadena_inversa + string(cadena[len(r)-1-pos])

	return esCapicua(cadena, nueva_cadena_inversa, pos+1)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	return esCapicua(cadena, "", 0)
}
