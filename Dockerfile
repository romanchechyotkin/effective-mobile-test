FROM golang:alpine as builder

RUN apk add --update --no-cache ca-certificates git openssl
ARG cert_location=/usr/local/share/ca-certificates

RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
RUN update-ca-certificates

ENV POSTGRES_HOST="" \
    POSTGRES_PORT="" \
    POSTGRES_DB="" \
    POSTGRES_USER="" \
    POSTGRES_PASSWORD="" \
    ENVIRONMENT=""

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/bin cmd/main/main.go

FROM alpine:latest

RUN apk add --no-cache bash

WORKDIR /app
COPY --from=builder /app/bin .
EXPOSE 8080
CMD ["/app/bin"]