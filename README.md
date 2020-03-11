# Engine

This service provides the business logic including changing account data in the database.
Runs on port: 6000

### First time setup

```
go get ./...
```

### Running the application

```
go run .\engine\cmd\server.go
```

### Secret Token needed to contact this service
Authorization: SecretInternalKey