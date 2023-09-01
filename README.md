# Thanos-go

Thanos-go is a highly scalable and robust base project, designed and developed using Golang, Echo, Hexagonal
Architecture, and SOLID principles. It comes with pre-configured support for WebSocket, gRPC, and GraphQL to provide
seamless communication and data management in modern applications.

![Hexagonal Architecture](https://github.com/moein9gh/Thanos/blob/main/hexagonal-arch.png?raw=true)

## Table of Contents

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Project Structure](#project-structure)
- [Tests](#tests)
- [Contributing](#contributing)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes. See [Installation](#installation) for notes on how to deploy the project on a live system.

## Prerequisites

Before you can run the project, you need to have the following software installed on your system:

- Docker

You can download and install Docker from the [official Docker website](https://docker.com/).

## Installation

Clone the project repository:

```bash
git clone https://github.com/moein9gh/thanos-go.git
```

Navigate to the project directory:

```bash
cd thanos-go
```

## Running the Application

go to scripts folder

```bash
cd ./docker/scripts
```

Then run run.sh file to run in development mode

```bash
./run.sh
```

run log.sh file to see logs

```bash
./log.sh
```

run down.sh file to makes containers down

```bash
./down.sh
```

* Please make sure that .sh files in scripts folder have right permissions for execution.

## Project Structure

Below is an overview of the project structure:

```
├──Thanos-go/
│   ├── adapter/                  # Dependency Injection folder
│   ├── build/                    # Binary folder, possibly containing scripts or executables
│   ├── build_info/               # Binary folder, possibly containing scripts or executables
│   ├── cmd/                      # Documentation folder
│   ├── config/                    # Documentation folder
│   ├── contract/                 # Docker-related files folder
│   ├── delivery/                 # Docker-related files folder
│   ├── deployment/               # Data Transfer Objects folder
│   ├── docker/                   # Entity folder, for defining domain entities
│   ├── docs/                     # Environment configuration folder
│   ├── e/                        # Gateway folder, for external communication and services
│   ├── interactor/               # Interactor folder, for use cases and application logic
│   ├── log/                      # Localization folder, containing translations or localization configs
│   ├── logs/                     # Logging folder, for log files and logging configuration
│   ├── model/                    # Model folder, for domain models
│   ├── param/                    # Ports folder, for defining interfaces and adapters
│   ├── pkg/                      # Protocol Buffers folder, for defining gRPC services
│   ├── service/                  # Repository folder, for data access and persistence
│   ├── storage/                  # Schema folder, for defining data schemas or validation
│   ├── store/                    # Store folder, for managing application database's
│   ├── vendor/                   # Test folder
│   └── wsutil/                   # Utility functions folder
│
├── .pre-commit-config.yml         # ESLint ignore configuration file
├── .commitlint.config.js          # ESLint configuration file
├── go.mod                    # Git ignore configuration file
├── main.go                   # Prettier configuration file
│
├── Makefile                 # Package.json file for npm project
└── README.md                 # TypeScript configuration file

```

## Tests

To run the test suite, execute the following command:

```bash
go test
```

## Contributing

Contributions are always welcome! Feel free to submit a PR, create an issue, or get in touch with the project
maintainer moein9gh to discuss new ideas or improvements.
