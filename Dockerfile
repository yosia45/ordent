FROM golang:1.23 as build

WORKDIR /app

COPY . .

RUN go build -o main main.go

RUN ls -l /app

CMD ["./main"]