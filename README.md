# Observability Demo

## Overview

This project demonstrates a complete observability setup using a variety of open-source tools. The system is designed to collect, process, and visualize metrics, logs, and traces from a web application.

The observability stack includes:

- Grafana for visualization
- Prometheus for metrics collection
- Loki for log aggregation
- Tempo for tracing
- OTel Collector for receiving and processing telemetry data
- Beyla for auto-instrumentation
- Alloy for log writing and OTLP exporting

## Service Structure

The following diagram illustrates the architecture of the observability system:

<p align="center">
    <img src="./figure/structure.png" width="900">
</p>

## Getting Started

### Start Service

To start the service using Docker Compose, run:

```sh
docker-compose up -d
```

### Access the Web Application

1. Open `http://localhost:3001` and click the button on the website.

2. Open `http://localhost:4000` to view the Grafana dashboards.

### API Endpoints

You can interact with the web application using the following API endpoints:

1. Create Task

```sh
curl http://localhost:8080/task -X POST -d '{"description": "Learning Observability"}'
```

2. Get Task List

```sh
curl http://localhost:8080/task
```

3. Get Joke

```sh
curl http://localhost:8080/joke
```

## Kubernetes Deployment

To deploy the observability stack in a Kubernetes cluster, follow these steps:

### Apply all resources

```sh
kubectl apply -f K8s

# check pod and service
kubectl get pods
kubectl get svc
kubectl logs <your-pod-id>
```

### Port Forwarding for Testing

1. Forward the port to access the web application:

```sh
kubectl port-forward svc/observability-demo-tellme-app-service 3000:80
```

2. Open `http://localhost:3000` and click the button on the website.
