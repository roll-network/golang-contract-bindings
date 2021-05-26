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

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Referral           common.Address
	ReferralPercentage uint8
}

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalApprove\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"address\",\"name\":\"tokenAddress\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalImported\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"string\",\"name\":\"symbol\",\"indexed\":false},{\"type\":\"address\",\"name\":\"proposer\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"hashIndex\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalReferralSubmit\",\"inputs\":[{\"type\":\"address\",\"name\":\"referral\",\"indexed\":false},{\"type\":\"uint8\",\"name\":\"referralPercentage\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"hashIndex\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"LogProposalSubmit\",\"inputs\":[{\"type\":\"string\",\"name\":\"name\",\"indexed\":false},{\"type\":\"string\",\"name\":\"symbol\",\"indexed\":false},{\"type\":\"address\",\"name\":\"proposer\",\"indexed\":false},{\"type\":\"bytes32\",\"name\":\"hashIndex\",\"indexed\":true}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true}]},{\"type\":\"function\",\"name\":\"approveProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_token\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"creatorReferral\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"referral\"},{\"type\":\"uint8\",\"name\":\"referralPercentage\"}]},{\"type\":\"function\",\"name\":\"getCreatorByIndex\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getCreatorByName\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getCreatorBySymbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_symbol\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getCreatorReferralByIndex\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"}],\"outputs\":[{\"type\":\"tuple\",\"components\":[{\"type\":\"address\",\"name\":\"referral\"},{\"type\":\"uint8\",\"name\":\"referralPercentage\"}]}]},{\"type\":\"function\",\"name\":\"getIndexByName\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"}],\"outputs\":[{\"type\":\"bytes32\"}]},{\"type\":\"function\",\"name\":\"getIndexBySymbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_symbol\"}],\"outputs\":[{\"type\":\"bytes32\"}]},{\"type\":\"function\",\"name\":\"importByIndex\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_oldRegistry\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"initialize\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"owner\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"address\"}]},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[],\"outputs\":[]},{\"type\":\"function\",\"name\":\"rolodex\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"bytes32\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"token\"},{\"type\":\"string\",\"name\":\"name\"},{\"type\":\"string\",\"name\":\"symbol\"},{\"type\":\"uint8\",\"name\":\"decimals\"},{\"type\":\"uint256\",\"name\":\"totalSupply\"},{\"type\":\"address\",\"name\":\"proposer\"},{\"type\":\"address\",\"name\":\"vestingBeneficiary\"},{\"type\":\"uint8\",\"name\":\"initialPercentage\"},{\"type\":\"uint256\",\"name\":\"vestingPeriodInDays\"},{\"type\":\"bool\",\"name\":\"approved\"},{\"type\":\"uint8\",\"name\":\"initialPlatformPercentage\"}]},{\"type\":\"function\",\"name\":\"submitProposal\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_totalSupply\"},{\"type\":\"uint8\",\"name\":\"_initialPercentage\"},{\"type\":\"uint256\",\"name\":\"_vestingPeriodInDays\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"address\",\"name\":\"_proposer\"},{\"type\":\"uint8\",\"name\":\"_initialPlatformPercentage\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"hashIndex\"}]},{\"type\":\"function\",\"name\":\"submitProposalReferral\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_hashIndex\"},{\"type\":\"address\",\"name\":\"_referral\"},{\"type\":\"uint8\",\"name\":\"_referralPercentage\"}],\"outputs\":[]},{\"type\":\"function\",\"name\":\"transferOwnership\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\"}],\"outputs\":[]}]"

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

// CreatorReferral is a free data retrieval call binding the contract method 0xdd79506d.
//
// Solidity: function creatorReferral(bytes32 ) view returns(address referral, uint8 referralPercentage)
func (_Contracts *ContractsCaller) CreatorReferral(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Referral           common.Address
	ReferralPercentage uint8
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "creatorReferral", arg0)

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
func (_Contracts *ContractsSession) CreatorReferral(arg0 [32]byte) (struct {
	Referral           common.Address
	ReferralPercentage uint8
}, error) {
	return _Contracts.Contract.CreatorReferral(&_Contracts.CallOpts, arg0)
}

// CreatorReferral is a free data retrieval call binding the contract method 0xdd79506d.
//
// Solidity: function creatorReferral(bytes32 ) view returns(address referral, uint8 referralPercentage)
func (_Contracts *ContractsCallerSession) CreatorReferral(arg0 [32]byte) (struct {
	Referral           common.Address
	ReferralPercentage uint8
}, error) {
	return _Contracts.Contract.CreatorReferral(&_Contracts.CallOpts, arg0)
}

// GetCreatorByIndex is a free data retrieval call binding the contract method 0x969033a4.
//
// Solidity: function getCreatorByIndex(bytes32 _hashIndex) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsCaller) GetCreatorByIndex(opts *bind.CallOpts, _hashIndex [32]byte) (Struct1, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCreatorByIndex", _hashIndex)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetCreatorByIndex is a free data retrieval call binding the contract method 0x969033a4.
//
// Solidity: function getCreatorByIndex(bytes32 _hashIndex) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsSession) GetCreatorByIndex(_hashIndex [32]byte) (Struct1, error) {
	return _Contracts.Contract.GetCreatorByIndex(&_Contracts.CallOpts, _hashIndex)
}

// GetCreatorByIndex is a free data retrieval call binding the contract method 0x969033a4.
//
// Solidity: function getCreatorByIndex(bytes32 _hashIndex) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsCallerSession) GetCreatorByIndex(_hashIndex [32]byte) (Struct1, error) {
	return _Contracts.Contract.GetCreatorByIndex(&_Contracts.CallOpts, _hashIndex)
}

// GetCreatorByName is a free data retrieval call binding the contract method 0x0aed6660.
//
// Solidity: function getCreatorByName(string _name) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsCaller) GetCreatorByName(opts *bind.CallOpts, _name string) (Struct1, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCreatorByName", _name)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetCreatorByName is a free data retrieval call binding the contract method 0x0aed6660.
//
// Solidity: function getCreatorByName(string _name) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsSession) GetCreatorByName(_name string) (Struct1, error) {
	return _Contracts.Contract.GetCreatorByName(&_Contracts.CallOpts, _name)
}

// GetCreatorByName is a free data retrieval call binding the contract method 0x0aed6660.
//
// Solidity: function getCreatorByName(string _name) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsCallerSession) GetCreatorByName(_name string) (Struct1, error) {
	return _Contracts.Contract.GetCreatorByName(&_Contracts.CallOpts, _name)
}

// GetCreatorBySymbol is a free data retrieval call binding the contract method 0x775e3477.
//
// Solidity: function getCreatorBySymbol(string _symbol) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsCaller) GetCreatorBySymbol(opts *bind.CallOpts, _symbol string) (Struct1, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCreatorBySymbol", _symbol)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetCreatorBySymbol is a free data retrieval call binding the contract method 0x775e3477.
//
// Solidity: function getCreatorBySymbol(string _symbol) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsSession) GetCreatorBySymbol(_symbol string) (Struct1, error) {
	return _Contracts.Contract.GetCreatorBySymbol(&_Contracts.CallOpts, _symbol)
}

// GetCreatorBySymbol is a free data retrieval call binding the contract method 0x775e3477.
//
// Solidity: function getCreatorBySymbol(string _symbol) view returns((address,string,string,uint8,uint256,address,address,uint8,uint256,bool,uint8))
func (_Contracts *ContractsCallerSession) GetCreatorBySymbol(_symbol string) (Struct1, error) {
	return _Contracts.Contract.GetCreatorBySymbol(&_Contracts.CallOpts, _symbol)
}

// GetCreatorReferralByIndex is a free data retrieval call binding the contract method 0x05e1ac7b.
//
// Solidity: function getCreatorReferralByIndex(bytes32 _hashIndex) view returns((address,uint8))
func (_Contracts *ContractsCaller) GetCreatorReferralByIndex(opts *bind.CallOpts, _hashIndex [32]byte) (Struct0, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getCreatorReferralByIndex", _hashIndex)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetCreatorReferralByIndex is a free data retrieval call binding the contract method 0x05e1ac7b.
//
// Solidity: function getCreatorReferralByIndex(bytes32 _hashIndex) view returns((address,uint8))
func (_Contracts *ContractsSession) GetCreatorReferralByIndex(_hashIndex [32]byte) (Struct0, error) {
	return _Contracts.Contract.GetCreatorReferralByIndex(&_Contracts.CallOpts, _hashIndex)
}

// GetCreatorReferralByIndex is a free data retrieval call binding the contract method 0x05e1ac7b.
//
// Solidity: function getCreatorReferralByIndex(bytes32 _hashIndex) view returns((address,uint8))
func (_Contracts *ContractsCallerSession) GetCreatorReferralByIndex(_hashIndex [32]byte) (Struct0, error) {
	return _Contracts.Contract.GetCreatorReferralByIndex(&_Contracts.CallOpts, _hashIndex)
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(string _name) view returns(bytes32)
func (_Contracts *ContractsCaller) GetIndexByName(opts *bind.CallOpts, _name string) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getIndexByName", _name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(string _name) view returns(bytes32)
func (_Contracts *ContractsSession) GetIndexByName(_name string) ([32]byte, error) {
	return _Contracts.Contract.GetIndexByName(&_Contracts.CallOpts, _name)
}

// GetIndexByName is a free data retrieval call binding the contract method 0xc0af8582.
//
// Solidity: function getIndexByName(string _name) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetIndexByName(_name string) ([32]byte, error) {
	return _Contracts.Contract.GetIndexByName(&_Contracts.CallOpts, _name)
}

// GetIndexBySymbol is a free data retrieval call binding the contract method 0x9515d3ea.
//
// Solidity: function getIndexBySymbol(string _symbol) view returns(bytes32)
func (_Contracts *ContractsCaller) GetIndexBySymbol(opts *bind.CallOpts, _symbol string) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getIndexBySymbol", _symbol)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetIndexBySymbol is a free data retrieval call binding the contract method 0x9515d3ea.
//
// Solidity: function getIndexBySymbol(string _symbol) view returns(bytes32)
func (_Contracts *ContractsSession) GetIndexBySymbol(_symbol string) ([32]byte, error) {
	return _Contracts.Contract.GetIndexBySymbol(&_Contracts.CallOpts, _symbol)
}

// GetIndexBySymbol is a free data retrieval call binding the contract method 0x9515d3ea.
//
// Solidity: function getIndexBySymbol(string _symbol) view returns(bytes32)
func (_Contracts *ContractsCallerSession) GetIndexBySymbol(_symbol string) ([32]byte, error) {
	return _Contracts.Contract.GetIndexBySymbol(&_Contracts.CallOpts, _symbol)
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

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex(bytes32 ) view returns(address token, string name, string symbol, uint8 decimals, uint256 totalSupply, address proposer, address vestingBeneficiary, uint8 initialPercentage, uint256 vestingPeriodInDays, bool approved, uint8 initialPlatformPercentage)
func (_Contracts *ContractsCaller) Rolodex(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Contracts.contract.Call(opts, &out, "rolodex", arg0)

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
func (_Contracts *ContractsSession) Rolodex(arg0 [32]byte) (struct {
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
	return _Contracts.Contract.Rolodex(&_Contracts.CallOpts, arg0)
}

// Rolodex is a free data retrieval call binding the contract method 0x469dfdc1.
//
// Solidity: function rolodex(bytes32 ) view returns(address token, string name, string symbol, uint8 decimals, uint256 totalSupply, address proposer, address vestingBeneficiary, uint8 initialPercentage, uint256 vestingPeriodInDays, bool approved, uint8 initialPlatformPercentage)
func (_Contracts *ContractsCallerSession) Rolodex(arg0 [32]byte) (struct {
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
	return _Contracts.Contract.Rolodex(&_Contracts.CallOpts, arg0)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(bytes32 _hashIndex, address _token) returns(bool)
func (_Contracts *ContractsTransactor) ApproveProposal(opts *bind.TransactOpts, _hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveProposal", _hashIndex, _token)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(bytes32 _hashIndex, address _token) returns(bool)
func (_Contracts *ContractsSession) ApproveProposal(_hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveProposal(&_Contracts.TransactOpts, _hashIndex, _token)
}

// ApproveProposal is a paid mutator transaction binding the contract method 0x26c52914.
//
// Solidity: function approveProposal(bytes32 _hashIndex, address _token) returns(bool)
func (_Contracts *ContractsTransactorSession) ApproveProposal(_hashIndex [32]byte, _token common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveProposal(&_Contracts.TransactOpts, _hashIndex, _token)
}

// ImportByIndex is a paid mutator transaction binding the contract method 0x397a88a1.
//
// Solidity: function importByIndex(bytes32 _hashIndex, address _oldRegistry) returns()
func (_Contracts *ContractsTransactor) ImportByIndex(opts *bind.TransactOpts, _hashIndex [32]byte, _oldRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "importByIndex", _hashIndex, _oldRegistry)
}

// ImportByIndex is a paid mutator transaction binding the contract method 0x397a88a1.
//
// Solidity: function importByIndex(bytes32 _hashIndex, address _oldRegistry) returns()
func (_Contracts *ContractsSession) ImportByIndex(_hashIndex [32]byte, _oldRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ImportByIndex(&_Contracts.TransactOpts, _hashIndex, _oldRegistry)
}

// ImportByIndex is a paid mutator transaction binding the contract method 0x397a88a1.
//
// Solidity: function importByIndex(bytes32 _hashIndex, address _oldRegistry) returns()
func (_Contracts *ContractsTransactorSession) ImportByIndex(_hashIndex [32]byte, _oldRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ImportByIndex(&_Contracts.TransactOpts, _hashIndex, _oldRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contracts *ContractsSession) Initialize() (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contracts *ContractsTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts)
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

// SubmitProposal is a paid mutator transaction binding the contract method 0xd490cb74.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, address _proposer, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Contracts *ContractsTransactor) SubmitProposal(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _proposer common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitProposal", _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _proposer, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd490cb74.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, address _proposer, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Contracts *ContractsSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _proposer common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitProposal(&_Contracts.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _proposer, _initialPlatformPercentage)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0xd490cb74.
//
// Solidity: function submitProposal(string _name, string _symbol, uint8 _decimals, uint256 _totalSupply, uint8 _initialPercentage, uint256 _vestingPeriodInDays, address _vestingBeneficiary, address _proposer, uint8 _initialPlatformPercentage) returns(bytes32 hashIndex)
func (_Contracts *ContractsTransactorSession) SubmitProposal(_name string, _symbol string, _decimals uint8, _totalSupply *big.Int, _initialPercentage uint8, _vestingPeriodInDays *big.Int, _vestingBeneficiary common.Address, _proposer common.Address, _initialPlatformPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitProposal(&_Contracts.TransactOpts, _name, _symbol, _decimals, _totalSupply, _initialPercentage, _vestingPeriodInDays, _vestingBeneficiary, _proposer, _initialPlatformPercentage)
}

// SubmitProposalReferral is a paid mutator transaction binding the contract method 0xb13873f0.
//
// Solidity: function submitProposalReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Contracts *ContractsTransactor) SubmitProposalReferral(opts *bind.TransactOpts, _hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitProposalReferral", _hashIndex, _referral, _referralPercentage)
}

// SubmitProposalReferral is a paid mutator transaction binding the contract method 0xb13873f0.
//
// Solidity: function submitProposalReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Contracts *ContractsSession) SubmitProposalReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitProposalReferral(&_Contracts.TransactOpts, _hashIndex, _referral, _referralPercentage)
}

// SubmitProposalReferral is a paid mutator transaction binding the contract method 0xb13873f0.
//
// Solidity: function submitProposalReferral(bytes32 _hashIndex, address _referral, uint8 _referralPercentage) returns()
func (_Contracts *ContractsTransactorSession) SubmitProposalReferral(_hashIndex [32]byte, _referral common.Address, _referralPercentage uint8) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitProposalReferral(&_Contracts.TransactOpts, _hashIndex, _referral, _referralPercentage)
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

// ContractsLogProposalApproveIterator is returned from FilterLogProposalApprove and is used to iterate over the raw logs and unpacked data for LogProposalApprove events raised by the Contracts contract.
type ContractsLogProposalApproveIterator struct {
	Event *ContractsLogProposalApprove // Event containing the contract specifics and raw log

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
func (it *ContractsLogProposalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogProposalApprove)
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
		it.Event = new(ContractsLogProposalApprove)
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
func (it *ContractsLogProposalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogProposalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogProposalApprove represents a LogProposalApprove event raised by the Contracts contract.
type ContractsLogProposalApprove struct {
	Name         string
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogProposalApprove is a free log retrieval operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: event LogProposalApprove(string name, address indexed tokenAddress)
func (_Contracts *ContractsFilterer) FilterLogProposalApprove(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractsLogProposalApproveIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogProposalApprove", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogProposalApproveIterator{contract: _Contracts.contract, event: "LogProposalApprove", logs: logs, sub: sub}, nil
}

// WatchLogProposalApprove is a free log subscription operation binding the contract event 0x19e9fff5bfdce65750e43ae23fbd302202cd96c991da16e86fdf447b87bfa368.
//
// Solidity: event LogProposalApprove(string name, address indexed tokenAddress)
func (_Contracts *ContractsFilterer) WatchLogProposalApprove(opts *bind.WatchOpts, sink chan<- *ContractsLogProposalApprove, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogProposalApprove", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogProposalApprove)
				if err := _Contracts.contract.UnpackLog(event, "LogProposalApprove", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLogProposalApprove(log types.Log) (*ContractsLogProposalApprove, error) {
	event := new(ContractsLogProposalApprove)
	if err := _Contracts.contract.UnpackLog(event, "LogProposalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogProposalImportedIterator is returned from FilterLogProposalImported and is used to iterate over the raw logs and unpacked data for LogProposalImported events raised by the Contracts contract.
type ContractsLogProposalImportedIterator struct {
	Event *ContractsLogProposalImported // Event containing the contract specifics and raw log

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
func (it *ContractsLogProposalImportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogProposalImported)
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
		it.Event = new(ContractsLogProposalImported)
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
func (it *ContractsLogProposalImportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogProposalImportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogProposalImported represents a LogProposalImported event raised by the Contracts contract.
type ContractsLogProposalImported struct {
	Name      string
	Symbol    string
	Proposer  common.Address
	HashIndex [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProposalImported is a free log retrieval operation binding the contract event 0x82c68a30819d3b8f0c181d716ed3fafa8551cad68eefe911f471516e958421b3.
//
// Solidity: event LogProposalImported(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Contracts *ContractsFilterer) FilterLogProposalImported(opts *bind.FilterOpts, hashIndex [][32]byte) (*ContractsLogProposalImportedIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogProposalImported", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogProposalImportedIterator{contract: _Contracts.contract, event: "LogProposalImported", logs: logs, sub: sub}, nil
}

// WatchLogProposalImported is a free log subscription operation binding the contract event 0x82c68a30819d3b8f0c181d716ed3fafa8551cad68eefe911f471516e958421b3.
//
// Solidity: event LogProposalImported(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Contracts *ContractsFilterer) WatchLogProposalImported(opts *bind.WatchOpts, sink chan<- *ContractsLogProposalImported, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogProposalImported", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogProposalImported)
				if err := _Contracts.contract.UnpackLog(event, "LogProposalImported", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLogProposalImported(log types.Log) (*ContractsLogProposalImported, error) {
	event := new(ContractsLogProposalImported)
	if err := _Contracts.contract.UnpackLog(event, "LogProposalImported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogProposalReferralSubmitIterator is returned from FilterLogProposalReferralSubmit and is used to iterate over the raw logs and unpacked data for LogProposalReferralSubmit events raised by the Contracts contract.
type ContractsLogProposalReferralSubmitIterator struct {
	Event *ContractsLogProposalReferralSubmit // Event containing the contract specifics and raw log

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
func (it *ContractsLogProposalReferralSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogProposalReferralSubmit)
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
		it.Event = new(ContractsLogProposalReferralSubmit)
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
func (it *ContractsLogProposalReferralSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogProposalReferralSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogProposalReferralSubmit represents a LogProposalReferralSubmit event raised by the Contracts contract.
type ContractsLogProposalReferralSubmit struct {
	Referral           common.Address
	ReferralPercentage uint8
	HashIndex          [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLogProposalReferralSubmit is a free log retrieval operation binding the contract event 0x924a2d7699539105c99283e511a2a67d57aecd39d59697c4a76cb2187fddb6ce.
//
// Solidity: event LogProposalReferralSubmit(address referral, uint8 referralPercentage, bytes32 indexed hashIndex)
func (_Contracts *ContractsFilterer) FilterLogProposalReferralSubmit(opts *bind.FilterOpts, hashIndex [][32]byte) (*ContractsLogProposalReferralSubmitIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogProposalReferralSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogProposalReferralSubmitIterator{contract: _Contracts.contract, event: "LogProposalReferralSubmit", logs: logs, sub: sub}, nil
}

// WatchLogProposalReferralSubmit is a free log subscription operation binding the contract event 0x924a2d7699539105c99283e511a2a67d57aecd39d59697c4a76cb2187fddb6ce.
//
// Solidity: event LogProposalReferralSubmit(address referral, uint8 referralPercentage, bytes32 indexed hashIndex)
func (_Contracts *ContractsFilterer) WatchLogProposalReferralSubmit(opts *bind.WatchOpts, sink chan<- *ContractsLogProposalReferralSubmit, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogProposalReferralSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogProposalReferralSubmit)
				if err := _Contracts.contract.UnpackLog(event, "LogProposalReferralSubmit", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLogProposalReferralSubmit(log types.Log) (*ContractsLogProposalReferralSubmit, error) {
	event := new(ContractsLogProposalReferralSubmit)
	if err := _Contracts.contract.UnpackLog(event, "LogProposalReferralSubmit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogProposalSubmitIterator is returned from FilterLogProposalSubmit and is used to iterate over the raw logs and unpacked data for LogProposalSubmit events raised by the Contracts contract.
type ContractsLogProposalSubmitIterator struct {
	Event *ContractsLogProposalSubmit // Event containing the contract specifics and raw log

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
func (it *ContractsLogProposalSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogProposalSubmit)
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
		it.Event = new(ContractsLogProposalSubmit)
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
func (it *ContractsLogProposalSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogProposalSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogProposalSubmit represents a LogProposalSubmit event raised by the Contracts contract.
type ContractsLogProposalSubmit struct {
	Name      string
	Symbol    string
	Proposer  common.Address
	HashIndex [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogProposalSubmit is a free log retrieval operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: event LogProposalSubmit(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Contracts *ContractsFilterer) FilterLogProposalSubmit(opts *bind.FilterOpts, hashIndex [][32]byte) (*ContractsLogProposalSubmitIterator, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogProposalSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogProposalSubmitIterator{contract: _Contracts.contract, event: "LogProposalSubmit", logs: logs, sub: sub}, nil
}

// WatchLogProposalSubmit is a free log subscription operation binding the contract event 0xc59aae764581590a3a647db2078529d429d40d2a95b2ebdb110b5515023a23c1.
//
// Solidity: event LogProposalSubmit(string name, string symbol, address proposer, bytes32 indexed hashIndex)
func (_Contracts *ContractsFilterer) WatchLogProposalSubmit(opts *bind.WatchOpts, sink chan<- *ContractsLogProposalSubmit, hashIndex [][32]byte) (event.Subscription, error) {

	var hashIndexRule []interface{}
	for _, hashIndexItem := range hashIndex {
		hashIndexRule = append(hashIndexRule, hashIndexItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogProposalSubmit", hashIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogProposalSubmit)
				if err := _Contracts.contract.UnpackLog(event, "LogProposalSubmit", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseLogProposalSubmit(log types.Log) (*ContractsLogProposalSubmit, error) {
	event := new(ContractsLogProposalSubmit)
	if err := _Contracts.contract.UnpackLog(event, "LogProposalSubmit", log); err != nil {
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
