package main

import (
	"fmt"
	"os"
	"strings"
	TDACola "tdas/cola"
)

type Posicion struct {
	fila    int
	columna int
}

type Direccion struct {
	signo        string
	deltaFila    int
	deltaColumna int
}

const (
	_SIGNO_INICIO byte   = 'S'
	_SIGNO_FIN    byte   = 'E'
	_SIGNO_PARED  byte   = '#'
	_ARRIBA       string = "ARRIBA"
	_ABAJO        string = "ABAJO"
	_IZQUIERDA    string = "IZQUIERDA"
	_DERECHA      string = "DERECHA"
)

var _DIRECCIONES = []Direccion{
	{_ARRIBA, -1, 0},
	{_ABAJO, 1, 0},
	{_IZQUIERDA, 0, -1},
	{_DERECHA, 0, 1},
}

func esTransitable(laberinto []string, filas int, columnas int, pos Posicion) bool {
	if pos.fila < 0 || pos.fila >= filas {
		return false
	}

	if pos.columna < 0 || pos.columna >= columnas {
		return false
	}

	if laberinto[pos.fila][pos.columna] == _SIGNO_PARED {
		return false
	}

	return true
}

func resolverLaberinto(laberinto []string, filas int, columnas int, inicio Posicion, fin Posicion) ([]string, bool) {
	visitado := make([][]bool, filas)
	origen := make([][]Posicion, filas)
	direccion := make([][]string, filas)

	for i := range filas {
		visitado[i] = make([]bool, columnas)
		origen[i] = make([]Posicion, columnas)
		direccion[i] = make([]string, columnas)
	}

	pendientes := TDACola.CrearColaEnlazada[Posicion]()

	visitado[inicio.fila][inicio.columna] = true
	pendientes.Encolar(inicio)

	for !pendientes.EstaVacia() {
		actual := pendientes.Desencolar()

		if actual == fin {
			break
		}

		for _, dir := range _DIRECCIONES {
			deltaFila, deltaColumna := dir.deltaFila, dir.deltaColumna
			nuevaFila := actual.fila + deltaFila
			nuevaColumna := actual.columna + deltaColumna
			vecino := Posicion{
				nuevaFila,
				nuevaColumna,
			}

			if esTransitable(laberinto, filas, columnas, vecino) {
				if !visitado[vecino.fila][vecino.columna] {
					visitado[vecino.fila][vecino.columna] = true
					origen[vecino.fila][vecino.columna] = actual
					direccion[vecino.fila][vecino.columna] = dir.signo
					pendientes.Encolar(vecino)
				}
			}
		}
	}

	if !visitado[fin.fila][fin.columna] {
		return nil, false
	}

	pasos := []string{}
	actual := fin

	for actual != inicio {
		pasos = append(pasos, direccion[actual.fila][actual.columna])
		actual = origen[actual.fila][actual.columna]
	}

	for i := 0; i < len(pasos)/2; i++ {
		j := len(pasos) - 1 - i
		aux := pasos[i]
		pasos[i] = pasos[j]
		pasos[j] = aux
	}

	return pasos, true
}

func main() {
	for {
		var filas, columnas int
		if _, err := fmt.Fscan(os.Stdin, &filas, &columnas); err != nil {
			break
		}

		laberinto := make([]string, filas)

		var inicio Posicion
		var fin Posicion

		for i := 0; i < filas; i++ {
			fmt.Fscan(os.Stdin, &laberinto[i])

			for j := 0; j < columnas; j++ {
				switch laberinto[i][j] {
				case _SIGNO_INICIO:
					inicio = Posicion{i, j}
				case _SIGNO_FIN:
					fin = Posicion{i, j}
				}
			}
		}

		pasos, encontrado := resolverLaberinto(laberinto, filas, columnas, inicio, fin)
		if encontrado {
			fmt.Println(len(pasos))
			fmt.Println(strings.Join(pasos, " "))
		} else {
			fmt.Println("ERROR")
		}
	}
}
