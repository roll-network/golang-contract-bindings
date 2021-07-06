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

// TokenFactoryABI is the input ABI used to generate the binding from.
const TokenFactoryABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogTokenCreated\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"string\",\"name\":\"symbol\",\"indexed\":false},{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"TokenVestingInstance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"calculateProportions\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[{\"type\":\"uint256[4]\",\"name\":\"proportions\"}]},{\"type\":\"function\",\"name\":\"createToken\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"},{\"type\":\"address\",\"name\":\"_referral\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"token\"}]},{\"type\":\"function\",\"name\":\"getTokenVesting\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_tokenVesting\"},{\"type\":\"address\",\"name\":\"_rollWallet\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"migrateTokenFactory\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenFactory\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"rollWallet\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"scalePercentages\",\"constant\":true,\"stateMutability\":\"pure\",\"payable\":false,\"inputs\":[{\"type\":\"uint8\",\"name\":\"p0\"},{\"type\":\"uint8\",\"name\":\"p1\"},{\"type\":\"uint8\",\"name\":\"p2\"}],\"outputs\":[{\"type\":\"uint256[3]\",\"name\":\"proportions\"}]},{\"type\":\"function\",\"name\":\"setPlatformWallet\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newPlatformWallet\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setTokenVesting\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_newTokenVesting\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingAddress\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_newVestingBeneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingReferral\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"validateTokenVestingOwner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"a1\"},{\"type\":\"address\",\"name\":\"a2\"}],\"outputs\":[]}]"

// TokenFactory is an auto generated Go binding around an Ethereum contract.
type TokenFactory struct {
	TokenFactoryCaller     // Read-only binding to the contract
	TokenFactoryTransactor // Write-only binding to the contract
	TokenFactoryFilterer   // Log filterer for contract events
}

// TokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFactorySession struct {
	Contract     *TokenFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFactoryCallerSession struct {
	Contract *TokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFactoryTransactorSession struct {
	Contract     *TokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFactoryRaw struct {
	Contract *TokenFactory // Generic contract binding to access the raw methods on
}

// TokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFactoryCallerRaw struct {
	Contract *TokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFactoryTransactorRaw struct {
	Contract *TokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFactory creates a new instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactory(address common.Address, backend bind.ContractBackend) (*TokenFactory, error) {
	contract, err := bindTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// NewTokenFactoryCaller creates a new read-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenFactoryCaller, error) {
	contract, err := bindTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryCaller{contract: contract}, nil
}

// NewTokenFactoryTransactor creates a new write-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFactoryTransactor, error) {
	contract, err := bindTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTransactor{contract: contract}, nil
}

// NewTokenFactoryFilterer creates a new log filterer instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFactoryFilterer, error) {
	contract, err := bindTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryFilterer{contract: contract}, nil
}

// bindTokenFactory binds a generic wrapper to an already deployed contract.
func bindTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.TokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transact(opts, method, params...)
}

// TokenVestingInstance is a free data retrieval call binding the contract method 0x151ad610.
//
// Solidity: function TokenVestingInstance() view returns(address)
func (_TokenFactory *TokenFactoryCaller) TokenVestingInstance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "TokenVestingInstance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenVestingInstance is a free data retrieval call binding the contract method 0x151ad610.
//
// Solidity: function TokenVestingInstance() view returns(address)
func (_TokenFactory *TokenFactorySession) TokenVestingInstance() (common.Address, error) {
	return _TokenFactory.Contract.TokenVestingInstance(&_TokenFactory.CallOpts)
}

// TokenVestingInstance is a free data retrieval call binding the contract method 0x151ad610.
//
// Solidity: function TokenVestingInstance() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) TokenVestingInstance() (common.Address, error) {
	return _TokenFactory.Contract.TokenVestingInstance(&_TokenFactory.CallOpts)
}

// CalculateProportions is a free data retrieval call binding the contract method 0xaaef10e7.
//
// Solidity: function calculateProportions(uint256 _totalSupply, uint8 _initialPercentage, uint8 _initialPlatformPercentage, uint8 _referralPercentage) pure returns(uint256[4] proportions)
func (_TokenFactory *TokenFactoryCaller) CalculateProportions(opts *bind.CallOpts, _totalSupply *big.Int, _initialPercentage uint8, _initialPlatformPercentage uint8, _referralPercentage uint8) ([4]*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "calculateProportions", _totalSupply, _initialPercentage, _initialPlatformPercentage, _referralPercentage)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CalculateProportions is a free data retrieval call binding the contract method 0xaaef10e7.
//
// Solidity: function calculateProportions(uint256 _totalSupply, uint8 _initialPercentage, uint8 _initialPlatformPercentage, uint8 _referralPercentage) pure returns(uint256[4] proportions)
func (_TokenFactory *TokenFactorySession) CalculateProportions(_totalSupply *big.Int, _initialPercentage uint8, _initialPlatformPercentage uint8, _referralPercentage uint8) ([4]*big.Int, error) {
	return _TokenFactory.Contract.CalculateProportions(&_TokenFactory.CallOpts, _totalSupply, _initialPercentage, _initialPlatformPercentage, _referralPercentage)
}

// CalculateProportions is a free data retrieval call binding the contract method 0xaaef10e7.
//
// Solidity: function calculateProportions(uint256 _totalSupply, uint8 _initialPercentage, uint8 _initialPlatformPercentage, uint8 _referralPercentage) pure returns(uint256[4] proportions)
func (_TokenFactory *TokenFactoryCallerSession) CalculateProportions(_totalSupply *big.Int, _initialPercentage uint8, _initialPlatformPercentage uint8, _referralPercentage uint8) ([4]*big.Int, error) {
	return _TokenFactory.Contract.CalculateProportions(&_TokenFactory.CallOpts, _totalSupply, _initialPercentage, _initialPlatformPercentage, _referralPercentage)
}

// GetTokenVesting is a free data retrieval call binding the contract method 0x72172364.
//
// Solidity: function getTokenVesting() view returns(address)
func (_TokenFactory *TokenFactoryCaller) GetTokenVesting(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "getTokenVesting")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenVesting is a free data retrieval call binding the contract method 0x72172364.
//
// Solidity: function getTokenVesting() view returns(address)
func (_TokenFactory *TokenFactorySession) GetTokenVesting() (common.Address, error) {
	return _TokenFactory.Contract.GetTokenVesting(&_TokenFactory.CallOpts)
}

// GetTokenVesting is a free data retrieval call binding the contract method 0x72172364.
//
// Solidity: function getTokenVesting() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) GetTokenVesting() (common.Address, error) {
	return _TokenFactory.Contract.GetTokenVesting(&_TokenFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactorySession) Owner() (common.Address, error) {
	return _TokenFactory.Contract.Owner(&_TokenFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) Owner() (common.Address, error) {
	return _TokenFactory.Contract.Owner(&_TokenFactory.CallOpts)
}

// RollWallet is a free data retrieval call binding the contract method 0x1494b0c9.
//
// Solidity: function rollWallet() view returns(address)
func (_TokenFactory *TokenFactoryCaller) RollWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "rollWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollWallet is a free data retrieval call binding the contract method 0x1494b0c9.
//
// Solidity: function rollWallet() view returns(address)
func (_TokenFactory *TokenFactorySession) RollWallet() (common.Address, error) {
	return _TokenFactory.Contract.RollWallet(&_TokenFactory.CallOpts)
}

// RollWallet is a free data retrieval call binding the contract method 0x1494b0c9.
//
// Solidity: function rollWallet() view returns(address)
func (_TokenFactory *TokenFactoryCallerSession) RollWallet() (common.Address, error) {
	return _TokenFactory.Contract.RollWallet(&_TokenFactory.CallOpts)
}

// ScalePercentages is a free data retrieval call binding the contract method 0x630c8b51.
//
// Solidity: function scalePercentages(uint8 p0, uint8 p1, uint8 p2) pure returns(uint256[3] proportions)
func (_TokenFactory *TokenFactoryCaller) ScalePercentages(opts *bind.CallOpts, p0 uint8, p1 uint8, p2 uint8) ([3]*big.Int, error) {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "scalePercentages", p0, p1, p2)

	if err != nil {
		return *new([3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3]*big.Int)).(*[3]*big.Int)

	return out0, err

}

// ScalePercentages is a free data retrieval call binding the contract method 0x630c8b51.
//
// Solidity: function scalePercentages(uint8 p0, uint8 p1, uint8 p2) pure returns(uint256[3] proportions)
func (_TokenFactory *TokenFactorySession) ScalePercentages(p0 uint8, p1 uint8, p2 uint8) ([3]*big.Int, error) {
	return _TokenFactory.Contract.ScalePercentages(&_TokenFactory.CallOpts, p0, p1, p2)
}

// ScalePercentages is a free data retrieval call binding the contract method 0x630c8b51.
//
// Solidity: function scalePercentages(uint8 p0, uint8 p1, uint8 p2) pure returns(uint256[3] proportions)
func (_TokenFactory *TokenFactoryCallerSession) ScalePercentages(p0 uint8, p1 uint8, p2 uint8) ([3]*big.Int, error) {
	return _TokenFactory.Contract.ScalePercentages(&_TokenFactory.CallOpts, p0, p1, p2)
}

// ValidateTokenVestingOwner is a free data retrieval call binding the contract method 0x84e7d7ac.
//
// Solidity: function validateTokenVestingOwner(address a1, address a2) view returns()
func (_TokenFactory *TokenFactoryCaller) ValidateTokenVestingOwner(opts *bind.CallOpts, a1 common.Address, a2 common.Address) error {
	var out []interface{}
	err := _TokenFactory.contract.Call(opts, &out, "validateTokenVestingOwner", a1, a2)

	if err != nil {
		return err
	}

	return err

}

// ValidateTokenVestingOwner is a free data retrieval call binding the contract method 0x84e7d7ac.
//
// Solidity: function validateTokenVestingOwner(address a1, address a2) view returns()
func (_TokenFactory *TokenFactorySession) ValidateTokenVestingOwner(a1 common.Address, a2 common.Address) error {
	return _TokenFactory.Contract.ValidateTokenVestingOwner(&_TokenFactory.CallOpts, a1, a2)
}

// ValidateTokenVestingOwner is a free data retrieval call binding the contract method 0x84e7d7ac.
//
// Solidity: function validateTokenVestingOwner(address a1, address a2) view returns()
func (_TokenFactory *TokenFactoryCallerSession) ValidateTokenVestingOwner(a1 common.Address, a2 common.Address) error {
	return _TokenFactory.Contract.ValidateTokenVestingOwner(&_TokenFactory.CallOpts, a1, a2)
}

// CreateToken is a paid mutator transaction binding the contract method 0x44ae5e45.
//
// Solidity: function createToken(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage, address _referral, uint8 _referralPercentage) returns(address token)
func (_TokenFactory *TokenFactoryTransactor) CreateToken(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createToken", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage, _referral, _referralPercentage)
}

// CreateToken is a paid mutator transaction binding the contract method 0x44ae5e45.
//
// Solidity: function createToken(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage, address _referral, uint8 _referralPercentage) returns(address token)
func (_TokenFactory *TokenFactorySession) CreateToken(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage, _referral, _referralPercentage)
}

// CreateToken is a paid mutator transaction binding the contract method 0x44ae5e45.
//
// Solidity: function createToken(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, uint8 _initialPlatformPercentage, address _referral, uint8 _referralPercentage) returns(address token)
func (_TokenFactory *TokenFactoryTransactorSession) CreateToken(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _initialPlatformPercentage uint8, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _initialPlatformPercentage, _referral, _referralPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tokenVesting, address _rollWallet) returns()
func (_TokenFactory *TokenFactoryTransactor) Initialize(opts *bind.TransactOpts, _tokenVesting common.Address, _rollWallet common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "initialize", _tokenVesting, _rollWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tokenVesting, address _rollWallet) returns()
func (_TokenFactory *TokenFactorySession) Initialize(_tokenVesting common.Address, _rollWallet common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.Initialize(&_TokenFactory.TransactOpts, _tokenVesting, _rollWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _tokenVesting, address _rollWallet) returns()
func (_TokenFactory *TokenFactoryTransactorSession) Initialize(_tokenVesting common.Address, _rollWallet common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.Initialize(&_TokenFactory.TransactOpts, _tokenVesting, _rollWallet)
}

// MigrateTokenFactory is a paid mutator transaction binding the contract method 0x57796797.
//
// Solidity: function migrateTokenFactory(address _newTokenFactory) returns()
func (_TokenFactory *TokenFactoryTransactor) MigrateTokenFactory(opts *bind.TransactOpts, _newTokenFactory common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "migrateTokenFactory", _newTokenFactory)
}

// MigrateTokenFactory is a paid mutator transaction binding the contract method 0x57796797.
//
// Solidity: function migrateTokenFactory(address _newTokenFactory) returns()
func (_TokenFactory *TokenFactorySession) MigrateTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.MigrateTokenFactory(&_TokenFactory.TransactOpts, _newTokenFactory)
}

// MigrateTokenFactory is a paid mutator transaction binding the contract method 0x57796797.
//
// Solidity: function migrateTokenFactory(address _newTokenFactory) returns()
func (_TokenFactory *TokenFactoryTransactorSession) MigrateTokenFactory(_newTokenFactory common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.MigrateTokenFactory(&_TokenFactory.TransactOpts, _newTokenFactory)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFactory.Contract.RenounceOwnership(&_TokenFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFactory *TokenFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFactory.Contract.RenounceOwnership(&_TokenFactory.TransactOpts)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_TokenFactory *TokenFactoryTransactor) SetPlatformWallet(opts *bind.TransactOpts, _newPlatformWallet common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setPlatformWallet", _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_TokenFactory *TokenFactorySession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetPlatformWallet(&_TokenFactory.TransactOpts, _newPlatformWallet)
}

// SetPlatformWallet is a paid mutator transaction binding the contract method 0x8831e9cf.
//
// Solidity: function setPlatformWallet(address _newPlatformWallet) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetPlatformWallet(_newPlatformWallet common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetPlatformWallet(&_TokenFactory.TransactOpts, _newPlatformWallet)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_TokenFactory *TokenFactoryTransactor) SetTokenVesting(opts *bind.TransactOpts, _newTokenVesting common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setTokenVesting", _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_TokenFactory *TokenFactorySession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetTokenVesting(&_TokenFactory.TransactOpts, _newTokenVesting)
}

// SetTokenVesting is a paid mutator transaction binding the contract method 0x43e34696.
//
// Solidity: function setTokenVesting(address _newTokenVesting) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetTokenVesting(_newTokenVesting common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetTokenVesting(&_TokenFactory.TransactOpts, _newTokenVesting)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0xd2e836f8.
//
// Solidity: function setVestingAddress(address _vestingBeneficiary, address _token, address _newVestingBeneficiary) returns()
func (_TokenFactory *TokenFactoryTransactor) SetVestingAddress(opts *bind.TransactOpts, _vestingBeneficiary common.Address, _token common.Address, _newVestingBeneficiary common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setVestingAddress", _vestingBeneficiary, _token, _newVestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0xd2e836f8.
//
// Solidity: function setVestingAddress(address _vestingBeneficiary, address _token, address _newVestingBeneficiary) returns()
func (_TokenFactory *TokenFactorySession) SetVestingAddress(_vestingBeneficiary common.Address, _token common.Address, _newVestingBeneficiary common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetVestingAddress(&_TokenFactory.TransactOpts, _vestingBeneficiary, _token, _newVestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0xd2e836f8.
//
// Solidity: function setVestingAddress(address _vestingBeneficiary, address _token, address _newVestingBeneficiary) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetVestingAddress(_vestingBeneficiary common.Address, _token common.Address, _newVestingBeneficiary common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetVestingAddress(&_TokenFactory.TransactOpts, _vestingBeneficiary, _token, _newVestingBeneficiary)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0x0c0cf20b.
//
// Solidity: function setVestingReferral(address _vestingBeneficiary, address _token, address _vestingReferral) returns()
func (_TokenFactory *TokenFactoryTransactor) SetVestingReferral(opts *bind.TransactOpts, _vestingBeneficiary common.Address, _token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "setVestingReferral", _vestingBeneficiary, _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0x0c0cf20b.
//
// Solidity: function setVestingReferral(address _vestingBeneficiary, address _token, address _vestingReferral) returns()
func (_TokenFactory *TokenFactorySession) SetVestingReferral(_vestingBeneficiary common.Address, _token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetVestingReferral(&_TokenFactory.TransactOpts, _vestingBeneficiary, _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0x0c0cf20b.
//
// Solidity: function setVestingReferral(address _vestingBeneficiary, address _token, address _vestingReferral) returns()
func (_TokenFactory *TokenFactoryTransactorSession) SetVestingReferral(_vestingBeneficiary common.Address, _token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.SetVestingReferral(&_TokenFactory.TransactOpts, _vestingBeneficiary, _token, _vestingReferral)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.TransferOwnership(&_TokenFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFactory *TokenFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.TransferOwnership(&_TokenFactory.TransactOpts, newOwner)
}

// TokenFactoryLogTokenCreatedIterator is returned from FilterLogTokenCreated and is used to iterate over the raw logs and unpacked data for LogTokenCreated events raised by the TokenFactory contract.
type TokenFactoryLogTokenCreatedIterator struct {
	Event *TokenFactoryLogTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryLogTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryLogTokenCreated)
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
		it.Event = new(TokenFactoryLogTokenCreated)
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
func (it *TokenFactoryLogTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryLogTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryLogTokenCreated represents a LogTokenCreated event raised by the TokenFactory contract.
type TokenFactoryLogTokenCreated struct {
	Name               string
	Symbol             string
	Token              common.Address
	VestingBeneficiary common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogTokenCreated is a free log retrieval operation binding the contract event 0xa3813ec7fba2e70e20df54e4f52ec626b028802742c8757d5be33c4a3742fb45.
//
// Solidity: event LogTokenCreated(string name, string symbol, address indexed token, address vestingBeneficiary)
func (_TokenFactory *TokenFactoryFilterer) FilterLogTokenCreated(opts *bind.FilterOpts, token []common.Address) (*TokenFactoryLogTokenCreatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "LogTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryLogTokenCreatedIterator{contract: _TokenFactory.contract, event: "LogTokenCreated", logs: logs, sub: sub}, nil
}

// WatchLogTokenCreated is a free log subscription operation binding the contract event 0xa3813ec7fba2e70e20df54e4f52ec626b028802742c8757d5be33c4a3742fb45.
//
// Solidity: event LogTokenCreated(string name, string symbol, address indexed token, address vestingBeneficiary)
func (_TokenFactory *TokenFactoryFilterer) WatchLogTokenCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryLogTokenCreated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "LogTokenCreated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryLogTokenCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "LogTokenCreated", log); err != nil {
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
func (_TokenFactory *TokenFactoryFilterer) ParseLogTokenCreated(log types.Log) (*TokenFactoryLogTokenCreated, error) {
	event := new(TokenFactoryLogTokenCreated)
	if err := _TokenFactory.contract.UnpackLog(event, "LogTokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenFactory contract.
type TokenFactoryOwnershipTransferredIterator struct {
	Event *TokenFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryOwnershipTransferred)
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
		it.Event = new(TokenFactoryOwnershipTransferred)
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
func (it *TokenFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the TokenFactory contract.
type TokenFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactory *TokenFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryOwnershipTransferredIterator{contract: _TokenFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFactory *TokenFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryOwnershipTransferred)
				if err := _TokenFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenFactory *TokenFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TokenFactoryOwnershipTransferred, error) {
	event := new(TokenFactoryOwnershipTransferred)
	if err := _TokenFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
