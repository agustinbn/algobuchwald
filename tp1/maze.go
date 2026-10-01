package main

import (
	"fmt"
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

type Laberinto struct {
	paredes  []string
	filas    int
	columnas int
	inicio   Posicion
	fin      Posicion
}

type registro struct {
	visitado  bool
	origen    Posicion
	direccion string
}

const (
	_SIGNO_INICIO byte = 'S'
	_SIGNO_FIN    byte = 'E'
	_SIGNO_PARED  byte = '#'
)

var _DIRECCIONES = []Direccion{
	{"ARRIBA", -1, 0},
	{"ABAJO", 1, 0},
	{"IZQUIERDA", 0, -1},
	{"DERECHA", 0, 1},
}

func (laberinto *Laberinto) esTransitable(pos Posicion) bool {
	if pos.fila < 0 || pos.fila >= laberinto.filas || pos.columna < 0 || pos.columna >= laberinto.columnas {
		return false
	}

	if laberinto.paredes[pos.fila][pos.columna] == _SIGNO_PARED {
		return false
	}

	return true
}

func (laberinto *Laberinto) registrarSignos(fila int) {
	for columna := range laberinto.columnas {
		switch laberinto.paredes[fila][columna] {
		case _SIGNO_INICIO:
			laberinto.inicio = Posicion{fila, columna}
		case _SIGNO_FIN:
			laberinto.fin = Posicion{fila, columna}
		}
	}
}

func leerLaberinto() (Laberinto, bool) {
	var laberinto Laberinto

	if _, err := fmt.Scan(&laberinto.filas, &laberinto.columnas); err != nil {
		return laberinto, false
	}

	laberinto.paredes = make([]string, laberinto.filas)

	for fila := range laberinto.filas {
		fmt.Scan(&laberinto.paredes[fila])
		laberinto.registrarSignos(fila)
	}

	return laberinto, true
}

func crearRegistros(filas int, columnas int) [][]registro {
	registros := make([][]registro, filas)

	for i := range filas {
		registros[i] = make([]registro, columnas)
	}

	return registros
}

func visitarVecinos(laberinto Laberinto, actual Posicion, registros [][]registro, pendientes TDACola.Cola[Posicion]) {
	for _, dir := range _DIRECCIONES {
		vecino := Posicion{actual.fila + dir.deltaFila, actual.columna + dir.deltaColumna}

		if !laberinto.esTransitable(vecino) || registros[vecino.fila][vecino.columna].visitado {
			continue
		}

		registros[vecino.fila][vecino.columna] = registro{true, actual, dir.signo}
		pendientes.Encolar(vecino)
	}
}

func recorrerLaberinto(laberinto Laberinto) [][]registro {
	registros := crearRegistros(laberinto.filas, laberinto.columnas)
	pendientes := TDACola.CrearColaEnlazada[Posicion]()

	registros[laberinto.inicio.fila][laberinto.inicio.columna].visitado = true
	pendientes.Encolar(laberinto.inicio)

	for !pendientes.EstaVacia() {
		actual := pendientes.Desencolar()

		if actual == laberinto.fin {
			break
		}

		visitarVecinos(laberinto, actual, registros, pendientes)
	}

	return registros
}

func reconstruirPasos(laberinto Laberinto, registros [][]registro) []string {
	pasos := []string{}

	for actual := laberinto.fin; actual != laberinto.inicio; {
		reg := registros[actual.fila][actual.columna]
		pasos = append(pasos, reg.direccion)
		actual = reg.origen
	}

	return pasos
}

func invertir(pasos []string) {
	for i := 0; i < len(pasos)/2; i++ {
		j := len(pasos) - 1 - i
		pasos[i], pasos[j] = pasos[j], pasos[i]
	}
}

func resolverLaberinto(laberinto Laberinto) ([]string, bool) {
	registros := recorrerLaberinto(laberinto)

	if !registros[laberinto.fin.fila][laberinto.fin.columna].visitado {
		return nil, false
	}

	pasos := reconstruirPasos(laberinto, registros)
	invertir(pasos)

	return pasos, true
}

func imprimirSolucion(pasos []string, encontrado bool) {
	if !encontrado {
		fmt.Println("ERROR")
		return
	}

	fmt.Println(len(pasos))
	fmt.Println(strings.Join(pasos, " "))
}

func main() {
	for {
		laberinto, exito := leerLaberinto()
		if !exito {
			break
		}

		pasos, encontrado := resolverLaberinto(laberinto)
		imprimirSolucion(pasos, encontrado)
	}
}
