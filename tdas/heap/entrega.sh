#!/bin/bash
set -e
cd "$(dirname "$0")/.."
# la corrección espera el directorio cola_prioridad/ (no heap/) e imports tdas/cola_prioridad
tmp=$(mktemp -d)
mkdir "$tmp/cola_prioridad"
cp go.mod "$tmp"
cp heap/*.go "$tmp/cola_prioridad/"
sed -i.bak 's#"tdas/heap"#"tdas/cola_prioridad"#' "$tmp"/cola_prioridad/*.go
rm "$tmp"/cola_prioridad/*.bak
rm -f heap/heap.zip
(cd "$tmp" && zip -r "$OLDPWD/heap/heap.zip" go.mod cola_prioridad)
rm -rf "$tmp"
