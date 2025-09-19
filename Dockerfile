FROM golang:1.24.1-alpine

WORKDIR /events-api

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -a -o ./bin ./cmd

CMD ["/events-api/api"]
EXPOSE 8080
