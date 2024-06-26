[![CircleCI](https://circleci.com/gh/alexfalkowski/status.svg?style=svg)](https://circleci.com/gh/alexfalkowski/status)
[![codecov](https://codecov.io/gh/alexfalkowski/status/graph/badge.svg?token=G6T3OIWUFK)](https://codecov.io/gh/alexfalkowski/status)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/status)](https://goreportcard.com/report/github.com/alexfalkowski/status)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/status.svg)](https://pkg.go.dev/github.com/alexfalkowski/status)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Status

This is just an alternative to using https://httpstat.us/

## Background

As the alternative suffers from stability reason, I wanted to stop using it.

### Why a service?

The service is needed for testing.

## Server

Below we outline all the endpoints:

### Status Code

Allows to set the status code of the response.

#### Request

```http
GET /v1/status/{code}
```
> [!NOTE]
> `code` is a number, e.g 200, 400, etc.

| Parameter | Description                                                                  |
| --------- | ---------------------------------------------------------------------------- |
| sleep     | The duration to sleep please check out https://pkg.go.dev/time#ParseDuration |

#### Response

The status with a description as text.

Example:

```http
200 OK
```

## Health

The system defines a way to monitor all of it's dependencies.

To configure we just need the have the following configuration:

```yaml
health:
  duration: 1s (how often to check)
  timeout: 1s (when we should timeout the check)
```

## Deployment

Since we are advocating building microservices, you would normally use a [container orchestration system](https://newrelic.com/blog/best-practices/container-orchestration-explained).

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- Golang

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

The get yourself setup, please run the following:

```sh
make setup
```

### Binaries

To make sure everything compiles for the app, please run the following:

```sh
make build-test
```

### Features

To run all the features, please run the following:

```sh
make features
```

### Changes

To see what has changed, please have a look at `CHANGELOG.md`
