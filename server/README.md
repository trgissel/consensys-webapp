# SERVER

## BUILD

## Prerequisite

Follow the BUILD AND DEPLOY steps found at in the contract [README.md](../contracts/README.md)

### ENV SETUP

* cd into server directory
* run: `export GOPATH=$(pwd)`
* run: `export export GOBIN=${GOPATH}/bin`

### External Dependency Changes

* run: `go get`
* run: `go install`

### Compile after src change

* run: `go build`

## USAGE

Login is required before use of non-login rest endpoints.

### Login

The `/login` endpoint returns a JWT use the JWT on subsequent requests to verify identity

1. POST to login endpoint and capture the returned JWT

```shell
JWT=$(curl -H "Content-Type: application/json" -X POST -d '{"username":"jim","password":"password"}' http://192.168.1.85:8000/login)
```

### CRUD

Always supply JWT on subsequent requests as `Authorization: Bearer` Header

```shell
curl -H "Authorization: Bearer ${JWT}" -X POST http://192.168.1.85:8000/contract
{"id":"a1d2ca66-4ab8-22f2-3875-a98924c66eb9"}
```

### Current state

The server is a work in progress. Please see [TODO List](TODO.md) for more details