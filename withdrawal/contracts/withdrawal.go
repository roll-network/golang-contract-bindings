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
	_ = abi.ConvertType
)

// WithdrawalOrder is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalOrder struct {
	Signer     common.Address
	Holder     common.Address
	To         common.Address
	Token      common.Address
	Amount     *big.Int
	Expiration *big.Int
	Key        *big.Int
}

// WithdrawalSig is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalSig struct {
	V uint8
	R [32]byte
	S [32]byte
}

// WithdrawalMetaData contains all meta data concerning the Withdrawal contract.
var WithdrawalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauseOperator_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CancelledOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessedOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CancelEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"UpdatePauseOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"UpdateSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedSigners\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawal.Sig\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getCompletedKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumWithdrawal.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAllowedSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauseOperator_\",\"type\":\"address\"}],\"name\":\"setPauseOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawal.Sig\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WithdrawalABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalMetaData.ABI instead.
var WithdrawalABI = WithdrawalMetaData.ABI

// Withdrawal is an auto generated Go binding around an Ethereum contract.
type Withdrawal struct {
	WithdrawalCaller     // Read-only binding to the contract
	WithdrawalTransactor // Write-only binding to the contract
	WithdrawalFilterer   // Log filterer for contract events
}

// WithdrawalCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalSession struct {
	Contract     *Withdrawal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalCallerSession struct {
	Contract *WithdrawalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WithdrawalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalTransactorSession struct {
	Contract     *WithdrawalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WithdrawalRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalRaw struct {
	Contract *Withdrawal // Generic contract binding to access the raw methods on
}

// WithdrawalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalCallerRaw struct {
	Contract *WithdrawalCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalTransactorRaw struct {
	Contract *WithdrawalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawal creates a new instance of Withdrawal, bound to a specific deployed contract.
func NewWithdrawal(address common.Address, backend bind.ContractBackend) (*Withdrawal, error) {
	contract, err := bindWithdrawal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawal{WithdrawalCaller: WithdrawalCaller{contract: contract}, WithdrawalTransactor: WithdrawalTransactor{contract: contract}, WithdrawalFilterer: WithdrawalFilterer{contract: contract}}, nil
}

// NewWithdrawalCaller creates a new read-only instance of Withdrawal, bound to a specific deployed contract.
func NewWithdrawalCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalCaller, error) {
	contract, err := bindWithdrawal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalCaller{contract: contract}, nil
}

// NewWithdrawalTransactor creates a new write-only instance of Withdrawal, bound to a specific deployed contract.
func NewWithdrawalTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalTransactor, error) {
	contract, err := bindWithdrawal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalTransactor{contract: contract}, nil
}

// NewWithdrawalFilterer creates a new log filterer instance of Withdrawal, bound to a specific deployed contract.
func NewWithdrawalFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalFilterer, error) {
	contract, err := bindWithdrawal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalFilterer{contract: contract}, nil
}

// bindWithdrawal binds a generic wrapper to an already deployed contract.
func bindWithdrawal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawal *WithdrawalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawal.Contract.WithdrawalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawal *WithdrawalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawal.Contract.WithdrawalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawal *WithdrawalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawal.Contract.WithdrawalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawal *WithdrawalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawal *WithdrawalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawal *WithdrawalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawal.Contract.contract.Transact(opts, method, params...)
}

// AllowedSigners is a free data retrieval call binding the contract method 0xf8b4d864.
//
// Solidity: function allowedSigners(address ) view returns(bool)
func (_Withdrawal *WithdrawalCaller) AllowedSigners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "allowedSigners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedSigners is a free data retrieval call binding the contract method 0xf8b4d864.
//
// Solidity: function allowedSigners(address ) view returns(bool)
func (_Withdrawal *WithdrawalSession) AllowedSigners(arg0 common.Address) (bool, error) {
	return _Withdrawal.Contract.AllowedSigners(&_Withdrawal.CallOpts, arg0)
}

// AllowedSigners is a free data retrieval call binding the contract method 0xf8b4d864.
//
// Solidity: function allowedSigners(address ) view returns(bool)
func (_Withdrawal *WithdrawalCallerSession) AllowedSigners(arg0 common.Address) (bool, error) {
	return _Withdrawal.Contract.AllowedSigners(&_Withdrawal.CallOpts, arg0)
}

// GetCompletedKey is a free data retrieval call binding the contract method 0x46c368c2.
//
// Solidity: function getCompletedKey((address,address,address,address,uint256,uint256,uint256) order) pure returns(uint256)
func (_Withdrawal *WithdrawalCaller) GetCompletedKey(opts *bind.CallOpts, order WithdrawalOrder) (*big.Int, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "getCompletedKey", order)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompletedKey is a free data retrieval call binding the contract method 0x46c368c2.
//
// Solidity: function getCompletedKey((address,address,address,address,uint256,uint256,uint256) order) pure returns(uint256)
func (_Withdrawal *WithdrawalSession) GetCompletedKey(order WithdrawalOrder) (*big.Int, error) {
	return _Withdrawal.Contract.GetCompletedKey(&_Withdrawal.CallOpts, order)
}

// GetCompletedKey is a free data retrieval call binding the contract method 0x46c368c2.
//
// Solidity: function getCompletedKey((address,address,address,address,uint256,uint256,uint256) order) pure returns(uint256)
func (_Withdrawal *WithdrawalCallerSession) GetCompletedKey(order WithdrawalOrder) (*big.Int, error) {
	return _Withdrawal.Contract.GetCompletedKey(&_Withdrawal.CallOpts, order)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 key) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Withdrawal *WithdrawalCaller) GetOrder(opts *bind.CallOpts, key *big.Int) (WithdrawalOrder, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "getOrder", key)

	if err != nil {
		return *new(WithdrawalOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(WithdrawalOrder)).(*WithdrawalOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 key) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Withdrawal *WithdrawalSession) GetOrder(key *big.Int) (WithdrawalOrder, error) {
	return _Withdrawal.Contract.GetOrder(&_Withdrawal.CallOpts, key)
}

// GetOrder is a free data retrieval call binding the contract method 0xd09ef241.
//
// Solidity: function getOrder(uint256 key) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Withdrawal *WithdrawalCallerSession) GetOrder(key *big.Int) (WithdrawalOrder, error) {
	return _Withdrawal.Contract.GetOrder(&_Withdrawal.CallOpts, key)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 key) view returns(uint8)
func (_Withdrawal *WithdrawalCaller) GetStatus(opts *bind.CallOpts, key *big.Int) (uint8, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "getStatus", key)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 key) view returns(uint8)
func (_Withdrawal *WithdrawalSession) GetStatus(key *big.Int) (uint8, error) {
	return _Withdrawal.Contract.GetStatus(&_Withdrawal.CallOpts, key)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 key) view returns(uint8)
func (_Withdrawal *WithdrawalCallerSession) GetStatus(key *big.Int) (uint8, error) {
	return _Withdrawal.Contract.GetStatus(&_Withdrawal.CallOpts, key)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address signer, address holder, address to, address token, uint256 amount, uint256 expiration, uint256 key)
func (_Withdrawal *WithdrawalCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Signer     common.Address
	Holder     common.Address
	To         common.Address
	Token      common.Address
	Amount     *big.Int
	Expiration *big.Int
	Key        *big.Int
}, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Signer     common.Address
		Holder     common.Address
		To         common.Address
		Token      common.Address
		Amount     *big.Int
		Expiration *big.Int
		Key        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Holder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Expiration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Key = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address signer, address holder, address to, address token, uint256 amount, uint256 expiration, uint256 key)
func (_Withdrawal *WithdrawalSession) Orders(arg0 *big.Int) (struct {
	Signer     common.Address
	Holder     common.Address
	To         common.Address
	Token      common.Address
	Amount     *big.Int
	Expiration *big.Int
	Key        *big.Int
}, error) {
	return _Withdrawal.Contract.Orders(&_Withdrawal.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address signer, address holder, address to, address token, uint256 amount, uint256 expiration, uint256 key)
func (_Withdrawal *WithdrawalCallerSession) Orders(arg0 *big.Int) (struct {
	Signer     common.Address
	Holder     common.Address
	To         common.Address
	Token      common.Address
	Amount     *big.Int
	Expiration *big.Int
	Key        *big.Int
}, error) {
	return _Withdrawal.Contract.Orders(&_Withdrawal.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Withdrawal *WithdrawalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Withdrawal *WithdrawalSession) Owner() (common.Address, error) {
	return _Withdrawal.Contract.Owner(&_Withdrawal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Withdrawal *WithdrawalCallerSession) Owner() (common.Address, error) {
	return _Withdrawal.Contract.Owner(&_Withdrawal.CallOpts)
}

// PauseOperator is a free data retrieval call binding the contract method 0x4afdcbde.
//
// Solidity: function pauseOperator() view returns(address)
func (_Withdrawal *WithdrawalCaller) PauseOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "pauseOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseOperator is a free data retrieval call binding the contract method 0x4afdcbde.
//
// Solidity: function pauseOperator() view returns(address)
func (_Withdrawal *WithdrawalSession) PauseOperator() (common.Address, error) {
	return _Withdrawal.Contract.PauseOperator(&_Withdrawal.CallOpts)
}

// PauseOperator is a free data retrieval call binding the contract method 0x4afdcbde.
//
// Solidity: function pauseOperator() view returns(address)
func (_Withdrawal *WithdrawalCallerSession) PauseOperator() (common.Address, error) {
	return _Withdrawal.Contract.PauseOperator(&_Withdrawal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Withdrawal *WithdrawalCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Withdrawal *WithdrawalSession) Paused() (bool, error) {
	return _Withdrawal.Contract.Paused(&_Withdrawal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Withdrawal *WithdrawalCallerSession) Paused() (bool, error) {
	return _Withdrawal.Contract.Paused(&_Withdrawal.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xebe73ca4.
//
// Solidity: function cancel((address,address,address,address,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Withdrawal *WithdrawalTransactor) Cancel(opts *bind.TransactOpts, order WithdrawalOrder, sig WithdrawalSig) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "cancel", order, sig)
}

// Cancel is a paid mutator transaction binding the contract method 0xebe73ca4.
//
// Solidity: function cancel((address,address,address,address,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Withdrawal *WithdrawalSession) Cancel(order WithdrawalOrder, sig WithdrawalSig) (*types.Transaction, error) {
	return _Withdrawal.Contract.Cancel(&_Withdrawal.TransactOpts, order, sig)
}

// Cancel is a paid mutator transaction binding the contract method 0xebe73ca4.
//
// Solidity: function cancel((address,address,address,address,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Withdrawal *WithdrawalTransactorSession) Cancel(order WithdrawalOrder, sig WithdrawalSig) (*types.Transaction, error) {
	return _Withdrawal.Contract.Cancel(&_Withdrawal.TransactOpts, order, sig)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Withdrawal *WithdrawalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Withdrawal *WithdrawalSession) Pause() (*types.Transaction, error) {
	return _Withdrawal.Contract.Pause(&_Withdrawal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Withdrawal *WithdrawalTransactorSession) Pause() (*types.Transaction, error) {
	return _Withdrawal.Contract.Pause(&_Withdrawal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Withdrawal *WithdrawalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Withdrawal *WithdrawalSession) RenounceOwnership() (*types.Transaction, error) {
	return _Withdrawal.Contract.RenounceOwnership(&_Withdrawal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Withdrawal *WithdrawalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Withdrawal.Contract.RenounceOwnership(&_Withdrawal.TransactOpts)
}

// SetAllowedSigner is a paid mutator transaction binding the contract method 0xf9213fc8.
//
// Solidity: function setAllowedSigner(address signer, bool value) returns()
func (_Withdrawal *WithdrawalTransactor) SetAllowedSigner(opts *bind.TransactOpts, signer common.Address, value bool) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "setAllowedSigner", signer, value)
}

// SetAllowedSigner is a paid mutator transaction binding the contract method 0xf9213fc8.
//
// Solidity: function setAllowedSigner(address signer, bool value) returns()
func (_Withdrawal *WithdrawalSession) SetAllowedSigner(signer common.Address, value bool) (*types.Transaction, error) {
	return _Withdrawal.Contract.SetAllowedSigner(&_Withdrawal.TransactOpts, signer, value)
}

// SetAllowedSigner is a paid mutator transaction binding the contract method 0xf9213fc8.
//
// Solidity: function setAllowedSigner(address signer, bool value) returns()
func (_Withdrawal *WithdrawalTransactorSession) SetAllowedSigner(signer common.Address, value bool) (*types.Transaction, error) {
	return _Withdrawal.Contract.SetAllowedSigner(&_Withdrawal.TransactOpts, signer, value)
}

// SetPauseOperator is a paid mutator transaction binding the contract method 0x1ad4aa42.
//
// Solidity: function setPauseOperator(address pauseOperator_) returns()
func (_Withdrawal *WithdrawalTransactor) SetPauseOperator(opts *bind.TransactOpts, pauseOperator_ common.Address) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "setPauseOperator", pauseOperator_)
}

// SetPauseOperator is a paid mutator transaction binding the contract method 0x1ad4aa42.
//
// Solidity: function setPauseOperator(address pauseOperator_) returns()
func (_Withdrawal *WithdrawalSession) SetPauseOperator(pauseOperator_ common.Address) (*types.Transaction, error) {
	return _Withdrawal.Contract.SetPauseOperator(&_Withdrawal.TransactOpts, pauseOperator_)
}

// SetPauseOperator is a paid mutator transaction binding the contract method 0x1ad4aa42.
//
// Solidity: function setPauseOperator(address pauseOperator_) returns()
func (_Withdrawal *WithdrawalTransactorSession) SetPauseOperator(pauseOperator_ common.Address) (*types.Transaction, error) {
	return _Withdrawal.Contract.SetPauseOperator(&_Withdrawal.TransactOpts, pauseOperator_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Withdrawal *WithdrawalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Withdrawal *WithdrawalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Withdrawal.Contract.TransferOwnership(&_Withdrawal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Withdrawal *WithdrawalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Withdrawal.Contract.TransferOwnership(&_Withdrawal.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Withdrawal *WithdrawalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Withdrawal *WithdrawalSession) Unpause() (*types.Transaction, error) {
	return _Withdrawal.Contract.Unpause(&_Withdrawal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Withdrawal *WithdrawalTransactorSession) Unpause() (*types.Transaction, error) {
	return _Withdrawal.Contract.Unpause(&_Withdrawal.TransactOpts)
}

// Withdrawal is a paid mutator transaction binding the contract method 0xa7678b18.
//
// Solidity: function withdrawal((address,address,address,address,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Withdrawal *WithdrawalTransactor) Withdrawal(opts *bind.TransactOpts, order WithdrawalOrder, sig WithdrawalSig) (*types.Transaction, error) {
	return _Withdrawal.contract.Transact(opts, "withdrawal", order, sig)
}

// Withdrawal is a paid mutator transaction binding the contract method 0xa7678b18.
//
// Solidity: function withdrawal((address,address,address,address,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Withdrawal *WithdrawalSession) Withdrawal(order WithdrawalOrder, sig WithdrawalSig) (*types.Transaction, error) {
	return _Withdrawal.Contract.Withdrawal(&_Withdrawal.TransactOpts, order, sig)
}

// Withdrawal is a paid mutator transaction binding the contract method 0xa7678b18.
//
// Solidity: function withdrawal((address,address,address,address,uint256,uint256,uint256) order, (uint8,bytes32,bytes32) sig) returns()
func (_Withdrawal *WithdrawalTransactorSession) Withdrawal(order WithdrawalOrder, sig WithdrawalSig) (*types.Transaction, error) {
	return _Withdrawal.Contract.Withdrawal(&_Withdrawal.TransactOpts, order, sig)
}

// WithdrawalCancelEventIterator is returned from FilterCancelEvent and is used to iterate over the raw logs and unpacked data for CancelEvent events raised by the Withdrawal contract.
type WithdrawalCancelEventIterator struct {
	Event *WithdrawalCancelEvent // Event containing the contract specifics and raw log

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
func (it *WithdrawalCancelEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalCancelEvent)
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
		it.Event = new(WithdrawalCancelEvent)
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
func (it *WithdrawalCancelEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalCancelEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalCancelEvent represents a CancelEvent event raised by the Withdrawal contract.
type WithdrawalCancelEvent struct {
	Key    *big.Int
	Signer common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelEvent is a free log retrieval operation binding the contract event 0x18351c9e93b203d861f9bb52616afdc76a08cf3bb7c23d5d7a9f742fd2f309c1.
//
// Solidity: event CancelEvent(uint256 indexed key, address signer, address to, address indexed token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) FilterCancelEvent(opts *bind.FilterOpts, key []*big.Int, token []common.Address) (*WithdrawalCancelEventIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "CancelEvent", keyRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalCancelEventIterator{contract: _Withdrawal.contract, event: "CancelEvent", logs: logs, sub: sub}, nil
}

// WatchCancelEvent is a free log subscription operation binding the contract event 0x18351c9e93b203d861f9bb52616afdc76a08cf3bb7c23d5d7a9f742fd2f309c1.
//
// Solidity: event CancelEvent(uint256 indexed key, address signer, address to, address indexed token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) WatchCancelEvent(opts *bind.WatchOpts, sink chan<- *WithdrawalCancelEvent, key []*big.Int, token []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "CancelEvent", keyRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalCancelEvent)
				if err := _Withdrawal.contract.UnpackLog(event, "CancelEvent", log); err != nil {
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

// ParseCancelEvent is a log parse operation binding the contract event 0x18351c9e93b203d861f9bb52616afdc76a08cf3bb7c23d5d7a9f742fd2f309c1.
//
// Solidity: event CancelEvent(uint256 indexed key, address signer, address to, address indexed token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) ParseCancelEvent(log types.Log) (*WithdrawalCancelEvent, error) {
	event := new(WithdrawalCancelEvent)
	if err := _Withdrawal.contract.UnpackLog(event, "CancelEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Withdrawal contract.
type WithdrawalOwnershipTransferredIterator struct {
	Event *WithdrawalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WithdrawalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalOwnershipTransferred)
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
		it.Event = new(WithdrawalOwnershipTransferred)
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
func (it *WithdrawalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalOwnershipTransferred represents a OwnershipTransferred event raised by the Withdrawal contract.
type WithdrawalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Withdrawal *WithdrawalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WithdrawalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalOwnershipTransferredIterator{contract: _Withdrawal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Withdrawal *WithdrawalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WithdrawalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalOwnershipTransferred)
				if err := _Withdrawal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Withdrawal *WithdrawalFilterer) ParseOwnershipTransferred(log types.Log) (*WithdrawalOwnershipTransferred, error) {
	event := new(WithdrawalOwnershipTransferred)
	if err := _Withdrawal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Withdrawal contract.
type WithdrawalPausedIterator struct {
	Event *WithdrawalPaused // Event containing the contract specifics and raw log

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
func (it *WithdrawalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalPaused)
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
		it.Event = new(WithdrawalPaused)
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
func (it *WithdrawalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalPaused represents a Paused event raised by the Withdrawal contract.
type WithdrawalPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Withdrawal *WithdrawalFilterer) FilterPaused(opts *bind.FilterOpts) (*WithdrawalPausedIterator, error) {

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &WithdrawalPausedIterator{contract: _Withdrawal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Withdrawal *WithdrawalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *WithdrawalPaused) (event.Subscription, error) {

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalPaused)
				if err := _Withdrawal.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Withdrawal *WithdrawalFilterer) ParsePaused(log types.Log) (*WithdrawalPaused, error) {
	event := new(WithdrawalPaused)
	if err := _Withdrawal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Withdrawal contract.
type WithdrawalUnpausedIterator struct {
	Event *WithdrawalUnpaused // Event containing the contract specifics and raw log

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
func (it *WithdrawalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalUnpaused)
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
		it.Event = new(WithdrawalUnpaused)
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
func (it *WithdrawalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalUnpaused represents a Unpaused event raised by the Withdrawal contract.
type WithdrawalUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Withdrawal *WithdrawalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*WithdrawalUnpausedIterator, error) {

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &WithdrawalUnpausedIterator{contract: _Withdrawal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Withdrawal *WithdrawalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *WithdrawalUnpaused) (event.Subscription, error) {

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalUnpaused)
				if err := _Withdrawal.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Withdrawal *WithdrawalFilterer) ParseUnpaused(log types.Log) (*WithdrawalUnpaused, error) {
	event := new(WithdrawalUnpaused)
	if err := _Withdrawal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalUpdatePauseOperatorIterator is returned from FilterUpdatePauseOperator and is used to iterate over the raw logs and unpacked data for UpdatePauseOperator events raised by the Withdrawal contract.
type WithdrawalUpdatePauseOperatorIterator struct {
	Event *WithdrawalUpdatePauseOperator // Event containing the contract specifics and raw log

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
func (it *WithdrawalUpdatePauseOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalUpdatePauseOperator)
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
		it.Event = new(WithdrawalUpdatePauseOperator)
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
func (it *WithdrawalUpdatePauseOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalUpdatePauseOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalUpdatePauseOperator represents a UpdatePauseOperator event raised by the Withdrawal contract.
type WithdrawalUpdatePauseOperator struct {
	OldOperator common.Address
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePauseOperator is a free log retrieval operation binding the contract event 0x49af92e8e702a1adb06ae79c01a1dfed719800edb7f85d9b79e2622ea33f4cf9.
//
// Solidity: event UpdatePauseOperator(address oldOperator, address newOperator)
func (_Withdrawal *WithdrawalFilterer) FilterUpdatePauseOperator(opts *bind.FilterOpts) (*WithdrawalUpdatePauseOperatorIterator, error) {

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "UpdatePauseOperator")
	if err != nil {
		return nil, err
	}
	return &WithdrawalUpdatePauseOperatorIterator{contract: _Withdrawal.contract, event: "UpdatePauseOperator", logs: logs, sub: sub}, nil
}

// WatchUpdatePauseOperator is a free log subscription operation binding the contract event 0x49af92e8e702a1adb06ae79c01a1dfed719800edb7f85d9b79e2622ea33f4cf9.
//
// Solidity: event UpdatePauseOperator(address oldOperator, address newOperator)
func (_Withdrawal *WithdrawalFilterer) WatchUpdatePauseOperator(opts *bind.WatchOpts, sink chan<- *WithdrawalUpdatePauseOperator) (event.Subscription, error) {

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "UpdatePauseOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalUpdatePauseOperator)
				if err := _Withdrawal.contract.UnpackLog(event, "UpdatePauseOperator", log); err != nil {
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

// ParseUpdatePauseOperator is a log parse operation binding the contract event 0x49af92e8e702a1adb06ae79c01a1dfed719800edb7f85d9b79e2622ea33f4cf9.
//
// Solidity: event UpdatePauseOperator(address oldOperator, address newOperator)
func (_Withdrawal *WithdrawalFilterer) ParseUpdatePauseOperator(log types.Log) (*WithdrawalUpdatePauseOperator, error) {
	event := new(WithdrawalUpdatePauseOperator)
	if err := _Withdrawal.contract.UnpackLog(event, "UpdatePauseOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalUpdateSignerIterator is returned from FilterUpdateSigner and is used to iterate over the raw logs and unpacked data for UpdateSigner events raised by the Withdrawal contract.
type WithdrawalUpdateSignerIterator struct {
	Event *WithdrawalUpdateSigner // Event containing the contract specifics and raw log

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
func (it *WithdrawalUpdateSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalUpdateSigner)
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
		it.Event = new(WithdrawalUpdateSigner)
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
func (it *WithdrawalUpdateSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalUpdateSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalUpdateSigner represents a UpdateSigner event raised by the Withdrawal contract.
type WithdrawalUpdateSigner struct {
	Signer common.Address
	Value  bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateSigner is a free log retrieval operation binding the contract event 0x863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce.
//
// Solidity: event UpdateSigner(address signer, bool value)
func (_Withdrawal *WithdrawalFilterer) FilterUpdateSigner(opts *bind.FilterOpts) (*WithdrawalUpdateSignerIterator, error) {

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "UpdateSigner")
	if err != nil {
		return nil, err
	}
	return &WithdrawalUpdateSignerIterator{contract: _Withdrawal.contract, event: "UpdateSigner", logs: logs, sub: sub}, nil
}

// WatchUpdateSigner is a free log subscription operation binding the contract event 0x863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce.
//
// Solidity: event UpdateSigner(address signer, bool value)
func (_Withdrawal *WithdrawalFilterer) WatchUpdateSigner(opts *bind.WatchOpts, sink chan<- *WithdrawalUpdateSigner) (event.Subscription, error) {

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "UpdateSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalUpdateSigner)
				if err := _Withdrawal.contract.UnpackLog(event, "UpdateSigner", log); err != nil {
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

// ParseUpdateSigner is a log parse operation binding the contract event 0x863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce.
//
// Solidity: event UpdateSigner(address signer, bool value)
func (_Withdrawal *WithdrawalFilterer) ParseUpdateSigner(log types.Log) (*WithdrawalUpdateSigner, error) {
	event := new(WithdrawalUpdateSigner)
	if err := _Withdrawal.contract.UnpackLog(event, "UpdateSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalWithdrawalEventIterator is returned from FilterWithdrawalEvent and is used to iterate over the raw logs and unpacked data for WithdrawalEvent events raised by the Withdrawal contract.
type WithdrawalWithdrawalEventIterator struct {
	Event *WithdrawalWithdrawalEvent // Event containing the contract specifics and raw log

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
func (it *WithdrawalWithdrawalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalWithdrawalEvent)
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
		it.Event = new(WithdrawalWithdrawalEvent)
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
func (it *WithdrawalWithdrawalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalWithdrawalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalWithdrawalEvent represents a WithdrawalEvent event raised by the Withdrawal contract.
type WithdrawalWithdrawalEvent struct {
	Key    *big.Int
	Signer common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0x427a5c054bf4858fa5eaf4fbc1125a64160741ef59682db0fdc9687b38580944.
//
// Solidity: event WithdrawalEvent(uint256 indexed key, address signer, address to, address indexed token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts, key []*big.Int, token []common.Address) (*WithdrawalWithdrawalEventIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "WithdrawalEvent", keyRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalWithdrawalEventIterator{contract: _Withdrawal.contract, event: "WithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0x427a5c054bf4858fa5eaf4fbc1125a64160741ef59682db0fdc9687b38580944.
//
// Solidity: event WithdrawalEvent(uint256 indexed key, address signer, address to, address indexed token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *WithdrawalWithdrawalEvent, key []*big.Int, token []common.Address) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "WithdrawalEvent", keyRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalWithdrawalEvent)
				if err := _Withdrawal.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0x427a5c054bf4858fa5eaf4fbc1125a64160741ef59682db0fdc9687b38580944.
//
// Solidity: event WithdrawalEvent(uint256 indexed key, address signer, address to, address indexed token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) ParseWithdrawalEvent(log types.Log) (*WithdrawalWithdrawalEvent, error) {
	event := new(WithdrawalWithdrawalEvent)
	if err := _Withdrawal.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

