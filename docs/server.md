# Building a RESTful API in Go with a PostgreSQL Database and Kubernetes Deployment

## Project Overview

This project aims to develop a RESTful API server in Go, designed to manage a virtual library system. Utilizing a PostgreSQL database for data persistence, the project will leverage GORM, a powerful ORM library for Go, to facilitate seamless database interactions. Enhanced by clear and concise API documentation and code generation via Swagger, the server architecture is designed for maintainability and scalability. To ensure reliability and quality, the server's functionality will be validated through comprehensive testing with Gomega and Ginkgo. The development and testing workflows are streamlined using Makefile and Skipper. Furthermore, the incorporation of Logrus for structured logging will improve the server's observability, allowing for better monitoring and troubleshooting capabilities. The deployment strategy is Docker/Podman and then k8s using minikube.


## Tools and Libraries

- **Swagger:** Utilized for API documentation and code generation, enabling clear communication of API structures and expectations.
- **Gomega & Ginkgo:** A matcher/assertion library and a BDD testing framework, respectively, used for comprehensive testing of Go applications.
- **PostgreSQL:** An open-source relational database for persisting library data.
- **GORM:** An ORM (Object-Relational Mapping) library for Go, designed to interact with databases like PostgreSQL in an idiomatic Go way. It simplifies CRUD operations and database schema migration, making it easier to work with relational databases.
- **Docker/Podman:** For containerizing the application.
- **Makefile/Skipper:** Utilized to automate the build, test, and deployment processes. Makefile serves as a tool to define a set of tasks to be executed, simplifying the development process. Skipper, often used in conjunction with Makefile, provides a containerized environment for running these tasks.
- **Logrus:** A structured logger for Go, offering advanced logging capabilities like structured logs in JSON format, which are invaluable for modern application debugging and monitoring. Its use is critical for maintaining high-quality, readable, and searchable logs.
- **Minikube:** A tool that enables the running of Kubernetes locally.

## Server Functionality and Endpoints

The API will manage a collection of books, allowing interaction through the following endpoints:

1. **Add a New Book**
   - **Method:** POST
   - **Endpoint:** `/books`
   - **Description:** Adds a new book to the library, including details like title, author and publication date.

2. **Get a List of All Books**
   - **Method:** GET
   - **Endpoint:** `/books`
   - **Description:** Retrieves a comprehensive list of all books, with options to filter by author, title, etc.

3. **Get Details of a Single Book**
   - **Method:** GET
   - **Endpoint:** `/books/{id}`
   - **Description:** Fetches detailed information about a book using its unique ID.

4. **Update Book Information**
   - **Method:** PUT
   - **Endpoint:** `/books/{id}`
   - **Description:** Updates existing book details like title and author.

5. **Delete a Book**
   - **Method:** DELETE
   - **Endpoint:** `/books/{id}`
   - **Description:** Removes a book from the library based on its ID.
