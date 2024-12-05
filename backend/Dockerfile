FROM golang:1.22-alpine

WORKDIR /app

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download

COPY backend/ .

RUN go build -o main cmd/main.go
CMD ["./main"]
