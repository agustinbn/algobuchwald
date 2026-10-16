#!/bin/bash
set -e
cd "$(dirname "$0")/.."
zip -r diccionario/abb.zip go.mod diccionario/diccionario.go diccionario/diccionario_ordenado.go diccionario/abb.go diccionario/diccionario_ordenado_test.go pila/ -x 'pila/entrega.sh' 'pila/pila.zip'