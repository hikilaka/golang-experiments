[![CI](https://github.com/hikilaka/golang-experiments/actions/workflows/ci.yaml/badge.svg)](https://github.com/hikilaka/golang-experiments/actions/workflows/ci.yaml)

# Golang Experiments

This repository is a personal project created to learn and experiment with the Go programming language. It contains various implementations and tests to explore Go's features, including generics, interfaces, and data structures.

## Project Structure

The repository is organized as follows:

- **`cmd/main`**: Contains the entry point of the application. Demonstrates the usage of the `LinkedList` implementation.
- **`types`**: Contains the implementation of a generic `LinkedList` and its supporting types.
- **`test`**: Contains unit tests for the `LinkedList` implementation using the `testify` library.

## Features

- **Generic Linked List**: A custom implementation of a linked list that supports generic types.
- **Unit Tests**: Comprehensive tests to validate the functionality of the linked list.
- **Go Modules**: The project uses Go modules for dependency management.

## Prerequisites

- Go 1.24 or later is required to run this project.

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/hikilaka/golang-experiments.git
   cd golang-experiments
   ```

2. Install dependencies:
    ```bash
    go get .
    ```

3. Run the application:
    ```bash
    go run ./cmd/main
    ```

4. Run the tests:
    ```bash
    go test ./test -v
    ```

## Learning Goals
This project is designed to help understand:
* How to use Go's generics to create reusable data structures.
* How to leverage Go's built-in concurrency
* How to write and structure unit tests in Go.
* How to manage dependencies using Go modules.
* How to use GitHub Actions for continuous integration.

## Contributing
This is a personal learning project, but feedback and suggestions are welcome! Feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License. See the LICENSE file for details.