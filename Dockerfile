FROM golang:1.23.1 AS builder
WORKDIR /app
COPY . .
COPY vendor ./vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o main

FROM alpine
WORKDIR /test
COPY my-chart /test/my-chart
COPY --from=builder /app/main .
CMD ["./main"]