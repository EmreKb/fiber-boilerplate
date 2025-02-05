# Fiber Boilerplate

A boilerplate project for building backend services using the Fiber framework in Go. This project includes Docker support, database integration, and monitoring setup.

## Description

This project serves as a starting point for developing a RESTful API using Go with Fiber. It includes:

- Docker and Docker Compose configuration for containerization
- Database integration with SQLC
- Migration setup
- Monitoring with Prometheus and Grafana

## Installation

To set up the project, follow these steps:

1. **Clone the repository:**

```bash
git clone https://github.com/EmreKb/fiber-boilerplate
cd fiber-boilerplate
rm -rf .git
```

2. **Install Go and dependencies:**

Ensure you have Go installed on your system. Then, install the project dependencies:

```bash
go mod tidy
```

3. **Set up Docker and Docker Compose:**

Install Docker and Docker Compose if you haven't already.

4. **Copy the environment file:**

```bash
cp .env.example .env
```

Update the `.env` file with your specific configurations.

5. **Install SQLC:**

Install SQLC for database operations:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

6. **Start the server:**

Use the provided Makefile to start the server:

```bash
make run
```

## Usage

### Running the Server

To start the server:

```bash
make run
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with clear commit messages.
4. Open a Pull Request against the `main` branch.


