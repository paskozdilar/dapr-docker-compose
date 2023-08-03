# Builder
FROM golang:latest AS builder
WORKDIR /build/

COPY vendor/ /build/vendor/
COPY go.mod go.sum *.go /build/

RUN go build -tags netgo -mod vendor -o client .

# Runner
FROM scratch
WORKDIR /run/

COPY --from=builder /build/client /run/

ENTRYPOINT ["/run/client"]
