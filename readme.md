# Getting Started with Create React App

This project is skeleton for ev backend microservice. This project is using grpc for communication between microservices and user REST API. This project also. We use air for hot reload and mockery to mock file for test.
This project have submodules for protobuf. To use this project, you need to clone this project with `--recurse-submodules` flag.
for more information about submodules, you can read [here](https://git-scm.com/book/en/v2/Git-Tools-Submodules).
We recommend to use go 1.21 or higher.

## Pre-requisites

- [Grpc](https://grpc.io/docs/languages/go/quickstart/)
- [Air](https://github.com/cosmtrek/air)
- [Mockery](https://vektra.github.io/mockery/latest/)

## Installation

```bash
  git clone --recurse-submodules git@gitlab.com:robinhood_ppv/ev-rental/backend/ev-skeleton.git
  cd ev-skeleton
  go mod tidy
```

## Pull submodules

To push the submodules, run the following command:

```bash
  git submodule update --remote --merge
```

## Update submodules

To update the submodules, run the following command:

```bash
  git push --recurse-submodules=on-demand
```

## Generate protobuf

To generate the protobuf files, run the following command:

```bash
  protoc --go_out=./internal/grpc --go_opt=paths=source_relative \
    --go-grpc_out=./internal/grpc --go-grpc_opt=paths=source_relative \
    protobuf/hello.proto
```

## Generate mock

To generate the mock files, run the following command:

```bash
  mockery --all
```

## Run

To run the application, run the following command:

```bash
  export GO_ENV=local && air
```

## Test

To run the tests, run the following command:

```bash
  go test ./... -cover
```

or

```bash
  go test -coverprofile=coverage.out ./... ;
  go tool cover -html=coverage.out
```
