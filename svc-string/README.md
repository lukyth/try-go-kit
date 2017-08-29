This is based on [go-kit/kit/examples/stringsvc2](https://github.com/go-kit/kit/tree/v0.5.0/examples/stringsvc2).

## Installation
```bash
dep ensure
```

## Usage
```bash
cd cmd
go build
./cmd
```
```bash
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8080/count
curl localhost:8080/metrics
```