#Build stage
FROM golang:1.23.9-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
ENV GIN_MODE=release

EXPOSE 9090
CMD ["/app/main"]
