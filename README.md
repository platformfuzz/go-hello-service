# Go Hello Service

A minimal Go HTTP service with health checks, built with `ko` and deployed via GitHub Actions.

## 🚀 Features

- **HTTP Server**: Simple Go HTTP server with `/` and `/health` endpoints
- **Health Checks**: Built-in health check endpoint for container orchestration
- **Containerized**: Built with `ko` using Alpine Linux base image
- **CI/CD**: Automated build, test, and deployment pipeline
- **Multi-platform**: Supports Linux AMD64 and ARM64
- **Security**: Automated vulnerability scanning and linting

## 📋 Prerequisites

- **Go 1.24+**: For local development
- **Docker**: For container operations
- **GitHub**: For CI/CD pipeline

## 🏗️ Local Development

### Using DevContainer (Recommended)

1. **Open in VS Code**: Clone and open the repository
2. **DevContainer**: VS Code will prompt to reopen in container
3. **Ready**: All tools pre-installed in the container

### Manual Setup

```bash
# Clone the repository
git clone git@github.com:platformfuzz/go-hello-service.git
cd go-hello-service

# Install dependencies
go mod tidy

# Run locally
go run ./cmd/server

# Test
go test ./cmd/server
```

## 🐳 Container Build

### Using ko (Recommended)

```bash
# Build locally
ko build --local ./cmd/server

# Build and push
ko build ./cmd/server --platform=linux/amd64,linux/arm64
```

### Container Details

- **Base Image**: `alpine:latest`
- **Binary**: `/server`
- **Port**: `8080` (configurable via `PORT` env var)
- **Health Check**: Available at `/health`

## 🏥 Health Checks

### HTTP Health Check

The service provides a `/health` endpoint that returns:

```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "version": "1.0.0"
}
```

### Container Health Check

For container orchestration (ECS, Kubernetes, etc.), use:

```json
{
  "healthCheck": {
    "command": [
      "CMD-SHELL",
      "wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1"
    ],
    "interval": 30,
    "timeout": 5,
    "retries": 3,
    "startPeriod": 60
  }
}
```

**Note**: Uses `wget` (available in Alpine base image) for HTTP health checks.

## 🔄 CI/CD Pipeline

### GitHub Actions Workflow

The `.github/workflows/ci.yml` workflow provides:

#### **PR Validation**

- ✅ **Go Module Check**: Ensures `go.mod` is clean
- ✅ **Binary Build**: Validates Go compilation
- ✅ **Container Build**: Tests `ko` build process
- ✅ **Multi-platform**: Validates AMD64/ARM64 builds
- ✅ **Security Scan**: Basic security checks
- ✅ **Linting**: `golangci-lint` integration
- ✅ **Vulnerability Check**: `govulncheck` integration
- ✅ **Tests**: Automated test suite

#### **Main Branch Deployment**

- 🚀 **Build & Push**: Publishes to GitHub Container Registry
- 🏷️ **Tagging**: Automatic `latest` and versioned tags
- 📦 **Multi-platform**: AMD64 and ARM64 images

### Registry

Images are published to: `ghcr.io/platformfuzz/go-hello-service/server:latest`

## 🛠️ Configuration

### Environment Variables

- `PORT`: Server port (default: `8080`)

### ko Configuration

See `ko.yaml` for build configuration:

- **Base Image**: Alpine Linux
- **Platforms**: AMD64, ARM64
- **Labels**: OCI metadata

## 📊 API Endpoints

### GET `/`

Returns a hello message with hostname and timestamp.

**Response**:

```json
{
  "message": "Hello, World!",
  "timestamp": "2024-01-01T12:00:00Z",
  "hostname": "container-hostname"
}
```

### GET `/health`

Returns service health status.

**Response**:

```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "version": "1.0.0"
}
```

## 🔧 Development

### Project Structure

```
.
├── cmd/
│   └── server/
│       ├── main.go          # Main server implementation
│       └── main_test.go     # Unit tests
├── .github/
│   └── workflows/
│       └── ci.yml           # CI/CD pipeline
├── .devcontainer/
│   └── devcontainer.json    # VS Code DevContainer
├── .vscode/
│   ├── extensions.json      # Recommended extensions
│   └── settings.json        # Go development settings
├── ko.yaml                  # ko build configuration
├── go.mod                   # Go module definition
└── README.md               # This file
```

### Adding Dependencies

```bash
# Add a new dependency
go get github.com/example/package

# Update go.mod
go mod tidy
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test ./cmd/server
```

## 🚀 Deployment

### ECS Fargate

```json
{
  "family": "go-hello-service",
  "containerDefinitions": [
    {
      "name": "server",
      "image": "ghcr.io/platformfuzz/go-hello-service/server:latest",
      "portMappings": [
        {
          "containerPort": 8080,
          "protocol": "tcp"
        }
      ],
      "healthCheck": {
        "command": [
          "CMD-SHELL",
          "wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1"
        ],
        "interval": 30,
        "timeout": 5,
        "retries": 3,
        "startPeriod": 60
      }
    }
  ]
}
```

### Docker Compose

```yaml
version: '3.8'
services:
  server:
    image: ghcr.io/platformfuzz/go-hello-service/server:latest
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
    healthcheck:
      test: ["CMD-SHELL", "wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1"]
      interval: 30s
      timeout: 5s
      retries: 3
      start_period: 60s
```

## 🔒 Security

- **Vulnerability Scanning**: Automated `govulncheck` integration
- **Code Quality**: `golangci-lint` enforcement
- **Minimal Base Image**: Alpine Linux for reduced attack surface
- **Graceful Shutdown**: Proper signal handling
- **Error Handling**: Comprehensive error management

## 📝 License

MIT License - see LICENSE file for details.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Ensure all tests pass
5. Submit a pull request

The CI/CD pipeline will automatically validate your changes before merging.
