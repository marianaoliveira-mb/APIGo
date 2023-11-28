# APIGo

O objetivo do projeto é a criação de uma API voltada para a área de vendas(arquitetura simples)

## Pré-requisitos

- Go
- Docker
- PostgreSQL

## Instalação

- Go: https://go.dev/doc/install
- Docker: https://docs.docker.com/get-docker/
- ProstgreSQL: https://www.postgresql.org/download/

## Setup

*1.* Clonar o repositório
```git clone git@github.com:marianaoliveira-mb/APIGo.git```

*2.* Instalar imagem docker-hub
```docker pull postgres:14.3-alpine```
```docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14.3-alpine```

*3.* Instalar migration
-- https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

*4.* Rodar as migrations
```migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up```

*5.* Executar o docker
```docker-compose up```

*6.* Executar o servidor
```go run main.go```
