FROM golang:1.21 AS builder
WORKDIR /app
COPY . ./

ENV GIN_MODE=release

RUN CGO_ENABLED=0 go build -v -o learning-o11y cmd/server/main.go


FROM debian:bullseye
WORKDIR /app

COPY --from=builder /app/learning-o11y /app/learning-o11y

ENV GIN_MODE=release

CMD ["/app/learning-o11y"]
