# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20.4

WORKDIR /app
COPY . ./

# Build the Go app
RUN go mod download
RUN go build -o receipt-processor .

EXPOSE ${PORT}
CMD ["./receipt-processor"]