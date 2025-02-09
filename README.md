# Authentication Service

The **Authentication Service** is a core component of the Pharmakart platform, responsible for user authentication, authorization, and token management. It provides secure endpoints for user registration, login, token validation, and password management.

---

## Table of Contents
1. [Overview](#overview)
2. [Features](#features)
3. [Prerequisites](#prerequisites)
4. [Setup and Installation](#setup-and-installation)
5. [Running the Service](#running-the-service)
6. [Environment Variables](#environment-variables)
7. [Contributing](#contributing)
8. [License](#license)

---

## Overview

The Authentication Service handles:
- User registration and login.
- JWT token generation and validation.
- Password management and recovery.

It is built using **gRPC** for communication and **PostgreSQL** for data storage.

---

## Features

- **User Registration**: Create new user accounts.
- **Login and Token Generation**: Generate JWT tokens for authenticated users.
- **Token Validation**: Validate JWT tokens for secure access.
- **Password Management**: Secure password storage and recovery.

---

## Prerequisites

Before setting up the service, ensure you have the following installed:
- **Docker**
- **Go** (for building and running the service).
- **Protobuf Compiler** (`protoc`) for generating gRPC/protobuf files.

---

## Setup and Installation

### 1. Clone the Repository
Clone the repository and navigate to the authentication service directory:
```bash
git clone https://github.com/PharmaKart/authentication-svc.git
cd authentication-svc
```

### 2. Generate Protobuf Files
Generate the protobuf files using the provided `Makefile`:
```bash
make proto
```

### 3. Install Dependencies
Run the following command to tidy up Go modules:
```bash
go mod tidy
```

### 4. Build the Service
Build the Docker image for the service:
```bash
docker build -t authentication-service .
```

---

## Running the Service

### Start the Service
You can run the service using one of the following methods:

#### Using Docker
```bash
docker run -p 50051:50051 --env-file .env authentication-service
```

#### Using Makefile
```bash
make run
```

The service will be available at:
- **gRPC**: `localhost:50051`

### Stop the Service
To stop the service, simply terminate the process running the container or use:
```bash
docker stop <container_id>
```

---

## Environment Variables

The service requires the following environment variables. Create a `.env` file in the `authentication-svc` directory with the following:

```env
AUTH_DB_HOST=postgres
AUTH_DB_PORT=5432
AUTH_DB_USER=postgres
AUTH_DB_PASSWORD=yourpassword
AUTH_DB_NAME=pharmakartdb
JWT_SECRET=your-jwt-secret
```

---

## Contributing

Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a detailed description of your changes.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Support

For any questions or issues, please open an issue in the repository or contact the maintainers.

