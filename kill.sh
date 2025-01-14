#!/bin/bash

# Porta onde o servidor está rodando
PORT=8080

# Encontrar o PID do processo que está usando a porta
PID=$(lsof -t -i :$PORT)

if [ -n "$PID" ]; then
    echo "Matando processo na porta $PORT (PID: $PID)"
    kill -9 $PID
else
    echo "Nenhum processo encontrado na porta $PORT"
fi
