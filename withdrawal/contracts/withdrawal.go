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
	Bin: "0x60806040523480156200001157600080fd5b50604051620021ca380380620021ca83398181016040528101906200003791906200019d565b60008060006101000a81548160ff0219169083151502179055506200007162000065620000b960201b60201c565b620000c160201b60201c565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000222565b600033905090565b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620001978162000208565b92915050565b600060208284031215620001b657620001b562000203565b5b6000620001c68482850162000186565b91505092915050565b6000620001dc82620001e3565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200021381620001cf565b81146200021f57600080fd5b50565b611f9880620002326000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638da5cb5b11610097578063ebe73ca411610066578063ebe73ca41461027b578063f2fde38b14610297578063f8b4d864146102b3578063f9213fc8146102e357610100565b80638da5cb5b146101db578063a7678b18146101f9578063a85c38ef14610215578063d09ef2411461024b57610100565b80635c622a0e116100d35780635c622a0e146101795780635c975abb146101a9578063715018a6146101c75780638456cb59146101d157610100565b80631ad4aa42146101055780633f4ba83a1461012157806346c368c21461012b5780634afdcbde1461015b575b600080fd5b61011f600480360381019061011a91906116bc565b6102ff565b005b610129610420565b005b61014560048036038101906101409190611756565b6104b1565b6040516101529190611c66565b60405180910390f35b6101636104d3565b6040516101709190611a3d565b60405180910390f35b610193600480360381019061018e91906117f1565b6104f9565b6040516101a09190611bb0565b60405180910390f35b6101b1610523565b6040516101be9190611b50565b60405180910390f35b6101cf610539565b005b6101d96105c1565b005b6101e3610652565b6040516101f09190611a3d565b60405180910390f35b610213600480360381019061020e9190611783565b61067b565b005b61022f600480360381019061022a91906117f1565b6109bf565b6040516102429796959493929190611a81565b60405180910390f35b610265600480360381019061026091906117f1565b610a81565b6040516102729190611c4b565b60405180910390f35b61029560048036038101906102909190611783565b610c24565b005b6102b160048036038101906102ac91906116bc565b610e62565b005b6102cd60048036038101906102c891906116bc565b610f5a565b6040516102da9190611b50565b60405180910390f35b6102fd60048036038101906102f891906116e9565b610f7a565b005b61030761108a565b73ffffffffffffffffffffffffffffffffffffffff16610325610652565b73ffffffffffffffffffffffffffffffffffffffff161461037b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037290611c2b565b60405180910390fd5b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f49af92e8e702a1adb06ae79c01a1dfed719800edb7f85d9b79e2622ea33f4cf98183604051610414929190611a58565b60405180910390a15050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104a7576040517f48f5c3ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104af611092565b565b60006104cc828036038101906104c791906117c4565b611133565b9050919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006001600083815260200190815260200160002060009054906101000a900460ff169050919050565b60008060009054906101000a900460ff16905090565b61054161108a565b73ffffffffffffffffffffffffffffffffffffffff1661055f610652565b73ffffffffffffffffffffffffffffffffffffffff16146105b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ac90611c2b565b60405180910390fd5b6105bf6000611141565b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610648576040517f48f5c3ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610650611206565b565b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610683610523565b156106c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ba90611c0b565b60405180910390fd5b6106cd82826112a8565b428260a001351161070a576040517f45fa2d0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006107258380360381019061072091906117c4565b611133565b90506001600281111561073b5761073a611d81565b5b6001600083815260200190815260200160002060009054906101000a900460ff16600281111561076e5761076d611d81565b5b14156107a6576040517f1f368c2d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002808111156107b9576107b8611d81565b5b6001600083815260200190815260200160002060009054906101000a900460ff1660028111156107ec576107eb611d81565b5b1415610824576040517f228f2fb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180600083815260200190815260200160002060006101000a81548160ff0219169083600281111561085a57610859611d81565b5b021790555082606001602081019061087291906116bc565b73ffffffffffffffffffffffffffffffffffffffff166323b872dd8460200160208101906108a091906116bc565b8560400160208101906108b391906116bc565b86608001356040518463ffffffff1660e01b81526004016108d693929190611af0565b602060405180830381600087803b1580156108f057600080fd5b505af1158015610904573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109289190611729565b5082606001602081019061093c91906116bc565b73ffffffffffffffffffffffffffffffffffffffff168360c001357f427a5c054bf4858fa5eaf4fbc1125a64160741ef59682db0fdc9687b3858094485600001602081019061098b91906116bc565b86604001602081019061099e91906116bc565b87608001356040516109b293929190611af0565b60405180910390a3505050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050154908060060154905087565b610a896114cc565b600260008381526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815250509050919050565b610c2c610523565b15610c6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6390611c0b565b60405180910390fd5b610c7682826112a8565b6000610c9183803603810190610c8c91906117c4565b611133565b905060016002811115610ca757610ca6611d81565b5b6001600083815260200190815260200160002060009054906101000a900460ff166002811115610cda57610cd9611d81565b5b1415610d12576040517f1f368c2d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600280811115610d2557610d24611d81565b5b6001600083815260200190815260200160002060009054906101000a900460ff166002811115610d5857610d57611d81565b5b1415610d90576040517f228f2fb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026001600083815260200190815260200160002060006101000a81548160ff02191690836002811115610dc757610dc6611d81565b5b0217905550826060016020810190610ddf91906116bc565b73ffffffffffffffffffffffffffffffffffffffff168360c001357f18351c9e93b203d861f9bb52616afdc76a08cf3bb7c23d5d7a9f742fd2f309c1856000016020810190610e2e91906116bc565b866040016020810190610e4191906116bc565b8760800135604051610e5593929190611af0565b60405180910390a3505050565b610e6a61108a565b73ffffffffffffffffffffffffffffffffffffffff16610e88610652565b73ffffffffffffffffffffffffffffffffffffffff1614610ede576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed590611c2b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4590611beb565b60405180910390fd5b610f5781611141565b50565b60036020528060005260406000206000915054906101000a900460ff1681565b610f8261108a565b73ffffffffffffffffffffffffffffffffffffffff16610fa0610652565b73ffffffffffffffffffffffffffffffffffffffff1614610ff6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fed90611c2b565b60405180910390fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f863d338cad74814b108a06288ad5e0e80d56495e0332238b1d2cdcfa0ca8e5ce828260405161107e929190611b27565b60405180910390a15050565b600033905090565b61109a610523565b6110d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d090611bcb565b60405180910390fd5b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61111c61108a565b6040516111299190611a3d565b60405180910390a1565b60008160c001519050919050565b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61120e610523565b1561124e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124590611c0b565b60405180910390fd5b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861129161108a565b60405161129e9190611a3d565b60405180910390a1565b60008260000160208101906112bd91906116bc565b8360200160208101906112d091906116bc565b8460400160208101906112e391906116bc565b8560600160208101906112f691906116bc565b86608001358760a001358860c0013560405160200161131b9796959493929190611a81565b6040516020818303038152906040528051906020012090506000816040516020016113469190611a17565b6040516020818303038152906040528051906020012090506000600182856000016020810190611376919061181e565b866020013587604001356040516000815260200160405260405161139d9493929190611b6b565b6020604051602081039080840390855afa1580156113bf573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168560000160208101906113f591906116bc565b73ffffffffffffffffffffffffffffffffffffffff1614611442576040517faf61069300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166114c5576040517f815e1d6400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b60008135905061157081611f06565b92915050565b60008135905061158581611f1d565b92915050565b60008151905061159a81611f1d565b92915050565b600060e082840312156115b6576115b5611ddf565b5b81905092915050565b600060e082840312156115d5576115d4611de4565b5b6115df60e0611c81565b905060006115ef84828501611561565b600083015250602061160384828501611561565b602083015250604061161784828501611561565b604083015250606061162b84828501611561565b606083015250608061163f84828501611692565b60808301525060a061165384828501611692565b60a08301525060c061166784828501611692565b60c08301525092915050565b60006060828403121561168957611688611ddf565b5b81905092915050565b6000813590506116a181611f34565b92915050565b6000813590506116b681611f4b565b92915050565b6000602082840312156116d2576116d1611de9565b5b60006116e084828501611561565b91505092915050565b60008060408385031215611700576116ff611de9565b5b600061170e85828601611561565b925050602061171f85828601611576565b9150509250929050565b60006020828403121561173f5761173e611de9565b5b600061174d8482850161158b565b91505092915050565b600060e0828403121561176c5761176b611de9565b5b600061177a848285016115a0565b91505092915050565b600080610140838503121561179b5761179a611de9565b5b60006117a9858286016115a0565b92505060e06117ba85828601611673565b9150509250929050565b600060e082840312156117da576117d9611de9565b5b60006117e8848285016115bf565b91505092915050565b60006020828403121561180757611806611de9565b5b600061181584828501611692565b91505092915050565b60006020828403121561183457611833611de9565b5b6000611842848285016116a7565b91505092915050565b61185481611cc2565b82525050565b61186381611cc2565b82525050565b61187281611cd4565b82525050565b61188181611ce0565b82525050565b61189861189382611ce0565b611d77565b82525050565b6118a781611d34565b82525050565b60006118ba601483611ca6565b91506118c582611dff565b602082019050919050565b60006118dd601c83611cb7565b91506118e882611e28565b601c82019050919050565b6000611900602683611ca6565b915061190b82611e51565b604082019050919050565b6000611923601083611ca6565b915061192e82611ea0565b602082019050919050565b6000611946602083611ca6565b915061195182611ec9565b602082019050919050565b60e082016000820151611972600085018261184b565b506020820151611985602085018261184b565b506040820151611998604085018261184b565b5060608201516119ab606085018261184b565b5060808201516119be60808501826119ea565b5060a08201516119d160a08501826119ea565b5060c08201516119e460c08501826119ea565b50505050565b6119f381611d1d565b82525050565b611a0281611d1d565b82525050565b611a1181611d27565b82525050565b6000611a22826118d0565b9150611a2e8284611887565b60208201915081905092915050565b6000602082019050611a52600083018461185a565b92915050565b6000604082019050611a6d600083018561185a565b611a7a602083018461185a565b9392505050565b600060e082019050611a96600083018a61185a565b611aa3602083018961185a565b611ab0604083018861185a565b611abd606083018761185a565b611aca60808301866119f9565b611ad760a08301856119f9565b611ae460c08301846119f9565b98975050505050505050565b6000606082019050611b05600083018661185a565b611b12602083018561185a565b611b1f60408301846119f9565b949350505050565b6000604082019050611b3c600083018561185a565b611b496020830184611869565b9392505050565b6000602082019050611b656000830184611869565b92915050565b6000608082019050611b806000830187611878565b611b8d6020830186611a08565b611b9a6040830185611878565b611ba76060830184611878565b95945050505050565b6000602082019050611bc5600083018461189e565b92915050565b60006020820190508181036000830152611be4816118ad565b9050919050565b60006020820190508181036000830152611c04816118f3565b9050919050565b60006020820190508181036000830152611c2481611916565b9050919050565b60006020820190508181036000830152611c4481611939565b9050919050565b600060e082019050611c60600083018461195c565b92915050565b6000602082019050611c7b60008301846119f9565b92915050565b6000611c8b611c9c565b9050611c978282611d46565b919050565b6000604051905090565b600082825260208201905092915050565b600081905092915050565b6000611ccd82611cfd565b9050919050565b60008115159050919050565b6000819050919050565b6000819050611cf882611ef2565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611d3f82611cea565b9050919050565b611d4f82611dee565b810181811067ffffffffffffffff82111715611d6e57611d6d611db0565b5b80604052505050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60038110611f0357611f02611d81565b5b50565b611f0f81611cc2565b8114611f1a57600080fd5b50565b611f2681611cd4565b8114611f3157600080fd5b50565b611f3d81611d1d565b8114611f4857600080fd5b50565b611f5481611d27565b8114611f5f57600080fd5b5056fea2646970667358221220af5116da8418cc23e4bf7a851a8222ce25ec6744c40efd88fd4060e45c213eef64736f6c63430008070033",
}

// WithdrawalABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalMetaData.ABI instead.
var WithdrawalABI = WithdrawalMetaData.ABI

// WithdrawalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WithdrawalMetaData.Bin instead.
var WithdrawalBin = WithdrawalMetaData.Bin

// DeployWithdrawal deploys a new Ethereum contract, binding an instance of Withdrawal to it.
func DeployWithdrawal(auth *bind.TransactOpts, backend bind.ContractBackend, pauseOperator_ common.Address) (common.Address, *types.Transaction, *Withdrawal, error) {
	parsed, err := WithdrawalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WithdrawalBin), backend, pauseOperator_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Withdrawal{WithdrawalCaller: WithdrawalCaller{contract: contract}, WithdrawalTransactor: WithdrawalTransactor{contract: contract}, WithdrawalFilterer: WithdrawalFilterer{contract: contract}}, nil
}

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
