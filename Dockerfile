FROM golang:1.23 as build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]