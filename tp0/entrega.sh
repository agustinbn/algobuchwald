#!/bin/bash
set -e
cd "$(dirname "$0")"
zip -r tp0.zip main.go go.mod ejercicios/tp0.go ejercicios/tp0_test.go
