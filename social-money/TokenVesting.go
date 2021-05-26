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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogBeneficiaryUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogTokenAdded\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Released\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"vestingBeneficiary\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}]},{\"type\":\"function\",\"name\":\"DAYS_IN_SECONDS\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"addToken\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address[3]\",\"name\":\"_beneficiaries\"},{\"type\":\"uint256[3]\",\"name\":\"_proportions\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"beneficiaries\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\"},{\"type\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"beneficiary\"},{\"type\":\"uint256\",\"name\":\"proportion\"},{\"type\":\"uint256\",\"name\":\"streamId\"},{\"type\":\"uint256\",\"name\":\"remaining\"}]},{\"type\":\"function\",\"name\":\"beneficiaryTokens\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\"},{\"type\":\"uint256\"}],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"getAllTokensByBeneficiary\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[{\"type\":\"address[]\"}]},{\"type\":\"function\",\"name\":\"getVestingInfo\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint256\",\"name\":\"totalBalance\"},{\"type\":\"uint256\",\"name\":\"beneficiariesCount\"},{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"stop\"}]}]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"release\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"releaseAll\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"releaseableAmount\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"releaseableAmountByAddress\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_beneficiary\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setSablier\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_sablier\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingAddress\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"setVestingReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"address\",\"name\":\"_vestingReferral\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"vestedAmount\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"vestingInfo\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint256\",\"name\":\"totalBalance\"},{\"type\":\"uint256\",\"name\":\"beneficiariesCount\"},{\"type\":\"uint256\",\"name\":\"start\"},{\"type\":\"uint256\",\"name\":\"stop\"}]}]"

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

// DAYSINSECONDS is a free data retrieval call binding the contract method 0xe1762ef3.
//
// Solidity: function DAYS_IN_SECONDS() view returns(uint256)
func (_Contracts *ContractsCaller) DAYSINSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DAYS_IN_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSINSECONDS is a free data retrieval call binding the contract method 0xe1762ef3.
//
// Solidity: function DAYS_IN_SECONDS() view returns(uint256)
func (_Contracts *ContractsSession) DAYSINSECONDS() (*big.Int, error) {
	return _Contracts.Contract.DAYSINSECONDS(&_Contracts.CallOpts)
}

// DAYSINSECONDS is a free data retrieval call binding the contract method 0xe1762ef3.
//
// Solidity: function DAYS_IN_SECONDS() view returns(uint256)
func (_Contracts *ContractsCallerSession) DAYSINSECONDS() (*big.Int, error) {
	return _Contracts.Contract.DAYSINSECONDS(&_Contracts.CallOpts)
}

// Beneficiaries is a free data retrieval call binding the contract method 0x529849e9.
//
// Solidity: function beneficiaries(address , uint256 ) view returns(address beneficiary, uint256 proportion, uint256 streamId, uint256 remaining)
func (_Contracts *ContractsCaller) Beneficiaries(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Beneficiary common.Address
	Proportion  *big.Int
	StreamId    *big.Int
	Remaining   *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "beneficiaries", arg0, arg1)

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
func (_Contracts *ContractsSession) Beneficiaries(arg0 common.Address, arg1 *big.Int) (struct {
	Beneficiary common.Address
	Proportion  *big.Int
	StreamId    *big.Int
	Remaining   *big.Int
}, error) {
	return _Contracts.Contract.Beneficiaries(&_Contracts.CallOpts, arg0, arg1)
}

// Beneficiaries is a free data retrieval call binding the contract method 0x529849e9.
//
// Solidity: function beneficiaries(address , uint256 ) view returns(address beneficiary, uint256 proportion, uint256 streamId, uint256 remaining)
func (_Contracts *ContractsCallerSession) Beneficiaries(arg0 common.Address, arg1 *big.Int) (struct {
	Beneficiary common.Address
	Proportion  *big.Int
	StreamId    *big.Int
	Remaining   *big.Int
}, error) {
	return _Contracts.Contract.Beneficiaries(&_Contracts.CallOpts, arg0, arg1)
}

// BeneficiaryTokens is a free data retrieval call binding the contract method 0x47a9d145.
//
// Solidity: function beneficiaryTokens(address , uint256 ) view returns(address)
func (_Contracts *ContractsCaller) BeneficiaryTokens(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "beneficiaryTokens", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeneficiaryTokens is a free data retrieval call binding the contract method 0x47a9d145.
//
// Solidity: function beneficiaryTokens(address , uint256 ) view returns(address)
func (_Contracts *ContractsSession) BeneficiaryTokens(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Contracts.Contract.BeneficiaryTokens(&_Contracts.CallOpts, arg0, arg1)
}

// BeneficiaryTokens is a free data retrieval call binding the contract method 0x47a9d145.
//
// Solidity: function beneficiaryTokens(address , uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) BeneficiaryTokens(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Contracts.Contract.BeneficiaryTokens(&_Contracts.CallOpts, arg0, arg1)
}

// GetAllTokensByBeneficiary is a free data retrieval call binding the contract method 0x6d18450a.
//
// Solidity: function getAllTokensByBeneficiary(address _beneficiary) view returns(address[])
func (_Contracts *ContractsCaller) GetAllTokensByBeneficiary(opts *bind.CallOpts, _beneficiary common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAllTokensByBeneficiary", _beneficiary)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllTokensByBeneficiary is a free data retrieval call binding the contract method 0x6d18450a.
//
// Solidity: function getAllTokensByBeneficiary(address _beneficiary) view returns(address[])
func (_Contracts *ContractsSession) GetAllTokensByBeneficiary(_beneficiary common.Address) ([]common.Address, error) {
	return _Contracts.Contract.GetAllTokensByBeneficiary(&_Contracts.CallOpts, _beneficiary)
}

// GetAllTokensByBeneficiary is a free data retrieval call binding the contract method 0x6d18450a.
//
// Solidity: function getAllTokensByBeneficiary(address _beneficiary) view returns(address[])
func (_Contracts *ContractsCallerSession) GetAllTokensByBeneficiary(_beneficiary common.Address) ([]common.Address, error) {
	return _Contracts.Contract.GetAllTokensByBeneficiary(&_Contracts.CallOpts, _beneficiary)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xfb897ce4.
//
// Solidity: function getVestingInfo(address _token) view returns((address,uint256,uint256,uint256,uint256))
func (_Contracts *ContractsCaller) GetVestingInfo(opts *bind.CallOpts, _token common.Address) (Struct0, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getVestingInfo", _token)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetVestingInfo is a free data retrieval call binding the contract method 0xfb897ce4.
//
// Solidity: function getVestingInfo(address _token) view returns((address,uint256,uint256,uint256,uint256))
func (_Contracts *ContractsSession) GetVestingInfo(_token common.Address) (Struct0, error) {
	return _Contracts.Contract.GetVestingInfo(&_Contracts.CallOpts, _token)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0xfb897ce4.
//
// Solidity: function getVestingInfo(address _token) view returns((address,uint256,uint256,uint256,uint256))
func (_Contracts *ContractsCallerSession) GetVestingInfo(_token common.Address) (Struct0, error) {
	return _Contracts.Contract.GetVestingInfo(&_Contracts.CallOpts, _token)
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

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(address _token) view returns(uint256)
func (_Contracts *ContractsCaller) ReleaseableAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "releaseableAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(address _token) view returns(uint256)
func (_Contracts *ContractsSession) ReleaseableAmount(_token common.Address) (*big.Int, error) {
	return _Contracts.Contract.ReleaseableAmount(&_Contracts.CallOpts, _token)
}

// ReleaseableAmount is a free data retrieval call binding the contract method 0x42e16161.
//
// Solidity: function releaseableAmount(address _token) view returns(uint256)
func (_Contracts *ContractsCallerSession) ReleaseableAmount(_token common.Address) (*big.Int, error) {
	return _Contracts.Contract.ReleaseableAmount(&_Contracts.CallOpts, _token)
}

// ReleaseableAmountByAddress is a free data retrieval call binding the contract method 0x4020925a.
//
// Solidity: function releaseableAmountByAddress(address _token, address _beneficiary) view returns(uint256)
func (_Contracts *ContractsCaller) ReleaseableAmountByAddress(opts *bind.CallOpts, _token common.Address, _beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "releaseableAmountByAddress", _token, _beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseableAmountByAddress is a free data retrieval call binding the contract method 0x4020925a.
//
// Solidity: function releaseableAmountByAddress(address _token, address _beneficiary) view returns(uint256)
func (_Contracts *ContractsSession) ReleaseableAmountByAddress(_token common.Address, _beneficiary common.Address) (*big.Int, error) {
	return _Contracts.Contract.ReleaseableAmountByAddress(&_Contracts.CallOpts, _token, _beneficiary)
}

// ReleaseableAmountByAddress is a free data retrieval call binding the contract method 0x4020925a.
//
// Solidity: function releaseableAmountByAddress(address _token, address _beneficiary) view returns(uint256)
func (_Contracts *ContractsCallerSession) ReleaseableAmountByAddress(_token common.Address, _beneficiary common.Address) (*big.Int, error) {
	return _Contracts.Contract.ReleaseableAmountByAddress(&_Contracts.CallOpts, _token, _beneficiary)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(address _token) view returns(uint256)
func (_Contracts *ContractsCaller) VestedAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "vestedAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(address _token) view returns(uint256)
func (_Contracts *ContractsSession) VestedAmount(_token common.Address) (*big.Int, error) {
	return _Contracts.Contract.VestedAmount(&_Contracts.CallOpts, _token)
}

// VestedAmount is a free data retrieval call binding the contract method 0x384711cc.
//
// Solidity: function vestedAmount(address _token) view returns(uint256)
func (_Contracts *ContractsCallerSession) VestedAmount(_token common.Address) (*big.Int, error) {
	return _Contracts.Contract.VestedAmount(&_Contracts.CallOpts, _token)
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo(address ) view returns(address vestingBeneficiary, uint256 totalBalance, uint256 beneficiariesCount, uint256 start, uint256 stop)
func (_Contracts *ContractsCaller) VestingInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "vestingInfo", arg0)

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
func (_Contracts *ContractsSession) VestingInfo(arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}, error) {
	return _Contracts.Contract.VestingInfo(&_Contracts.CallOpts, arg0)
}

// VestingInfo is a free data retrieval call binding the contract method 0xf78e633d.
//
// Solidity: function vestingInfo(address ) view returns(address vestingBeneficiary, uint256 totalBalance, uint256 beneficiariesCount, uint256 start, uint256 stop)
func (_Contracts *ContractsCallerSession) VestingInfo(arg0 common.Address) (struct {
	VestingBeneficiary common.Address
	TotalBalance       *big.Int
	BeneficiariesCount *big.Int
	Start              *big.Int
	Stop               *big.Int
}, error) {
	return _Contracts.Contract.VestingInfo(&_Contracts.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xc25e888f.
//
// Solidity: function addToken(address _token, address[3] _beneficiaries, uint256[3] _proportions, uint256 _vestingPeriodInDays) returns()
func (_Contracts *ContractsTransactor) AddToken(opts *bind.TransactOpts, _token common.Address, _beneficiaries [3]common.Address, _proportions [3]*big.Int, _vestingPeriodInDays *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addToken", _token, _beneficiaries, _proportions, _vestingPeriodInDays)
}

// AddToken is a paid mutator transaction binding the contract method 0xc25e888f.
//
// Solidity: function addToken(address _token, address[3] _beneficiaries, uint256[3] _proportions, uint256 _vestingPeriodInDays) returns()
func (_Contracts *ContractsSession) AddToken(_token common.Address, _beneficiaries [3]common.Address, _proportions [3]*big.Int, _vestingPeriodInDays *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddToken(&_Contracts.TransactOpts, _token, _beneficiaries, _proportions, _vestingPeriodInDays)
}

// AddToken is a paid mutator transaction binding the contract method 0xc25e888f.
//
// Solidity: function addToken(address _token, address[3] _beneficiaries, uint256[3] _proportions, uint256 _vestingPeriodInDays) returns()
func (_Contracts *ContractsTransactorSession) AddToken(_token common.Address, _beneficiaries [3]common.Address, _proportions [3]*big.Int, _vestingPeriodInDays *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddToken(&_Contracts.TransactOpts, _token, _beneficiaries, _proportions, _vestingPeriodInDays)
}

// Release is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address _token, address _beneficiary) returns()
func (_Contracts *ContractsTransactor) Release(opts *bind.TransactOpts, _token common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "release", _token, _beneficiary)
}

// Release is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address _token, address _beneficiary) returns()
func (_Contracts *ContractsSession) Release(_token common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Release(&_Contracts.TransactOpts, _token, _beneficiary)
}

// Release is a paid mutator transaction binding the contract method 0x48b75044.
//
// Solidity: function release(address _token, address _beneficiary) returns()
func (_Contracts *ContractsTransactorSession) Release(_token common.Address, _beneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Release(&_Contracts.TransactOpts, _token, _beneficiary)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address _beneficiary) returns()
func (_Contracts *ContractsTransactor) ReleaseAll(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "releaseAll", _beneficiary)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address _beneficiary) returns()
func (_Contracts *ContractsSession) ReleaseAll(_beneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ReleaseAll(&_Contracts.TransactOpts, _beneficiary)
}

// ReleaseAll is a paid mutator transaction binding the contract method 0x580fc80a.
//
// Solidity: function releaseAll(address _beneficiary) returns()
func (_Contracts *ContractsTransactorSession) ReleaseAll(_beneficiary common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ReleaseAll(&_Contracts.TransactOpts, _beneficiary)
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

// SetSablier is a paid mutator transaction binding the contract method 0x90096556.
//
// Solidity: function setSablier(address _sablier) returns()
func (_Contracts *ContractsTransactor) SetSablier(opts *bind.TransactOpts, _sablier common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSablier", _sablier)
}

// SetSablier is a paid mutator transaction binding the contract method 0x90096556.
//
// Solidity: function setSablier(address _sablier) returns()
func (_Contracts *ContractsSession) SetSablier(_sablier common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSablier(&_Contracts.TransactOpts, _sablier)
}

// SetSablier is a paid mutator transaction binding the contract method 0x90096556.
//
// Solidity: function setSablier(address _sablier) returns()
func (_Contracts *ContractsTransactorSession) SetSablier(_sablier common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetSablier(&_Contracts.TransactOpts, _sablier)
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

// ContractsLogBeneficiaryUpdatedIterator is returned from FilterLogBeneficiaryUpdated and is used to iterate over the raw logs and unpacked data for LogBeneficiaryUpdated events raised by the Contracts contract.
type ContractsLogBeneficiaryUpdatedIterator struct {
	Event *ContractsLogBeneficiaryUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsLogBeneficiaryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogBeneficiaryUpdated)
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
		it.Event = new(ContractsLogBeneficiaryUpdated)
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
func (it *ContractsLogBeneficiaryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogBeneficiaryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogBeneficiaryUpdated represents a LogBeneficiaryUpdated event raised by the Contracts contract.
type ContractsLogBeneficiaryUpdated struct {
	Token              common.Address
	VestingBeneficiary common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogBeneficiaryUpdated is a free log retrieval operation binding the contract event 0x276dd15f31f6b99d62254388504662bee969db7d6766172c0ff60f8576f09453.
//
// Solidity: event LogBeneficiaryUpdated(address indexed token, address vestingBeneficiary)
func (_Contracts *ContractsFilterer) FilterLogBeneficiaryUpdated(opts *bind.FilterOpts, token []common.Address) (*ContractsLogBeneficiaryUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogBeneficiaryUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogBeneficiaryUpdatedIterator{contract: _Contracts.contract, event: "LogBeneficiaryUpdated", logs: logs, sub: sub}, nil
}

// WatchLogBeneficiaryUpdated is a free log subscription operation binding the contract event 0x276dd15f31f6b99d62254388504662bee969db7d6766172c0ff60f8576f09453.
//
// Solidity: event LogBeneficiaryUpdated(address indexed token, address vestingBeneficiary)
func (_Contracts *ContractsFilterer) WatchLogBeneficiaryUpdated(opts *bind.WatchOpts, sink chan<- *ContractsLogBeneficiaryUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogBeneficiaryUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogBeneficiaryUpdated)
				if err := _Contracts.contract.UnpackLog(event, "LogBeneficiaryUpdated", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLogBeneficiaryUpdated(log types.Log) (*ContractsLogBeneficiaryUpdated, error) {
	event := new(ContractsLogBeneficiaryUpdated)
	if err := _Contracts.contract.UnpackLog(event, "LogBeneficiaryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogTokenAddedIterator is returned from FilterLogTokenAdded and is used to iterate over the raw logs and unpacked data for LogTokenAdded events raised by the Contracts contract.
type ContractsLogTokenAddedIterator struct {
	Event *ContractsLogTokenAdded // Event containing the contract specifics and raw log

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
func (it *ContractsLogTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogTokenAdded)
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
		it.Event = new(ContractsLogTokenAdded)
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
func (it *ContractsLogTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogTokenAdded represents a LogTokenAdded event raised by the Contracts contract.
type ContractsLogTokenAdded struct {
	Token               common.Address
	VestingBeneficiary  common.Address
	VestingPeriodInDays *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLogTokenAdded is a free log retrieval operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: event LogTokenAdded(address indexed token, address vestingBeneficiary, uint256 vestingPeriodInDays)
func (_Contracts *ContractsFilterer) FilterLogTokenAdded(opts *bind.FilterOpts, token []common.Address) (*ContractsLogTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogTokenAddedIterator{contract: _Contracts.contract, event: "LogTokenAdded", logs: logs, sub: sub}, nil
}

// WatchLogTokenAdded is a free log subscription operation binding the contract event 0x83ded1c585a48e261782bbbc8a77aeb32caff94dc4803fc6c8094721165a5490.
//
// Solidity: event LogTokenAdded(address indexed token, address vestingBeneficiary, uint256 vestingPeriodInDays)
func (_Contracts *ContractsFilterer) WatchLogTokenAdded(opts *bind.WatchOpts, sink chan<- *ContractsLogTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogTokenAdded)
				if err := _Contracts.contract.UnpackLog(event, "LogTokenAdded", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLogTokenAdded(log types.Log) (*ContractsLogTokenAdded, error) {
	event := new(ContractsLogTokenAdded)
	if err := _Contracts.contract.UnpackLog(event, "LogTokenAdded", log); err != nil {
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

// ContractsReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the Contracts contract.
type ContractsReleasedIterator struct {
	Event *ContractsReleased // Event containing the contract specifics and raw log

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
func (it *ContractsReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsReleased)
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
		it.Event = new(ContractsReleased)
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
func (it *ContractsReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsReleased represents a Released event raised by the Contracts contract.
type ContractsReleased struct {
	Token              common.Address
	VestingBeneficiary common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed token, address vestingBeneficiary, uint256 amount)
func (_Contracts *ContractsFilterer) FilterReleased(opts *bind.FilterOpts, token []common.Address) (*ContractsReleasedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ContractsReleasedIterator{contract: _Contracts.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed token, address vestingBeneficiary, uint256 amount)
func (_Contracts *ContractsFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *ContractsReleased, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Released", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsReleased)
				if err := _Contracts.contract.UnpackLog(event, "Released", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseReleased(log types.Log) (*ContractsReleased, error) {
	event := new(ContractsReleased)
	if err := _Contracts.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
