FROM golang:1.20.1-alpine3.17@sha256:87d0a3309b34e2ca732efd69fb899d3c420d3382370fd6e7e6d2cb5c930f27f9 AS builder

WORKDIR /workspace

RUN mkdir /workspace/build
RUN mkdir /workspace/build/file-store-api

#copy and build common module
RUN mkdir /workspace/build/customer-os-common-module

COPY customer-os-common-module/go.mod /workspace/build/customer-os-common-module/go.mod
COPY customer-os-common-module/go.sum /workspace/build/customer-os-common-module/go.sum

RUN cd /workspace/build/customer-os-common-module && go mod download && go mod verify

COPY customer-os-common-module/. /workspace/build/customer-os-common-module/
RUN cd /workspace/build/customer-os-common-module && go build

#copy and build customer-os-api-sdk
RUN mkdir /workspace/build/customer-os-api-sdk

COPY customer-os-api-sdk/go.mod /workspace/build/customer-os-api-sdk/go.mod
COPY customer-os-api-sdk/go.sum /workspace/build/customer-os-api-sdk/go.sum

RUN cd /workspace/build/customer-os-api-sdk && go mod download && go mod verify

COPY customer-os-api-sdk/. /workspace/build/customer-os-api-sdk/
RUN cd /workspace/build/customer-os-api-sdk && go build

#copy and build file storage api
COPY file-store-api/go.mod /workspace/build/file-store-api/go.mod
COPY file-store-api/go.sum /workspace/build/file-store-api/go.sum

RUN cd /workspace/build/file-store-api && go mod download && go mod verify

COPY file-store-api/. /workspace/build/file-store-api/

RUN cd /workspace/build/file-store-api && go build -v -o /go/bin/app

FROM alpine:3.14@sha256:0f2d5c38dd7a4f4f733e688e3a6733cb5ab1ac6e3cb4603a5dd564e5bfb80eed

COPY --chown=65534:65534 --from=builder /go/bin/app .
COPY --chown=65534:65534 --from=builder /workspace/build/file-store-api/.env .env
USER 65534

ENTRYPOINT [ "./app" ]