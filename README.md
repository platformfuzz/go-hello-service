# Go Hello Service

A minimal Go HTTP service with health checks, designed for containerized deployment on platforms like ECS, Kubernetes, and Cloud Run.

## 🚀 Features

- **HTTP Server**: Simple Go HTTP server with JSON responses
- **Health Checks**: Built-in `/health` endpoint for monitoring
- **Graceful Shutdown**: Handles SIGTERM and SIGINT signals
- **Structured Logging**: Request logging with timing
- **Container Ready**: Optimized for container deployment
- **Multi-platform**: Supports linux/amd64 and linux/arm64

## 📋 Endpoints

### `GET /`

Returns a hello message with hostname and timestamp.

**Response:**

```json
{
  "message": "Hello, World!",
  "timestamp": "2024-01-01T12:00:00Z",
  "hostname": "container-hostname"
}
```

### `GET /health`

Health check endpoint for monitoring and load balancers.

**Response:**

```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "version": "1.0.0"
}
```

## 🐳 Container Deployment

### Environment Variables

- `PORT`: Server port (default: `8080`)

### Docker Run

```bash
# Run with default port
docker run -p 8080:8080 ghcr.io/platformfuzz/go-hello-service/server-967d5646a4ce288d6928f233b912d34d:latest

# Run with custom port
docker run -p 3000:3000 -e PORT=3000 ghcr.io/platformfuzz/go-hello-service/server-967d5646a4ce288d6928f233b912d34d:latest
```

## ☁️ ECS Deployment

### Task Definition Health Check

```json
{
  "healthCheck": {
    "command": [
      "CMD-SHELL",
      "wget --quiet --spider http://localhost:8080/health || exit 1"
    ],
    "interval": 30,
    "timeout": 5,
    "retries": 3,
    "startPeriod": 60
  }
}
```

### For Custom Port

If you set a custom `PORT` environment variable, update the health check accordingly:

```json
{
  "healthCheck": {
    "command": [
      "CMD-SHELL",
      "wget --quiet --spider http://localhost:${PORT}/health || exit 1"
    ],
    "interval": 30,
    "timeout": 5,
    "retries": 3,
    "startPeriod": 60
  }
}
```

## 🏗️ Development

### Prerequisites

- Go 1.24+
- Docker
- ko (for building containers)

### Local Development

```bash
# Run locally
go run ./cmd/server

# Run with custom port
PORT=3000 go run ./cmd/server

# Build binary
go build -o server ./cmd/server

# Test
go test ./cmd/server
```

### DevContainer

This project includes a DevContainer configuration for VS Code:

1. Install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)
2. Open the project in VS Code
3. When prompted, click "Reopen in Container"

## 🔧 CI/CD

The project uses GitHub Actions for automated:

- **Testing**: Unit tests and integration tests
- **Linting**: golangci-lint for code quality
- **Security**: govulncheck for vulnerability scanning
- **Building**: ko for container image building
- **Publishing**: Automatic image publishing to GHCR

### Workflow Triggers

- **Pull Requests**: Validation, linting, security scanning
- **Main Branch**: Build and publish container images

## 📦 Container Images

Images are published to GitHub Container Registry:

- **Repository**: `ghcr.io/platformfuzz/go-hello-service/server-967d5646a4ce288d6928f233b912d34d`
- **Tags**: `latest`, `v0.0.1`, etc.
- **Platforms**: linux/amd64, linux/arm64
- **Base Image**: Alpine Linux (for health check compatibility)

## 🛠️ Technology Stack

- **Language**: Go 1.24
- **Framework**: Gorilla Mux (HTTP routing)
- **Container**: ko (Go-native container building)
- **Base Image**: Alpine Linux
- **CI/CD**: GitHub Actions
- **Registry**: GitHub Container Registry (GHCR)

## 📄 License

MIT License - see LICENSE file for details.
