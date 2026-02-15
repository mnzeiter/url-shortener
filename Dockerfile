FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/server .

ENV DB_HOST=db
ENV DB_PORT=3306
ENV DB_USER=root
ENV DB_PASSWORD=password
ENV DB_NAME=url_shortener

EXPOSE 8080

CMD ["./server"]
