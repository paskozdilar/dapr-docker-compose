# Builder
FROM golang:latest AS builder
WORKDIR /build/

COPY vendor/ /build/vendor/
COPY go.mod go.sum *.go /build/

RUN go build -tags netgo -mod vendor -o server .

# Runner
FROM scratch
WORKDIR /run/

COPY --from=builder /build/server /run/

ENTRYPOINT ["/run/server"]
