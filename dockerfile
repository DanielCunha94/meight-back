
FROM golang:1.22.5 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o /meight ./cmd/app

EXPOSE 8080

CMD ["/meight"]