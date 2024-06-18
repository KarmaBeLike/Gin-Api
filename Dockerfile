FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN apk add build-base && go build -o gin-api cmd/main.go

FROM alpine:3.6
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8080:8080
CMD ["/app/gin-api"]