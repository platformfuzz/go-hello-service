# 🚀 Go Hello Service

A minimal Go-based HTTP service with health checks, designed for cloud-native deployment with multiple deployment targets.

## ✨ Features

- **Minimal HTTP Server**: Exposes `/` and `/health` endpoints
- **Graceful Shutdown**: Handles SIGTERM/SIGINT signals properly
- **Structured Logging**: Request logging with timing information
- **Health Checks**: JSON health endpoint for monitoring
- **Containerized**: Built with `ko` for OCI-compliant images
- **Multi-Platform**: Supports linux/amd64 and linux/arm64
- **Cloud Ready**: Deployable to ECS Fargate, Cloud Run, and more

## 🏗️ Architecture

```plaintext
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Load Balancer │───▶│  Hello Service  │───▶│   Health Check  │
│   (ALB/Nginx)   │    │   (Port 8080)   │    │   (/health)     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
```

## 🚀 Quick Start

### Development Container (Recommended)

1. **Open in VS Code with Dev Containers**:
   - Install the "Dev Containers" extension
   - Open the project folder
   - Click "Reopen in Container" when prompted

### Local Development

1. **Clone and build**:

   ```bash
   git clone git@github.com:platformfuzz/go-hello-service.git
   cd go-hello-service
   go mod download
   go run ./cmd/server
   ```

2. **Test the endpoints**:

   ```bash
   curl http://localhost:8080/
   curl http://localhost:8080/health
   ```

### Docker Compose (Local)

```bash
docker-compose up --build
```

Access the service at `http://localhost:8080`

## 📡 API Endpoints

### GET `/`

Returns a hello message with hostname and timestamp.

**Response:**

```json
{
  "message": "Hello, World!",
  "timestamp": "2024-01-15T10:30:00Z",
  "hostname": "server-123"
}
```

### GET `/health`

Health check endpoint for monitoring systems.

**Response:**

```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "version": "1.0.0"
}
```

## 🐳 Containerization with `ko`

This project uses [`ko`](https://github.com/ko-build/ko) for building OCI-compliant images without Dockerfiles.

### Build and Push

```bash
# Install ko
go install github.com/ko-build/ko@latest

# Build and push to registry
KO_DOCKER_REPO=ghcr.io/platformfuzz/go-hello-service ko build ./cmd/server

# Build locally without pushing
ko build --local ./cmd/server

# Build for multiple platforms
ko build --platform=linux/amd64,linux/arm64 ./cmd/server
```

### Local Build

```bash
# Build locally without pushing
ko build --local ./cmd/server
```

## ☁️ Deployment Options

### A. Cloud Run (Google Cloud)

```bash
gcloud run deploy go-hello-service \
  --image ghcr.io/platformfuzz/go-hello-service:latest \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --port 8080
```

### B. App Runner (AWS)

```bash
aws apprunner create-service \
  --service-name go-hello-service \
  --source-configuration '{
    "ImageRepository": {
      "ImageIdentifier": "ghcr.io/platformfuzz/go-hello-service:latest",
      "ImageConfiguration": {
        "Port": "8080"
      }
    }
  }'
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT`   | `8080`  | Server port |

### Health Check Configuration

The service includes built-in health checks:

- **Path**: `/health`
- **Expected Status**: `200 OK`
- **Response Time**: < 5 seconds
- **Interval**: 30 seconds (configurable per platform)

## 🧪 Testing & Linting

```bash
# Run unit tests
go test ./...

# Run with coverage
go test -cover ./...

# Run linting locally
golangci-lint run

# Run all checks
go test ./... && golangci-lint run
```

## 📊 Monitoring & Observability

### Health Checks

The `/health` endpoint is designed for:

- Load balancer health checks
- Kubernetes liveness/readiness probes
- Container orchestration platforms

### Logging

Request logging includes:

- HTTP method
- Request URI
- Remote address
- Response time

### Metrics (Optional)

To add Prometheus metrics, uncomment the metrics code in `main.go`.

## 🔄 CI/CD Pipeline

The GitHub Actions workflow (`/.github/workflows/ci.yml`) provides:

- **Validation**: Go modules, tests, builds, and security checks
- **Linting**: golangci-lint for code quality
- **Security**: govulncheck for vulnerability scanning
- **Building**: Multi-platform container images
- **Publishing**: Automatic image publishing on push to main
- **Versioning**: Semantic versioning based on commit messages

### Automatic Versioning

The CI/CD pipeline automatically handles versioning:

- **Manual Tags**: Push a tag (e.g., `v1.0.0`) for specific releases
- **Automatic Semantic Versioning**: Based on commit messages:
  - `feat:` or `feature:` → Minor version bump
  - `fix:` or `bug:` → Patch version bump
  - `breaking:` or `major:` → Major version bump
  - Any other commit → Patch version bump

```bash
# Examples of automatic versioning:
git commit -m "feat: add new health endpoint"  # v1.1.0
git commit -m "fix: resolve memory leak"       # v1.0.1
git commit -m "breaking: change API response"  # v2.0.0
```

## 🛠️ Development

### Project Structure

```plaintext
.
├── cmd/server/          # Main application
│   ├── main.go         # Server implementation
│   └── main_test.go    # Unit tests
├── .github/workflows/  # CI/CD pipeline
├── ko.yaml            # Container build config
├── .devcontainer/     # Development container
└── README.md          # This file
```

### Adding Features

1. **New Endpoints**: Add handlers in `cmd/server/main.go`
2. **Configuration**: Use environment variables or flags
3. **Testing**: Add tests in `main_test.go`
4. **Documentation**: Update this README

## 🔒 Security

- **Base Image**: Uses `gcr.io/distroless/static:nonroot`
- **Non-root User**: Runs as non-root user (UID 65532)
- **Minimal Attack Surface**: Distroless base with no shell
- **HTTPS Ready**: Configure TLS termination at load balancer

## 📈 Performance

- **Startup Time**: < 1 second
- **Memory Usage**: ~10MB baseline
- **CPU Usage**: Minimal for HTTP serving
- **Concurrent Requests**: Handles hundreds of concurrent connections

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

1. **Port Already in Use**:

   ```bash
   # Change port via environment
   PORT=8081 go run ./cmd/server
   ```

2. **Container Won't Start**:

   ```bash
   # Check logs
   docker logs <container-id>
   ```

3. **Health Check Failing**:

   ```bash
   # Test health endpoint
   curl -f http://localhost:8080/health
   ```

### Debug Mode

```bash
# Run with debug logging
DEBUG=true go run ./cmd/server
```

---

Built with ❤️ using Go and modern cloud-native practices
