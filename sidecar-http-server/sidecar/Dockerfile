FROM golang:1.21-alpine

WORKDIR /app

COPY run.go .
COPY go.mod .

RUN go mod download

RUN go build -o ./out/go-sample-app .

EXPOSE 9000

CMD ["./out/go-sample-app"]