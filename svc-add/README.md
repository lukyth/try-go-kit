This is based on [go-kit/kit/examples/addsvc](https://github.com/go-kit/kit/tree/v0.5.0/examples/addsvc).

## Installation
```bash
dep ensure
```

## Usage
### Server
```bash
cd cmd/svc-add
go build
./svc-add
```
```
curl -XPOST -d'{"a":1, "b":2}' localhost:8081/sum
curl -XPOST -d'{"a":"1", "b":"2"}' localhost:8081/concat
```