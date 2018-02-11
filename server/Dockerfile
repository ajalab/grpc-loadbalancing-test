FROM alpine:3.6

RUN apk --no-cache add ca-certificates openssl && update-ca-certificates

WORKDIR /service

COPY app app
