#Nesse primeiro estágio, que irei chamar de builder, irei empacotar todas as dependências do projeto e compilar o projeto deixando-o pronto para ser executado.

# Usando como base a imagem oficial do Golang na versão 1.23 com Alpine
FROM golang:1.23-alpine AS builder

# Definindo o diretório app como diretório de trabalho
WORKDIR /app

#Copiando os arquivos go.mod e go.sum, que são os arquivos que contém as dependências do projeto, para o diretório app
COPY go.mod go.sum ./

# Baixando as dependências do projeto
RUN go mod download

# Copiando todos os arquivos do diretório atual para o diretório app
COPY . .

# Compilando o projeto
RUN go build -o myapp

# Nesse segundo estágio, irei criar uma imagem mínima do Alpine Linux e copiar o binário do projeto compilado no primeiro estágio para a imagem.
FROM alpine:latest

# Definindo o diretório /root como diretório de trabalho
WORKDIR /root/

# Copiando o binário do projeto compilado no primeiro estágio para o diretório /root
COPY --from=builder /app/myapp .

# Expondo a porta 8080
EXPOSE 8080

# Executando o binário do projeto
CMD ["./myapp"]