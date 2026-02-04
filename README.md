# Greetings App (Go + Python)

This project is a microservices application running on Kubernetes.

### Local Setup

**1. Start Cluster**
```bash
minikube start --driver=docker
eval $(minikube docker-env)
```

**2. Build Image**
```bash
docker build --platform linux/amd64 -t go-backend:v1 ./backend
docker build --platform linux/amd64 -t streamlit-frontend:v1 ./frontend
```

**3. Deploy**
```bash
kubectl apply -f k8s/deployment.yaml
```

**3. Open App**
```bash
minikube service streamlit-service
```

Architecture Details
Frontend: Python Streamlit app (Port 8501).

Backend: Go REST API (Port 8080).

Internal Communication: The frontend connects to the backend via the internal DNS: http://go-backend-service:8080/api/data.

Services:

go-backend-service: ClusterIP (Internal only).

streamlit-service: LoadBalancer (External access).
