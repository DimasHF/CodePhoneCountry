# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go source code
COPY . .

# Build binary
RUN go build -o codephone-server .

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Copy binary and assets
COPY --from=builder /app/codephone-server .
COPY --from=builder /app/code_phone.json .

EXPOSE 3001

CMD ["./codephone-server"]
