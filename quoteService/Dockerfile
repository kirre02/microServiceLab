FROM golang:1.24 AS build-stage

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOPATH=/quoteService

WORKDIR /build

COPY . .

RUN go mod download

RUN go build ./cmd/quote/

FROM alpine:3.19.0

WORKDIR /api/

COPY --from=build-stage /build/quote ./bin/quote

EXPOSE 8083

CMD ["./bin/quote"]
