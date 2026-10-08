#!/bin/bash
set -e
cd "$(dirname "$0")/.."
zip -r diccionario/diccionario.zip go.mod diccionario/hash.go