FROM golang:1.7.3 AS builder

WORKDIR /root/go
COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest
WORKDIR /root/
COPY static static
COPY --from=builder /root/go/app .
EXPOSE 9000

CMD ["./app"]