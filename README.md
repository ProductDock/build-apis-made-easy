# Building apis made easy

## Prerequisites

- Typespec: `npm install -g @typespec/compiler`
- oapi-codegen: `go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest`
- Postman

## Run examples

After cloning this repository, check Makefile first. 

Choose server example:

- Run golang api (port 8084): `make run-go`

- Run spring boot api (port 8084): `make run-jvm`

Import `events-postman-collection.json` to postman to test the API.