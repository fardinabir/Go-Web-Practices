FROM golang:1.19 AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o ./nethttp
EXPOSE 8443
CMD ["./nethttp"]

