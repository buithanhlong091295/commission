## Requirements

Building and running requires:

- [docker](https://docs.docker.com) install docker
- [docker-compose](https://docs.docker.com/compose/) install docker compose
- [golang](https://golang.org/) version 1.11+
- [protobuf](https://github.com/protocolbuffers/protobuf/releases) version 3.7+
- [postgresql](https://www.postgresql.org/download/) version 11+

## Install Protobuf

```sh
# for linux
$ curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip

$ unzip protoc-3.7.0-linux-x86_64.zip -d protoc3
$ mv protoc3/bin/* /usr/local/bin/
$ mv protoc3/include/* /usr/local/include/
```

## Install protoc gen validate

```sh
# fetches this repo into $GOPATH
$ go get -d github.com/envoyproxy/protoc-gen-validate

# move point to dir protoc-gen-validate
$ cd $GOPATH/src/github.com/envoyproxy/protoc-gen-validate
# installs PGV into $GOPATH/bin
$ make build
```

## Getting started

#### Install dependencies (go proto, grpc gateway)

```sh
# install protoc-gen-go cli
$ go get -u github.com/grpc-ecosystem/grpc-gateway
$ go install github.com/golang/protobuf/proto
$ go install github.com/golang/protobuf/protoc-gen-go
$ go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

## Clone source

```sh
$ git clone --recursive https://gitlab.com/citicoin-sources/citicoin-backend.git $GOPATH/src/xtek/exchange/user
```

## Start, Run & Build

```sh
$ cd $GOPATH/src/xtek/exchange/user

# Use docker compose run redis, postgres, pbadmin, consul, jaeger
$ make docker-dev

# Generate protobuf to ./pb
$ make proto-gen

$ make build

# Used to start user service
$ ./user_srv start

# or enable sql auto migrate
$ ./user_srv --sql.auto_migrate start
```

## Postgres Database connection for docker compose development

#### Postgress connection

```json
{
	"host": "postgres",
	"port": 5432,
	"username": "user",
	"password": "password0000",
	"database": "citicoins"
}
```

#### PbAdmin postgres

```json
{
	"email": "pbadmin@citicoins.io",
	"password": "pass0000"
}
```
