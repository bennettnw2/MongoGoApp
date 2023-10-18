FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum main.go zscalerca.crt ./
RUN cp zscalerca.crt /usr/local/share/ca-certificates/zscalerca.crt

RUN update-ca-certificates
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /tasker

EXPOSE 8080

CMD ["/tasker"]
