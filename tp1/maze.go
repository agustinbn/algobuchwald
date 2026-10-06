package main

import (
	"fmt"
	"strings"
)

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
