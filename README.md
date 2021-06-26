<img align="right" width=200px height=200px src="https://cdn.discordapp.com/attachments/776153365452554301/786297555415859220/Tech-With-Tim.png" alt="Project logo">

<h1>Tech With Tim - CDN</h1>

<div>
  
[![Status](https://img.shields.io/website?url=https%3A%2F%2Fcdn.twtcodejam.net)](https://cdn.dev.twtcodejam.net)
[![codecov](https://codecov.io/gh/Tech-With-Tim/cdn/branch/main/graph/badge.svg?token=YKpXOrUO80)](https://codecov.io/gh/Tech-With-Tim/cdn)
[![Lint & Test](https://github.com/Tech-With-Tim/cdn/actions/workflows/lint-test.yml/badge.svg)](https://github.com/Tech-With-Tim/cdn/actions/workflows/lint-test.yml)
[![GitHub Issues](https://img.shields.io/github/issues/Tech-With-Tim/API.svg)](https://github.com/Tech-With-Tim/API/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/Tech-With-Tim/API.svg)](https://github.com/Tech-With-Tim/API/pulls)
[![Licence](https://img.shields.io/badge/licence-MIT-blue.svg)](/LICENCE)
[![Discord](https://discord.com/api/guilds/501090983539245061/widget.png?style=shield)](https://discord.gg/twt)

</div>

CDN for the Tech With Tim website using [Go](https://go.dev/)

## 📝 Table of Contents
- [🏁 Getting Started](#-getting-started)
  - [Environment variables](#environment-variables)
  - [Running](#running)
- [🐳 Running with Docker](#-running-with-docker)
- [🚨 Tests](#-tests)
- [📜 Licence](/LICENCE)
- [⛏️ Built Using](#️-built-using)
- [✍️ Authors](#️-authors)


## 🏁 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [Running with Docker](#-running-with-docker) if you want to setup the CDN faster with Docker. ( Docker is optional )

### Environment variables

Set the environment variables. Start by writing this in a file named `app.env` and `test.env`:
(test.env is required for running tests)

app.env and test.env should look like this:

```prolog
POSTGRES_USER=user
POSTGRES_PASSWORD=pwd
DB_NAME=db
DB_HOST=localhost
DB_PORT=5432
SECRET_KEY=secret
MAX_FILE_SIZE=30
```
- ``SECRET_KEY`` is the key used for the JWT token encoding.
- ``MAX_FILE_SIZE`` is the maxiumum file size allowed in asset upload (in mb)

### Running

- To create postgres container `make postgres`
- To create db `make createdb`
- To drop db `make dropdb`

#### run `go mod tidy` to install packages
#### cli commands 
```
go run main.go migrate_up
go run main.go dropdb
go run main.go migrate_steps --steps int
go run main.go runserver --host localhost --port port (localhost, 5000 are default)
```

#### to run migrations on the test database 
```
go run main.go migrate_up -t
go run main.go dropdb -t
go run main.go migrate_steps -t --steps int
```

### Use the make file its your best friend 🛠
#### Make commands
##### If you are on windows please use git bash or wsl also you would have to install make for windows
##### to install make for windows run `winget install GnuWin32.Make`

```shell
make postgres #creates docker container for postgres12
# reads env variables from app.env
make createdb #creates the db in the postgres container
make dropdb #drops the db
make migrate_up #migrates to the latest schema
make sqlc_generate #generates sqlc code if you write queries
make test # tests your code and shows coverage
#its a big output make sure to read it all
```

## 🐳 Running with Docker

Start the cdn `docker-compose up`

## 🚨 Tests
To test the cdn we can use two methods
```sh
make test
```
If you don't have make installed
```sh
go run main.go migrate_up -t
go test ./... -v 
```
**When you contribute, you need to add tests on the features you add.**

## ⛏️ Built Using

- [Go](https://go.dev/) - Language
- [go-chi](https://github.com/go-chi/chi) - Router
- [sqlc](https://github.com/kyleconroy/sqlc) - Database Query Helper

## ✍️ Authors
See the list of [contributors](https://github.com/Tech-With-Tim/cdn/contributors) who participated in this project.

