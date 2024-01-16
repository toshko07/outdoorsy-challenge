# outdoorsy-challenge

## Table of Contents

- [About](#about)
- [Requirements](#requirements)
- [General Structure](#general-structure)
- [Configuration](#configuration)
- [Running](#running)
- [Testing](#testing)


# About
This is a solution for the outdoorsy-challenge.

# Requirements
- Go 1.21.6
- Docker version 24.0.5
- Docker Compose version v2.17.2

# General Structure
The application is structured as follows:

```
.
├── api # swagger api definition and generated models
├── cmd # main entrypoint for the applications
├── internal # private packages
│   ├── configs
│   ├── controllers # http handlers
│   ├── db # database setup
│   ├── models # business models
│   ├── repositories # data access layer
|   ├── component_test
│   └── services  # business logic
├── Makefile # makefile for running make commands related to the application
```


# Configuration
The application is configured using environment variables. Copy `.env.example` to `.env` and edit the values if needed. Database configuration is set in the `docker-compose.yml` file.


# Running
To run the application, execute the following commands:

```
make run-db
make run-app
```

# Testing
Here are some example curl commands to test the application:

```
curl --request GET \
  --url http://localhost:8181/v1/rentals/1
```
```
curl --request GET \
  --url http://localhost:8181/v1/rentals/404
```

```
curl --request GET \
  --url 'http://localhost:8181/v1/rentals?near=33.64%2C-117.93&price_min=9000&price_max=75000&limit=3&offset=3&sort=price'
```

```
curl --request GET \
  --url 'http://localhost:8181/v1/rentals
```