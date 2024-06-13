FROM golang:1.21 AS builder
WORKDIR /app
COPY . ./

ENV GIN_MODE=release

RUN CGO_ENABLED=1 go build -v -o observability-demo cmd/server/main.go


FROM debian:bookworm
WORKDIR /app

RUN apt-get update && \
    apt-get install -y ca-certificates 

COPY --from=builder /app/observability-demo /app/observability-demo

ENV GIN_MODE=release

CMD ["/app/observability-demo"]
