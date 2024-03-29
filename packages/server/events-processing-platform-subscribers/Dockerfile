FROM golang:1.20.1-alpine3.17@sha256:87d0a3309b34e2ca732efd69fb899d3c420d3382370fd6e7e6d2cb5c930f27f9 AS builder

WORKDIR /workspace

RUN mkdir /workspace/build
RUN mkdir /workspace/build/customer-os-api
RUN mkdir /workspace/build/customer-os-common-ai
RUN mkdir /workspace/build/customer-os-common-module
RUN mkdir /workspace/build/events-processing-proto
RUN mkdir /workspace/build/events-processing-platform

#copy and build AI module
COPY customer-os-common-ai/go.mod /workspace/build/customer-os-common-ai/go.mod
COPY customer-os-common-ai/go.sum /workspace/build/customer-os-common-ai/go.sum
RUN cd /workspace/build/customer-os-common-ai && go mod download && go mod verify

COPY customer-os-common-ai/. /workspace/build/customer-os-common-ai/
RUN cd /workspace/build/customer-os-common-ai && go build

#copy and build common module
COPY customer-os-common-module/go.mod /workspace/build/customer-os-common-module/go.mod
COPY customer-os-common-module/go.sum /workspace/build/customer-os-common-module/go.sum
RUN cd /workspace/build/customer-os-common-module && go mod download && go mod verify

COPY customer-os-common-module/. /workspace/build/customer-os-common-module/
RUN cd /workspace/build/customer-os-common-module && go build

#copy and build customer-os-neo4j-repository
RUN mkdir /workspace/build/customer-os-neo4j-repository
COPY customer-os-neo4j-repository/go.mod /workspace/build/customer-os-neo4j-repository/go.mod
COPY customer-os-neo4j-repository/go.sum /workspace/build/customer-os-neo4j-repository/go.sum
RUN cd /workspace/build/customer-os-neo4j-repository && go mod download && go mod verify
COPY customer-os-neo4j-repository/. /workspace/build/customer-os-neo4j-repository/
RUN cd /workspace/build/customer-os-neo4j-repository && go build -v ./...

#copy events-processing-proto
COPY events-processing-proto/go.mod /workspace/build/events-processing-proto/go.mod
COPY events-processing-proto/go.sum /workspace/build/events-processing-proto/go.sum
COPY events-processing-proto/. /workspace/build/events-processing-proto/

#copy and build events-processing-platform
COPY events-processing-platform/go.mod /workspace/build/events-processing-platform/go.mod
COPY events-processing-platform/go.sum /workspace/build/events-processing-platform/go.sum
RUN cd /workspace/build/events-processing-platform && go mod download && go mod verify
COPY events-processing-platform/. /workspace/build/events-processing-platform/
RUN cd /workspace/build/events-processing-platform && go build -v ./...

#copy and build events-processing-platform-subscribers
COPY events-processing-platform-subscribers/go.mod /workspace/build/events-processing-platform-subscribers/go.mod
COPY events-processing-platform-subscribers/go.sum /workspace/build/events-processing-platform-subscribers/go.sum
RUN cd /workspace/build/events-processing-platform-subscribers && go mod download && go mod verify

COPY events-processing-platform-subscribers/. /workspace/build/events-processing-platform-subscribers/
RUN cd /workspace/build/events-processing-platform-subscribers && go build -v -o /go/bin/app

FROM alpine:3.14@sha256:0f2d5c38dd7a4f4f733e688e3a6733cb5ab1ac6e3cb4603a5dd564e5bfb80eed

COPY --chown=65534:65534 --from=builder /go/bin/app .
COPY --chown=65534:65534 --from=builder /workspace/build/events-processing-platform-subscribers/.env .env
COPY --chown=65534:65534 --from=builder /workspace/build/events-processing-platform-subscribers/subscriptions/invoice/pdf_template/* /subscriptions/invoice/pdf_template/
USER 65534

ENTRYPOINT [ "./app" ]
