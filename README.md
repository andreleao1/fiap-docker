# Aula Teste FIAP: Introdução ao Docker Compose

Bem-vindo à aula teste da FIAP sobre Docker Compose! Neste repositório, você encontrará exemplos práticos e explicações sobre como criar e utilizar um arquivo `docker-compose.yml` para gerenciar múltiplos contêineres Docker.

## Objetivo

O objetivo desta aula é fornecer uma introdução prática ao Docker Compose, mostrando como ele pode ser usado para definir e gerenciar ambientes de desenvolvimento multi-contêiner de forma simples e eficiente.

## Pré-requisitos

- Docker instalado na sua máquina. [Instruções de instalação](https://docs.docker.com/get-docker/)
- Docker Compose instalado. [Instruções de instalação](https://docs.docker.com/compose/install/)

## Conteúdo

1. [Introdução ao Docker Compose](#introdução-ao-docker-compose)
2. [Criando um arquivo docker-compose.yml](#criando-um-arquivo-docker-composeyml)
3. [Comandos básicos do Docker Compose](#comandos-básicos-do-docker-compose)
4. [Exemplo prático](#exemplo-prático)
5. [Referências](#referências)

## Introdução ao Docker Compose

Docker Compose é uma ferramenta para definir e gerenciar aplicativos Docker multi-contêiner. Com um simples arquivo YAML, você pode definir todos os serviços necessários para a sua aplicação e gerenciá-los com comandos simples.

## Criando um arquivo docker-compose.yml

Vamos criar um arquivo `docker-compose.yml` básico que define dois serviços: uma aplicação Go e um banco de dados PostgreSQL.

```yaml
version: '3.8'
services:
  go-app-users:
    image: poc_docker
    ports:
      - "8080:8080"
    depends_on:
      - postgres
    environment:
      - DB_HOST=postgres
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=postgres
      - DB_PORT=5432

  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: postgres
    ports:
      - "5432:5432"
```

## Explicação
- **version**: Define a versão do Docker Compose.
- **services**: Define os serviços que compõem a aplicação.
- **go-app-users**: Serviço que usa a imagem da aplicação Go.
- **ports**: Mapeia a porta 8080 do contêiner para a porta 8080 do host.
- **depends_on**: Define que o serviço depende do serviço `postgres`.
- **environment**: Define variáveis de ambiente para o contêiner.
- **postgres**: Serviço que usa a imagem do PostgreSQL.
- **environment**: Define variáveis de ambiente para o contêiner.
- **ports**: Mapeia a porta 5432 do contêiner para a porta 5432 do host.

## Comandos básicos do Docker Compose
- `docker-compose up`: Inicia todos os serviços definidos no arquivo `docker-compose.yml`.
- `docker-compose down`: Para e remove todos os contêineres, redes e volumes definidos no arquivo.
- `docker-compose ps`: Lista os contêineres em execução.
- `docker-compose logs`: Exibe os logs dos serviços.

## Exemplo prático
Clone este repositório:
```bash
git clone https://github.com/andreleao1/fiap-docker.git
cd fiap-docker
docker compose up -d
```