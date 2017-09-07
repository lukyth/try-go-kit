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
curl -X POST -d '{}' localhost:8080/get_messages
curl -X POST -d '{}' localhost:8080/get_message
curl -X POST -d '{}' localhost:8080/post_message
```