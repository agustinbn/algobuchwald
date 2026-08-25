package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

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
	procesar_archivo("archivo1.in", &vector_1)
	procesar_archivo("archivo2.in", &vector_2)
	if ejercicios.Comparar(vector_1, vector_2) == 1 {
		ejercicios.Seleccion(vector_1)
		for i := range vector_1 {
			fmt.Println(vector_1[i])
		}
	} else {
		ejercicios.Seleccion(vector_2)
		for i := range vector_2 {
			fmt.Println(vector_2[i])
		}
	}
}
