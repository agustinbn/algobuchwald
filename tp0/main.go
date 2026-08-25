package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const ARCHIVO_1 string = "archivo1.in"
const ARCHIVO_2 string = "archivo2.in"

func procesar_archivo(ruta string, vector *[]int) {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta, err)
		return
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	i := 0
	for s.Scan() {
		var ts int
		ts, err = strconv.Atoi(s.Text())
		*vector = append(*vector, ts)
		i++
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	vector_1 := []int{}
	vector_2 := []int{}
	procesar_archivo(ARCHIVO_1, &vector_1)
	procesar_archivo(ARCHIVO_2, &vector_2)
	mayor := vector_2
	if ejercicios.Comparar(vector_1, vector_2) == 1 {
		mayor = vector_1
	}
	ejercicios.Seleccion(mayor)
	for i := range mayor {
		fmt.Println(mayor[i])
	}
}
