# learning-o11y

## Service Structure

![image](./figure/structure.png)

## Run

1. Start Service

```sh
docker-compose up -d
```

2. (Optional) Start Backend Service Locally

```sh
go run cmd/server/main.go
```

### API

1. Create Task

```sh
curl http://localhost:8080/task -X POST -d '{"description": "Learning Observability"}'
```

2. Get Task List

```sh
curl http://localhost:8080/task
```

## Kubernetes

### Deployment

```sh
kubectl apply -f K8s/deployment.yaml
kubectl get pods
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
