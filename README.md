# fyfirman-blog-service

## How to run

```
  go run src/main.go
```

## How to run with live reload

1. Install gin
   
```
go install github.com/codegangsta/gin@latest
```

2. Run with

```
gin --appPort 8080 --path ./cmd run main.go
```