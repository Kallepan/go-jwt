# go-jwt

A simple JWT implementation in Go. This go project uses gin and raw sql to communicate with a postgres database. A simple login and register system is implemented.
I suggest using [Postman](https://www.postman.com/) to test the API. 

## Installation

Clone the repository and run the following command to install the dependencies.

```bash
go get
```

## Usage

- Edit the .env file to your liking.

```bash
docker-compose up -d
go run main.go
# or
bash launch.db.sh
bash launch.dev.sh
```
