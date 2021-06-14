// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenFactory

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
const ContractsABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogTokenCreated\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"string\",\"name\":\"symbol\",\"indexed\":false},{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"TokenVestingInstance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"calculateProportions\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[{\"type\":\"uint256[4]\",\"name\":\"proportions\"}]},{\"type\":\"function\",\"name\":\"createToken\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"},{\"type\":\"address\",\"name\":\"_referral\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"token\"}]},{\"type\":\"function\",\"name\":\"getTokenVesting\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_tokenVesting\"},{\"type\":\"address\",\"name\":\"_rollWallet\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"migrateTokenFactory\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenFactory\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"rollWallet\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"scalePercentages\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[{\"type\":\"uint8\",\"name\":\"p0\"},{\"type\":\"uint8\",\"name\":\"p1\"},{\"type\":\"uint8\",\"name\":\"p2\"}],\"outputs\":[{\"type\":\"uint256[3]\",\"name\":\"proportions\"}]},{\"type\":\"function\",\"name\":\"setPlatformWallet\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newPlatformWallet\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setTokenVesting\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenVesting\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"validateTokenVestingOwner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"a1\"},{\"type\":\"address\",\"name\":\"a2\"}],\"outputs\":[]}]"

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

// TokenVestingInstance is a free data retrieval call binding the contract method 0x151ad610.
//
// Solidity: function TokenVestingInstance() view returns(address)
func (_Contracts *ContractsCaller) TokenVestingInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "TokenVestingInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVestingInstance is a free data retrieval call binding the contract method 0x151ad610.
//
// Solidity: function TokenVestingInstance() view returns(address)
func (_Contracts *ContractsSession) TokenVestingInstance() (common.Address, error) {
	return _Contracts.Contract.TokenVestingInstance(&_Contracts.CallOpts)
}

// TokenVestingInstance is a free data retrieval call binding the contract method 0x151ad610.
//
// Solidity: function TokenVestingInstance() view returns(address)
func (_Contracts *ContractsCallerSession) TokenVestingInstance() (common.Address, error) {
	return _Contracts.Contract.TokenVestingInstance(&_Contracts.CallOpts)
}

// CalculateProportions is a free data retrieval call binding the contract method 0xaaef10e7.
//
// Solidity: function calculateProportions(uint256 _totalSupply, uint8 _initialPercentage, uint8 _initialPlatformPercentage, uint8 _referralPercentage) pure returns(uint256[4] proportions)
func (_Contracts *ContractsCaller) CalculateProportions(opts *bind.CallOpts, _totalSupply *big.Int, _initialPercentage uint8, _initialPlatformPercentage uint8, _referralPercentage uint8) ([4]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "calculateProportions", _totalSupply, _initialPercentage, _initialPlatformPercentage, _referralPercentage)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CalculateProportions is a free data retrieval call binding the contract method 0xaaef10e7.
//
// Solidity: function calculateProportions(uint256 _totalSupply, uint8 _initialPercentage, uint8 _initialPlatformPercentage, uint8 _referralPercentage) pure returns(uint256[4] proportions)
func (_Contracts *ContractsSession) CalculateProportions(_totalSupply *big.Int, _initialPercentage uint8, _initialPlatformPercentage uint8, _referralPercentage uint8) ([4]*big.Int, error) {
	return _Contracts.Contract.CalculateProportions(&_Contracts.CallOpts, _totalSupply, _initialPercentage, _initialPlatformPercentage, _referralPercentage)
}

// CalculateProportions is a free data retrieval call binding the contract method 0xaaef10e7.
//
// Solidity: function calculateProportions(uint256 _totalSupply, uint8 _initialPercentage, uint8 _initialPlatformPercentage, uint8 _referralPercentage) pure returns(uint256[4] proportions)
func (_Contracts *ContractsCallerSession) CalculateProportions(_totalSupply *big.Int, _initialPercentage uint8, _initialPlatformPercentage uint8, _referralPercentage uint8) ([4]*big.Int, error) {
	return _Contracts.Contract.CalculateProportions(&_Contracts.CallOpts, _totalSupply, _initialPercentage, _initialPlatformPercentage, _referralPercentage)
}

// GetTokenVesting is a free data retrieval call binding the contract method 0x72172364.
//
// Solidity: function getTokenVesting() view returns(address)
func (_Contracts *ContractsCaller) GetTokenVesting(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getTokenVesting")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenVesting is a free data retrieval call binding the contract method 0x72172364.
//
// Solidity: function getTokenVesting() view returns(address)
func (_Contracts *ContractsSession) GetTokenVesting() (common.Address, error) {
	return _Contracts.Contract.GetTokenVesting(&_Contracts.CallOpts)
}

// GetTokenVesting is a free data retrieval call binding the contract method 0x72172364.
//
// Solidity: function getTokenVesting() view returns(address)
func (_Contracts *ContractsCallerSession) GetTokenVesting() (common.Address, error) {
	return _Contracts.Contract.GetTokenVesting(&_Contracts.CallOpts)
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

// RollWallet is a free data retrieval call binding the contract method 0x1494b0c9.
//
// Solidity: function rollWallet() view returns(address)
func (_Contracts *ContractsCaller) RollWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "rollWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollWallet is a free data retrieval call binding the contract method 0x1494b0c9.
//
// Solidity: function rollWallet() view returns(address)
func (_Contracts *ContractsSession) RollWallet() (common.Address, error) {
	return _Contracts.Contract.RollWallet(&_Contracts.CallOpts)
}

// RollWallet is a free data retrieval call binding the contract method 0x1494b0c9.
//
// Solidity: function rollWallet() view returns(address)
func (_Contracts *ContractsCallerSession) RollWallet() (common.Address, error) {
	return _Contracts.Contract.RollWallet(&_Contracts.CallOpts)
}

// ScalePercentages is a free data retrieval call binding the contract method 0x630c8b51.
//
// Solidity: function scalePercentages(uint8 p0, uint8 p1, uint8 p2) pure returns(uint256[3] proportions)
func (_Contracts *ContractsCaller) ScalePercentages(opts *bind.CallOpts, p0 uint8, p1 uint8, p2 uint8) ([3]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "scalePercentages", p0, p1, p2)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// ScalePercentages is a free data retrieval call binding the contract method 0x630c8b51.
//
// Solidity: function scalePercentages(uint8 p0, uint8 p1, uint8 p2) pure returns(uint256[3] proportions)
func (_Contracts *ContractsSession) ScalePercentages(p0 uint8, p1 uint8, p2 uint8) ([3]*big.Int, error) {
	return _Contracts.Contract.ScalePercentages(&_Contracts.CallOpts, p0, p1, p2)
}

// ScalePercentages is a free data retrieval call binding the contract method 0x630c8b51.
//
// Solidity: function scalePercentages(uint8 p0, uint8 p1, uint8 p2) pure returns(uint256[3] proportions)
func (_Contracts *ContractsCallerSession) ScalePercentages(p0 uint8, p1 uint8, p2 uint8) ([3]*big.Int, error) {
	return _Contracts.Contract.ScalePercentages(&_Contracts.CallOpts, p0, p1, p2)
}

// ValidateTokenVestingOwner is a free data retrieval call binding the contract method 0x84e7d7ac.
//
// Solidity: function validateTokenVestingOwner(address a1, address a2) view returns()
func (_Contracts *ContractsCaller) ValidateTokenVestingOwner(opts *bind.CallOpts, a1 common.Address, a2 common.Address) error {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "validateTokenVestingOwner", a1, a2)

	if err != nil {
		return err
	}

	return err

}

// ValidateTokenVestingOwner is a free data retrieval call binding the contract method 0x84e7d7ac.
//
// Solidity: function validateTokenVestingOwner(address a1, address a2) view returns()
func (_Contracts *ContractsSession) ValidateTokenVestingOwner(a1 common.Address, a2 common.Address) error {
	return _Contracts.Contract.ValidateTokenVestingOwner(&_Contracts.CallOpts, a1, a2)
}

// ValidateTokenVestingOwner is a free data retrieval call binding the contract method 0x84e7d7ac.
//
// Solidity: function validateTokenVestingOwner(address a1, address a2) view returns()
func (_Contracts *ContractsCallerSession) ValidateTokenVestingOwner(a1 common.Address, a2 common.Address) error {
	return _Contracts.Contract.ValidateTokenVestingOwner(&_Contracts.CallOpts, a1, a2)
}

// CreateToken is a paid mutator transaction binding the contract method 0x44ae5e45.
//
// Solidity: function createToken(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage, address _referral, uint8 _referralPercentage) returns(address token)
func (_Contracts *ContractsTransactor) CreateToken(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createToken", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage, _referral, _referralPercentage)
}

// CreateToken is a paid mutator transaction binding the contract method 0x44ae5e45.
//
// Solidity: function createToken(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage, address _referral, uint8 _referralPercentage) returns(address token)
func (_Contracts *ContractsSession) CreateToken(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.CreateToken(&_Contracts.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage, _referral, _referralPercentage)
}

// CreateToken is a paid mutator transaction binding the contract method 0x44ae5e45.
//
// Solidity: function createToken(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage, address _referral, uint8 _referralPercentage) returns(address token)
func (_Contracts *ContractsTransactorSession) CreateToken(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.CreateToken(&_Contracts.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage, _referral, _referralPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tokenVesting, address _rollWallet) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, _tokenVesting common.Address, _rollWallet common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", _tokenVesting, _rollWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tokenVesting, address _rollWallet) returns()
func (_Contracts *ContractsSession) Initialize(_tokenVesting common.Address, _rollWallet common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, _tokenVesting, _rollWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tokenVesting, address _rollWallet) returns()
func (_Contracts *ContractsTransactorSession) Initialize(_tokenVesting common.Address, _rollWallet common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, _tokenVesting, _rollWallet)
}

// MigrateTokenFactory is a paid mutator transaction binding the contract method 0x57796797.
//
// Solidity: function migrateTokenFactory(address _newTokenFactory) returns()
func (_Contracts *ContractsTransactor) MigrateTokenFactory(opts *bind.TransactOpts, _newTokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "migrateTokenFactory", _newTokenFactory)
}

// MigrateTokenFactory is a paid mutator transaction binding the contract method 0x57796797.
//
// Solidity: function migrateTokenFactory(address _newTokenFactory) returns()
func (_Contracts *ContractsSession) MigrateTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.MigrateTokenFactory(&_Contracts.TransactOpts, _newTokenFactory)
}

// MigrateTokenFactory is a paid mutator transaction binding the contract method 0x57796797.
//
// Solidity: function migrateTokenFactory(address _newTokenFactory) returns()
func (_Contracts *ContractsTransactorSession) MigrateTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.MigrateTokenFactory(&_Contracts.TransactOpts, _newTokenFactory)
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

// ContractsLogTokenCreatedIterator is returned from FilterLogTokenCreated and is used to iterate over the raw logs and unpacked data for LogTokenCreated events raised by the Contracts contract.
type ContractsLogTokenCreatedIterator struct {
	Event *ContractsLogTokenCreated // Event containing the contract specifics and raw log

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
func (it *ContractsLogTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogTokenCreated)
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
		it.Event = new(ContractsLogTokenCreated)
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
func (it *ContractsLogTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogTokenCreated represents a LogTokenCreated event raised by the Contracts contract.
type ContractsLogTokenCreated struct {
	Name               string
	Symbol             string
	Token              common.Address
	VestingBeneficiary common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogTokenCreated is a free log retrieval operation binding the contract event 0xa3813ec7fba2e70e20df54e4f52ec626b028802742c8757d5be33c4a3742fb45.
//
// Solidity: event LogTokenCreated(string name, string symbol, address indexed token, address vestingBeneficiary)
func (_Contracts *ContractsFilterer) FilterLogTokenCreated(opts *bind.FilterOpts, token []common.Address) (*ContractsLogTokenCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogTokenCreatedIterator{contract: _Contracts.contract, event: "LogTokenCreated", logs: logs, sub: sub}, nil
}

// WatchLogTokenCreated is a free log subscription operation binding the contract event 0xa3813ec7fba2e70e20df54e4f52ec626b028802742c8757d5be33c4a3742fb45.
//
// Solidity: event LogTokenCreated(string name, string symbol, address indexed token, address vestingBeneficiary)
func (_Contracts *ContractsFilterer) WatchLogTokenCreated(opts *bind.WatchOpts, sink chan<- *ContractsLogTokenCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogTokenCreated)
				if err := _Contracts.contract.UnpackLog(event, "LogTokenCreated", log); err != nil {
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

// ParseLogTokenCreated is a log parse operation binding the contract event 0xa3813ec7fba2e70e20df54e4f52ec626b028802742c8757d5be33c4a3742fb45.
//
// Solidity: event LogTokenCreated(string name, string symbol, address indexed token, address vestingBeneficiary)
func (_Contracts *ContractsFilterer) ParseLogTokenCreated(log types.Log) (*ContractsLogTokenCreated, error) {
	event := new(ContractsLogTokenCreated)
	if err := _Contracts.contract.UnpackLog(event, "LogTokenCreated", log); err != nil {
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
