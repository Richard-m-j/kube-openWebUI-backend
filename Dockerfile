# Stage 1: The builder
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
# Build a static binary for a minimal final image
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Stage 2: The final image
FROM alpine:latest
WORKDIR /app
# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
