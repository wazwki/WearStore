FROM golang:1.23.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o product-service ./cmd/main.go

FROM alpine:3.18

WORKDIR /app
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/product-service /app/product-service

EXPOSE ${PRODUCT_PORT}

CMD ["/app/product-service"]