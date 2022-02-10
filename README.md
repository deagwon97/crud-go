# CRUD api on Go

## Development

### 1. docker-compose up
```shell
$ cd dev
$ docker-compose up
```
### 2. Attatch vscode
<div align="center">
    <img src="ref/attatch-vscode.png" style="width:60%;" />
</div>

### 3. Run main.go on debug

<div align="center">
    <img src="ref/run-debug.png" style="width:60%;"></img>
</div>

## Build deployment image

```
$ docker build -t go-api:latest .

$ docker run -p 8000:8000 --env-file .env go-api:latest
```

## Reference
  - Hands-On Full-Stack Development with Go
  - https://github.com/gin-gonic/gin


