// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controllers

import "strings"

// CarABI is the input ABI used to generate the binding from.
const CarABI = "[{\"inputs\":[{\"name\":\"_manufature\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CarBin is the compiled bytecode used for deploying new contracts.
const CarBin = `0x60606040523415600e57600080fd5b6040516101143803806101148339810160405280805190910190506000818051603a9291602001906040565b505060d1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10607f57805160ff191683800117855560a9565b8280016001018555821560a9579182015b8281111560a95782518255916020019190600101906090565b5060b392915060b7565b5090565b60ce91905b8082111560b3576000815560010160bc565b90565b6035806100df6000396000f3006060604052600080fd00a165627a7a72305820be34287574a0ab1dbb7d506d3d0076a97ddf2cded0411c6e29ccebe2165e36680029`

// DeployCar deploys a new Ethereum contract, binding an instance of Car to it.
func DeployCar(auth *bind.TransactOpts, backend bind.ContractBackend, _manufature string) (common.Address, *types.Transaction, *Car, error) {
	parsed, err := abi.JSON(strings.NewReader(CarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CarBin), backend, _manufature)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Car{CarCaller: CarCaller{contract: contract}, CarTransactor: CarTransactor{contract: contract}, CarFilterer: CarFilterer{contract: contract}}, nil
}

// Car is an auto generated Go binding around an Ethereum contract.
type Car struct {
	CarCaller     // Read-only binding to the contract
	CarTransactor // Write-only binding to the contract
	CarFilterer   // Log filterer for contract events
}

// CarCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarSession struct {
	Contract     *Car              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarCallerSession struct {
	Contract *CarCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarTransactorSession struct {
	Contract     *CarTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarRaw struct {
	Contract *Car // Generic contract binding to access the raw methods on
}

// CarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarCallerRaw struct {
	Contract *CarCaller // Generic read-only contract binding to access the raw methods on
}

// CarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarTransactorRaw struct {
	Contract *CarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCar creates a new instance of Car, bound to a specific deployed contract.
func NewCar(address common.Address, backend bind.ContractBackend) (*Car, error) {
	contract, err := bindCar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Car{CarCaller: CarCaller{contract: contract}, CarTransactor: CarTransactor{contract: contract}, CarFilterer: CarFilterer{contract: contract}}, nil
}

// NewCarCaller creates a new read-only instance of Car, bound to a specific deployed contract.
func NewCarCaller(address common.Address, caller bind.ContractCaller) (*CarCaller, error) {
	contract, err := bindCar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CarCaller{contract: contract}, nil
}

// NewCarTransactor creates a new write-only instance of Car, bound to a specific deployed contract.
func NewCarTransactor(address common.Address, transactor bind.ContractTransactor) (*CarTransactor, error) {
	contract, err := bindCar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CarTransactor{contract: contract}, nil
}

// NewCarFilterer creates a new log filterer instance of Car, bound to a specific deployed contract.
func NewCarFilterer(address common.Address, filterer bind.ContractFilterer) (*CarFilterer, error) {
	contract, err := bindCar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CarFilterer{contract: contract}, nil
}

// bindCar binds a generic wrapper to an already deployed contract.
func bindCar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Car *CarRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Car.Contract.CarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Car *CarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.Contract.CarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Car *CarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Car.Contract.CarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Car *CarCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Car.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Car *CarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Car *CarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Car.Contract.contract.Transact(opts, method, params...)
}
