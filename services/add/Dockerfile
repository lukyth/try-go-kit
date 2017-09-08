FROM golang:1.9.0 AS builder

RUN go get -u github.com/golang/dep/cmd/dep

RUN mkdir -p /go/src/github.com/lukyth/try-go-kit/services/add
WORKDIR /go/src/github.com/lukyth/try-go-kit/services/add

# copies the Gopkg.toml and Gopkg.lock to WORKDIR
COPY Gopkg.toml Gopkg.lock ./

# install the dependencies without checking for go code
RUN dep ensure -vendor-only

# copies any other required code to the WORKDIR
COPY . .

# the -ldflags '-d' -tags netgo -installsuffix netgo is used to avoid
# standard_init_linux.go:187: exec user process caused "no such file or directory"
# see: https://www.ianlewis.org/en/creating-smaller-docker-images-static-binaries
# and https://github.com/docker-library/golang/issues/152
# the -o flag is used to specify the name of the executable to be installed
RUN go build -ldflags '-d' -tags netgo -installsuffix netgo -o bin/add cmd/add/add.go

FROM alpine:3.6

WORKDIR /root/

# copies only the built executable to the new image
COPY --from=builder /go/src/github.com/lukyth/try-go-kit/services/add/bin/add  .

ENTRYPOINT ["./add"]