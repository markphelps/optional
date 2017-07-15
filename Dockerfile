# multi-stage builds require >= docker 17.05
FROM golang:1.8.1-alpine
MAINTAINER mark.aaron.phelps@gmail.com

# EXPOSE 8080

WORKDIR /go/src/github.com/markphelps/example
COPY . .
RUN go build -o ./bin/example ./cmd/example/...

FROM alpine:latest

# copy the binary from the `build-stage`
COPY --from=0 /go/src/github.com/markphelps/example/bin/example /bin
CMD "example"
