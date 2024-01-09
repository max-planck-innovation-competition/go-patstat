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

## Author

Sebastian Erhardt

## Documentation

The documentation of patstat can be
found [here](https://documents.epo.org/projects/babylon/eponot.nsf/0/9BB068EEB37E80BCC125878200565B60/$File/patstat_data_catalog_global_5_18_en.pdf)
.

## Environment Variables
```
SQL_HOST=localhost
SQL_USER=admin
SQL_PASSWORD=password
SQL_DATABASE=patstat
SQL_PORT=5432
SQL_TYPE=POSTGRES
```

These can be set in a `.env` file in the root directory.

## Deployments

* [PostgreSQL Docker-Compose](./deployments/postgres)
