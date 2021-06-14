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

// TokenVestingStruct0 is an auto generated low-level Go binding around an user-defined struct.
type TokenVestingStruct0 struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}

// TokenVestingABI is the input ABI used to generate the binding from.
const TokenVestingABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogBeneficiaryUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogTokenAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Released\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}]},{\"type\":\"function\",\"name\":\"DAYS_IN_SECONDS\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"addToken\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address[3]\",\"name\":\"_beneficiaries\"},{\"type\":\"uint256[3]\",\"name\":\"_proportions\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"beneficiaries\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\"},{\"type\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"beneficiary\"},{\"type\":\"uint256\",\"name\":\"proportion\"},{\"type\":\"uint256\",\"name\":\"streamId\"},{\"type\":\"uint256\",\"name\":\"remaining\"}]},{\"type\":\"function\",\"name\":\"beneficiaryTokens\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\"},{\"type\":\"uint256\"}],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"getAllTokensByBeneficiary\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[{\"type\":\"address[]\"}]},{\"type\":\"function\",\"name\":\"getVestingInfo\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint256\",\"name\":\"totalBalance\"},{\"type\":\"uint256\",\"name\":\"beneficiariesCount\"},{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"stop\"}]}]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"release\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"releaseAll\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"releaseableAmount\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"releaseableAmountByAddress\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setSablier\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_sablier\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingAddress\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingReferral\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"vestedAmount\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"vestingInfo\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint256\",\"name\":\"totalBalance\"},{\"type\":\"uint256\",\"name\":\"beneficiariesCount\"},{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"stop\"}]}]"

// TokenVesting is an auto generated Go binding around an Ethereum contract.
type TokenVesting struct {
	TokenVestingCaller     // Read-only binding to the contract
	TokenVestingTransactor // Write-only binding to the contract
	TokenVestingFilterer   // Log filterer for contract events
}

// TokenVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenVestingSession struct {
	Contract     *TokenVesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenVestingCallerSession struct {
	Contract *TokenVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenVestingTransactorSession struct {
	Contract     *TokenVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenVestingRaw struct {
	Contract *TokenVesting // Generic contract binding to access the raw methods on
}

// TokenVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenVestingCallerRaw struct {
	Contract *TokenVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokenVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenVestingTransactorRaw struct {
	Contract *TokenVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenVesting creates a new instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVesting(address common.Address, backend bind.ContractBackend) (*TokenVesting, error) {
	contract, err := bindTokenVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// NewTokenVestingCaller creates a new read-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingCaller(address common.Address, caller bind.ContractCaller) (*TokenVestingCaller, error) {
	contract, err := bindTokenVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingCaller{contract: contract}, nil
}

// NewTokenVestingTransactor creates a new write-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenVestingTransactor, error) {
	contract, err := bindTokenVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingTransactor{contract: contract}, nil
}

// NewTokenVestingFilterer creates a new log filterer instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenVestingFilterer, error) {
	contract, err := bindTokenVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenVestingFilterer{contract: contract}, nil
}

// bindTokenVesting binds a generic wrapper to an already deployed contract.
func bindTokenVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.TokenVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transact(opts, method, params...)
}

// DAYSINSECONDS is a free data retrieval call binding the contract method 0xe1762ef3.
//
// Solidity: function DAYS_IN_SECONDS() view returns(uint256)
func (_TokenVesting *TokenVestingCaller) DAYSINSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "DAYS_IN_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSINSECONDS is a free data retrieval call binding the contract method 0xe1762ef3.
//
// Solidity: function DAYS_IN_SECONDS() view returns(uint256)
func (_TokenVesting *TokenVestingSession) DAYSINSECONDS() (*big.Int, error) {
	return _TokenVesting.Contract.DAYSINSECONDS(&_TokenVesting.CallOpts)
}

// DAYSINSECONDS is a free data retrieval call binding the contract method 0xe1762ef3.
//
// Solidity: function DAYS_IN_SECONDS() view returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) DAYSINSECONDS() (*big.Int, error) {
	return _TokenVesting.Contract.DAYSINSECONDS(&_TokenVesting.CallOpts)
}

// Beneficiaries is a free data retrieval call binding the contract method 0x529849e9.
//
// Solidity: function beneficiaries(address , uint256 ) view returns(address beneficiary, uint256 proportion, uint256 streamId, uint256 remaining)
func (_TokenVesting *TokenVestingCaller) Beneficiaries(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Beneficiary common.Address
	Proportion  *big.Int
	StreamId    *big.Int
	Remaining   *big.Int
}, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "beneficiaries", arg0, arg1)

	outstruct := new(struct {
		Beneficiary common.Address
		Proportion  *big.Int
		StreamId    *big.Int
		Remaining   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Beneficiary = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Proportion = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StreamId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Remaining = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Beneficiaries is a free data retrieval call binding the contract method 0x529849e9.
//
// Solidity: function beneficiaries(address , uint256 ) view returns(address beneficiary, uint256 proportion, uint256 streamId, uint256 remaining)
func (_TokenVesting *TokenVestingSession) Beneficiaries(arg0 common.Address, arg1 *big.Int) (struct {
	Beneficiary common.Address
	Proportion  *big.Int
	StreamId    *big.Int
	Remaining   *big.Int
}, error) {
	return _TokenVesting.Contract.Beneficiaries(&_TokenVesting.CallOpts, arg0, arg1)
}

// Beneficiaries is a free data retrieval call binding the contract method 0x529849e9.
//
// Solidity: function beneficiaries(address , uint256 ) view returns(address beneficiary, uint256 proportion, uint256 streamId, uint256 remaining)
func (_TokenVesting *TokenVestingCallerSession) Beneficiaries(arg0 common.Address, arg1 *big.Int) (struct {
	Beneficiary common.Address
	Proportion  *big.Int
	StreamId    *big.Int
	Remaining   *big.Int
}, error) {
	return _TokenVesting.Contract.Beneficiaries(&_TokenVesting.CallOpts, arg0, arg1)
}

// BeneficiaryTokens is a free data retrieval call binding the contract method 0x47a9d145.
//
// Solidity: function beneficiaryTokens(address , uint256 ) view returns(address)
func (_TokenVesting *TokenVestingCaller) BeneficiaryTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "beneficiaryTokens", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeneficiaryTokens is a free data retrieval call binding the contract method 0x47a9d145.
//
// Solidity: function beneficiaryTokens(address , uint256 ) view returns(address)
func (_TokenVesting *TokenVestingSession) BeneficiaryTokens(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TokenVesting.Contract.BeneficiaryTokens(&_TokenVesting.CallOpts, arg0, arg1)
}

// BeneficiaryTokens is a free data retrieval call binding the contract method 0x47a9d145.
//
// Solidity: function beneficiaryTokens(address , uint256 ) view returns(address)
func (_TokenVesting *TokenVestingCallerSession) BeneficiaryTokens(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TokenVesting.Contract.BeneficiaryTokens(&_TokenVesting.CallOpts, arg0, arg1)
}

// GetAllTokensByBeneficiary is a free data retrieval call binding the contract method 0x6d18450a.
//
// Solidity: function getAllTokensByBeneficiary(address _beneficiary) view returns(address[])
func (_TokenVesting *TokenVestingCaller) GetAllTokensByBeneficiary(opts *bind.CallOpts, _beneficiary common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "getAllTokensByBeneficiary", _beneficiary)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllTokensByBeneficiary is a free data retrieval call binding the contract method 0x6d18450a.
//
// Solidity: function getAllTokensByBeneficiary(address _beneficiary) view returns(address[])
func (_TokenVesting *TokenVestingSession) GetAllTokensByBeneficiary(_beneficiary common.Address) ([]common.Address, error) {
	return _TokenVesting.Contract.GetAllTokensByBeneficiary(&_TokenVesting.CallOpts, _beneficiary)
}

// GetAllTokensByBeneficiary is a free data retrieval call binding the contract method 0x6d18450a.
//
// Solidity: function getAllTokensByBeneficiary(address _beneficiary) view returns(address[])
func (_TokenVesting *TokenVestingCallerSession) GetAllTokensByBeneficiary(_beneficiary common.Address) ([]common.Address, error) {
	return _TokenVesting.Contract.GetAllTokensByBeneficiary(&_TokenVesting.CallOpts, _beneficiary)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xfb897ce4.
//
// Solidity: function getVestingInfo(address _token) view returns((address,uint256,uint256,uint256,uint256))
func (_TokenVesting *TokenVestingCaller) GetVestingInfo(opts *bind.CallOpts, _token common.Address) (TokenVestingStruct0, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "getVestingInfo", _token)

	if err != nil {
		return *new(TokenVestingStruct0), err
	}

	out0 := *abi.ConvertType(out[0], new(TokenVestingStruct0)).(*TokenVestingStruct0)

	return out0, err

}

// GetVestingInfo is a free data retrieval call binding the contract method 0xfb897ce4.
//
// Solidity: function getVestingInfo(address _token) view returns((address,uint256,uint256,uint256,uint256))
func (_TokenVesting *TokenVestingSession) GetVestingInfo(_token common.Address) (TokenVestingStruct0, error) {
	return _TokenVesting.Contract.GetVestingInfo(&_TokenVesting.CallOpts, _token)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xfb897ce4.
//
// Solidity: function getVestingInfo(address _token) view returns((address,uint256,uint256,uint256,uint256))
func (_TokenVesting *TokenVestingCallerSession) GetVestingInfo(_token common.Address) (TokenVestingStruct0, error) {
	return _TokenVesting.Contract.GetVestingInfo(&_TokenVesting.CallOpts, _token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenVesting *TokenVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenVesting *TokenVestingSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenVesting *TokenVestingCallerSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(address _token) view returns(uint256)
func (_TokenVesting *TokenVestingCaller) ReleaseableAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "releaseableAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(address _token) view returns(uint256)
func (_TokenVesting *TokenVestingSession) ReleaseableAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleaseableAmount(&_TokenVesting.CallOpts, _token)
}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(address _token) view returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) ReleaseableAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleaseableAmount(&_TokenVesting.CallOpts, _token)
}

// ReleaseableAmountByAddress is a free data retrieval call binding the contract method 0x4020925a.
//
// Solidity: function releaseableAmountByAddress(address _token, address _beneficiary) view returns(uint256)
func (_TokenVesting *TokenVestingCaller) ReleaseableAmountByAddress(opts *bind.CallOpts, _token common.Address, _beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "releaseableAmountByAddress", _token, _beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseableAmountByAddress is a free data retrieval call binding the contract method 0x4020925a.
//
// Solidity: function releaseableAmountByAddress(address _token, address _beneficiary) view returns(uint256)
func (_TokenVesting *TokenVestingSession) ReleaseableAmountByAddress(_token common.Address, _beneficiary common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleaseableAmountByAddress(&_TokenVesting.CallOpts, _token, _beneficiary)
}

// ReleaseableAmountByAddress is a free data retrieval call binding the contract method 0x4020925a.
//
// Solidity: function releaseableAmountByAddress(address _token, address _beneficiary) view returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) ReleaseableAmountByAddress(_token common.Address, _beneficiary common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.ReleaseableAmountByAddress(&_TokenVesting.CallOpts, _token, _beneficiary)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(address _token) view returns(uint256)
func (_TokenVesting *TokenVestingCaller) VestedAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "vestedAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(address _token) view returns(uint256)
func (_TokenVesting *TokenVestingSession) VestedAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, _token)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(address _token) view returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) VestedAmount(_token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.VestedAmount(&_TokenVesting.CallOpts, _token)
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo(address ) view returns(address vestingBeneficiary, uint256 totalBalance, uint256 beneficiariesCount, uint256 start, uint256 stop)
func (_TokenVesting *TokenVestingCaller) VestingInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}, error) {
	var out []interface{}
	err := _TokenVesting.contract.Call(opts, &out, "vestingInfo", arg0)

	outstruct := new(struct {
		VestingBeneficiary common.Address
		TotalBalance       *big.Int
		BeneficiariesCount *big.Int
		Start              *big.Int
		Stop               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VestingBeneficiary = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TotalBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeneficiariesCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Stop = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo(address ) view returns(address vestingBeneficiary, uint256 totalBalance, uint256 beneficiariesCount, uint256 start, uint256 stop)
func (_TokenVesting *TokenVestingSession) VestingInfo(arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}, error) {
	return _TokenVesting.Contract.VestingInfo(&_TokenVesting.CallOpts, arg0)
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo(address ) view returns(address vestingBeneficiary, uint256 totalBalance, uint256 beneficiariesCount, uint256 start, uint256 stop)
func (_TokenVesting *TokenVestingCallerSession) VestingInfo(arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}, error) {
	return _TokenVesting.Contract.VestingInfo(&_TokenVesting.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xc25e888f.
//
// Solidity: function addToken(address _token, address[3] _beneficiaries, uint256[3] _proportions, uint256 _vestingPeriodInDays) returns()
func (_TokenVesting *TokenVestingTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _beneficiaries [3]common.Address, _proportions [3]*big.Int, _vestingPeriodInDays *big.Int) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "addToken", _token, _beneficiaries, _proportions, _vestingPeriodInDays)
}

// AddToken is a paid mutator transaction binding the contract method 0xc25e888f.
//
// Solidity: function addToken(address _token, address[3] _beneficiaries, uint256[3] _proportions, uint256 _vestingPeriodInDays) returns()
func (_TokenVesting *TokenVestingSession) AddToken(_token common.Address, _beneficiaries [3]common.Address, _proportions [3]*big.Int, _vestingPeriodInDays *big.Int) (*types.Transaction, error) {
	return _TokenVesting.Contract.AddToken(&_TokenVesting.TransactOpts, _token, _beneficiaries, _proportions, _vestingPeriodInDays)
}

// AddToken is a paid mutator transaction binding the contract method 0xc25e888f.
//
// Solidity: function addToken(address _token, address[3] _beneficiaries, uint256[3] _proportions, uint256 _vestingPeriodInDays) returns()
func (_TokenVesting *TokenVestingTransactorSession) AddToken(_token common.Address, _beneficiaries [3]common.Address, _proportions [3]*big.Int, _vestingPeriodInDays *big.Int) (*types.Transaction, error) {
	return _TokenVesting.Contract.AddToken(&_TokenVesting.TransactOpts, _token, _beneficiaries, _proportions, _vestingPeriodInDays)
}

// Release is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address _token, address _beneficiary) returns()
func (_TokenVesting *TokenVestingTransactor) Release(opts *bind.TransactOpts, _token common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "release", _token, _beneficiary)
}

// Release is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address _token, address _beneficiary) returns()
func (_TokenVesting *TokenVestingSession) Release(_token common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, _token, _beneficiary)
}

// Release is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address _token, address _beneficiary) returns()
func (_TokenVesting *TokenVestingTransactorSession) Release(_token common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, _token, _beneficiary)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address _beneficiary) returns()
func (_TokenVesting *TokenVestingTransactor) ReleaseAll(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "releaseAll", _beneficiary)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address _beneficiary) returns()
func (_TokenVesting *TokenVestingSession) ReleaseAll(_beneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.ReleaseAll(&_TokenVesting.TransactOpts, _beneficiary)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address _beneficiary) returns()
func (_TokenVesting *TokenVestingTransactorSession) ReleaseAll(_beneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.ReleaseAll(&_TokenVesting.TransactOpts, _beneficiary)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenVesting.Contract.RenounceOwnership(&_TokenVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenVesting.Contract.RenounceOwnership(&_TokenVesting.TransactOpts)
}

// SetSablier is a paid mutator transaction binding the contract method 0x90096556.
//
// Solidity: function setSablier(address _sablier) returns()
func (_TokenVesting *TokenVestingTransactor) SetSablier(opts *bind.TransactOpts, _sablier common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "setSablier", _sablier)
}

// SetSablier is a paid mutator transaction binding the contract method 0x90096556.
//
// Solidity: function setSablier(address _sablier) returns()
func (_TokenVesting *TokenVestingSession) SetSablier(_sablier common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.SetSablier(&_TokenVesting.TransactOpts, _sablier)
}

// SetSablier is a paid mutator transaction binding the contract method 0x90096556.
//
// Solidity: function setSablier(address _sablier) returns()
func (_TokenVesting *TokenVestingTransactorSession) SetSablier(_sablier common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.SetSablier(&_TokenVesting.TransactOpts, _sablier)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_TokenVesting *TokenVestingTransactor) SetVestingAddress(opts *bind.TransactOpts, _token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "setVestingAddress", _token, _vestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_TokenVesting *TokenVestingSession) SetVestingAddress(_token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.SetVestingAddress(&_TokenVesting.TransactOpts, _token, _vestingBeneficiary)
}

// SetVestingAddress is a paid mutator transaction binding the contract method 0x0876eae5.
//
// Solidity: function setVestingAddress(address _token, address _vestingBeneficiary) returns()
func (_TokenVesting *TokenVestingTransactorSession) SetVestingAddress(_token common.Address, _vestingBeneficiary common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.SetVestingAddress(&_TokenVesting.TransactOpts, _token, _vestingBeneficiary)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_TokenVesting *TokenVestingTransactor) SetVestingReferral(opts *bind.TransactOpts, _token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "setVestingReferral", _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_TokenVesting *TokenVestingSession) SetVestingReferral(_token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.SetVestingReferral(&_TokenVesting.TransactOpts, _token, _vestingReferral)
}

// SetVestingReferral is a paid mutator transaction binding the contract method 0xa8fc15c6.
//
// Solidity: function setVestingReferral(address _token, address _vestingReferral) returns()
func (_TokenVesting *TokenVestingTransactorSession) SetVestingReferral(_token common.Address, _vestingReferral common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.SetVestingReferral(&_TokenVesting.TransactOpts, _token, _vestingReferral)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenVesting *TokenVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenVesting *TokenVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenVesting *TokenVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TokenVestingLogBeneficiaryUpdatedIterator is returned from FilterLogBeneficiaryUpdated and is used to iterate over the raw logs and unpacked data for LogBeneficiaryUpdated events raised by the TokenVesting contract.
type TokenVestingLogBeneficiaryUpdatedIterator struct {
	Event *TokenVestingLogBeneficiaryUpdated // Event containing the contract specifics and raw log

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
func (it *TokenVestingLogBeneficiaryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingLogBeneficiaryUpdated)
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
		it.Event = new(TokenVestingLogBeneficiaryUpdated)
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
func (it *TokenVestingLogBeneficiaryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingLogBeneficiaryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingLogBeneficiaryUpdated represents a LogBeneficiaryUpdated event raised by the TokenVesting contract.
type TokenVestingLogBeneficiaryUpdated struct {
	Token              common.Address
	VestingBeneficiary common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogBeneficiaryUpdated is a free log retrieval operation binding the contract event 0x276dd15f31f6b99d62254388504662bee969db7d6766172c0ff60f8576f09453.
//
// Solidity: event LogBeneficiaryUpdated(address indexed token, address vestingBeneficiary)
func (_TokenVesting *TokenVestingFilterer) FilterLogBeneficiaryUpdated(opts *bind.FilterOpts, token []common.Address) (*TokenVestingLogBeneficiaryUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "LogBeneficiaryUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingLogBeneficiaryUpdatedIterator{contract: _TokenVesting.contract, event: "LogBeneficiaryUpdated", logs: logs, sub: sub}, nil
}

// WatchLogBeneficiaryUpdated is a free log subscription operation binding the contract event 0x276dd15f31f6b99d62254388504662bee969db7d6766172c0ff60f8576f09453.
//
// Solidity: event LogBeneficiaryUpdated(address indexed token, address vestingBeneficiary)
func (_TokenVesting *TokenVestingFilterer) WatchLogBeneficiaryUpdated(opts *bind.WatchOpts, sink chan<- *TokenVestingLogBeneficiaryUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "LogBeneficiaryUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingLogBeneficiaryUpdated)
				if err := _TokenVesting.contract.UnpackLog(event, "LogBeneficiaryUpdated", log); err != nil {
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

// ParseLogBeneficiaryUpdated is a log parse operation binding the contract event 0x276dd15f31f6b99d62254388504662bee969db7d6766172c0ff60f8576f09453.
//
// Solidity: event LogBeneficiaryUpdated(address indexed token, address vestingBeneficiary)
func (_TokenVesting *TokenVestingFilterer) ParseLogBeneficiaryUpdated(log types.Log) (*TokenVestingLogBeneficiaryUpdated, error) {
	event := new(TokenVestingLogBeneficiaryUpdated)
	if err := _TokenVesting.contract.UnpackLog(event, "LogBeneficiaryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVestingLogTokenAddedIterator is returned from FilterLogTokenAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdded events raised by the TokenVesting contract.
type TokenVestingLogTokenAddedIterator struct {
	Event *TokenVestingLogTokenAdded // Event containing the contract specifics and raw log

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
func (it *TokenVestingLogTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingLogTokenAdded)
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
		it.Event = new(TokenVestingLogTokenAdded)
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
func (it *TokenVestingLogTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingLogTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingLogTokenAdded represents a LogTokenAdded event raised by the TokenVesting contract.
type TokenVestingLogTokenAdded struct {
	Token               common.Address
	VestingBeneficiary  common.Address
	VestingPeriodInDays *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdded is a free log retrieval operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: event LogTokenAdded(address indexed token, address vestingBeneficiary, uint256 vestingPeriodInDays)
func (_TokenVesting *TokenVestingFilterer) FilterLogTokenAdded(opts *bind.FilterOpts, token []common.Address) (*TokenVestingLogTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "LogTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingLogTokenAddedIterator{contract: _TokenVesting.contract, event: "LogTokenAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdded is a free log subscription operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: event LogTokenAdded(address indexed token, address vestingBeneficiary, uint256 vestingPeriodInDays)
func (_TokenVesting *TokenVestingFilterer) WatchLogTokenAdded(opts *bind.WatchOpts, sink chan<- *TokenVestingLogTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "LogTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingLogTokenAdded)
				if err := _TokenVesting.contract.UnpackLog(event, "LogTokenAdded", log); err != nil {
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

// ParseLogTokenAdded is a log parse operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: event LogTokenAdded(address indexed token, address vestingBeneficiary, uint256 vestingPeriodInDays)
func (_TokenVesting *TokenVestingFilterer) ParseLogTokenAdded(log types.Log) (*TokenVestingLogTokenAdded, error) {
	event := new(TokenVestingLogTokenAdded)
	if err := _TokenVesting.contract.UnpackLog(event, "LogTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenVesting contract.
type TokenVestingOwnershipTransferredIterator struct {
	Event *TokenVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingOwnershipTransferred)
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
		it.Event = new(TokenVestingOwnershipTransferred)
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
func (it *TokenVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TokenVesting contract.
type TokenVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenVesting *TokenVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingOwnershipTransferredIterator{contract: _TokenVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenVesting *TokenVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingOwnershipTransferred)
				if err := _TokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenVesting *TokenVestingFilterer) ParseOwnershipTransferred(log types.Log) (*TokenVestingOwnershipTransferred, error) {
	event := new(TokenVestingOwnershipTransferred)
	if err := _TokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenVestingReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the TokenVesting contract.
type TokenVestingReleasedIterator struct {
	Event *TokenVestingReleased // Event containing the contract specifics and raw log

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
func (it *TokenVestingReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingReleased)
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
		it.Event = new(TokenVestingReleased)
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
func (it *TokenVestingReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingReleased represents a Released event raised by the TokenVesting contract.
type TokenVestingReleased struct {
	Token              common.Address
	VestingBeneficiary common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed token, address vestingBeneficiary, uint256 amount)
func (_TokenVesting *TokenVestingFilterer) FilterReleased(opts *bind.FilterOpts, token []common.Address) (*TokenVestingReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingReleasedIterator{contract: _TokenVesting.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed token, address vestingBeneficiary, uint256 amount)
func (_TokenVesting *TokenVestingFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *TokenVestingReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingReleased)
				if err := _TokenVesting.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed token, address vestingBeneficiary, uint256 amount)
func (_TokenVesting *TokenVestingFilterer) ParseReleased(log types.Log) (*TokenVestingReleased, error) {
	event := new(TokenVestingReleased)
	if err := _TokenVesting.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
