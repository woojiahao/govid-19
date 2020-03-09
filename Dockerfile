FROM golang:1.14.0-buster

EXPOSE 8080

WORKDIR /app

COPY . .

RUN go mod download
RUN go build cmd/main.go

CMD ./main
