FROM golang:1.21

WORKDIR /app

COPY server/ ./
WORKDIR /app/server
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /orders

EXPOSE 8080

CMD orders
