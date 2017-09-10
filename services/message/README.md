# Message Service

## Installation
### Install dependencies
```bash
dep ensure
```
### Build a service binary
```bash
./build.sh
```

## Usage
```bash
./bin/message
```
```bash
curl -X GET localhost:8080/messages
curl -X GET localhost:8080/messages/1
curl -X POST -d '{"m":{"id":"3", "body":"3 message"}}' localhost:8080/messages
```