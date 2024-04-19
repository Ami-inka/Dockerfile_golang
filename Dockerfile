FROM golang:1.18-alpine as builder

WORKDIR /app

COPY server.go .

RUN go build -o webserver server.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/webserver .

EXPOSE 8080

CMD ["./webserver"]

# Some changes
# Second commit in dev bracnh