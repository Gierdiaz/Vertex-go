# Etapa de build
FROM golang:1.23-alpine as builder

# Definir o diretório de trabalho no contêiner
WORKDIR /app

# Copiar go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod tidy

# Copiar o restante do código para o contêiner
COPY . ./

COPY .env .env

# Compilar a aplicação Go
RUN go build -o main cmd/main.go

# Etapa de execução
FROM alpine:latest

# Instalar dependências necessárias (como ca-certificates para comunicação segura)
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho
WORKDIR /root/

# Copiar o arquivo compilado do estágio anterior
COPY --from=builder /app/main .

# Expor a porta que a aplicação vai rodar
EXPOSE 3000

# Definir o comando para rodar a aplicação
CMD ["./main"]
