# learning-o11y

## Run

1. Start Prometheus

```sh
docker-compose up -d
```

2. Start Server

```sh
go run cmd/server/main.go
```

## Kubernetes

### Deployment

```sh
kubectl apply -f K8s/deployment.yaml
kubectl get pods
kubectl exec --stdin --tty learning-o11y-<your-pod-id> -- /bin/bash
kubectl logs learning-o11y-<your-pod-id>
```

### Service

```sh
kubectl apply -f K8s/service.yaml
kubectl get svc
```

### Service Monitor

```sh
kubectl apply -f K8s/service-monitor.yaml
```

### Test

```sh
kubectl port-forward svc/learning-o11y-service 8080:80
curl http://localhost:8080/hello
```
