# Aula Teste FIAP: Introdução ao Dockerfile

Bem-vindo à aula teste da FIAP sobre Dockerfile! Neste repositório, você encontrará exemplos práticos e explicações sobre como criar e utilizar um Dockerfile para construir imagens Docker personalizadas.

## Objetivo

O objetivo desta aula é fornecer uma introdução prática ao Dockerfile, mostrando como ele pode ser usado para definir e construir imagens Docker de forma simples e eficiente.

## Pré-requisitos

- Docker instalado na sua máquina. [Instruções de instalação](https://docs.docker.com/get-docker/)

## Conteúdo

1. [Introdução ao Dockerfile](#introdução-ao-dockerfile)
2. [Criando um Dockerfile](#criando-um-dockerfile)
3. [Comandos básicos do Docker](#comandos-básicos-do-docker)
4. [Exemplo prático](#exemplo-prático)
5. [Referências](#referências)

## Introdução ao Dockerfile

Dockerfile é um arquivo de texto que contém uma série de instruções para construir uma imagem Docker. Com ele, você pode definir o ambiente e as dependências necessárias para a sua aplicação, garantindo que ela seja executada de forma consistente em qualquer lugar.

## Criando um Dockerfile

Vamos criar um Dockerfile básico que define uma aplicação Node.js.

```dockerfile
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
```

## Explicação
- **FROM**: Define a imagem base que será usada.
- **WORKDIR**: Define o diretório de trabalho dentro do contêiner.
- **COPY**: Copia arquivos do host para o contêiner.
- **RUN**: Executa comandos no contêiner durante a construção da imagem.
- **EXPOSE**: Informa ao Docker que a aplicação irá rodar na porta especificada.
- **CMD**: Define o comando que será executado quando o contêiner iniciar.

## Comandos básicos do Docker
- `docker build -t nome-da-imagem .`: Constrói uma imagem Docker a partir do Dockerfile no diretório atual.
- `docker run -p 8080:8080 nome-da-imagem`: Executa um contêiner baseado na imagem criada, mapeando a porta 3000 do host para a porta 3000 do contêiner.
- `docker ps`: Lista os contêineres em execução.
- `docker stop id-do-conteiner`: Para um contêiner em execução.
- `docker rm id-do-conteiner`: Remove um contêiner parado.

## Exemplo prático
Clone este repositório:
```shell
git clone https://github.com/andreleao1/fiap-docker.git
cd fiap-dockerfile
docker build -t poc_docker .
docker run -p 8080:8080 poc_docker
```

## Referências
- [Documentação oficial do Docker](https://docs.docker.com/)
- [Guia de referência do Dockerfile](https://docs.docker.com/engine/reference/builder/)