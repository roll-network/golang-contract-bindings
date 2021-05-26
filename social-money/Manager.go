// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogManagerMigrated\",\"inputs\":[{\"type\":\"address\",\"name\":\"newManager\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogRegistryChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"oldR\",\"indexed\":false},{\"type\":\"address\",\"name\":\"newR\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogTokenFactoryChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"oldTF\",\"indexed\":false},{\"type\":\"address\",\"name\":\"newTF\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"RegistryInstance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"TokenFactoryInstance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"approveProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_registry\"},{\"type\":\"address\",\"name\":\"_tokenFactory\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"initialized\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"migrateManager\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newManager\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"parseAddr\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[{\"type\":\"bytes\",\"name\":\"data\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"parsed\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setPlatformWallet\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newPlatformWallet\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setRegistry\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newRegistry\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setTokenFactory\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenFactory\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setTokenVesting\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenVesting\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingAddress\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingReferral\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"submitProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"hashIndex\"}]},{\"type\":\"function\",\"name\":\"submitReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_referral\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() view returns(address)
func (_Contracts *ContractsCaller) RegistryInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "RegistryInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() view returns(address)
func (_Contracts *ContractsSession) RegistryInstance() (common.Address, error) {
	return _Contracts.Contract.RegistryInstance(&_Contracts.CallOpts)
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() view returns(address)
func (_Contracts *ContractsCallerSession) RegistryInstance() (common.Address, error) {
	return _Contracts.Contract.RegistryInstance(&_Contracts.CallOpts)
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() view returns(address)
func (_Contracts *ContractsCaller) TokenFactoryInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "TokenFactoryInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() view returns(address)
func (_Contracts *ContractsSession) TokenFactoryInstance() (common.Address, error) {
	return _Contracts.Contract.TokenFactoryInstance(&_Contracts.CallOpts)
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() view returns(address)
func (_Contracts *ContractsCallerSession) TokenFactoryInstance() (common.Address, error) {
	return _Contracts.Contract.TokenFactoryInstance(&_Contracts.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsSession) Initialized() (bool, error) {
	return _Contracts.Contract.Initialized(&_Contracts.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Contracts *ContractsCallerSession) Initialized() (bool, error) {
	return _Contracts.Contract.Initialized(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(bytes data) pure returns(address parsed)
func (_Contracts *ContractsCaller) ParseAddr(opts *bind.CallOpts, data []byte) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "parseAddr", data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(bytes data) pure returns(address parsed)
func (_Contracts *ContractsSession) ParseAddr(data []byte) (common.Address, error) {
	return _Contracts.Contract.ParseAddr(&_Contracts.CallOpts, data)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(bytes data) pure returns(address parsed)
func (_Contracts *ContractsCallerSession) ParseAddr(data []byte) (common.Address, error) {
	return _Contracts.Contract.ParseAddr(&_Contracts.CallOpts, data)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(bytes32 _hashIndex) returns()
func (_Contracts *ContractsTransactor) ApproveProposal(opts *bind.TransactOpts, _hashIndex [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveProposal", _hashIndex)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(bytes32 _hashIndex) returns()
func (_Contracts *ContractsSession) ApproveProposal(_hashIndex [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveProposal(&_Contracts.TransactOpts, _hashIndex)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(bytes32 _hashIndex) returns()
func (_Contracts *ContractsTransactorSession) ApproveProposal(_hashIndex [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveProposal(&_Contracts.TransactOpts, _hashIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _tokenFactory) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", _registry, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _tokenFactory) returns()
func (_Contracts *ContractsSession) Initialize(_registry common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, _registry, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _tokenFactory) returns()
func (_Contracts *ContractsTransactorSession) Initialize(_registry common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, _registry, _tokenFactory)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(address _newManager) returns()
func (_Contracts *ContractsTransactor) MigrateManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "migrateManager", _newManager)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(address _newManager) returns()
func (_Contracts *ContractsSession) MigrateManager(_newManager common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.MigrateManager(&_Contracts.TransactOpts, _newManager)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(address _newManager) returns()
func (_Contracts *ContractsTransactorSession) MigrateManager(_newManager common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.MigrateManager(&_Contracts.TransactOpts, _newManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_Contracts *ContractsTransactor) SetPlatformWallet(opts *bind.TransactOpts, _newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPlatformWallet", _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_Contracts *ContractsSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetPlatformWallet(&_Contracts.TransactOpts, _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_Contracts *ContractsTransactorSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetPlatformWallet(&_Contracts.TransactOpts, _newPlatformWallet)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _newRegistry) returns()
func (_Contracts *ContractsTransactor) SetRegistry(opts *bind.TransactOpts, _newRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setRegistry", _newRegistry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _newRegistry) returns()
func (_Contracts *ContractsSession) SetRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetRegistry(&_Contracts.TransactOpts, _newRegistry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _newRegistry) returns()
func (_Contracts *ContractsTransactorSession) SetRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetRegistry(&_Contracts.TransactOpts, _newRegistry)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _newTokenFactory) returns()
func (_Contracts *ContractsTransactor) SetTokenFactory(opts *bind.TransactOpts, _newTokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setTokenFactory", _newTokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _newTokenFactory) returns()
func (_Contracts *ContractsSession) SetTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetTokenFactory(&_Contracts.TransactOpts, _newTokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _newTokenFactory) returns()
func (_Contracts *ContractsTransactorSession) SetTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetTokenFactory(&_Contracts.TransactOpts, _newTokenFactory)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_Contracts *ContractsTransactor) SetTokenVesting(opts *bind.TransactOpts, _newTokenVesting common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setTokenVesting", _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_Contracts *ContractsSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetTokenVesting(&_Contracts.TransactOpts, _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_Contracts *ContractsTransactorSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetTokenVesting(&_Contracts.TransactOpts, _newTokenVesting)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_Contracts *ContractsTransactor) SetVestingAddress(opts *bind.TransactOpts, _token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setVestingAddress", _token, _vestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_Contracts *ContractsSession) SetVestingAddress(_token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetVestingAddress(&_Contracts.TransactOpts, _token, _vestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_Contracts *ContractsTransactorSession) SetVestingAddress(_token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetVestingAddress(&_Contracts.TransactOpts, _token, _vestingBeneficiary)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_Contracts *ContractsTransactor) SetVestingReferral(opts *bind.TransactOpts, _token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setVestingReferral", _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_Contracts *ContractsSession) SetVestingReferral(_token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetVestingReferral(&_Contracts.TransactOpts, _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_Contracts *ContractsTransactorSession) SetVestingReferral(_token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetVestingReferral(&_Contracts.TransactOpts, _token, _vestingReferral)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x6709a651.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Contracts *ContractsTransactor) SubmitProposal(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitProposal", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x6709a651.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Contracts *ContractsSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitProposal(&_Contracts.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x6709a651.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Contracts *ContractsTransactorSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitProposal(&_Contracts.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage)
}

// SubmitReferral is a paid mutator transaction binding the contract method 0x1fd49f4b.
//
// Solidity: function submitReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Contracts *ContractsTransactor) SubmitReferral(opts *bind.TransactOpts, _hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitReferral", _hashIndex, _referral, _referralPercentage)
}

// SubmitReferral is a paid mutator transaction binding the contract method 0x1fd49f4b.
//
// Solidity: function submitReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Contracts *ContractsSession) SubmitReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitReferral(&_Contracts.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// SubmitReferral is a paid mutator transaction binding the contract method 0x1fd49f4b.
//
// Solidity: function submitReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Contracts *ContractsTransactorSession) SubmitReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitReferral(&_Contracts.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// ContractsLogManagerMigratedIterator is returned from FilterLogManagerMigrated and is used to iterate over the raw logs and unpacked data for LogManagerMigrated events raised by the Contracts contract.
type ContractsLogManagerMigratedIterator struct {
	Event *ContractsLogManagerMigrated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogManagerMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogManagerMigrated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogManagerMigrated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogManagerMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogManagerMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogManagerMigrated represents a LogManagerMigrated event raised by the Contracts contract.
type ContractsLogManagerMigrated struct {
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogManagerMigrated is a free log retrieval operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: event LogManagerMigrated(address indexed newManager)
func (_Contracts *ContractsFilterer) FilterLogManagerMigrated(opts *bind.FilterOpts, newManager []common.Address) (*ContractsLogManagerMigratedIterator, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogManagerMigrated", newManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogManagerMigratedIterator{contract: _Contracts.contract, event: "LogManagerMigrated", logs: logs, sub: sub}, nil
}

// WatchLogManagerMigrated is a free log subscription operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: event LogManagerMigrated(address indexed newManager)
func (_Contracts *ContractsFilterer) WatchLogManagerMigrated(opts *bind.WatchOpts, sink chan<- *ContractsLogManagerMigrated, newManager []common.Address) (event.Subscription, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogManagerMigrated", newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogManagerMigrated)
				if err := _Contracts.contract.UnpackLog(event, "LogManagerMigrated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogManagerMigrated is a log parse operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: event LogManagerMigrated(address indexed newManager)
func (_Contracts *ContractsFilterer) ParseLogManagerMigrated(log types.Log) (*ContractsLogManagerMigrated, error) {
	event := new(ContractsLogManagerMigrated)
	if err := _Contracts.contract.UnpackLog(event, "LogManagerMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRegistryChangedIterator is returned from FilterLogRegistryChanged and is used to iterate over the raw logs and unpacked data for LogRegistryChanged events raised by the Contracts contract.
type ContractsLogRegistryChangedIterator struct {
	Event *ContractsLogRegistryChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRegistryChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogRegistryChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRegistryChanged represents a LogRegistryChanged event raised by the Contracts contract.
type ContractsLogRegistryChanged struct {
	OldR common.Address
	NewR common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRegistryChanged is a free log retrieval operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: event LogRegistryChanged(address oldR, address newR)
func (_Contracts *ContractsFilterer) FilterLogRegistryChanged(opts *bind.FilterOpts) (*ContractsLogRegistryChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsLogRegistryChangedIterator{contract: _Contracts.contract, event: "LogRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchLogRegistryChanged is a free log subscription operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: event LogRegistryChanged(address oldR, address newR)
func (_Contracts *ContractsFilterer) WatchLogRegistryChanged(opts *bind.WatchOpts, sink chan<- *ContractsLogRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRegistryChanged)
				if err := _Contracts.contract.UnpackLog(event, "LogRegistryChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogRegistryChanged is a log parse operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: event LogRegistryChanged(address oldR, address newR)
func (_Contracts *ContractsFilterer) ParseLogRegistryChanged(log types.Log) (*ContractsLogRegistryChanged, error) {
	event := new(ContractsLogRegistryChanged)
	if err := _Contracts.contract.UnpackLog(event, "LogRegistryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogTokenFactoryChangedIterator is returned from FilterLogTokenFactoryChanged and is used to iterate over the raw logs and unpacked data for LogTokenFactoryChanged events raised by the Contracts contract.
type ContractsLogTokenFactoryChangedIterator struct {
	Event *ContractsLogTokenFactoryChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogTokenFactoryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogTokenFactoryChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogTokenFactoryChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogTokenFactoryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogTokenFactoryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogTokenFactoryChanged represents a LogTokenFactoryChanged event raised by the Contracts contract.
type ContractsLogTokenFactoryChanged struct {
	OldTF common.Address
	NewTF common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogTokenFactoryChanged is a free log retrieval operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: event LogTokenFactoryChanged(address oldTF, address newTF)
func (_Contracts *ContractsFilterer) FilterLogTokenFactoryChanged(opts *bind.FilterOpts) (*ContractsLogTokenFactoryChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogTokenFactoryChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsLogTokenFactoryChangedIterator{contract: _Contracts.contract, event: "LogTokenFactoryChanged", logs: logs, sub: sub}, nil
}

// WatchLogTokenFactoryChanged is a free log subscription operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: event LogTokenFactoryChanged(address oldTF, address newTF)
func (_Contracts *ContractsFilterer) WatchLogTokenFactoryChanged(opts *bind.WatchOpts, sink chan<- *ContractsLogTokenFactoryChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogTokenFactoryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogTokenFactoryChanged)
				if err := _Contracts.contract.UnpackLog(event, "LogTokenFactoryChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogTokenFactoryChanged is a log parse operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: event LogTokenFactoryChanged(address oldTF, address newTF)
func (_Contracts *ContractsFilterer) ParseLogTokenFactoryChanged(log types.Log) (*ContractsLogTokenFactoryChanged, error) {
	event := new(ContractsLogTokenFactoryChanged)
	if err := _Contracts.contract.UnpackLog(event, "LogTokenFactoryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
