#!/bin/bash
set -e
cd "$(dirname "$0")/.."
zip -r lista/lista.zip go.mod lista/lista_enlazada.go lista/lista_test.go