# ExoplanetService
<div align="start">
<img src="golang.png" alt="Go Logo" width="100"/>
</div>

## Description

The problem statement for this service is explained in the PDF located at the root of this repository:

[Problem-Statement.pdf](Problem-Statement.pdf)

## Prerequisites

- Golang version go1.22.4 or later should be installed

## How to Run

### Using Makefile

To set up and run the service, use the following commands:

```bash
make setup
make run
```

## Using Docker

To build a Docker image and run the service in a Docker container, use the following commands.

```bash
    make docker
    make docker-run
```

## Running Unit Tests
To run unit tests, use the following command:
```bash
    make test
```
## API Endpoints

To access the api endpoints. please checkout the swagger documentation

http://localhost:8083/exoplanetservice/swagger/index.html

## Notes

## Author
Priya Kawane \
Software Engineer \
Thank You :)