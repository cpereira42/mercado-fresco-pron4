<h1 align="center">Mercado Fresco Pron4 </h1>
<br>
<br>
<br>

![Meli](https://anymarket.com.br/wp-content/uploads/2018/07/images.png)

<br>
<br>
<br>
<p align="center">
  <img alt="Go" src="https://img.shields.io/badge/technology-go-242D7B.svg">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/cpereira42/mercado-fresco-pron4?color=FFF159">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/cpereira42/mercado-fresco-pron4?color=242D7B">

  <a href="https://github.com/cpereira42/mercado-fresco-pron4/commits/master">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/cpereira42/mercado-fresco-pron4?color=FFF159">
  </a>

  <img alt="License" src="https://img.shields.io/badge/license-MIT-242D7B">

</p>

## RestAPI (Golang)

## Description

The objective of the “Final Project” is to implement a REST API within the structure of the utterance and
apply the contents worked during BOOTCAMP-GO MELI. (Git, GO, Storage and Quality).

## Features

The following routes are found in the API:

>- [Seller](docs/SELLER.md)
>- [Warehouse](docs/WAREHOUSE.md)
>- [Section](docs/SECTION.md)
>- [Product](docs/PRODUCT.md)
>- [Employee](docs/EMPLOYEE.md)
>- [Buyer](docs/BUYER.md)
>- [Locality](docs/LOCALITY.md)
>- [Carry](docs/CARRIES.md)
>- [Product Batch](docs/PRODUCTBATCHS.md)
>- [Product Record](docs/PRODUCTRECORDS.md)
>- [Inbound Order](docs/INBOUNDORDERS.md)
>- [Purchase Order](docs/PURCHASEORDERS.md)


## Technologies

The following tools were used in this project:

> - [Go](https://go.dev/)
> - [Godotenv](https://github.com/joho/godotenv)
> - [Testify](https://github.com/stretchr/testify)
> - [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)
> - [Mockery](https://github.com/vektra/mockery)
> - [Gin](https://gin-gonic.com/)
> - [Docker](https://www.docker.com/)
> - [Git](https://git-scm.com/)

## Installation

First you need to have [Go](https://go.dev/) version 1.17, [Git](https://git-scm.com/) and [Docker](https://www.docker.com/) installed.

#### Clone the repository  
```shell
 git clone git@github.com:cpereira42/mercado-fresco-pron4.git
```
#### Install the dependencies
```shell 
cd mercado-fresco-pron4
go mod tidy
```
#### Create database with docker
```shell
docker-compose up
```
#### Create .env file in the root of the project and set the following values
```dosini
HOST_DB=
USER_DB=
PASS_DB=
PORT_DB=
DATABASE=
```
#### Start the API
```shell
go run ./...
```

## Team

| Eduardo Araújo | Vinicius Oliveira | Cezar Pereira | Adriana Rosa |José Neto| Eneas Sena| 
| :---: | :---: | :---: | :---: | :---: | :---: |
|[<img src="https://avatars.githubusercontent.com/eduaraujogf" width=140><br><sub></sub>](https://github.com/eduaraujogf)|[<img src="https://avatars.githubusercontent.com/runnice" width=140><br><sub></sub>](https://github.com/runnice) |[<img src="https://avatars.githubusercontent.com/cpereira42" width=140><br><sub></sub>](https://github.com/cpereira42) |[<img src="https://avatars.githubusercontent.com/adikrosa" width=140><br><sub></sub>](https://github.com/adikrosa) |[<img src="https://avatars.githubusercontent.com/JAMNeto" width=140><br><sub></sub>](https://github.com/JAMNeto) |[<img src="https://avatars.githubusercontent.com/eneassena" width=140><br><sub></sub>](https://github.com/eneassena) 




