# Go Prometheus Monitoring

This repo contains a simple Golang application that exposes a few endpoints to demonstrate how to monitor a Golang application using Prometheus and Grafana. 

We will containerize the Golang application and connect it to Prometheus and Grafana using Docker Compose. We will also create a simple Grafana dashboard to visualize the metrics collected by Prometheus. The Golang application is built using the Gin web framework and exposes the following endpoints:

- `/v1/users` - Returns a simple JSON response.
- `/metrics` - Exposes the Prometheus metrics.
- `/health` - Returns the health status of the application.

<img width="1512" alt="grafana-panel" src="https://github.com/user-attachments/assets/2d6edc7b-14ab-4c27-ba6f-fd0e11dae554" />

## Project Structure

```text
go-prometheus-monitoring
├── CONTRIBUTING.md
├── Docker
│   ├── grafana.yml
│   └── prometheus.yml
├── Dockerfile
├── LICENSE
├── README.md
├── compose.yaml
├── go.mod
├── go.sum
└── main.go
```

- **main.go** - The entry point of the application.
- 
- **go.mod and go.sum** - Go module files.
- **Dockerfile** - Dockerfile used to build the app.
- **Docker/** - Contains the Docker Compose configuration files for Grafana and Prometheus.
- **compose.yaml** - Compose file to launch everything (Golang app, Prometheus, and Grafana).
- **dashboard.json** - Grafana dashboard configuration file.
- **Dockerfile** - Dockerfile used to build the Golang app.
- **compose.yaml** - Docker Compose file to launch everything (Golang app, Prometheus, and Grafana).
- Other files are for licensing and documentation purposes.

## Setup Instructions

### Running with Docker Compose

To run the application, you need to have Docker and Docker Compose installed on your machine. You can install Docker and Docker Compose by following the instructions in the official documentation:

To run the the server using Docker Compose, you'll need to create a Dockerfile for the server. Below is the [Dockerfile](Dockerfile) for the our server:

```Dockerfile
# Use the official Golang image as the base
FROM golang:1.24-alpine AS builder

# Set environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source
COPY . .

# Build the Go binary
RUN go build -o /app .

# Final lightweight stage
FROM alpine:3.17 AS final

# Copy the compiled binary from the builder stage
COPY --from=builder /app /bin/app

# Expose the application's port
EXPOSE 8000

# Run the application
CMD ["bin/app"]
```

Now, to run the application with all the services, you can use the following Docker Compose file:

```yaml
services:
  api:
    container_name: go-api
    build:
      context: .
      dockerfile: Dockerfile
    image: go-api:latest
    ports:
      - 8000:8000
    networks:
      - go-network
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]
      interval: 30s
      timeout: 10s
      retries: 5
    develop:
      watch:
        - path: .
          action: rebuild
      
  prometheus:
    container_name: prometheus
    image: prom/prometheus:v2.55.0
    volumes:
      - ./Docker/prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - 9090:9090
    networks:
      - go-network
  
  grafana:
    container_name: grafana
    image: grafana/grafana:11.3.0
    volumes:
      - ./Docker/grafana.yml:/etc/grafana/provisioning/datasources/datasource.yaml
      - grafana-data:/var/lib/grafana
    ports:
      - 3000:3000
    networks:
      - go-network
    environment:
      - GF_SECURITY_ADMIN_USER=admin
      - GF_SECURITY_ADMIN_PASSWORD=password

volumes:
  grafana-data:

networks:
  go-network:
    driver: bridge
```

To build and run the Docker image using Docker Compose, use the following command:

```bash
docker-compose up --build
```