# Kube OpenWebUI Backend

This repository contains the backend service for a model manager application that interacts with an Ollama AI engine. The application is designed to be deployed in a Kubernetes environment and includes a comprehensive CI/CD pipeline for automated builds, testing, and deployments.

## Features

  * **Go Backend**: A lightweight and efficient backend service built in Go that communicates with the Ollama AI engine.
  * **Dockerized**: The application is containerized using a multi-stage Dockerfile for a small and secure final image.
  * **Kubernetes-Native**: Includes a full set of Kubernetes manifests for deploying the application and its dependencies to a cluster.
  * **Automated CI/CD**: A complete CI/CD pipeline using GitHub Actions for automated builds, tests, security scans, and deployments.
  * **Horizontal Pod Autoscaling**: Automatically scales the Kubernetes deployments based on CPU and memory usage.
  * **Secure by Default**: Implements Kubernetes network policies to control traffic between services.

## Architecture

The application consists of the following services:

  * **ollama**: The core AI engine for running large language models.
  * **backend**: The model manager backend, which provides an API for interacting with Ollama.
  * **frontend**: A custom user interface for the model manager.
  * **open-webui**: A feature-rich, open-source web UI for Ollama.

These services are designed to work together, with the backend communicating with Ollama and the frontend interacting with the backend.

## Getting Started

To run the application locally using Docker Compose, you will need to have Docker and Docker Compose installed.

1.  Clone the repository:
    ```bash
    git clone https://github.com/Richard-m-j/kube-openWebUI-backend.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd kube-openWebUI-backend
    ```
3.  Start the application:
    ```bash
    docker-compose -f compose.yml up -d
    ```

This will start all the services in detached mode. You can then access the frontend and OpenWebUI in your browser.

## Kubernetes Deployment

To deploy the application to a Kubernetes cluster, you can use the manifests provided in the `mykubernetes` directory.

1.  Apply the manifests:
    ```bash
    kubectl apply -f mykubernetes/
    ```

This will create all the necessary Kubernetes resources, including deployments, services, persistent volumes, and network policies.

## CI/CD Pipeline

The repository has a comprehensive CI/CD pipeline that automates the build, test, and deployment process. The pipeline is defined in the `.github/workflows` directory and consists of the following workflows:

  * **`ci-cd.yml`**: This workflow is triggered on pushes and pull requests to the `main` branch. It performs linting, builds the application, runs unit tests, and conducts security scans. On a push to `main`, it also builds and pushes a new Docker image, updates the Kubernetes manifest with the new image tag, and commits the change back to the repository.
  * **`docker-image.yaml`**: This workflow builds and pushes the backend's Docker image to Docker Hub on pushes to the `main` branch.
  * **`code-turtle.yaml`**: This workflow uses the Code-Turtle AI Assistant to index the repository and review pull requests.

## API Endpoints

The backend service provides the following API endpoints:

  * `GET /api/models`: Lists the available models in the Ollama service.
  * `POST /api/models/pull`: Downloads a new model from Ollama.

## Configuration

The backend service can be configured using the following environment variable:

  * `OLLAMA_HOST`: The hostname and port of the Ollama service. The default is `http://localhost:11434`.

## Contributing

Contributions are welcome\! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License.