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
	Token                     common.Address
	Name                      string
	Symbol                    string
	Decimals                  uint8
	TotalSupply               *big.Int
	Proposer                  common.Address
	VestingBeneficiary        common.Address
	InitialPercentage         uint8
	VestingPeriodInDays       *big.Int
	Approved                  bool
	InitialPlatformPercentage uint8
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Referral           common.Address
	ReferralPercentage uint8
}

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalApprove\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"address\",\"name\":\"tokenAddress\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalImported\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"string\",\"name\":\"symbol\",\"indexed\":false},{\"type\":\"address\",\"name\":\"proposer\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"hashIndex\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalReferralSubmit\",\"inputs\":[{\"type\":\"address\",\"name\":\"referral\",\"indexed\":false},{\"type\":\"uint8\",\"name\":\"referralPercentage\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"hashIndex\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalSubmit\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"string\",\"name\":\"symbol\",\"indexed\":false},{\"type\":\"address\",\"name\":\"proposer\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"hashIndex\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"approveProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"creatorReferral\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"referral\"},{\"type\":\"uint8\",\"name\":\"referralPercentage\"}]},{\"type\":\"function\",\"name\":\"getCreatorByIndex\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getCreatorByName\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getCreatorBySymbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_symbol\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getCreatorReferralByIndex\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"referral\"},{\"type\":\"uint8\",\"name\":\"referralPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getIndexByName\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"}],\"outputs\":[{\"type\":\"bytes32\"}]},{\"type\":\"function\",\"name\":\"getIndexBySymbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_symbol\"}],\"outputs\":[{\"type\":\"bytes32\"}]},{\"type\":\"function\",\"name\":\"importByIndex\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_oldRegistry\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"rolodex\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]},{\"type\":\"function\",\"name\":\"submitProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"address\",\"name\":\"_proposer\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"hashIndex\"}]},{\"type\":\"function\",\"name\":\"submitProposalReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_referral\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]}]"

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// CreatorReferral is a free data retrieval call binding the contract method 0xdd79506d.
//
// Solidity: function creatorReferral(bytes32 ) view returns(address referral, uint8 referralPercentage)
func (_Registry *RegistryCaller) CreatorReferral(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Referral           common.Address
	ReferralPercentage uint8
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "creatorReferral", arg0)

	outstruct := new(struct {
		Referral           common.Address
		ReferralPercentage uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Referral = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ReferralPercentage = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// CreatorReferral is a free data retrieval call binding the contract method 0xdd79506d.
//
// Solidity: function creatorReferral(bytes32 ) view returns(address referral, uint8 referralPercentage)
func (_Registry *RegistrySession) CreatorReferral(arg0 [32]byte) (struct {
	Referral           common.Address
	ReferralPercentage uint8
}, error) {
	return _Registry.Contract.CreatorReferral(&_Registry.CallOpts, arg0)
}

// CreatorReferral is a free data retrieval call binding the contract method 0xdd79506d.
//
// Solidity: function creatorReferral(bytes32 ) view returns(address referral, uint8 referralPercentage)
func (_Registry *RegistryCallerSession) CreatorReferral(arg0 [32]byte) (struct {
	Referral           common.Address
	ReferralPercentage uint8
}, error) {
	return _Registry.Contract.CreatorReferral(&_Registry.CallOpts, arg0)
}

// GetCreatorByIndex is a free data retrieval call binding the contract method 0x969033a4.
//
// Solidity: function getCreatorByIndex(bytes32 _hashIndex) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistryCaller) GetCreatorByIndex(opts *bind.CallOpts, _hashIndex [32]byte) (Struct0, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getCreatorByIndex", _hashIndex)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetCreatorByIndex is a free data retrieval call binding the contract method 0x969033a4.
//
// Solidity: function getCreatorByIndex(bytes32 _hashIndex) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistrySession) GetCreatorByIndex(_hashIndex [32]byte) (Struct0, error) {
	return _Registry.Contract.GetCreatorByIndex(&_Registry.CallOpts, _hashIndex)
}

// GetCreatorByIndex is a free data retrieval call binding the contract method 0x969033a4.
//
// Solidity: function getCreatorByIndex(bytes32 _hashIndex) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistryCallerSession) GetCreatorByIndex(_hashIndex [32]byte) (Struct0, error) {
	return _Registry.Contract.GetCreatorByIndex(&_Registry.CallOpts, _hashIndex)
}

// GetCreatorByName is a free data retrieval call binding the contract method 0x0aed6660.
//
// Solidity: function getCreatorByName(string _name) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistryCaller) GetCreatorByName(opts *bind.CallOpts, _name string) (Struct0, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getCreatorByName", _name)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetCreatorByName is a free data retrieval call binding the contract method 0x0aed6660.
//
// Solidity: function getCreatorByName(string _name) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistrySession) GetCreatorByName(_name string) (Struct0, error) {
	return _Registry.Contract.GetCreatorByName(&_Registry.CallOpts, _name)
}

// GetCreatorByName is a free data retrieval call binding the contract method 0x0aed6660.
//
// Solidity: function getCreatorByName(string _name) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistryCallerSession) GetCreatorByName(_name string) (Struct0, error) {
	return _Registry.Contract.GetCreatorByName(&_Registry.CallOpts, _name)
}

// GetCreatorBySymbol is a free data retrieval call binding the contract method 0x775e3477.
//
// Solidity: function getCreatorBySymbol(string _symbol) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistryCaller) GetCreatorBySymbol(opts *bind.CallOpts, _symbol string) (Struct0, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getCreatorBySymbol", _symbol)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetCreatorBySymbol is a free data retrieval call binding the contract method 0x775e3477.
//
// Solidity: function getCreatorBySymbol(string _symbol) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistrySession) GetCreatorBySymbol(_symbol string) (Struct0, error) {
	return _Registry.Contract.GetCreatorBySymbol(&_Registry.CallOpts, _symbol)
}

// GetCreatorBySymbol is a free data retrieval call binding the contract method 0x775e3477.
//
// Solidity: function getCreatorBySymbol(string _symbol) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Registry *RegistryCallerSession) GetCreatorBySymbol(_symbol string) (Struct0, error) {
	return _Registry.Contract.GetCreatorBySymbol(&_Registry.CallOpts, _symbol)
}

// GetCreatorReferralByIndex is a free data retrieval call binding the contract method 0x05e1ac7b.
//
// Solidity: function getCreatorReferralByIndex(bytes32 _hashIndex) view returns((address,uint8))
func (_Registry *RegistryCaller) GetCreatorReferralByIndex(opts *bind.CallOpts, _hashIndex [32]byte) (Struct1, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getCreatorReferralByIndex", _hashIndex)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetCreatorReferralByIndex is a free data retrieval call binding the contract method 0x05e1ac7b.
//
// Solidity: function getCreatorReferralByIndex(bytes32 _hashIndex) view returns((address,uint8))
func (_Registry *RegistrySession) GetCreatorReferralByIndex(_hashIndex [32]byte) (Struct1, error) {
	return _Registry.Contract.GetCreatorReferralByIndex(&_Registry.CallOpts, _hashIndex)
}

// GetCreatorReferralByIndex is a free data retrieval call binding the contract method 0x05e1ac7b.
//
// Solidity: function getCreatorReferralByIndex(bytes32 _hashIndex) view returns((address,uint8))
func (_Registry *RegistryCallerSession) GetCreatorReferralByIndex(_hashIndex [32]byte) (Struct1, error) {
	return _Registry.Contract.GetCreatorReferralByIndex(&_Registry.CallOpts, _hashIndex)
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(string _name) view returns(bytes32)
func (_Registry *RegistryCaller) GetIndexByName(opts *bind.CallOpts, _name string) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getIndexByName", _name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(string _name) view returns(bytes32)
func (_Registry *RegistrySession) GetIndexByName(_name string) ([32]byte, error) {
	return _Registry.Contract.GetIndexByName(&_Registry.CallOpts, _name)
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(string _name) view returns(bytes32)
func (_Registry *RegistryCallerSession) GetIndexByName(_name string) ([32]byte, error) {
	return _Registry.Contract.GetIndexByName(&_Registry.CallOpts, _name)
}

// GetIndexBySymbol is a free data retrieval call binding the contract method 0x9515d3ea.
//
// Solidity: function getIndexBySymbol(string _symbol) view returns(bytes32)
func (_Registry *RegistryCaller) GetIndexBySymbol(opts *bind.CallOpts, _symbol string) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getIndexBySymbol", _symbol)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetIndexBySymbol is a free data retrieval call binding the contract method 0x9515d3ea.
//
// Solidity: function getIndexBySymbol(string _symbol) view returns(bytes32)
func (_Registry *RegistrySession) GetIndexBySymbol(_symbol string) ([32]byte, error) {
	return _Registry.Contract.GetIndexBySymbol(&_Registry.CallOpts, _symbol)
}

// GetIndexBySymbol is a free data retrieval call binding the contract method 0x9515d3ea.
//
// Solidity: function getIndexBySymbol(string _symbol) view returns(bytes32)
func (_Registry *RegistryCallerSession) GetIndexBySymbol(_symbol string) ([32]byte, error) {
	return _Registry.Contract.GetIndexBySymbol(&_Registry.CallOpts, _symbol)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex(bytes32 ) view returns(address token, string name, string symbol, uint8 decimals, uint256 totalSupply, address proposer, address vestingBeneficiary, uint8 initialPercentage, uint256 vestingPeriodInDays, bool approved, uint8 initialPlatformPercentage)
func (_Registry *RegistryCaller) Rolodex(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Token                     common.Address
	Name                      string
	Symbol                    string
	Decimals                  uint8
	TotalSupply               *big.Int
	Proposer                  common.Address
	VestingBeneficiary        common.Address
	InitialPercentage         uint8
	VestingPeriodInDays       *big.Int
	Approved                  bool
	InitialPlatformPercentage uint8
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "rolodex", arg0)

	outstruct := new(struct {
		Token                     common.Address
		Name                      string
		Symbol                    string
		Decimals                  uint8
		TotalSupply               *big.Int
		Proposer                  common.Address
		VestingBeneficiary        common.Address
		InitialPercentage         uint8
		VestingPeriodInDays       *big.Int
		Approved                  bool
		InitialPlatformPercentage uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.TotalSupply = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Proposer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.VestingBeneficiary = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.InitialPercentage = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.VestingPeriodInDays = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Approved = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.InitialPlatformPercentage = *abi.ConvertType(out[10], new(uint8)).(*uint8)

	return *outstruct, err

}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex(bytes32 ) view returns(address token, string name, string symbol, uint8 decimals, uint256 totalSupply, address proposer, address vestingBeneficiary, uint8 initialPercentage, uint256 vestingPeriodInDays, bool approved, uint8 initialPlatformPercentage)
func (_Registry *RegistrySession) Rolodex(arg0 [32]byte) (struct {
	Token                     common.Address
	Name                      string
	Symbol                    string
	Decimals                  uint8
	TotalSupply               *big.Int
	Proposer                  common.Address
	VestingBeneficiary        common.Address
	InitialPercentage         uint8
	VestingPeriodInDays       *big.Int
	Approved                  bool
	InitialPlatformPercentage uint8
}, error) {
	return _Registry.Contract.Rolodex(&_Registry.CallOpts, arg0)
}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex(bytes32 ) view returns(address token, string name, string symbol, uint8 decimals, uint256 totalSupply, address proposer, address vestingBeneficiary, uint8 initialPercentage, uint256 vestingPeriodInDays, bool approved, uint8 initialPlatformPercentage)
func (_Registry *RegistryCallerSession) Rolodex(arg0 [32]byte) (struct {
	Token                     common.Address
	Name                      string
	Symbol                    string
	Decimals                  uint8
	TotalSupply               *big.Int
	Proposer                  common.Address
	VestingBeneficiary        common.Address
	InitialPercentage         uint8
	VestingPeriodInDays       *big.Int
	Approved                  bool
	InitialPlatformPercentage uint8
}, error) {
	return _Registry.Contract.Rolodex(&_Registry.CallOpts, arg0)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(bytes32 _hashIndex, address _token) returns(bool)
func (_Registry *RegistryTransactor) ApproveProposal(opts *bind.TransactOpts, _hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "approveProposal", _hashIndex, _token)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(bytes32 _hashIndex, address _token) returns(bool)
func (_Registry *RegistrySession) ApproveProposal(_hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ApproveProposal(&_Registry.TransactOpts, _hashIndex, _token)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(bytes32 _hashIndex, address _token) returns(bool)
func (_Registry *RegistryTransactorSession) ApproveProposal(_hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ApproveProposal(&_Registry.TransactOpts, _hashIndex, _token)
}

// ImportByIndex is a paid mutator transaction binding the contract method 0x397a88a1.
//
// Solidity: function importByIndex(bytes32 _hashIndex, address _oldRegistry) returns()
func (_Registry *RegistryTransactor) ImportByIndex(opts *bind.TransactOpts, _hashIndex [32]byte, _oldRegistry common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "importByIndex", _hashIndex, _oldRegistry)
}

// ImportByIndex is a paid mutator transaction binding the contract method 0x397a88a1.
//
// Solidity: function importByIndex(bytes32 _hashIndex, address _oldRegistry) returns()
func (_Registry *RegistrySession) ImportByIndex(_hashIndex [32]byte, _oldRegistry common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ImportByIndex(&_Registry.TransactOpts, _hashIndex, _oldRegistry)
}

// ImportByIndex is a paid mutator transaction binding the contract method 0x397a88a1.
//
// Solidity: function importByIndex(bytes32 _hashIndex, address _oldRegistry) returns()
func (_Registry *RegistryTransactorSession) ImportByIndex(_hashIndex [32]byte, _oldRegistry common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ImportByIndex(&_Registry.TransactOpts, _hashIndex, _oldRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Registry *RegistrySession) Initialize() (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Registry *RegistryTransactorSession) Initialize() (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd490cb74.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, address _proposer, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Registry *RegistryTransactor) SubmitProposal(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _proposer common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "submitProposal", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _proposer, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd490cb74.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, address _proposer, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Registry *RegistrySession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _proposer common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Registry.Contract.SubmitProposal(&_Registry.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _proposer, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd490cb74.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, address _proposer, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Registry *RegistryTransactorSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _proposer common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Registry.Contract.SubmitProposal(&_Registry.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _proposer, _initialPlatformPercentage)
}

// SubmitProposalReferral is a paid mutator transaction binding the contract method 0xb13873f0.
//
// Solidity: function submitProposalReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Registry *RegistryTransactor) SubmitProposalReferral(opts *bind.TransactOpts, _hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "submitProposalReferral", _hashIndex, _referral, _referralPercentage)
}

// SubmitProposalReferral is a paid mutator transaction binding the contract method 0xb13873f0.
//
// Solidity: function submitProposalReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Registry *RegistrySession) SubmitProposalReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Registry.Contract.SubmitProposalReferral(&_Registry.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// SubmitProposalReferral is a paid mutator transaction binding the contract method 0xb13873f0.
//
// Solidity: function submitProposalReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Registry *RegistryTransactorSession) SubmitProposalReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Registry.Contract.SubmitProposalReferral(&_Registry.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryLogProposalApproveIterator is returned from FilterLogProposalApprove and is used to iterate over the raw logs and unpacked data for LogProposalApprove events raised by the Registry contract.
type RegistryLogProposalApproveIterator struct {
	Event *RegistryLogProposalApprove // Event containing the contract specifics and raw log

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
func (it *RegistryLogProposalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLogProposalApprove)
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
		it.Event = new(RegistryLogProposalApprove)
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
func (it *RegistryLogProposalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLogProposalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLogProposalApprove represents a LogProposalApprove event raised by the Registry contract.
type RegistryLogProposalApprove struct {
	Name         string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogProposalApprove is a free log retrieval operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: event LogProposalApprove(string name, address indexed tokenAddress)
func (_Registry *RegistryFilterer) FilterLogProposalApprove(opts *bind.FilterOpts, tokenAddress []common.Address) (*RegistryLogProposalApproveIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LogProposalApprove", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLogProposalApproveIterator{contract: _Registry.contract, event: "LogProposalApprove", logs: logs, sub: sub}, nil
}

// WatchLogProposalApprove is a free log subscription operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: event LogProposalApprove(string name, address indexed tokenAddress)
func (_Registry *RegistryFilterer) WatchLogProposalApprove(opts *bind.WatchOpts, sink chan<- *RegistryLogProposalApprove, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LogProposalApprove", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLogProposalApprove)
				if err := _Registry.contract.UnpackLog(event, "LogProposalApprove", log); err != nil {
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

// ParseLogProposalApprove is a log parse operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: event LogProposalApprove(string name, address indexed tokenAddress)
func (_Registry *RegistryFilterer) ParseLogProposalApprove(log types.Log) (*RegistryLogProposalApprove, error) {
	event := new(RegistryLogProposalApprove)
	if err := _Registry.contract.UnpackLog(event, "LogProposalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryLogProposalImportedIterator is returned from FilterLogProposalImported and is used to iterate over the raw logs and unpacked data for LogProposalImported events raised by the Registry contract.
type RegistryLogProposalImportedIterator struct {
	Event *RegistryLogProposalImported // Event containing the contract specifics and raw log

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
func (it *RegistryLogProposalImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLogProposalImported)
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
		it.Event = new(RegistryLogProposalImported)
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
func (it *RegistryLogProposalImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLogProposalImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLogProposalImported represents a LogProposalImported event raised by the Registry contract.
type RegistryLogProposalImported struct {
	Name      string
	Symbol    string
	Proposer  common.Address
	HashIndex [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProposalImported is a free log retrieval operation binding the contract event 0x82c68a30819d3b8f0c181d716ed3fafa8551cad68eefe911f471516e958421b3.
//
// Solidity: event LogProposalImported(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) FilterLogProposalImported(opts *bind.FilterOpts, hashIndex [][32]byte) (*RegistryLogProposalImportedIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LogProposalImported", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLogProposalImportedIterator{contract: _Registry.contract, event: "LogProposalImported", logs: logs, sub: sub}, nil
}

// WatchLogProposalImported is a free log subscription operation binding the contract event 0x82c68a30819d3b8f0c181d716ed3fafa8551cad68eefe911f471516e958421b3.
//
// Solidity: event LogProposalImported(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) WatchLogProposalImported(opts *bind.WatchOpts, sink chan<- *RegistryLogProposalImported, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LogProposalImported", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLogProposalImported)
				if err := _Registry.contract.UnpackLog(event, "LogProposalImported", log); err != nil {
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

// ParseLogProposalImported is a log parse operation binding the contract event 0x82c68a30819d3b8f0c181d716ed3fafa8551cad68eefe911f471516e958421b3.
//
// Solidity: event LogProposalImported(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) ParseLogProposalImported(log types.Log) (*RegistryLogProposalImported, error) {
	event := new(RegistryLogProposalImported)
	if err := _Registry.contract.UnpackLog(event, "LogProposalImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryLogProposalReferralSubmitIterator is returned from FilterLogProposalReferralSubmit and is used to iterate over the raw logs and unpacked data for LogProposalReferralSubmit events raised by the Registry contract.
type RegistryLogProposalReferralSubmitIterator struct {
	Event *RegistryLogProposalReferralSubmit // Event containing the contract specifics and raw log

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
func (it *RegistryLogProposalReferralSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLogProposalReferralSubmit)
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
		it.Event = new(RegistryLogProposalReferralSubmit)
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
func (it *RegistryLogProposalReferralSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLogProposalReferralSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLogProposalReferralSubmit represents a LogProposalReferralSubmit event raised by the Registry contract.
type RegistryLogProposalReferralSubmit struct {
	Referral           common.Address
	ReferralPercentage uint8
	HashIndex          [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogProposalReferralSubmit is a free log retrieval operation binding the contract event 0x924a2d7699539105c99283e511a2a67d57aecd39d59697c4a76cb2187fddb6ce.
//
// Solidity: event LogProposalReferralSubmit(address referral, uint8 referralPercentage, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) FilterLogProposalReferralSubmit(opts *bind.FilterOpts, hashIndex [][32]byte) (*RegistryLogProposalReferralSubmitIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LogProposalReferralSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLogProposalReferralSubmitIterator{contract: _Registry.contract, event: "LogProposalReferralSubmit", logs: logs, sub: sub}, nil
}

// WatchLogProposalReferralSubmit is a free log subscription operation binding the contract event 0x924a2d7699539105c99283e511a2a67d57aecd39d59697c4a76cb2187fddb6ce.
//
// Solidity: event LogProposalReferralSubmit(address referral, uint8 referralPercentage, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) WatchLogProposalReferralSubmit(opts *bind.WatchOpts, sink chan<- *RegistryLogProposalReferralSubmit, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LogProposalReferralSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLogProposalReferralSubmit)
				if err := _Registry.contract.UnpackLog(event, "LogProposalReferralSubmit", log); err != nil {
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

// ParseLogProposalReferralSubmit is a log parse operation binding the contract event 0x924a2d7699539105c99283e511a2a67d57aecd39d59697c4a76cb2187fddb6ce.
//
// Solidity: event LogProposalReferralSubmit(address referral, uint8 referralPercentage, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) ParseLogProposalReferralSubmit(log types.Log) (*RegistryLogProposalReferralSubmit, error) {
	event := new(RegistryLogProposalReferralSubmit)
	if err := _Registry.contract.UnpackLog(event, "LogProposalReferralSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryLogProposalSubmitIterator is returned from FilterLogProposalSubmit and is used to iterate over the raw logs and unpacked data for LogProposalSubmit events raised by the Registry contract.
type RegistryLogProposalSubmitIterator struct {
	Event *RegistryLogProposalSubmit // Event containing the contract specifics and raw log

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
func (it *RegistryLogProposalSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryLogProposalSubmit)
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
		it.Event = new(RegistryLogProposalSubmit)
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
func (it *RegistryLogProposalSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryLogProposalSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryLogProposalSubmit represents a LogProposalSubmit event raised by the Registry contract.
type RegistryLogProposalSubmit struct {
	Name      string
	Symbol    string
	Proposer  common.Address
	HashIndex [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProposalSubmit is a free log retrieval operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: event LogProposalSubmit(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) FilterLogProposalSubmit(opts *bind.FilterOpts, hashIndex [][32]byte) (*RegistryLogProposalSubmitIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "LogProposalSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &RegistryLogProposalSubmitIterator{contract: _Registry.contract, event: "LogProposalSubmit", logs: logs, sub: sub}, nil
}

// WatchLogProposalSubmit is a free log subscription operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: event LogProposalSubmit(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) WatchLogProposalSubmit(opts *bind.WatchOpts, sink chan<- *RegistryLogProposalSubmit, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "LogProposalSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryLogProposalSubmit)
				if err := _Registry.contract.UnpackLog(event, "LogProposalSubmit", log); err != nil {
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

// ParseLogProposalSubmit is a log parse operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: event LogProposalSubmit(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Registry *RegistryFilterer) ParseLogProposalSubmit(log types.Log) (*RegistryLogProposalSubmit, error) {
	event := new(RegistryLogProposalSubmit)
	if err := _Registry.contract.UnpackLog(event, "LogProposalSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}