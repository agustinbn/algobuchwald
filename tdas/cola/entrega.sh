#!/bin/bash
set -e
cd "$(dirname "$0")/.."
zip -r cola/cola.zip go.mod cola/cola_enlazada.go cola/cola_test.go
