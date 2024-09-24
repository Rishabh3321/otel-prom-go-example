# Prometheus OTel Go Example

## Overview
This repository provides a Go application instrumented with OpenTelemetry for collecting and exporting metrics to Prometheus. The metrics can then be visualized using any Prometheus-compatible monitoring tool, such as Grafana.

## Prerequisites
- Go installed (version 1.22 or later)
- Docker and Docker Compose installed


## Setup
- Clone the repository:
```bash
git clone https://github.com/Rishabh3321/otel-prom-go-example.git
```

- Start the application and Prometheus:
```bash
cd otel-go-example
docker-compose up -d
```

## Accessing Metrics
### Prometheus:
- Open your web browser and navigate to http://localhost:9090.
- Use the query language to explore and visualize the collected metrics.

## Customization
- **Metrics**: Modify the main.go file to add or remove metrics as needed.
- **Exporters**: Adjust the OpenTelemetry configuration in the Dockerfile to use different exporters (e.g., Jaeger, Zipkin).
- **Prometheus Configuration**: Customize the prometheus.yml file to configure Prometheus' behavior.

## Contributing
Feel free to open issues or submit pull requests to improve the example or add new features.