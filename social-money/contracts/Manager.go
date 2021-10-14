// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ManagerMetaData contains all meta data concerning the Manager contract.
var ManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogManagerMigrated\",\"inputs\":[{\"type\":\"address\",\"name\":\"newManager\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogRegistryChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"oldR\",\"indexed\":false},{\"type\":\"address\",\"name\":\"newR\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogTokenFactoryChanged\",\"inputs\":[{\"type\":\"address\",\"name\":\"oldTF\",\"indexed\":false},{\"type\":\"address\",\"name\":\"newTF\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"RegistryInstance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"TokenFactoryInstance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"approveProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_registry\"},{\"type\":\"address\",\"name\":\"_tokenFactory\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"initialized\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"migrateManager\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newManager\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"parseAddr\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[{\"type\":\"bytes\",\"name\":\"data\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"parsed\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setPlatformWallet\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newPlatformWallet\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setRegistry\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newRegistry\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setTokenFactory\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenFactory\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setTokenVesting\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenVesting\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingAddress\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingReferral\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"submitProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"hashIndex\"}]},{\"type\":\"function\",\"name\":\"submitReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_referral\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]}]",
}

// ManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ManagerMetaData.ABI instead.
var ManagerABI = ManagerMetaData.ABI

// Manager is an auto generated Go binding around an Ethereum contract.
type Manager struct {
	ManagerCaller     // Read-only binding to the contract
	ManagerTransactor // Write-only binding to the contract
	ManagerFilterer   // Log filterer for contract events
}

// ManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ManagerSession struct {
	Contract     *Manager          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ManagerCallerSession struct {
	Contract *ManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ManagerTransactorSession struct {
	Contract     *ManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ManagerRaw struct {
	Contract *Manager // Generic contract binding to access the raw methods on
}

// ManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ManagerCallerRaw struct {
	Contract *ManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ManagerTransactorRaw struct {
	Contract *ManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewManager creates a new instance of Manager, bound to a specific deployed contract.
func NewManager(address common.Address, backend bind.ContractBackend) (*Manager, error) {
	contract, err := bindManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Manager{ManagerCaller: ManagerCaller{contract: contract}, ManagerTransactor: ManagerTransactor{contract: contract}, ManagerFilterer: ManagerFilterer{contract: contract}}, nil
}

// NewManagerCaller creates a new read-only instance of Manager, bound to a specific deployed contract.
func NewManagerCaller(address common.Address, caller bind.ContractCaller) (*ManagerCaller, error) {
	contract, err := bindManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerCaller{contract: contract}, nil
}

// NewManagerTransactor creates a new write-only instance of Manager, bound to a specific deployed contract.
func NewManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ManagerTransactor, error) {
	contract, err := bindManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ManagerTransactor{contract: contract}, nil
}

// NewManagerFilterer creates a new log filterer instance of Manager, bound to a specific deployed contract.
func NewManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ManagerFilterer, error) {
	contract, err := bindManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ManagerFilterer{contract: contract}, nil
}

// bindManager binds a generic wrapper to an already deployed contract.
func bindManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.ManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.ManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Manager *ManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Manager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Manager *ManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Manager *ManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Manager.Contract.contract.Transact(opts, method, params...)
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() view returns(address)
func (_Manager *ManagerCaller) RegistryInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "RegistryInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() view returns(address)
func (_Manager *ManagerSession) RegistryInstance() (common.Address, error) {
	return _Manager.Contract.RegistryInstance(&_Manager.CallOpts)
}

// RegistryInstance is a free data retrieval call binding the contract method 0xb75ceb71.
//
// Solidity: function RegistryInstance() view returns(address)
func (_Manager *ManagerCallerSession) RegistryInstance() (common.Address, error) {
	return _Manager.Contract.RegistryInstance(&_Manager.CallOpts)
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() view returns(address)
func (_Manager *ManagerCaller) TokenFactoryInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "TokenFactoryInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() view returns(address)
func (_Manager *ManagerSession) TokenFactoryInstance() (common.Address, error) {
	return _Manager.Contract.TokenFactoryInstance(&_Manager.CallOpts)
}

// TokenFactoryInstance is a free data retrieval call binding the contract method 0xbbba6e92.
//
// Solidity: function TokenFactoryInstance() view returns(address)
func (_Manager *ManagerCallerSession) TokenFactoryInstance() (common.Address, error) {
	return _Manager.Contract.TokenFactoryInstance(&_Manager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Manager *ManagerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Manager *ManagerSession) Initialized() (bool, error) {
	return _Manager.Contract.Initialized(&_Manager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Manager *ManagerCallerSession) Initialized() (bool, error) {
	return _Manager.Contract.Initialized(&_Manager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manager *ManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manager *ManagerSession) Owner() (common.Address, error) {
	return _Manager.Contract.Owner(&_Manager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Manager *ManagerCallerSession) Owner() (common.Address, error) {
	return _Manager.Contract.Owner(&_Manager.CallOpts)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(bytes data) pure returns(address parsed)
func (_Manager *ManagerCaller) ParseAddr(opts *bind.CallOpts, data []byte) (common.Address, error) {
	var out []interface{}
	err := _Manager.contract.Call(opts, &out, "parseAddr", data)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(bytes data) pure returns(address parsed)
func (_Manager *ManagerSession) ParseAddr(data []byte) (common.Address, error) {
	return _Manager.Contract.ParseAddr(&_Manager.CallOpts, data)
}

// ParseAddr is a free data retrieval call binding the contract method 0xb718f83a.
//
// Solidity: function parseAddr(bytes data) pure returns(address parsed)
func (_Manager *ManagerCallerSession) ParseAddr(data []byte) (common.Address, error) {
	return _Manager.Contract.ParseAddr(&_Manager.CallOpts, data)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(bytes32 _hashIndex) returns()
func (_Manager *ManagerTransactor) ApproveProposal(opts *bind.TransactOpts, _hashIndex [32]byte) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "approveProposal", _hashIndex)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(bytes32 _hashIndex) returns()
func (_Manager *ManagerSession) ApproveProposal(_hashIndex [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.ApproveProposal(&_Manager.TransactOpts, _hashIndex)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0xf20b5b2c.
//
// Solidity: function approveProposal(bytes32 _hashIndex) returns()
func (_Manager *ManagerTransactorSession) ApproveProposal(_hashIndex [32]byte) (*types.Transaction, error) {
	return _Manager.Contract.ApproveProposal(&_Manager.TransactOpts, _hashIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _tokenFactory) returns()
func (_Manager *ManagerTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "initialize", _registry, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _tokenFactory) returns()
func (_Manager *ManagerSession) Initialize(_registry common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.Contract.Initialize(&_Manager.TransactOpts, _registry, _tokenFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registry, address _tokenFactory) returns()
func (_Manager *ManagerTransactorSession) Initialize(_registry common.Address, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.Contract.Initialize(&_Manager.TransactOpts, _registry, _tokenFactory)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(address _newManager) returns()
func (_Manager *ManagerTransactor) MigrateManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "migrateManager", _newManager)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(address _newManager) returns()
func (_Manager *ManagerSession) MigrateManager(_newManager common.Address) (*types.Transaction, error) {
	return _Manager.Contract.MigrateManager(&_Manager.TransactOpts, _newManager)
}

// MigrateManager is a paid mutator transaction binding the contract method 0xc7a8db19.
//
// Solidity: function migrateManager(address _newManager) returns()
func (_Manager *ManagerTransactorSession) MigrateManager(_newManager common.Address) (*types.Transaction, error) {
	return _Manager.Contract.MigrateManager(&_Manager.TransactOpts, _newManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manager *ManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manager *ManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manager.Contract.RenounceOwnership(&_Manager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Manager *ManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Manager.Contract.RenounceOwnership(&_Manager.TransactOpts)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_Manager *ManagerTransactor) SetPlatformWallet(opts *bind.TransactOpts, _newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setPlatformWallet", _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_Manager *ManagerSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetPlatformWallet(&_Manager.TransactOpts, _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_Manager *ManagerTransactorSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetPlatformWallet(&_Manager.TransactOpts, _newPlatformWallet)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _newRegistry) returns()
func (_Manager *ManagerTransactor) SetRegistry(opts *bind.TransactOpts, _newRegistry common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setRegistry", _newRegistry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _newRegistry) returns()
func (_Manager *ManagerSession) SetRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetRegistry(&_Manager.TransactOpts, _newRegistry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _newRegistry) returns()
func (_Manager *ManagerTransactorSession) SetRegistry(_newRegistry common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetRegistry(&_Manager.TransactOpts, _newRegistry)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _newTokenFactory) returns()
func (_Manager *ManagerTransactor) SetTokenFactory(opts *bind.TransactOpts, _newTokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setTokenFactory", _newTokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _newTokenFactory) returns()
func (_Manager *ManagerSession) SetTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenFactory(&_Manager.TransactOpts, _newTokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _newTokenFactory) returns()
func (_Manager *ManagerTransactorSession) SetTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenFactory(&_Manager.TransactOpts, _newTokenFactory)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_Manager *ManagerTransactor) SetTokenVesting(opts *bind.TransactOpts, _newTokenVesting common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setTokenVesting", _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_Manager *ManagerSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenVesting(&_Manager.TransactOpts, _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_Manager *ManagerTransactorSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetTokenVesting(&_Manager.TransactOpts, _newTokenVesting)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_Manager *ManagerTransactor) SetVestingAddress(opts *bind.TransactOpts, _token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setVestingAddress", _token, _vestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_Manager *ManagerSession) SetVestingAddress(_token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetVestingAddress(&_Manager.TransactOpts, _token, _vestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_Manager *ManagerTransactorSession) SetVestingAddress(_token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetVestingAddress(&_Manager.TransactOpts, _token, _vestingBeneficiary)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_Manager *ManagerTransactor) SetVestingReferral(opts *bind.TransactOpts, _token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "setVestingReferral", _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_Manager *ManagerSession) SetVestingReferral(_token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetVestingReferral(&_Manager.TransactOpts, _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_Manager *ManagerTransactorSession) SetVestingReferral(_token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _Manager.Contract.SetVestingReferral(&_Manager.TransactOpts, _token, _vestingReferral)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x6709a651.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Manager *ManagerTransactor) SubmitProposal(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "submitProposal", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x6709a651.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Manager *ManagerSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Manager.Contract.SubmitProposal(&_Manager.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x6709a651.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Manager *ManagerTransactorSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Manager.Contract.SubmitProposal(&_Manager.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage)
}

// SubmitReferral is a paid mutator transaction binding the contract method 0x1fd49f4b.
//
// Solidity: function submitReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Manager *ManagerTransactor) SubmitReferral(opts *bind.TransactOpts, _hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "submitReferral", _hashIndex, _referral, _referralPercentage)
}

// SubmitReferral is a paid mutator transaction binding the contract method 0x1fd49f4b.
//
// Solidity: function submitReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Manager *ManagerSession) SubmitReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Manager.Contract.SubmitReferral(&_Manager.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// SubmitReferral is a paid mutator transaction binding the contract method 0x1fd49f4b.
//
// Solidity: function submitReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Manager *ManagerTransactorSession) SubmitReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Manager.Contract.SubmitReferral(&_Manager.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manager *ManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Manager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manager *ManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manager.Contract.TransferOwnership(&_Manager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Manager *ManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Manager.Contract.TransferOwnership(&_Manager.TransactOpts, newOwner)
}

// ManagerLogManagerMigratedIterator is returned from FilterLogManagerMigrated and is used to iterate over the raw logs and unpacked data for LogManagerMigrated events raised by the Manager contract.
type ManagerLogManagerMigratedIterator struct {
	Event *ManagerLogManagerMigrated // Event containing the contract specifics and raw log

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
func (it *ManagerLogManagerMigratedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerLogManagerMigrated)
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
		it.Event = new(ManagerLogManagerMigrated)
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
func (it *ManagerLogManagerMigratedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerLogManagerMigratedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerLogManagerMigrated represents a LogManagerMigrated event raised by the Manager contract.
type ManagerLogManagerMigrated struct {
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogManagerMigrated is a free log retrieval operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: event LogManagerMigrated(address indexed newManager)
func (_Manager *ManagerFilterer) FilterLogManagerMigrated(opts *bind.FilterOpts, newManager []common.Address) (*ManagerLogManagerMigratedIterator, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "LogManagerMigrated", newManagerRule)
	if err != nil {
		return nil, err
	}
	return &ManagerLogManagerMigratedIterator{contract: _Manager.contract, event: "LogManagerMigrated", logs: logs, sub: sub}, nil
}

// WatchLogManagerMigrated is a free log subscription operation binding the contract event 0x829487e3aa7912bd48776fb898d596977df4debceb186d7575e27a94f70be174.
//
// Solidity: event LogManagerMigrated(address indexed newManager)
func (_Manager *ManagerFilterer) WatchLogManagerMigrated(opts *bind.WatchOpts, sink chan<- *ManagerLogManagerMigrated, newManager []common.Address) (event.Subscription, error) {

	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "LogManagerMigrated", newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerLogManagerMigrated)
				if err := _Manager.contract.UnpackLog(event, "LogManagerMigrated", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseLogManagerMigrated(log types.Log) (*ManagerLogManagerMigrated, error) {
	event := new(ManagerLogManagerMigrated)
	if err := _Manager.contract.UnpackLog(event, "LogManagerMigrated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerLogRegistryChangedIterator is returned from FilterLogRegistryChanged and is used to iterate over the raw logs and unpacked data for LogRegistryChanged events raised by the Manager contract.
type ManagerLogRegistryChangedIterator struct {
	Event *ManagerLogRegistryChanged // Event containing the contract specifics and raw log

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
func (it *ManagerLogRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerLogRegistryChanged)
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
		it.Event = new(ManagerLogRegistryChanged)
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
func (it *ManagerLogRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerLogRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerLogRegistryChanged represents a LogRegistryChanged event raised by the Manager contract.
type ManagerLogRegistryChanged struct {
	OldR common.Address
	NewR common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRegistryChanged is a free log retrieval operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: event LogRegistryChanged(address oldR, address newR)
func (_Manager *ManagerFilterer) FilterLogRegistryChanged(opts *bind.FilterOpts) (*ManagerLogRegistryChangedIterator, error) {

	logs, sub, err := _Manager.contract.FilterLogs(opts, "LogRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &ManagerLogRegistryChangedIterator{contract: _Manager.contract, event: "LogRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchLogRegistryChanged is a free log subscription operation binding the contract event 0xec10a16af385903532506d1567380fd8d93d880e8ed4bcacc42b0848e781406a.
//
// Solidity: event LogRegistryChanged(address oldR, address newR)
func (_Manager *ManagerFilterer) WatchLogRegistryChanged(opts *bind.WatchOpts, sink chan<- *ManagerLogRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _Manager.contract.WatchLogs(opts, "LogRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerLogRegistryChanged)
				if err := _Manager.contract.UnpackLog(event, "LogRegistryChanged", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseLogRegistryChanged(log types.Log) (*ManagerLogRegistryChanged, error) {
	event := new(ManagerLogRegistryChanged)
	if err := _Manager.contract.UnpackLog(event, "LogRegistryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerLogTokenFactoryChangedIterator is returned from FilterLogTokenFactoryChanged and is used to iterate over the raw logs and unpacked data for LogTokenFactoryChanged events raised by the Manager contract.
type ManagerLogTokenFactoryChangedIterator struct {
	Event *ManagerLogTokenFactoryChanged // Event containing the contract specifics and raw log

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
func (it *ManagerLogTokenFactoryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerLogTokenFactoryChanged)
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
		it.Event = new(ManagerLogTokenFactoryChanged)
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
func (it *ManagerLogTokenFactoryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerLogTokenFactoryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerLogTokenFactoryChanged represents a LogTokenFactoryChanged event raised by the Manager contract.
type ManagerLogTokenFactoryChanged struct {
	OldTF common.Address
	NewTF common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogTokenFactoryChanged is a free log retrieval operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: event LogTokenFactoryChanged(address oldTF, address newTF)
func (_Manager *ManagerFilterer) FilterLogTokenFactoryChanged(opts *bind.FilterOpts) (*ManagerLogTokenFactoryChangedIterator, error) {

	logs, sub, err := _Manager.contract.FilterLogs(opts, "LogTokenFactoryChanged")
	if err != nil {
		return nil, err
	}
	return &ManagerLogTokenFactoryChangedIterator{contract: _Manager.contract, event: "LogTokenFactoryChanged", logs: logs, sub: sub}, nil
}

// WatchLogTokenFactoryChanged is a free log subscription operation binding the contract event 0x4b517b4fb873b53bdb7db1289eb8ec3d2c92288dda09e8e68b3db0efade30b7e.
//
// Solidity: event LogTokenFactoryChanged(address oldTF, address newTF)
func (_Manager *ManagerFilterer) WatchLogTokenFactoryChanged(opts *bind.WatchOpts, sink chan<- *ManagerLogTokenFactoryChanged) (event.Subscription, error) {

	logs, sub, err := _Manager.contract.WatchLogs(opts, "LogTokenFactoryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerLogTokenFactoryChanged)
				if err := _Manager.contract.UnpackLog(event, "LogTokenFactoryChanged", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseLogTokenFactoryChanged(log types.Log) (*ManagerLogTokenFactoryChanged, error) {
	event := new(ManagerLogTokenFactoryChanged)
	if err := _Manager.contract.UnpackLog(event, "LogTokenFactoryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Manager contract.
type ManagerOwnershipTransferredIterator struct {
	Event *ManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ManagerOwnershipTransferred)
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
		it.Event = new(ManagerOwnershipTransferred)
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
func (it *ManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ManagerOwnershipTransferred represents a OwnershipTransferred event raised by the Manager contract.
type ManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manager *ManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ManagerOwnershipTransferredIterator{contract: _Manager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Manager *ManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Manager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ManagerOwnershipTransferred)
				if err := _Manager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Manager *ManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ManagerOwnershipTransferred, error) {
	event := new(ManagerOwnershipTransferred)
	if err := _Manager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
