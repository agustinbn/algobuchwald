#!/bin/bash
set -e
cd "$(dirname "$0")/.."
rm -f tp1/tp1.zip
zip -r tp1/tp1.zip \
	tp1/go.mod tp1/go.sum tp1/maze.go tp1/pruebas \
	tdas/go.mod tdas/go.sum tdas/cola \
	-x 'tdas/cola/entrega.sh' 'tdas/cola/cola.zip'
