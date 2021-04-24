FROM golang:1.16-alpine AS build_base

# change working dir to /go/src/avaas
# copy all project files to this directory
WORKDIR /go/src/tsubot
COPY . .

# install dependencies
# build the binary
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o bot cmd/main.go

# run the built binary
ENTRYPOINT ["./bot", "-c=cmd/config.yaml"]