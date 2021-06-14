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

// SocialMoneyABI is the input ABI used to generate the binding from.
const SocialMoneyABI = "[{\"type\":\"constructor\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256[4]\",\"name\":\"_proportions\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"address\",\"name\":\"_platformWallet\"},{\"type\":\"address\",\"name\":\"_tokenVestingInstance\"},{\"type\":\"address\",\"name\":\"_referral\"}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}]},{\"type\":\"function\",\"name\":\"allowance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"approve\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"balanceOf\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"account\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"decimals\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint8\"}]},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"subtractedValue\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"addedValue\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"name\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"symbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"totalSupply\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"transfer\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"transferFrom\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"outputs\":[{\"type\":\"bool\"}]}]"

// SocialMoney is an auto generated Go binding around an Ethereum contract.
type SocialMoney struct {
	SocialMoneyCaller     // Read-only binding to the contract
	SocialMoneyTransactor // Write-only binding to the contract
	SocialMoneyFilterer   // Log filterer for contract events
}

// SocialMoneyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SocialMoneyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocialMoneyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SocialMoneyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocialMoneyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SocialMoneyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocialMoneySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SocialMoneySession struct {
	Contract     *SocialMoney      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SocialMoneyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SocialMoneyCallerSession struct {
	Contract *SocialMoneyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SocialMoneyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SocialMoneyTransactorSession struct {
	Contract     *SocialMoneyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SocialMoneyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SocialMoneyRaw struct {
	Contract *SocialMoney // Generic contract binding to access the raw methods on
}

// SocialMoneyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SocialMoneyCallerRaw struct {
	Contract *SocialMoneyCaller // Generic read-only contract binding to access the raw methods on
}

// SocialMoneyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SocialMoneyTransactorRaw struct {
	Contract *SocialMoneyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSocialMoney creates a new instance of SocialMoney, bound to a specific deployed contract.
func NewSocialMoney(address common.Address, backend bind.ContractBackend) (*SocialMoney, error) {
	contract, err := bindSocialMoney(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SocialMoney{SocialMoneyCaller: SocialMoneyCaller{contract: contract}, SocialMoneyTransactor: SocialMoneyTransactor{contract: contract}, SocialMoneyFilterer: SocialMoneyFilterer{contract: contract}}, nil
}

// NewSocialMoneyCaller creates a new read-only instance of SocialMoney, bound to a specific deployed contract.
func NewSocialMoneyCaller(address common.Address, caller bind.ContractCaller) (*SocialMoneyCaller, error) {
	contract, err := bindSocialMoney(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SocialMoneyCaller{contract: contract}, nil
}

// NewSocialMoneyTransactor creates a new write-only instance of SocialMoney, bound to a specific deployed contract.
func NewSocialMoneyTransactor(address common.Address, transactor bind.ContractTransactor) (*SocialMoneyTransactor, error) {
	contract, err := bindSocialMoney(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SocialMoneyTransactor{contract: contract}, nil
}

// NewSocialMoneyFilterer creates a new log filterer instance of SocialMoney, bound to a specific deployed contract.
func NewSocialMoneyFilterer(address common.Address, filterer bind.ContractFilterer) (*SocialMoneyFilterer, error) {
	contract, err := bindSocialMoney(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SocialMoneyFilterer{contract: contract}, nil
}

// bindSocialMoney binds a generic wrapper to an already deployed contract.
func bindSocialMoney(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SocialMoneyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocialMoney *SocialMoneyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocialMoney.Contract.SocialMoneyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocialMoney *SocialMoneyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocialMoney.Contract.SocialMoneyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocialMoney *SocialMoneyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocialMoney.Contract.SocialMoneyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocialMoney *SocialMoneyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocialMoney.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocialMoney *SocialMoneyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocialMoney.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocialMoney *SocialMoneyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocialMoney.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SocialMoney *SocialMoneyCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SocialMoney.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SocialMoney *SocialMoneySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.Allowance(&_SocialMoney.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SocialMoney *SocialMoneyCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.Allowance(&_SocialMoney.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SocialMoney *SocialMoneyCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SocialMoney.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SocialMoney *SocialMoneySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.BalanceOf(&_SocialMoney.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SocialMoney *SocialMoneyCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.BalanceOf(&_SocialMoney.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SocialMoney *SocialMoneyCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SocialMoney.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SocialMoney *SocialMoneySession) Decimals() (uint8, error) {
	return _SocialMoney.Contract.Decimals(&_SocialMoney.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SocialMoney *SocialMoneyCallerSession) Decimals() (uint8, error) {
	return _SocialMoney.Contract.Decimals(&_SocialMoney.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SocialMoney *SocialMoneyCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SocialMoney.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SocialMoney *SocialMoneySession) Name() (string, error) {
	return _SocialMoney.Contract.Name(&_SocialMoney.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SocialMoney *SocialMoneyCallerSession) Name() (string, error) {
	return _SocialMoney.Contract.Name(&_SocialMoney.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SocialMoney *SocialMoneyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SocialMoney.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SocialMoney *SocialMoneySession) Symbol() (string, error) {
	return _SocialMoney.Contract.Symbol(&_SocialMoney.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SocialMoney *SocialMoneyCallerSession) Symbol() (string, error) {
	return _SocialMoney.Contract.Symbol(&_SocialMoney.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SocialMoney *SocialMoneyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SocialMoney.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SocialMoney *SocialMoneySession) TotalSupply() (*big.Int, error) {
	return _SocialMoney.Contract.TotalSupply(&_SocialMoney.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SocialMoney *SocialMoneyCallerSession) TotalSupply() (*big.Int, error) {
	return _SocialMoney.Contract.TotalSupply(&_SocialMoney.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneyTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneySession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.Approve(&_SocialMoney.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneyTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.Approve(&_SocialMoney.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SocialMoney *SocialMoneyTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SocialMoney.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SocialMoney *SocialMoneySession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.DecreaseAllowance(&_SocialMoney.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SocialMoney *SocialMoneyTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.DecreaseAllowance(&_SocialMoney.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SocialMoney *SocialMoneyTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SocialMoney.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SocialMoney *SocialMoneySession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.IncreaseAllowance(&_SocialMoney.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SocialMoney *SocialMoneyTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.IncreaseAllowance(&_SocialMoney.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneyTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneySession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.Transfer(&_SocialMoney.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneyTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.Transfer(&_SocialMoney.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneyTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneySession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.TransferFrom(&_SocialMoney.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SocialMoney *SocialMoneyTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SocialMoney.Contract.TransferFrom(&_SocialMoney.TransactOpts, sender, recipient, amount)
}

// SocialMoneyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SocialMoney contract.
type SocialMoneyApprovalIterator struct {
	Event *SocialMoneyApproval // Event containing the contract specifics and raw log

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
func (it *SocialMoneyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocialMoneyApproval)
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
		it.Event = new(SocialMoneyApproval)
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
func (it *SocialMoneyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocialMoneyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocialMoneyApproval represents a Approval event raised by the SocialMoney contract.
type SocialMoneyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SocialMoney *SocialMoneyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SocialMoneyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SocialMoney.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SocialMoneyApprovalIterator{contract: _SocialMoney.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SocialMoney *SocialMoneyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SocialMoneyApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SocialMoney.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocialMoneyApproval)
				if err := _SocialMoney.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SocialMoney *SocialMoneyFilterer) ParseApproval(log types.Log) (*SocialMoneyApproval, error) {
	event := new(SocialMoneyApproval)
	if err := _SocialMoney.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SocialMoneyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SocialMoney contract.
type SocialMoneyTransferIterator struct {
	Event *SocialMoneyTransfer // Event containing the contract specifics and raw log

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
func (it *SocialMoneyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocialMoneyTransfer)
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
		it.Event = new(SocialMoneyTransfer)
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
func (it *SocialMoneyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocialMoneyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocialMoneyTransfer represents a Transfer event raised by the SocialMoney contract.
type SocialMoneyTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SocialMoney *SocialMoneyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SocialMoneyTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SocialMoney.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SocialMoneyTransferIterator{contract: _SocialMoney.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SocialMoney *SocialMoneyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SocialMoneyTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SocialMoney.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocialMoneyTransfer)
				if err := _SocialMoney.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SocialMoney *SocialMoneyFilterer) ParseTransfer(log types.Log) (*SocialMoneyTransfer, error) {
	event := new(SocialMoneyTransfer)
	if err := _SocialMoney.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
