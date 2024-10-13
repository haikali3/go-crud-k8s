# Go CRUD API with Kubernetes

This project is a simple CRUD (Create, Read, Update, Delete) API built with Go, Dockerized, and deployed on a Kubernetes cluster. The API allows basic operations on "items" and demonstrates a simple microservice architecture.

## Table of Contents
- [Features](#features)
- [Technologies](#technologies)
- [Usage](#usage)
  - [API Endpoints](#api-endpoints)
  - [Running Locally](#running-locally)
- [Deploying to Kubernetes](#deploying-to-kubernetes)
- [Accessing the API](#accessing-the-api)

## Features
- RESTful CRUD API for managing items.
- Dockerized for easy containerization.
- Kubernetes deployment for scaling and orchestration.

## Technologies
- Go (Golang)
- Gorilla Mux (for routing)
- Docker
- Kubernetes (Minikube or cloud-based clusters)

### Prerequisites
- Docker installed on your system.
- Kubernetes installed (use Minikube for local development or any cloud-based cluster).
- Go installed on your system.

### 1. Clone the repository

```bash
git clone https://github.com/haikali3/go-crud-k8s.git
cd go-crud-k8s
```


### 2. Build Docker Image
```bash
docker build -t haikali3/go-crud-k8s:latest .
```

### 3. Push the Docker Image to Docker Hub
```bash
docker tag go-crud-k8s haikali3/go-crud-k8s:latest
docker push haikali3/go-crud-k8s:latest
```

### Usage
API Endpoints
The following endpoints are available:

Method	Endpoint	Description
| Method | Endpoint      | Description              |
|--------|---------------|--------------------------|
| POST   | /items        | Create a new item        |
| GET    | /items        | Get all items            |
| PUT    | /items/{id}   | Update an item by ID     |
| DELETE | /items/{id}   | Delete an item by ID     |


### Example

#### Create an item:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"id":"1","name":"Item 1"}' http://<EXTERNAL-IP>:8000/items
```

#### Get all item:
```bash
curl http://<EXTERNAL-IP>:8000/items
```

#### Update an item:
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Updated Item"}' http://<EXTERNAL-IP>:8000/items/1
```

#### Delete an item:
```bash
curl -X DELETE http://<EXTERNAL-IP>:8000/items/1
```

## Deploying to Kubernetes
### 1. Apply Kubernetes deployment:

Make sure you are connected to your Kubernetes cluster (Minikube or cloud provider), then apply the deployment and service manifest:

```bash
kubectl apply -f k8s-deployment.yaml
```

### 2. Check the status of your deployment:

```bash
kubectl get pods
kubectl get svc
```

Make sure the pods are running, and the service has been assigned an external IP or NodePort.

## Accessing the API
#### Option 1: Using Minikube
For Minikube, you can access the service using the minikube service command:

```bash
minikube service go-crud-service
```
This will provide you with a URL to access the API.

#### Option 2: Using NodePort
If you're using NodePort, find the external IP and NodePort:

```bash
minikube ip
kubectl get svc go-crud-service
```
Access the API with the Minikube IP and NodePort, e.g.
`http://<minikube-ip>:<NodePort>`.

#### Option 3: Using LoadBalancer
If your service is exposed as a LoadBalancer (cloud setup), check the external IP:

```bash
kubectl get svc go-crud-service
```

Use the external IP to access the API.