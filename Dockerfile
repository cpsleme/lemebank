# Imagem base
FROM golang:1.16-alpine

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos necessários para o diretório de trabalho
COPY . .

# Instala as dependências
RUN go mod download

# Compila o código
RUN go build -o app

# Expõe a porta necessária para a aplicação
EXPOSE 8080

# Executa o aplicativo
CMD ["./app"]
