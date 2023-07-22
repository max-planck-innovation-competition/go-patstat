# Patstat

[![Go Report Card](https://goreportcard.com/badge/github.com/max-planck-innovation-competition/go-patstat)](https://goreportcard.com/report/github.com/max-planck-innovation-competition/go-patstat)
[![Go Reference](https://pkg.go.dev/badge/github.com/max-planck-innovation-competition/go-patstat.svg)](https://pkg.go.dev/github.com/max-planck-innovation-competition/go-patstat)

This repository generates the PATSTAT Database using Object Relational Mapping (ORM) and Docker.

This allows for creation of PATSTAT in:

* MySQL
* PostgreSQL
* SQLite

## Status

Under development

**⚠️ Experimental - Not ready for production.**

Tested with:
* Fall 2019
* Spring 2023


## Author

Sebastian Erhardt

## Documentation

The documentation of patstat can be
found [here](https://documents.epo.org/projects/babylon/eponot.nsf/0/9BB068EEB37E80BCC125878200565B60/$File/patstat_data_catalog_global_5_18_en.pdf)
.

## Deployments

* [MySQL Docker-Compose](./deployments/mysql)
* [PostgreSQL Docker-Compose](./deployments/postgres)
* [SQLite Docker-Compose](./deployments/sqlite)

## Go Module

```shell
go get -u github.com/max-planck-innovation-competition/go-patstat
```