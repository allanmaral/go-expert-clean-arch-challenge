# Desafio Go - Clean Architecture

Bem-vindo ao Desafio de Clean Architecture da Pós-Graduação Go Expert! Este projeto consiste em um serviço de cadastro e listagem de `orders` que expõe uma API REST, um servidor GraphQL e um serviço gRPC.

## Pré-requisitos

Antes de começar, certifique-se de ter instalado os seguintes requisitos:

- [Go SDK](https://golang.org/dl/): Linguagem de programação Go.
- [Docker](https://docs.docker.com/get-docker/): Plataforma de conteinerização.
- [Make](https://www.gnu.org/software/make/): Utilizado para automatização de tarefas.

## Executando o Projeto

1. Clone este repositório em sua máquina local:

   ```bash
   git clone https://github.com/allanmaral/go-expert-clean-arch-challenge.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd go-expert-clean-arch-challenge
   ```

3. Execute o seguinte comando para subir o banco de dados MySQL e uma instância do RabbitMQ:

   ```bash
   docker compose up -d
   ```

4. Instale as dependências do projeto:

   ```bash
   go mod tidy
   ```

5. Finalmente, suba o serviço executando:

   ```bash
   make run
   ```

## Acesso aos Serviços

Após subir o serviço, você poderá acessar as seguintes interfaces:

- **API REST**: http://localhost:8000
- **Servidor GraphQL**: http://localhost:8080
- **Serviço gRPC**: Porta 50051

## Documentação da API REST

A documentação das rotas do servidor HTTP está disponível no arquivo `./api/api.http`.
