## Tech-With-Tim CDN

[![codecov](https://codecov.io/gh/Tech-With-Tim/cdn/branch/main/graph/badge.svg?token=YKpXOrUO80)](https://codecov.io/gh/Tech-With-Tim/cdn)
[![Lint & Test](https://github.com/Tech-With-Tim/cdn/actions/workflows/lint-test.yml/badge.svg)](https://github.com/Tech-With-Tim/cdn/actions/workflows/lint-test.yml)

app.env and test.env should look like this
(test.env is required for running tests)

```
POSTGRES_USER=user
POSTGRES_PASSWORD=pwd
DB_NAME=db
DB_HOST=localhost
DB_PORT=5432
SECRET_KEY=secret
MAX_FILE_SIZE=30
```
to create postgres container `make postgres`
to create db `make createdb`
to drop db `make dropdb`

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
