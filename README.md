# Aula FIAP: Introdução ao Docker

Bem-vindo à aula da FIAP sobre Docker! Este repositório contém materiais e exemplos práticos para ajudar você a entender e utilizar o Docker de forma eficiente.

## Objetivo

O objetivo desta aula é fornecer uma introdução prática ao Docker, abordando os conceitos fundamentais e mostrando como utilizá-los no dia a dia. Você aprenderá sobre:

- O que é Docker
- O que são Imagens Docker
- Docker Registry
- Dockerfile
- Docker Compose

## Pré-requisitos

- Docker instalado na sua máquina. [Instruções de instalação](https://docs.docker.com/get-docker/)

## Conteúdo

1. [O que é Docker](#o-que-é-docker)
2. [O que são Imagens Docker](#o-que-são-imagens-docker)
3. [Docker Registry](#docker-registry)
4. [Dockerfile](#dockerfile)
5. [Docker Compose](#docker-compose)
6. [Exemplos Práticos](#exemplos-práticos)
7. [Referências](#referências)

## O que é Docker

Docker é uma plataforma que permite criar, implantar e executar aplicações em contêineres. Contêineres são leves, portáteis e garantem que a aplicação rode de forma consistente em qualquer ambiente.

## O que são Imagens Docker

Imagens Docker são templates imutáveis que contêm tudo o que é necessário para rodar uma aplicação, incluindo o código, runtime, bibliotecas e dependências.

## Docker Registry

Docker Registry é um serviço para armazenar e distribuir imagens Docker. O Docker Hub é um registro público popular, mas você também pode configurar registros privados.

## Dockerfile

Dockerfile é um arquivo de texto que contém uma série de instruções para construir uma imagem Docker. Com ele, você pode definir o ambiente e as dependências necessárias para a sua aplicação.

### Exemplo de Dockerfile

Confira o exemplo prático de um Dockerfile [aqui](https://github.com/andreleao1/fiap-docker/blob/main/dockerfile).

## Docker Compose

Docker Compose é uma ferramenta para definir e gerenciar multi-contêineres Docker. Com ele, você pode usar um arquivo YAML para configurar os serviços da sua aplicação.

### Exemplo de Docker Compose

Confira o exemplo prático de um Docker Compose [aqui](https://github.com/andreleao1/fiap-docker/blob/main/docker-compose.yaml).

## Exemplos Práticos

Clone este repositório e siga os passos abaixo para executar os exemplos práticos:

```shell
git clone https://github.com/andreleao1/fiap-docker.git
cd fiap-docker
```

## Executando o Dockerfile

```shell
docker build -t poc_docker .
docker run -p 8080:8080 poc_docker
```

## Executando o Docker Compose

```shell
docker-compose up
```

## Material da Aula

O PDF da aula está disponível no diretório material-da-aula com o nome do arquivo FIAP_AULA_TESTE_DOCKER.pdf. Você pode acessar o material [aqui](https://github.com/andreleao1/fiap-docker/blob/main/docker-compose.yaml).


## Referências
- [Documentação oficial do Docker](https://docs.docker.com/)
- [Guia de referência do Dockerfile](https://docs.docker.com/engine/reference/builder/)
- [Documentação do Docker Compose](https://docs.docker.com/compose/)