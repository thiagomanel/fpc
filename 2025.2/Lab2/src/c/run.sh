#!/usr/bin/env bash
# ==========================================
# run.sh — Executa o programa compilado
# ==========================================

set -e

if [ "$#" -ne 2 ]; then
    echo "Uso: bash run.sh <qtd_turmas> <qtd_alunos_por_turma>"
    exit 1
fi

BUILD_DIR=".build"
QTD_TURMAS=$1
QTD_ALUNOS=$2

if [ -f "$BUILD_DIR/concurrent" ]; then
    EXEC="$BUILD_DIR/concurrent"
    echo "⚡ Rodando versão CONCORRENTE..."
elif [ -f "$BUILD_DIR/serial" ]; then
    EXEC="$BUILD_DIR/serial"
    echo "▶️ Rodando versão SERIAL..."
else
    echo "Nenhum build encontrado. Rode primeiro: bash build.sh <S|C>"
    exit 1
fi

time "$EXEC" "$QTD_TURMAS" "$QTD_ALUNOS"

