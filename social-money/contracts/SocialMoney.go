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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SocialMoneyABI is the input ABI used to generate the binding from.
const SocialMoneyABI = "[{\"type\":\"constructor\",\"payable\":false,\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint8\",\"name\":\"_decimals\"},{\"type\":\"uint256[4]\",\"name\":\"_proportions\"},{\"type\":\"address\",\"name\":\"_vestingBeneficiary\"},{\"type\":\"address\",\"name\":\"_platformWallet\"},{\"type\":\"address\",\"name\":\"_tokenVestingInstance\"},{\"type\":\"address\",\"name\":\"_referral\"}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}]},{\"type\":\"event\",\"anonymous\":false,\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}]},{\"type\":\"function\",\"name\":\"allowance\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"owner\"},{\"type\":\"address\",\"name\":\"spender\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"approve\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"balanceOf\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"account\"}],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"decimals\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint8\"}]},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"subtractedValue\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"spender\"},{\"type\":\"uint256\",\"name\":\"addedValue\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"name\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"symbol\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"string\"}]},{\"type\":\"function\",\"name\":\"totalSupply\",\"constant\":true,\"stateMutability\":\"view\",\"payable\":false,\"inputs\":[],\"outputs\":[{\"type\":\"uint256\"}]},{\"type\":\"function\",\"name\":\"transfer\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"outputs\":[{\"type\":\"bool\"}]},{\"type\":\"function\",\"name\":\"transferFrom\",\"constant\":false,\"payable\":false,\"inputs\":[{\"type\":\"address\",\"name\":\"sender\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"outputs\":[{\"type\":\"bool\"}]}]"

// SocialMoneyBin is the compiled bytecode used for deploying new contracts.
var SocialMoneyBin = "0x60806040523480156200001157600080fd5b5060405162001b2338038062001b238339818101604052810190620000379190620006f2565b8787816003908051906020019062000051929190620004ef565b5080600490805190602001906200006a929190620004ef565b506012600560006101000a81548160ff021916908360ff16021790555050506200009a866200025b60201b60201c565b60006200013086600360048110620000ae57fe5b60200201516200011c88600260048110620000c557fe5b6020020151620001088a600160048110620000dc57fe5b60200201518b600060048110620000ef57fe5b60200201516200027960201b6200073f1790919060201c565b6200027960201b6200073f1790919060201c565b6200027960201b6200073f1790919060201c565b90506200015685876000600481106200014557fe5b60200201516200030260201b60201c565b6200017a84876001600481106200016957fe5b60200201516200030260201b60201c565b6200019e83876002600481106200018d57fe5b60200201516200030260201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614620001f857620001f78287600360048110620001e657fe5b60200201516200030260201b60201c565b5b62000208620004e060201b60201c565b81146200024c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000243906200082c565b60405180910390fd5b505050505050505050620009c0565b80600560006101000a81548160ff021916908360ff16021790555050565b600080828401905083811015620002f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620003a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b620003ba60008383620004ea60201b60201c565b620003d6816002546200027960201b6200073f1790919060201c565b60028190555062000434816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546200027960201b6200073f1790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b6000600254905090565b505050565b828054600181600116156101000203166002900490600052602060002090601f01602090048101928262000527576000855562000573565b82601f106200054257805160ff191683800117855562000573565b8280016001018555821562000573579182015b828111156200057257825182559160200191906001019062000555565b5b50905062000582919062000586565b5090565b5b80821115620005a157600081600090555060010162000587565b5090565b6000620005bc620005b68462000882565b6200084e565b90508082856020860282011115620005d357600080fd5b60005b85811015620006075781620005ec8882620006c4565b845260208401935060208301925050600181019050620005d6565b5050509392505050565b6000620006286200062284620008ab565b6200084e565b9050828152602081018484840111156200064157600080fd5b6200064e8482856200093a565b509392505050565b600081519050620006678162000972565b92915050565b600082601f8301126200067f57600080fd5b60046200068e848285620005a5565b91505092915050565b600082601f830112620006a957600080fd5b8151620006bb84826020860162000611565b91505092915050565b600081519050620006d5816200098c565b92915050565b600081519050620006ec81620009a6565b92915050565b600080600080600080600080610160898b0312156200071057600080fd5b600089015167ffffffffffffffff8111156200072b57600080fd5b620007398b828c0162000697565b985050602089015167ffffffffffffffff8111156200075757600080fd5b620007658b828c0162000697565b9750506040620007788b828c01620006db565b96505060606200078b8b828c016200066d565b95505060e06200079e8b828c0162000656565b945050610100620007b28b828c0162000656565b935050610120620007c68b828c0162000656565b925050610140620007da8b828c0162000656565b9150509295985092959890939650565b6000620007f9601483620008de565b91507f4572726f72206f6e20746f74616c537570706c790000000000000000000000006000830152602082019050919050565b600060208201905081810360008301526200084781620007ea565b9050919050565b6000604051905081810181811067ffffffffffffffff8211171562000878576200087762000970565b5b8060405250919050565b600067ffffffffffffffff821115620008a0576200089f62000970565b5b602082029050919050565b600067ffffffffffffffff821115620008c957620008c862000970565b5b601f19601f8301169050602081019050919050565b600082825260208201905092915050565b6000620008fc8262000903565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b838110156200095a5780820151818401526020810190506200093d565b838111156200096a576000848401525b50505050565bfe5b6200097d81620008ef565b81146200098957600080fd5b50565b620009978162000923565b8114620009a357600080fd5b50565b620009b1816200092d565b8114620009bd57600080fd5b50565b61115380620009d06000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610ee1565b60405180910390f35b6100e660048036038101906100e19190610e24565b610318565b6040516100f39190610ec6565b60405180910390f35b610104610336565b6040516101119190610f03565b60405180910390f35b610134600480360381019061012f9190610dd5565b610340565b6040516101419190610ec6565b60405180910390f35b610152610419565b60405161015f9190610f1e565b60405180910390f35b610182600480360381019061017d9190610e24565b610430565b60405161018f9190610ec6565b60405180910390f35b6101b260048036038101906101ad9190610d70565b6104e3565b6040516101bf9190610f03565b60405180910390f35b6101d061052b565b6040516101dd9190610ee1565b60405180910390f35b61020060048036038101906101fb9190610e24565b6105cd565b60405161020d9190610ec6565b60405180910390f35b610230600480360381019061022b9190610e24565b61069a565b60405161023d9190610ec6565b60405180910390f35b610260600480360381019061025b9190610d99565b6106b8565b60405161026d9190610f03565b60405180910390f35b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561030e5780601f106102e35761010080835404028352916020019161030e565b820191906000526020600020905b8154815290600101906020018083116102f157829003601f168201915b5050505050905090565b600061032c6103256107c7565b84846107cf565b6001905092915050565b6000600254905090565b600061034d8484846109c6565b61040e846103596107c7565b6104098560405180606001604052806028815260200161108860289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006103bf6107c7565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c879092919063ffffffff16565b6107cf565b600190509392505050565b6000600560009054906101000a900460ff16905090565b60006104d961043d6107c7565b846104d4856001600061044e6107c7565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461073f90919063ffffffff16565b6107cf565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105c35780601f10610598576101008083540402835291602001916105c3565b820191906000526020600020905b8154815290600101906020018083116105a657829003601f168201915b5050505050905090565b60006106906105da6107c7565b8461068b856040518060600160405280602581526020016110f960259139600160006106046107c7565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c879092919063ffffffff16565b6107cf565b6001905092915050565b60006106ae6106a76107c7565b84846109c6565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000808284019050838110156107bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610855576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806110d56024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806110406022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610a4c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806110b06025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ad2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061101d6023913960400191505060405180910390fd5b610add838383610d41565b610b4881604051806060016040528060268152602001611062602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610c879092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610bdb816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461073f90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290610d34576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610cf9578082015181840152602081019050610cde565b50505050905090810190601f168015610d265780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5082840390509392505050565b505050565b600081359050610d5581610fee565b92915050565b600081359050610d6a81611005565b92915050565b600060208284031215610d8257600080fd5b6000610d9084828501610d46565b91505092915050565b60008060408385031215610dac57600080fd5b6000610dba85828601610d46565b9250506020610dcb85828601610d46565b9150509250929050565b600080600060608486031215610dea57600080fd5b6000610df886828701610d46565b9350506020610e0986828701610d46565b9250506040610e1a86828701610d5b565b9150509250925092565b60008060408385031215610e3757600080fd5b6000610e4585828601610d46565b9250506020610e5685828601610d5b565b9150509250929050565b610e6981610f67565b82525050565b6000610e7a82610f39565b610e848185610f44565b9350610e94818560208601610faa565b610e9d81610fdd565b840191505092915050565b610eb181610f93565b82525050565b610ec081610f9d565b82525050565b6000602082019050610edb6000830184610e60565b92915050565b60006020820190508181036000830152610efb8184610e6f565b905092915050565b6000602082019050610f186000830184610ea8565b92915050565b6000602082019050610f336000830184610eb7565b92915050565b600081519050919050565b600082825260208201905092915050565b6000610f6082610f73565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015610fc8578082015181840152602081019050610fad565b83811115610fd7576000848401525b50505050565b6000601f19601f8301169050919050565b610ff781610f55565b811461100257600080fd5b50565b61100e81610f93565b811461101957600080fd5b5056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212204f13a783594f0bdac0c5feb455b895bec507f77741635e25507ce5fade66810f64736f6c63430007060033"

// DeploySocialMoney deploys a new Ethereum contract, binding an instance of SocialMoney to it.
func DeploySocialMoney(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, _proportions [4]*big.Int, _vestingBeneficiary common.Address, _platformWallet common.Address, _tokenVestingInstance common.Address, _referral common.Address) (common.Address, *types.Transaction, *SocialMoney, error) {
	parsed, err := abi.JSON(strings.NewReader(SocialMoneyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SocialMoneyBin), backend, _name, _symbol, _decimals, _proportions, _vestingBeneficiary, _platformWallet, _tokenVestingInstance, _referral)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SocialMoney{SocialMoneyCaller: SocialMoneyCaller{contract: contract}, SocialMoneyTransactor: SocialMoneyTransactor{contract: contract}, SocialMoneyFilterer: SocialMoneyFilterer{contract: contract}}, nil
}

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
func (_SocialMoney *SocialMoneyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_SocialMoney *SocialMoneyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SocialMoney *SocialMoneyCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SocialMoney.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SocialMoney *SocialMoneySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.Allowance(&_SocialMoney.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_SocialMoney *SocialMoneyCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.Allowance(&_SocialMoney.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SocialMoney *SocialMoneyCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SocialMoney.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SocialMoney *SocialMoneySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.BalanceOf(&_SocialMoney.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_SocialMoney *SocialMoneyCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SocialMoney.Contract.BalanceOf(&_SocialMoney.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SocialMoney *SocialMoneyCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SocialMoney.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SocialMoney *SocialMoneySession) Decimals() (uint8, error) {
	return _SocialMoney.Contract.Decimals(&_SocialMoney.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SocialMoney *SocialMoneyCallerSession) Decimals() (uint8, error) {
	return _SocialMoney.Contract.Decimals(&_SocialMoney.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SocialMoney *SocialMoneyCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SocialMoney.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SocialMoney *SocialMoneySession) Name() (string, error) {
	return _SocialMoney.Contract.Name(&_SocialMoney.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SocialMoney *SocialMoneyCallerSession) Name() (string, error) {
	return _SocialMoney.Contract.Name(&_SocialMoney.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SocialMoney *SocialMoneyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SocialMoney.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SocialMoney *SocialMoneySession) Symbol() (string, error) {
	return _SocialMoney.Contract.Symbol(&_SocialMoney.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SocialMoney *SocialMoneyCallerSession) Symbol() (string, error) {
	return _SocialMoney.Contract.Symbol(&_SocialMoney.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SocialMoney *SocialMoneyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SocialMoney.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SocialMoney *SocialMoneySession) TotalSupply() (*big.Int, error) {
	return _SocialMoney.Contract.TotalSupply(&_SocialMoney.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
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
	return event, nil
}
