FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o ./m3go
EXPOSE 8080

CMD ./m3go