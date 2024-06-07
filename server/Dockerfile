# Dockerfile
FROM golang:1.22.3-alpine3.18

WORKDIR /app
COPY . .
# ホットリロード air
RUN go install github.com/cosmtrek/air@v1.52.1

RUN go mod download