This is based on [go-kit/kit/examples/addsvc](https://github.com/go-kit/kit/tree/v0.5.0/examples/addsvc).

## Installation
```bash
dep ensure
```

## Usage
### Server
```bash
cd cmd/addsvc
go build
./addsvc
```
### Client
```bash
cd cmd/addcli
go build
./addcli
```
```bash
./addcli -http-addr=:8081 -method=concat 1 2
./addcli -http-addr=:8081 -method=sum 1 2
```