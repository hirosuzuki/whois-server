FROM golang:1.20.4-alpine3.17 as builder

WORKDIR /app
COPY go.* *.go /app/
RUN go build

FROM alpine:3.17

WORKDIR /app
COPY --from=builder /app/whois-server /app/
EXPOSE 8080

CMD ["/app/whois-server"]
