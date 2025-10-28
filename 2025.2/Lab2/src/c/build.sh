#!/usr/bin/env bash
# ==========================================
# build.sh — Compila o programa Serial (S)
# ou Concorrente (C) e coloca dentro de .build/
# ==========================================

set -e

if [ "$#" -ne 1 ]; then
    echo "Uso: bash build.sh <S|C>"
    exit 1
fi

MODE=$1
BUILD_DIR=".build"

echo "Limpando diretório de build..."
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

if [ "$MODE" = "S" ]; then
    SRC="serial/main.c"
    OUT="$BUILD_DIR/serial"
    echo "▶️  Compilando versão SERIAL..."
elif [ "$MODE" = "C" ]; then
    SRC="concurrent/main.c"
    OUT="$BUILD_DIR/concurrent"
    echo "⚡ Compilando versão CONCORRENTE..."
else
    echo "Modo inválido. Use 'S' (serial) ou 'C' (concorrente)."
    exit 1
fi

if [ "$MODE" = "C" ]; then
    gcc -Wall -O2 -pthread "$SRC" -o "$OUT"
else
    gcc -Wall -O2 "$SRC" -o "$OUT"
fi
echo "Build concluído: $OUT"
