# Pokkne

## Database

### Generate query client

```
sqlc generate
```

### Migrate Up

```
goose postgres "host=localhost port=5432 user=postgresdev password=postgresdev dbname=postgresdev sslmode=disable" up
```

### Migrate Down

```
goose postgres "host=localhost port=5432 user=postgresdev password=postgresdev dbname=postgresdev sslmode=disable" down
```
