# Smart Contracts

## Prerequisite

* install `solc` or `solcjs` to compile the contracts. See [Solidarity Documention](https://solidity.readthedocs.io/en/latest/installing-solidity.html#binary-packages)
* install go etherium binaries `geth` and `abigen`.  Installastion instructions can be found here [go-ethereum](https://github.com/ethereum/go-ethereum)

## BUILD AND DEPLOY

Perfom the Build and Deploy cycle whenever changes are made to the `Car.sol` file

### BUILD

* compile `Car.sol`.

```shell
  solcjs --bin Car.sol
  solcjs --abi Car.sol
```

* generate the go bindings via: `abigen --abi Car_sol_Car.abi  --pkg contracts -type Car -out car.go --bin Car_sol_Car.bin`


### DEPLOY

* If you don't have an ethereum account on target then network create one
  * See [Managing your accounts](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) for more information
* Deploy Car Smart Contract, see [How To: Deploy Smart Contracts Onto The Ethereum Blockchain](https://medium.com/mercuryprotocol/dev-highlights-of-this-week-cb33e58c745f)

