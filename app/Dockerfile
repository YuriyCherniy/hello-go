FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o hello-go main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/hello-go .
EXPOSE 8080
CMD ["./hello-go"]