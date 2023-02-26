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

## Build docker image 

1. Run this command
   
```
docker build -t fyfirman-blog-service
```

2. Run the image with this command

```
docker run -p 8080:8080 fyfirman-blog-service 
```