// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CarABI is the input ABI used to generate the binding from.
const CarABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMiles\",\"outputs\":[{\"name\":\"miles\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSafetyInspectionsId\",\"outputs\":[{\"name\":\"safetyId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"setEmmisionInspectionsId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isScrapped\",\"type\":\"bool\"}],\"name\":\"setScrapped\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"milesToAdd\",\"type\":\"uint32\"}],\"name\":\"addMiles\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getScrapped\",\"outputs\":[{\"name\":\"isScrapped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEmmisionInspectionsId\",\"outputs\":[{\"name\":\"emmissionsId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"setSafetyInspectionsId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"inPrimaryName\",\"type\":\"bytes32\"},{\"name\":\"inSecondaryName\",\"type\":\"bytes32\"},{\"name\":\"inOwnerAddress\",\"type\":\"address\"}],\"name\":\"changeOwners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"inManufature\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CarBin is the compiled bytecode used for deploying new contracts.
const CarBin = `6060604052341561000f57600080fd5b6040516020806105ca83398101604052808051906020019091905050806000816000191690555033600360020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506105408061008a6000396000f300606060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063090e195e1461009e5780631df8668b146100db5780632f37a0b51461010c57806389a1e5d11461013357806394c7cfb8146101585780639624dfaf14610181578063a6069c53146101ae578063b4212b5d146101df578063b5705da414610206575b600080fd5b34156100a957600080fd5b6100b1610259565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b34156100e657600080fd5b6100ee610277565b60405180826000191660001916815260200191505060405180910390f35b341561011757600080fd5b610131600480803560001916906020019091905050610284565b005b341561013e57600080fd5b610156600480803515159060200190919050506102f4565b005b341561016357600080fd5b61017f600480803563ffffffff16906020019091905050610371565b005b341561018c57600080fd5b6101946103bb565b604051808215151515815260200191505060405180910390f35b34156101b957600080fd5b6101c16103d2565b60405180826000191660001916815260200191505060405180910390f35b34156101ea57600080fd5b6102046004808035600019169060200190919050506103de565b005b341561021157600080fd5b6102576004808035600019169060200190919080356000191690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061044f565b005b6000600660019054906101000a900467ffffffffffffffff16905090565b6000600160000154905090565b600360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102e3576102f1565b806001800181600019169055505b50565b600360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103535761036e565b80600660006101000a81548160ff0219169083151502179055505b50565b8063ffffffff16600660019054906101000a900467ffffffffffffffff1601600660016101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050565b6000600660009054906101000a900460ff16905090565b60006001800154905090565b600360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561043d5761044c565b80600160000181600019169055505b50565b600360020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156104ae5761050f565b8260036000018160001916905550816003600101816000191690555080600360020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050505600a165627a7a72305820f1754ab5cc4402ae2857ec26bd8623007ce0bc21b16b260b55787e2d714616200029`

// DeployCar deploys a new Ethereum contract, binding an instance of Car to it.
func DeployCar(auth *bind.TransactOpts, backend bind.ContractBackend, inManufature [32]byte) (common.Address, *types.Transaction, *Car, error) {
	parsed, err := abi.JSON(strings.NewReader(CarABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CarBin), backend, inManufature)
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

// GetEmmisionInspectionsId is a free data retrieval call binding the contract method 0xa6069c53.
//
// Solidity: function getEmmisionInspectionsId() constant returns(emmissionsId bytes32)
func (_Car *CarCaller) GetEmmisionInspectionsId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Car.contract.Call(opts, out, "getEmmisionInspectionsId")
	return *ret0, err
}

// GetEmmisionInspectionsId is a free data retrieval call binding the contract method 0xa6069c53.
//
// Solidity: function getEmmisionInspectionsId() constant returns(emmissionsId bytes32)
func (_Car *CarSession) GetEmmisionInspectionsId() ([32]byte, error) {
	return _Car.Contract.GetEmmisionInspectionsId(&_Car.CallOpts)
}

// GetEmmisionInspectionsId is a free data retrieval call binding the contract method 0xa6069c53.
//
// Solidity: function getEmmisionInspectionsId() constant returns(emmissionsId bytes32)
func (_Car *CarCallerSession) GetEmmisionInspectionsId() ([32]byte, error) {
	return _Car.Contract.GetEmmisionInspectionsId(&_Car.CallOpts)
}

// GetMiles is a free data retrieval call binding the contract method 0x090e195e.
//
// Solidity: function getMiles() constant returns(miles uint64)
func (_Car *CarCaller) GetMiles(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Car.contract.Call(opts, out, "getMiles")
	return *ret0, err
}

// GetMiles is a free data retrieval call binding the contract method 0x090e195e.
//
// Solidity: function getMiles() constant returns(miles uint64)
func (_Car *CarSession) GetMiles() (uint64, error) {
	return _Car.Contract.GetMiles(&_Car.CallOpts)
}

// GetMiles is a free data retrieval call binding the contract method 0x090e195e.
//
// Solidity: function getMiles() constant returns(miles uint64)
func (_Car *CarCallerSession) GetMiles() (uint64, error) {
	return _Car.Contract.GetMiles(&_Car.CallOpts)
}

// GetSafetyInspectionsId is a free data retrieval call binding the contract method 0x1df8668b.
//
// Solidity: function getSafetyInspectionsId() constant returns(safetyId bytes32)
func (_Car *CarCaller) GetSafetyInspectionsId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Car.contract.Call(opts, out, "getSafetyInspectionsId")
	return *ret0, err
}

// GetSafetyInspectionsId is a free data retrieval call binding the contract method 0x1df8668b.
//
// Solidity: function getSafetyInspectionsId() constant returns(safetyId bytes32)
func (_Car *CarSession) GetSafetyInspectionsId() ([32]byte, error) {
	return _Car.Contract.GetSafetyInspectionsId(&_Car.CallOpts)
}

// GetSafetyInspectionsId is a free data retrieval call binding the contract method 0x1df8668b.
//
// Solidity: function getSafetyInspectionsId() constant returns(safetyId bytes32)
func (_Car *CarCallerSession) GetSafetyInspectionsId() ([32]byte, error) {
	return _Car.Contract.GetSafetyInspectionsId(&_Car.CallOpts)
}

// GetScrapped is a free data retrieval call binding the contract method 0x9624dfaf.
//
// Solidity: function getScrapped() constant returns(isScrapped bool)
func (_Car *CarCaller) GetScrapped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Car.contract.Call(opts, out, "getScrapped")
	return *ret0, err
}

// GetScrapped is a free data retrieval call binding the contract method 0x9624dfaf.
//
// Solidity: function getScrapped() constant returns(isScrapped bool)
func (_Car *CarSession) GetScrapped() (bool, error) {
	return _Car.Contract.GetScrapped(&_Car.CallOpts)
}

// GetScrapped is a free data retrieval call binding the contract method 0x9624dfaf.
//
// Solidity: function getScrapped() constant returns(isScrapped bool)
func (_Car *CarCallerSession) GetScrapped() (bool, error) {
	return _Car.Contract.GetScrapped(&_Car.CallOpts)
}

// AddMiles is a paid mutator transaction binding the contract method 0x94c7cfb8.
//
// Solidity: function addMiles(milesToAdd uint32) returns()
func (_Car *CarTransactor) AddMiles(opts *bind.TransactOpts, milesToAdd uint32) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "addMiles", milesToAdd)
}

// AddMiles is a paid mutator transaction binding the contract method 0x94c7cfb8.
//
// Solidity: function addMiles(milesToAdd uint32) returns()
func (_Car *CarSession) AddMiles(milesToAdd uint32) (*types.Transaction, error) {
	return _Car.Contract.AddMiles(&_Car.TransactOpts, milesToAdd)
}

// AddMiles is a paid mutator transaction binding the contract method 0x94c7cfb8.
//
// Solidity: function addMiles(milesToAdd uint32) returns()
func (_Car *CarTransactorSession) AddMiles(milesToAdd uint32) (*types.Transaction, error) {
	return _Car.Contract.AddMiles(&_Car.TransactOpts, milesToAdd)
}

// ChangeOwners is a paid mutator transaction binding the contract method 0xb5705da4.
//
// Solidity: function changeOwners(inPrimaryName bytes32, inSecondaryName bytes32, inOwnerAddress address) returns()
func (_Car *CarTransactor) ChangeOwners(opts *bind.TransactOpts, inPrimaryName [32]byte, inSecondaryName [32]byte, inOwnerAddress common.Address) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "changeOwners", inPrimaryName, inSecondaryName, inOwnerAddress)
}

// ChangeOwners is a paid mutator transaction binding the contract method 0xb5705da4.
//
// Solidity: function changeOwners(inPrimaryName bytes32, inSecondaryName bytes32, inOwnerAddress address) returns()
func (_Car *CarSession) ChangeOwners(inPrimaryName [32]byte, inSecondaryName [32]byte, inOwnerAddress common.Address) (*types.Transaction, error) {
	return _Car.Contract.ChangeOwners(&_Car.TransactOpts, inPrimaryName, inSecondaryName, inOwnerAddress)
}

// ChangeOwners is a paid mutator transaction binding the contract method 0xb5705da4.
//
// Solidity: function changeOwners(inPrimaryName bytes32, inSecondaryName bytes32, inOwnerAddress address) returns()
func (_Car *CarTransactorSession) ChangeOwners(inPrimaryName [32]byte, inSecondaryName [32]byte, inOwnerAddress common.Address) (*types.Transaction, error) {
	return _Car.Contract.ChangeOwners(&_Car.TransactOpts, inPrimaryName, inSecondaryName, inOwnerAddress)
}

// SetEmmisionInspectionsId is a paid mutator transaction binding the contract method 0x2f37a0b5.
//
// Solidity: function setEmmisionInspectionsId(id bytes32) returns()
func (_Car *CarTransactor) SetEmmisionInspectionsId(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "setEmmisionInspectionsId", id)
}

// SetEmmisionInspectionsId is a paid mutator transaction binding the contract method 0x2f37a0b5.
//
// Solidity: function setEmmisionInspectionsId(id bytes32) returns()
func (_Car *CarSession) SetEmmisionInspectionsId(id [32]byte) (*types.Transaction, error) {
	return _Car.Contract.SetEmmisionInspectionsId(&_Car.TransactOpts, id)
}

// SetEmmisionInspectionsId is a paid mutator transaction binding the contract method 0x2f37a0b5.
//
// Solidity: function setEmmisionInspectionsId(id bytes32) returns()
func (_Car *CarTransactorSession) SetEmmisionInspectionsId(id [32]byte) (*types.Transaction, error) {
	return _Car.Contract.SetEmmisionInspectionsId(&_Car.TransactOpts, id)
}

// SetSafetyInspectionsId is a paid mutator transaction binding the contract method 0xb4212b5d.
//
// Solidity: function setSafetyInspectionsId(id bytes32) returns()
func (_Car *CarTransactor) SetSafetyInspectionsId(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "setSafetyInspectionsId", id)
}

// SetSafetyInspectionsId is a paid mutator transaction binding the contract method 0xb4212b5d.
//
// Solidity: function setSafetyInspectionsId(id bytes32) returns()
func (_Car *CarSession) SetSafetyInspectionsId(id [32]byte) (*types.Transaction, error) {
	return _Car.Contract.SetSafetyInspectionsId(&_Car.TransactOpts, id)
}

// SetSafetyInspectionsId is a paid mutator transaction binding the contract method 0xb4212b5d.
//
// Solidity: function setSafetyInspectionsId(id bytes32) returns()
func (_Car *CarTransactorSession) SetSafetyInspectionsId(id [32]byte) (*types.Transaction, error) {
	return _Car.Contract.SetSafetyInspectionsId(&_Car.TransactOpts, id)
}

// SetScrapped is a paid mutator transaction binding the contract method 0x89a1e5d1.
//
// Solidity: function setScrapped(isScrapped bool) returns()
func (_Car *CarTransactor) SetScrapped(opts *bind.TransactOpts, isScrapped bool) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "setScrapped", isScrapped)
}

// SetScrapped is a paid mutator transaction binding the contract method 0x89a1e5d1.
//
// Solidity: function setScrapped(isScrapped bool) returns()
func (_Car *CarSession) SetScrapped(isScrapped bool) (*types.Transaction, error) {
	return _Car.Contract.SetScrapped(&_Car.TransactOpts, isScrapped)
}

// SetScrapped is a paid mutator transaction binding the contract method 0x89a1e5d1.
//
// Solidity: function setScrapped(isScrapped bool) returns()
func (_Car *CarTransactorSession) SetScrapped(isScrapped bool) (*types.Transaction, error) {
	return _Car.Contract.SetScrapped(&_Car.TransactOpts, isScrapped)
}
