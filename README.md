# Building a Virtual Library System

This project focuses on creating a Virtual Library System. The goal is to build, test, and deploy a RESTful API server in Go, backed by a PostgreSQL database for persistent storage of library data. Additionally, a Go client will be developed to interact with the server. The entire system, including the server, client, and PostgreSQL database, will be deployed to a Kubernetes cluster.

## Exercise Topics

This exercise contains several key topics and technologies, specifically focusing on building and deploying an app with relevant tools used by OpenShift Assisted Installer team. Here's a summary of the topics this project covers:

1. **RESTful API Development**: Designing and implementing a server in Go that provides a CRUD (Create, Read, Update, Delete) interface for managing library resources, demonstrating principles of RESTful service architecture.

2. **Go Programming**: Utilizing Go (Golang) for server and client development, highlighting its efficiency, ease of use for concurrent programming, and suitability for building high-performance web services.

3. **PostgreSQL Integration**: Leveraging PostgreSQL for persistent storage, showcasing how to integrate a relational database with a Go application to manage complex data relationships and ensure data integrity.

4. **Containerization with Docker/Podman**: Employing container technology to package the server and client applications along with their dependencies, facilitating consistency across different development and deployment environments.

5. **Kubernetes Orchestration**: Deploying the containerized applications to a Kubernetes cluster, covering aspects like deployments, services, ConfigMaps, Secrets, and PVCs, to manage and scale applications dynamically.

6. **Continuous Integration (CI) with Prow**: Implementing a Continuous Integration pipeline for this project by integrating Prow, a Kubernetes-based CI/CD system developed by the Kubernetes project.

## Project Components

### RESTful Go Server
- **Functionality**: The server will implement a CRUD API for managing library resources such as books, authors, and users. It will handle requests to add, retrieve, update, and delete information in the library system.
- **Technology**: Written in Go for its efficiency and suitability for concurrent tasks (concurrent is optional in this project).
- **Database Integration**: Utilizes PostgreSQL, a powerful open-source relational database, to store and manage library data. The Go server will interact with PostgreSQL.

### Go Client
- **Purpose**: A command-line client application written in Go, designed to interact with the server's RESTful API. This client will demonstrate how to authenticate, send requests, and process responses from the server.

### Kubernetes Deployment
- **Containerization**: Both the Go server and the client will be containerized using Docker/Podman, making them portable and easy to deploy.
- **Kubernetes Orchestration**: The containers will be deployed to a Kubernetes cluster, showcasing how to manage containerized applications at scale. This includes setting up Deployments for the server and database, ConfigMaps and Secrets for configuration and sensitive information, and Persistent Volume Claims (PVCs) for database storage needs.
- **Networking**: Services will be defined in Kubernetes to enable communication between the Go server and PostgreSQL database, as well as to expose the Go server to external traffic, demonstrating how to handle networking within a k8s cluster.

## Development and Deployment Workflow

1. **Local Development**: Initial development of the server and client will be done locally, with the Go server connecting to a local or containerized instance of PostgreSQL.
2. **Testing**: Unit tests will be written using Go testing tools like the standard `testing` package, Ginkgo for BDD testing, and Gomega for assertions. This ensures the reliability of both server and client functionalities. The tests should be able to be executed by calling a Makefile target using Skipper.
3. **Containerization**: Docker will be used to create containers for the server, client, and a PostgreSQL instance. Dockerfiles will specify the build process and runtime configuration of each component.
4. **Kubernetes Deployment**: The project will include YAML files for deploying the Docker containers to Kubernetes. This involves setting up the necessary k8s resources to ensure the system is scalable, maintainable, and secure.
