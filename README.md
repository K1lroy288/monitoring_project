# Monitoring Project

A comprehensive monitoring solution built with Go, Prometheus, and Grafana to demonstrate modern observability practices. This project showcases a complete monitoring stack with custom metrics collection, visualization, and system monitoring capabilities.

## 🚀 Features

- **Custom Metrics Collection**: Implemented using Prometheus client library with various metric types (Counter, Gauge, Histogram)
- **System Monitoring**: CPU load tracking and other system metrics via Node Exporter
- **Complete Observability Stack**: Prometheus for metrics storage, Grafana for visualization
- **Containerized Deployment**: Docker and Docker Compose for easy deployment and scaling
- **Automated Provisioning**: Grafana dashboards and datasources automatically configured
- **Health Checks**: Built-in health check endpoint with monitoring
- **HTTP Request Monitoring**: Tracking of request count, duration, and status codes

## 🛠️ Tech Stack

- **Backend**: Go with Prometheus client library
- **Metrics Storage**: Prometheus
- **Visualization**: Grafana
- **Containerization**: Docker, Docker Compose
- **System Metrics**: Node Exporter
- **Monitoring**: Custom metrics and system-level monitoring

## 📊 Metrics Overview

The application exposes the following custom metrics:

- `http_requests_total` - Total number of HTTP requests with labels for method, path, and status
- `http_request_duration_seconds` - Request duration histogram with method and path labels
- `CPU_loaded` - Current CPU load percentage
- System metrics from Node Exporter including:
  - Network traffic (transmit/receive bytes)
  - Disk usage
  - System uptime
  - Process metrics

The application demonstrates proper metric labeling and the use of different metric types (Counter, Gauge, Histogram) for different use cases.

## 🏗️ Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Application   │────│   Prometheus     │────│     Grafana     │
│   (Go + Metrics)│    │   (Metrics DB)   │    │   (Dashboard)   │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                        ┌──────────────────┐
                        │  Node Exporter   │
                        │ (System Metrics) │
                        └──────────────────┘
```

## 🚀 Quick Start

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd monitoring_project
   ```

2. Build and start the services:
   ```bash
   docker-compose up --build
   ```

### Access Services

- **Application**: http://localhost:3425
- **Metrics endpoint**: http://localhost:3425/metrics
- **Health check**: http://localhost:3425/health
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/secret)
- **Node Exporter**: http://localhost:9100

## 📈 Available Endpoints

- `/health` - Health check endpoint
- `/custom_duration` - Endpoint with simulated processing time (500ms delay)
- `/metrics` - Prometheus metrics endpoint

## 🧪 Testing

The application automatically exposes metrics that can be tested by making requests to different endpoints:

```bash
# Health check
curl http://localhost:3425/health

# Custom duration endpoint
curl http://localhost:3425/custom_duration

# View metrics
curl http://localhost:3425/metrics
```

## 🔧 Configuration

### Prometheus Configuration

The Prometheus configuration is located at `./prometheus/prometheus.yml` and includes:
- Scraping configuration for the application
- Scraping configuration for Prometheus itself
- Scraping configuration for Node Exporter

### Grafana Dashboards

Grafana is pre-configured with:
- Custom dashboard for application metrics
- System metrics visualization
- Pre-configured data source for Prometheus

Dashboard includes panels for:
- Goroutines count
- CPU load
- Network traffic
- HTTP request duration
- System uptime
- Disk usage

## 🐳 Docker Images

The project uses multi-service Docker Compose with the following services:

- **backend**: Go application with custom metrics
- **prometheus**: Metrics collection and storage
- **grafana**: Visualization and dashboarding
- **node-exporter**: System metrics collection

## 🚨 Production Considerations

While this project demonstrates core monitoring concepts, a production deployment would require additional considerations:

- **Security**: Currently, containers run as root users and Grafana uses a hardcoded password. Production setup should use non-root users and proper secret management
- **Authentication**: Implement proper authentication and authorization mechanisms
- **Alerting**: Configure Prometheus alerting rules and integrate with notification systems
- **Logging**: Implement structured logging with centralized log aggregation
- **Monitoring**: Monitor the monitoring stack itself (blackbox exporter, etc.)
- **TLS**: Enable HTTPS/TLS for secure communication between services
- **Configuration**: Use environment variables and configuration files for flexible deployment
- **Resource Management**: Define resource limits and requests for containers
- **Health Checks**: Implement comprehensive liveness and readiness probes
- **Backup & Recovery**: Set up backup strategies for Prometheus and Grafana data

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 📞 Contact

Your Name - your.email@example.com

Project Link: [https://github.com/yourusername/monitoring_project](https://github.com/yourusername/monitoring_project)
