FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o cart-service ./cmd/main.go

FROM alpine:3.18

WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/cart-service /app/cart-service
COPY .env /app/.env

EXPOSE ${PORT}

CMD ["/app/cart-service"]