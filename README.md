# Patstat

[![Go Report Card](https://goreportcard.com/badge/github.com/max-planck-innovation-competition/go-patstat)](https://goreportcard.com/report/github.com/max-planck-innovation-competition/go-patstat)
[![Go Reference](https://pkg.go.dev/badge/github.com/max-planck-innovation-competition/go-patstat.svg)](https://pkg.go.dev/github.com/max-planck-innovation-competition/go-patstat)

Updated to Patstat Version Spring 2023

This repository generates the PATSTAT Database using Object Relational Mapping (ORM) and Docker.

This allows for creation of PATSTAT for various databases:
* MySQL
* PostgreSQL
* SQLite

## Status

Under development

**⚠️ Experimental - Not ready for production.**

## Author

Sebastian Erhardt

## Documentation

The documentation of patstat can be
found [here](https://documents.epo.org/projects/babylon/eponot.nsf/0/9BB068EEB37E80BCC125878200565B60/$File/patstat_data_catalog_global_5_18_en.pdf)
.

## Deployments

* [PostgreSQL Docker-Compose](./deployments/postgres)

**TODO:**
* [SQLite Docker-Compose](./deployments/sqlite)
* [MySQL Docker-Compose](./deployments/mysql)

## Go Module

```shell
go get -u github.com/max-planck-innovation-competition/go-patstat
```

## Binary

1) Download Patstat via [EPO Bulk Data Download](https://publication-bdds.apps.epo.org/raw-data/products)
2) Set-Up Database Docker e.g. Postgres [docker-compose.yml](deployments%2Fpostgres%2Fdocker-compose.yml)
3) Move Patstat Data to Database Folder
4) Unzip Patstat Data
```shell
# unzip part files 
unzip *.zip
# unzip tls files
unzip tls*.zip
```
5) Create `.env` file with following content
```
SQL_HOST=localhost
SQL_USER=postgres
SQL_PASSWORD=changeme
SQL_PORT=5432
SQL_DATABASE=patstat
SQL_TYPE=POSTGRES
```
6) Build the binary
```shell
echo BUILD BINARY
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o patstat ./cmd/main.go
```

7) Run Patstat installation Script

```shell
./patstat install --db=patstat_2024_spring --directory=./ --postgres-directory=/var/lib/postgresql/data/ingest/
```