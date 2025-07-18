FROM golang:1.24

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

# Copiar os arquivos do microservice
COPY . .

# Instalar dependências
RUN go mod tidy

CMD ["go", "run", "infra/cmd/main.go"]