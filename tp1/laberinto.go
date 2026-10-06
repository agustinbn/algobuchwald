package main

type Posicion struct {
	fila    int
	columna int
}

type Laberinto struct {
	paredes  []string
	filas    int
	columnas int
	inicio   Posicion
	fin      Posicion
}

const (
	_SIGNO_INICIO byte = 'S'
	_SIGNO_FIN    byte = 'E'
	_SIGNO_PARED  byte = '#'
)

type Direccion struct {
	signo        string
	deltaFila    int
	deltaColumna int
}

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
