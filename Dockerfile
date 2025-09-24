# Stage 1: The builder (remains the same)
FROM golang:1.25-alpine AS builder
# Set build-time metadata arguments
ARG APP_VERSION="v0.1.0-default"
ARG TARGETOS=linux
ARG TARGETARCH=amd64

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .
# Build a static binary for a minimal final image
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build \
    -ldflags="-w -s -X main.Version=${APP_VERSION}" -o /app/main .

#-------------------------------------------------------------------------

# Stage 2: The final hardened image
FROM alpine:3.20.1

# Add metadata labels using OCI standard
ARG APP_VERSION
LABEL org.opencontainers.image.source="https://github.com/Richard-m-j/kube-openWebUI-backend" \
      org.opencontainers.image.version=${APP_VERSION} \
      org.opencontainers.image.title="Kube OpenWebUI Backend"

# Create a dedicated group and user for the application
# -S creates a system user/group, which is more secure (no password, no login shell)
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /app

# Copy only the compiled binary from the builder stage
# --chown sets the ownership of the copied file to the new user and group
COPY --from=builder --chown=appuser:appgroup /app/main .

# Switch to the non-root user
USER appuser

EXPOSE 8080

ENTRYPOINT ["./main"]