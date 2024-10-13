# build stage
FROM golang:1.23.2-alpine3.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/main.go

# run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY ./upload upload

EXPOSE 8050
CMD ["./main"]