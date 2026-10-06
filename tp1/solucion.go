package main

import TDACola "tdas/cola"

type registro struct {
	visitado  bool
	origen    Posicion
	direccion string
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
