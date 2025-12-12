FROM golang:1.21-alpine AS builder

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build for Lambda (Linux)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap cmd/lambda/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /build/bootstrap .

ENTRYPOINT ["./bootstrap"]
