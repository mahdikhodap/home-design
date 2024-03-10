FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o url_checker .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/url_checker .

CMD ["./url_checker"]
