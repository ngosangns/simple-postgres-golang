# Simple Setup for PostgreSQL in Golang

This repository provides a streamlined setup for integrating PostgreSQL with a Golang application.

## Dependencies

- Generate type-safe code from SQL: https://github.com/sqlc-dev/sqlc
- Migration tool: https://github.com/pressly/goose
- PostgresSQL driver: https://github.com/jackc/pgx

## Usage

```
make cp-env
make mig-up
```

---

For more commands, refer to the `makefile`.