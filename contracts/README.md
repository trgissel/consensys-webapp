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

* generate the go bindings via: `abigen --abi Car_sol_Car.abi  --pkg contracts -type Car -out car.go --bin Car_sol_Car.bin`
* copy `car.go` to the server/controller directory: `cp car.go ../server/contracts/`