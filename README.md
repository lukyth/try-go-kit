This is me, trying out [go-kit](https://github.com/go-kit/kit)

Here's the plan

This'll be a server with two endpoints, `/add` and `/string`. Each endpoint will be responsible by its own service. There'll be a gateway to manage those endpoints.

This project use [dep](https://github.com/golang/dep) as a dependencies management.
```
brew install dep
```

## Usage
```bash
docker-compose up
```

```bash
curl -XPOST -d'{"s":"hello, world"}' localhost:8000/string/count
curl -XPOST -d'{"s":"hello, world"}' localhost:8000/string/uppercase
curl -XPOST -d'{"a":1,"b":2}' localhost:8000/add/sum
curl -XPOST -d'{"a":"1","b":"2"}' localhost:8000/add/concat
```