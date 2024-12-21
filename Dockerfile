FROM golang:1.22.2-alpine

WORKDIR /app

COPY agenda-backend/mvc/go.mod agenda-backend/mvc/go.sum ./
RUN go mod download

COPY agenda-backend/mvc .
RUN ls /app  # Verificar o conteúdo do diretório
RUN go build -o app .

CMD ["./app"]
