# microservices-mockup
This repository contains some microservices application that will be deployed to Kubernetes cluster. It use Github Action for CI/CD pipeline.

## How It Works?
### CI/CD
* Each backend service has its own CI/CD pipeline (2 pipelines for now).
* Every change on directory `backend-go` and `backend-node` will trigger the CI/CD pipelines. 
* CI/CD pipeline will build the container image --> push image to Github Container Registry --> deploy to Kubernetes cluster.
* Done! :)

###  🚀 Exposing backend services to internet 🚀
* Setup NGINX ingress controller for traffic routing and HTTPS termination. Manifest can be found on [here](add-ons/ingress-nginx/deploy.yaml)
* Setup Cert Manager for certificate management. Manifest can be found on [here](add-ons/cert-manager/cert-manager.yaml)
* Setup Cluster Issuer and use Let's Encrypt as certificate provider. Manifest can be found [here](add-ons/cert-manager/cluster-issuer.yaml)
* Create Ingress for backend-go and backend-node. Manifest can be found on [here](add-ons/ingress-nginx/ingress.yaml)

### Services Endpoints
* Backend Golang:
  * https://go.rahmatawe.com/ping
  * https://go.rahmatawe.com/home
  * https://go.rahmatawe.com/ (will return 404)

* Backend Node:
  * https://node.rahmatawe.com/
  * https://node.rahmatawe.com/health/sync
  * https://node.rahmatawe.com/health/async

## Directory Structure

```
.
├── README.md
├── add-ons
│   ├── cert-manager
│   └── ingress-nginx
├── backend-go
│   ├── Dockerfile
│   ├── README.md
│   ├── go.mod
│   ├── go.sum
│   ├── kubernetes
│   ├── main.go
│   └── main_test.go
└── backend-node
    ├── Dockerfile
    ├── LICENSE
    ├── README.md
    ├── kubernetes
    ├── package-lock.json
    ├── package.json
    ├── server.js
    └── src
```

### add-ons
There are some manifest to setup `cert-manager` and `ingress-nginx`. 
* `cert-manager` will be used to manage TLS certificate using Let's Encrypt provider.
* `ingress-nginx` will be used for traffic routing from public to service inside Kubernetes.

### backend-go
Contain source code to build Golang based backend. Inside `kubernetes` directory, there are manifest used to deploy the service inside Kubernetes cluster.

### backend-node
Contain source code to build NodeJS based backend. Inside `kubernetes` directory, there are manifest used to deploy the service inside Kubernetes cluster. This service need MongoDB to run, example manifest can be found [here](backend-node/kubernetes/statefulset.yaml)

## Technology Stack
* Digital Ocean Kubernetes v1.24.4
* NGINX Ingress Controller v1.3.1
* Cert-Manager 1.10.0
* Github Container Registry