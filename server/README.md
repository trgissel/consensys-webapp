# SERVER

## BUILD

## Prerequisite

Follow the Build AND Deploy steps found at in the smart_contract [README.md](../smart_contracts/README.md)

### ENV SETUP

* cd into server directory
* run: `export GOPATH=$(pwd)`
* run: `export export GOBIN=${GOPATH}/bin`

### External Dependency Changes

* run: `go get`
* run: `go install`

### Compile after src change

* run: `go build`

### Update Config

* replace `UTC--2018-04-09T13-01-22.708849993Z--064444c4cc39e4d2df6eca18fb4420574a075bb0` file with valid keystore for account
* update the values in `ethereumConfig.json` to the appropraite values

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
$JWT=$(curl -H "Content-Type: application/json" -X POST -d '{"username":"jim","password":"password"}' http://192.168.1.85:8000/login)
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   192  100   152  100    40  57099  15026 --:--:-- --:--:-- --:--:-- 76000
$ echo $JWT
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjMzODYwMDgsIm5hbWUiOiJqaW0iLCJwYXNzd29yZCI6InBhc3N3b3JkIn0.LbFq_0TlQnMcJH2aDFjBV9Kscwt9-yuvx7DVlxcDgZU
$ curl -H "Authorization: Bearer ${JWT}" -X POST http://192.168.1.85:8000/contract
{"address":"0xe57e01183ef46fa4c48576dc05fa4870e7d0ca56","txID":"0xbcc9f4257a436ae01ccfef9b5b92659bad5678746fee0261185a6c60f902105b"}
$ curl -H "Authorization: Bearer ${JWT}" http://192.168.1.85:8000/contract/0xe57e01183ef46fa4c48576dc05fa4870e7d0ca56
{"address":"0xe57e01183ef46fa4c48576dc05fa4870e7d0ca56","txID":"0xbcc9f4257a436ae01ccfef9b5b92659bad5678746fee0261185a6c60f902105b"}
$ curl -H "Authorization: Bearer ${JWT}" http://192.168.1.85:8000/contract/0xe57e01183ef46fa4c48576dc05fa4870e7d0ca56/mileage
0
```

### Current state

The server is a work in progress. Please see [TODO List](TODO.md) for more details