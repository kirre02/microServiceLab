FROM golang:1.24 AS build-stage

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPATH=/echoService

WORKDIR /build

COPY . .

RUN go mod download

RUN go build ./cmd/echo/

FROM alpine:3.19.0

WORKDIR /api/

COPY --from=Build-stage /build/echo ./bin/echo

EXPOSE 8082

CMD ["./bin/echo"]
