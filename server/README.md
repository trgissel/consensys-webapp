# SERVER

## BUILD

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

1. POST to login endpoint e.g. `curl -H "Content-Type: application/json" -X POST -d '{"username":"jim","password":"password"}' http://localhost:8000/login`
1. capture the returned JWT
1. supply JWT on subsequent requests as `Authorization: Bearer` Header, e.g. `"Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMxOTU0OTIsIm5hbWUiOiJqaW0iLCJwYXNzd29yZCI6InBhc3N3b3JkIn0.idxk9EEkFChmmxys4BRJLUM17w3ucGvNVM9m19HA3zI"`
```
$ curl -H "Content-Type: application/json" -X POST -d '{"username":"jim","password":"password"}' http://localhost:8000/login
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMxOTY5NzEsIm5hbWUiOiJqaW0iLCJwYXNzd29yZCI6InBhc3N3b3JkIn0.AHBO0cP9OQg3SYJR1caBy_2xPyyKI_KYEVJAHrrCl1c
$ curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMxOTY5NzEsIm5hbWUiOiJqaW0iLCJwYXNzd29yZCI6InBhc3N3b3JkIn0.AHBO0cP9OQg3SYJR1caBy_2xPyyKI_KYEVJAHrrCl1c" -X POST http://localhost:8000/contract
{"id":"a1d2ca66-4ab8-22f2-3875-a98924c66eb9"}
```
