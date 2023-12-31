# Rapigo Auth Microservice

![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/codegshinde/rapigo-auth-microservice/go.yml)
![Static Badge](https://img.shields.io/badge/GO-v1.21.4-green)
![GitHub License](https://img.shields.io/github/license/codegshinde/rapigo-auth-microservice?style=flat)
![GitHub last commit (by committer)](https://img.shields.io/github/last-commit/codegshinde/rapigo-auth-microservice)

Rapigo Auth Microservice is a Golang-based microservice for authentication.

## Features

- JWT-based authentication
- User and Admin management
- Password hashing using bcrypt

## Prerequisites

- [Go](https://golang.org/dl/) installed
- MongoDB set up and running

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/codegshinde/rapigo-auth-microservice.git
   ```

2. **Navigate to project directory:**

   ```bash
   cd rapigo-auth-microservice
   ```

3. **Create a .env file and set the required environment variables (See .env.sample).**

4. **Build the microservice:**

   ```bash
   go build -o bin/rapigo-auth-microservice cmd/main.go
   ```

5. **Run the microservice:**

   ```bash
   bin/rapigo-auth-microservice
   ```

## Usage

Your microservice is now running at [http://localhost:8080](http://localhost:8080). You can use the API for user and admin authentication.

Refer to the API documentation for detailed usage instructions.

## API Documentation

Describe where users can find detailed API documentation. You may use tools like Swagger or provide a Postman collection.

## Contributing

We welcome contributions! Please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This project is licensed under the MIT License.
