#!/bin/bash
set -e
cd "$(dirname "$0")/.."
zip -r pila/pila.zip go.mod pila/pila_dinamica.go pila/pila_test.go
