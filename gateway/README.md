## Prerequisite
svc-add and svc-string need to be ran first.

## Installation
```bash
dep ensure
go build
./gateway
```

## Usage
```bash
curl -XPOST -d'{"a":"1","b":"2"}' localhost:8000/add/concat
curl -XPOST -d'{"a":1,"b":2}' localhost:8000/add/sum
curl -XPOST -d'{"s":"hello, world"}' localhost:8000/string/uppercase
curl -XPOST -d'{"s":"hello, world"}' localhost:8000/string/count
```