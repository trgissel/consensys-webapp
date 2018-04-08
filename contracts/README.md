# Smart Contracts

## Pre-requisite

* install `solc` or `solcjs` to compile the contracts. See [Solidarity Documention](https://solidity.readthedocs.io/en/latest/installing-solidity.html#binary-packages)
* install go etherium binaries `geth` and `abigen`.  Installastion instructions can be found here [go-ethereum](https://github.com/ethereum/go-ethereum)
  * Note that on Mac OS X follow the `Building the source` instructions

## BUILD

* compile `Car.sol`.

```shell
  solcjs --bin Car.sol
  solcjs --abi Car.sol
```

* generate the go bindings via: `abigen --sol Car.sol --pkg contracts --out car.go`
* copy `car.go` to the server/controller directory: `cp car.go ../server/contracts/`