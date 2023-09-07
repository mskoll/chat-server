FROM golang:1.21.0-alpine3.18

COPY . .
RUN go mod tidy
RUN go build -o app ./cmd/main.go

EXPOSE 9000

CMD ["./app"]