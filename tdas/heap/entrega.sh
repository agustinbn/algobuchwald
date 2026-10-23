#!/bin/bash
set -e
cd "$(dirname "$0")/.."
zip -r heap/heap.zip go.mod heap/ -x 'heap/entrega.sh' 'heap/heap.zip'
