# Stage 1: The builder
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
# Stage 2: The final image
FROM scratch

ARG APP_VERSION
# Add metadata labels using OCI standard
LABEL org.opencontainers.image.source="https://github.com/Richard-m-j/kube-openWebUI-backend" \
      org.opencontainers.image.version=${APP_VERSION} \
      org.opencontainers.image.title="Kube OpenWebUI Backend"

WORKDIR /app
# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8080

USER 10001

ENTRYPOINT ["./main"]
