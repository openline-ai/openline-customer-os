FROM golang:1.20.1-alpine3.17@sha256:87d0a3309b34e2ca732efd69fb899d3c420d3382370fd6e7e6d2cb5c930f27f9 AS builder

WORKDIR /workspace

RUN mkdir /workspace/packages
RUN mkdir /workspace/packages/runner
RUN mkdir /workspace/packages/runner/sync-gmail-raw
RUN mkdir /workspace/packages/server
RUN mkdir /workspace/packages/server/customer-os-common-module
RUN mkdir /workspace/packages/server/customer-os-common-auth

#copy and build common module
COPY customer-os-common-module/go.mod /workspace/packages/server/customer-os-common-module/go.mod
COPY customer-os-common-module/go.sum /workspace/packages/server/customer-os-common-module/go.sum
RUN cd /workspace/packages/server/customer-os-common-module && go mod download && go mod verify

COPY customer-os-common-module/. /workspace/packages/server/customer-os-common-module/
RUN cd /workspace/packages/server/customer-os-common-module && go build

#copy and build common auth module
COPY customer-os-common-auth/go.mod /workspace/packages/server/customer-os-common-auth/go.mod
COPY customer-os-common-auth/go.sum /workspace/packages/server/customer-os-common-auth/go.sum
RUN cd /workspace/packages/server/customer-os-common-auth && go mod download && go mod verify

COPY customer-os-common-auth/. /workspace/packages/server/customer-os-common-auth/
RUN cd /workspace/packages/server/customer-os-common-auth && go build

#copy and build sync-gmail-raw
COPY sync-gmail-raw/go.mod /workspace/packages/runner/sync-gmail-raw/go.mod
COPY sync-gmail-raw/go.sum /workspace/packages/runner/sync-gmail-raw/go.sum
RUN cd /workspace/packages/runner/sync-gmail-raw && go mod download && go mod verify

COPY sync-gmail-raw/. /workspace/packages/runner/sync-gmail-raw/
RUN cd /workspace/packages/runner/sync-gmail-raw && go build -v -o /go/bin/app

FROM alpine:3.14@sha256:0f2d5c38dd7a4f4f733e688e3a6733cb5ab1ac6e3cb4603a5dd564e5bfb80eed

COPY --chown=65534:65534 --from=builder /go/bin/app .
COPY --chown=65534:65534 --from=builder /workspace/packages/runner/sync-gmail-raw/.env .env
USER 65534

ENTRYPOINT [ "./app" ]