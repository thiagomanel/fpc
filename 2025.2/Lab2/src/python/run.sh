#!/usr/bin/env bash
# ==========================================
# Script para rodar a versão Serial (S)
# ou Concorrente (C) do programa de turmas.
# Uso:
#   bash run.sh S $1 $2   -> Serial
#   bash run.sh C $1 $2   -> Concorrente
# ==========================================

set -e

if [ "$#" -ne 3 ]; then
    echo "Uso: bash run.sh <S|C> <qtd_turmas> <qtd_alunos_por_turma>"
    exit 1
fi

MODE=$1
QTD_TURMAS=$2
QTD_ALUNOS=$3

if [ "$MODE" = "S" ]; then
    FILE="serial/main.py"
    echo "▶️  Rodando versão SERIAL..."
elif [ "$MODE" = "C" ]; then
    FILE="concurrent/main.py"
    echo "⚡ Rodando versão CONCORRENTE..."
else
    echo "Modo inválido. Use 'S' para serial ou 'C' para concorrente."
    exit 1
fi

time python3 "$FILE" "$QTD_TURMAS" "$QTD_ALUNOS"

