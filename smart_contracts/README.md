# Smart Contracts

## General Prerequisite

* install `solc` or `solcjs` to compile the contracts. See [Solidarity Documention](https://solidity.readthedocs.io/en/latest/installing-solidity.html#binary-packages)
* install go etherium binaries `geth` and `abigen`.  Installastion instructions can be found here [go-ethereum](https://github.com/ethereum/go-ethereum)

## Test

These steps for for testing Car contract with truffle

### Test Prerequisite

* nodejs 8 or newer.  See instructions at [nodejs.org](https://nodejs.org/en/)
* truffle: `npm install -g truffle`

### Test Build

* cd into migrations subdirectory: `cd <smart_contracts_dir>/migrations`
* run `truffle compile`

#### Run Automated Test

* cd into smart_contracts directory
* run `truffle test`

#### Interactive Testing

* cd into migrations subdirectory: `cd <smart_contracts_dir>/migrations`
* start interactive testing session with: `truffle develop`

Example session

```shell
$ truffle develop
truffle(develop)> migrate --reset
Using network 'develop'.

Running migration: 1_initial_migration.js
  Replacing Migrations...
  ... 0x178f8fbe39c130c2bc607d7388ae105731ca68fd87025de694054b79f11375d0
  Migrations: 0x8cdaf0cd259887258bc13a92c0a6da92698644c0
Saving successful migration to network...
  ... 0xd7bc86d31bee32fa3988f1c1eabce403a1b5d570340a3a9cdba53a472ee8c956
Saving artifacts...
Running migration: 2_deploy_contracts.js
  Replacing Car...
  ... 0xd581700b317c70d03af8155a577914d2e5a82a9e7fa7fb63db0d315727f811a9
  Car: 0x345ca3e014aaf5dca488057592ee47305d9b3e10
Saving successful migration to network...
  ... 0xf36163615f41ef7ed8f4a8f192149a0bf633fe1a2398ce001bf44c43dc7bdda0
Saving artifacts...

truffle(develop)> Car.deployed().then(function(instance){return instance.getMiles();});
BigNumber { s: 1, e: 0, c: [ 0 ] }

truffle(develop)> Car.deployed().then(function(instance){return instance.addMiles(1);});
{ tx: '0x021f635e6937b8251939c4caf997e90df28108e836aa68248f3f6261d42f7115',
  receipt: 
   { transactionHash: '0x021f635e6937b8251939c4caf997e90df28108e836aa68248f3f6261d42f7115',
     transactionIndex: 0,
     blockHash: '0xd3e5e0b45c5b78a81cbb4b2d0c635bb60620c725b5fc25888644db8309b91333',
     blockNumber: 9,
     gasUsed: 42315,
     cumulativeGasUsed: 42315,
     contractAddress: null,
     logs: [],
     status: '0x01',
     logsBloom: '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' },
  logs: [] }

truffle(develop)> Car.deployed().then(function(instance){return instance.getMiles();});
BigNumber { s: 1, e: 0, c: [ 1 ] }
```

## Build and Deploy

Perfom the Build and Deploy cycle whenever changes are made to the `Car.sol` file to ensure the server is up to date

## Build

These steps are for deployment into the ethereum network

* cd into the `contracts` directory

* compile `Car.sol`.

```shell
  solcjs --bin Car.sol
  solcjs --abi Car.sol
```

* generate the go bindings via: `abigen --abi Car_sol_Car.abi  --pkg contracts -type Car -out car.go --bin Car_sol_Car.bin`

* copy `car.go` to `../../server/contracts/`

```shell
cp car.go ../../server/contracts/
```

### Deploy into Ethereum Network

* If you don't have an ethereum account on target network then create one
  * See [Managing your accounts](https://github.com/ethereum/go-ethereum/wiki/Managing-your-accounts) for more information
* Deploy Car Smart Contract, see [How To: Deploy Smart Contracts Onto The Ethereum Blockchain](https://medium.com/mercuryprotocol/dev-highlights-of-this-week-cb33e58c745f)
