FROM golang:1.19.4-alpine3.16 AS builder

WORKDIR /app  

COPY . /app

RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app .

EXPOSE 7070

CMD ["/app/main"]
