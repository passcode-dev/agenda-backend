FROM golang:1.23-alpine

# Configurar o diretório de trabalho
WORKDIR /app

# Instalar ferramentas necessárias, incluindo uma versão estável do 'air'
RUN apk add --no-cache git \
    && go install github.com/cosmtrek/air@v1.41.0

# Copiar os arquivos do projeto para o container
COPY . .

# Verificar o conteúdo do diretório
RUN ls /app

# Comando para iniciar o servidor com o 'air'
CMD ["air", "-c", ".air.toml"]
