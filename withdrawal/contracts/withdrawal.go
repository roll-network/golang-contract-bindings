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
	Salt       *big.Int
}

// WithdrawalSig is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalSig struct {
	V uint8
	R [32]byte
	S [32]byte
}

// WithdrawalMetaData contains all meta data concerning the Withdrawal contract.
var WithdrawalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauseOperator_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CancelledOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpiredOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProcessedOrder\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CancelEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalEvent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawal.Sig\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getCompletedKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumWithdrawal.OrderStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAllowedSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauseOperator_\",\"type\":\"address\"}],\"name\":\"setPauseOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawal.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawal.Sig\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001cd538038062001cd583398181016040528101906200003791906200019d565b60008060006101000a81548160ff0219169083151502179055506200007162000065620000b960201b60201c565b620000c160201b60201c565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000222565b600033905090565b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620001978162000208565b92915050565b600060208284031215620001b657620001b562000203565b5b6000620001c68482850162000186565b91505092915050565b6000620001dc82620001e3565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200021381620001cf565b81146200021f57600080fd5b50565b611aa380620002326000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063715018a61161008c578063a7678b1811610066578063a7678b18146101c8578063ebe73ca4146101e4578063f2fde38b14610200578063f9213fc81461021c576100cf565b8063715018a6146101965780638456cb59146101a05780638da5cb5b146101aa576100cf565b80631ad4aa42146100d45780633f4ba83a146100f057806346c368c2146100fa5780634afdcbde1461012a5780635c975abb146101485780635de28ae014610166575b600080fd5b6100ee60048036038101906100e99190611248565b610238565b005b6100f86102f8565b005b610114600480360381019061010f919061130f565b610389565b6040516101219190611670565b60405180910390f35b6101326103ab565b60405161013f919061154f565b60405180910390f35b6101506103d1565b60405161015d9190611655565b60405180910390f35b610180600480360381019061017b91906112e2565b6103e7565b60405161018d91906116d0565b60405180910390f35b61019e610411565b005b6101a8610499565b005b6101b261052a565b6040516101bf919061154f565b60405180910390f35b6101e260048036038101906101dd919061133c565b610553565b005b6101fe60048036038101906101f9919061133c565b61087d565b005b61021a60048036038101906102159190611248565b610aa1565b005b61023660048036038101906102319190611275565b610b99565b005b610240610c70565b73ffffffffffffffffffffffffffffffffffffffff1661025e61052a565b73ffffffffffffffffffffffffffffffffffffffff16146102b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ab9061174b565b60405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461037f576040517f48f5c3ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610387610c78565b565b60006103a48280360381019061039f919061137d565b610d19565b9050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900460ff16905090565b60006001600083815260200190815260200160002060009054906101000a900460ff169050919050565b610419610c70565b73ffffffffffffffffffffffffffffffffffffffff1661043761052a565b73ffffffffffffffffffffffffffffffffffffffff161461048d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104849061174b565b60405180910390fd5b6104976000610d4d565b565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610520576040517f48f5c3ed00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610528610e12565b565b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61055b6103d1565b1561059b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105929061172b565b60405180910390fd5b6105a58282610eb4565b428260a00135116105e2576040517f45fa2d0600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105fd838036038101906105f8919061137d565b610d19565b90506001600281111561061357610612611875565b5b6001600083815260200190815260200160002060009054906101000a900460ff16600281111561064657610645611875565b5b141561067e576040517f1f368c2d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028081111561069157610690611875565b5b6001600083815260200190815260200160002060009054906101000a900460ff1660028111156106c4576106c3611875565b5b14156106fc576040517f228f2fb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180600083815260200190815260200160002060006101000a81548160ff0219169083600281111561073257610731611875565b5b021790555082606001602081019061074a9190611248565b73ffffffffffffffffffffffffffffffffffffffff166323b872dd8460200160208101906107789190611248565b85604001602081019061078b9190611248565b86608001356040518463ffffffff1660e01b81526004016107ae9392919061161e565b602060405180830381600087803b1580156107c857600080fd5b505af11580156107dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080091906112b5565b507fac1b5146d64cce65efba97457b49a2d72f9be2601587498d2be35fc679d229e18360000160208101906108359190611248565b8460400160208101906108489190611248565b85606001602081019061085b9190611248565b866080013560405161087094939291906115d9565b60405180910390a1505050565b6108856103d1565b156108c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bc9061172b565b60405180910390fd5b6108cf8282610eb4565b60006108ea838036038101906108e5919061137d565b610d19565b905060016002811115610900576108ff611875565b5b6001600083815260200190815260200160002060009054906101000a900460ff16600281111561093357610932611875565b5b141561096b576040517f1f368c2d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028081111561097e5761097d611875565b5b6001600083815260200190815260200160002060009054906101000a900460ff1660028111156109b1576109b0611875565b5b14156109e9576040517f228f2fb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026001600083815260200190815260200160002060006101000a81548160ff02191690836002811115610a2057610a1f611875565b5b02179055507fef4a2fffd8d48cda4f0eb1d1ce2a38ea8b52f8751c3897c5eccce16f1d18f234836000016020810190610a599190611248565b846040016020810190610a6c9190611248565b856060016020810190610a7f9190611248565b8660800135604051610a9494939291906115d9565b60405180910390a1505050565b610aa9610c70565b73ffffffffffffffffffffffffffffffffffffffff16610ac761052a565b73ffffffffffffffffffffffffffffffffffffffff1614610b1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b149061174b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610b8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b849061170b565b60405180910390fd5b610b9681610d4d565b50565b610ba1610c70565b73ffffffffffffffffffffffffffffffffffffffff16610bbf61052a565b73ffffffffffffffffffffffffffffffffffffffff1614610c15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0c9061174b565b60405180910390fd5b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600033905090565b610c806103d1565b610cbf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cb6906116eb565b60405180910390fd5b60008060006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610d02610c70565b604051610d0f919061154f565b60405180910390a1565b60008160c00151604051602001610d309190611534565b604051602081830303815290604052805190602001209050919050565b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610e1a6103d1565b15610e5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e519061172b565b60405180910390fd5b60016000806101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e9d610c70565b604051610eaa919061154f565b60405180910390a1565b6000826000016020810190610ec99190611248565b836020016020810190610edc9190611248565b846040016020810190610eef9190611248565b856060016020810190610f029190611248565b86608001358760a001358860c00135604051602001610f27979695949392919061156a565b604051602081830303815290604052805190602001209050600081604051602001610f52919061150e565b6040516020818303038152906040528051906020012090506000600182856000016020810190610f8291906113aa565b8660200135876040013560405160008152602001604052604051610fa9949392919061168b565b6020604051602081039080840390855afa158015610fcb573d6000803e3d6000fd5b5050506020604051035190508073ffffffffffffffffffffffffffffffffffffffff168560000160208101906110019190611248565b73ffffffffffffffffffffffffffffffffffffffff161461104e576040517faf61069300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166110d1576040517f815e1d6400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b6000813590506110e7816119fa565b92915050565b6000813590506110fc81611a11565b92915050565b60008151905061111181611a11565b92915050565b60008135905061112681611a28565b92915050565b600060e08284031215611142576111416118d3565b5b81905092915050565b600060e08284031215611161576111606118d8565b5b61116b60e061176b565b9050600061117b848285016110d8565b600083015250602061118f848285016110d8565b60208301525060406111a3848285016110d8565b60408301525060606111b7848285016110d8565b60608301525060806111cb8482850161121e565b60808301525060a06111df8482850161121e565b60a08301525060c06111f38482850161121e565b60c08301525092915050565b600060608284031215611215576112146118d3565b5b81905092915050565b60008135905061122d81611a3f565b92915050565b60008135905061124281611a56565b92915050565b60006020828403121561125e5761125d6118dd565b5b600061126c848285016110d8565b91505092915050565b6000806040838503121561128c5761128b6118dd565b5b600061129a858286016110d8565b92505060206112ab858286016110ed565b9150509250929050565b6000602082840312156112cb576112ca6118dd565b5b60006112d984828501611102565b91505092915050565b6000602082840312156112f8576112f76118dd565b5b600061130684828501611117565b91505092915050565b600060e08284031215611325576113246118dd565b5b60006113338482850161112c565b91505092915050565b6000806101408385031215611354576113536118dd565b5b60006113628582860161112c565b92505060e0611373858286016111ff565b9150509250929050565b600060e08284031215611393576113926118dd565b5b60006113a18482850161114b565b91505092915050565b6000602082840312156113c0576113bf6118dd565b5b60006113ce84828501611233565b91505092915050565b6113e0816117ac565b82525050565b6113ef816117be565b82525050565b6113fe816117ca565b82525050565b611415611410826117ca565b611861565b82525050565b6114248161181e565b82525050565b6000611437601483611790565b9150611442826118f3565b602082019050919050565b600061145a601c836117a1565b91506114658261191c565b601c82019050919050565b600061147d602683611790565b915061148882611945565b604082019050919050565b60006114a0601083611790565b91506114ab82611994565b602082019050919050565b60006114c3602083611790565b91506114ce826119bd565b602082019050919050565b6114e281611807565b82525050565b6114f96114f482611807565b61186b565b82525050565b61150881611811565b82525050565b60006115198261144d565b91506115258284611404565b60208201915081905092915050565b600061154082846114e8565b60208201915081905092915050565b600060208201905061156460008301846113d7565b92915050565b600060e08201905061157f600083018a6113d7565b61158c60208301896113d7565b61159960408301886113d7565b6115a660608301876113d7565b6115b360808301866114d9565b6115c060a08301856114d9565b6115cd60c08301846114d9565b98975050505050505050565b60006080820190506115ee60008301876113d7565b6115fb60208301866113d7565b61160860408301856113d7565b61161560608301846114d9565b95945050505050565b600060608201905061163360008301866113d7565b61164060208301856113d7565b61164d60408301846114d9565b949350505050565b600060208201905061166a60008301846113e6565b92915050565b600060208201905061168560008301846113f5565b92915050565b60006080820190506116a060008301876113f5565b6116ad60208301866114ff565b6116ba60408301856113f5565b6116c760608301846113f5565b95945050505050565b60006020820190506116e5600083018461141b565b92915050565b600060208201905081810360008301526117048161142a565b9050919050565b6000602082019050818103600083015261172481611470565b9050919050565b6000602082019050818103600083015261174481611493565b9050919050565b60006020820190508181036000830152611764816114b6565b9050919050565b6000611775611786565b90506117818282611830565b919050565b6000604051905090565b600082825260208201905092915050565b600081905092915050565b60006117b7826117e7565b9050919050565b60008115159050919050565b6000819050919050565b60008190506117e2826119e6565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b6000611829826117d4565b9050919050565b611839826118e2565b810181811067ffffffffffffffff82111715611858576118576118a4565b5b80604052505050565b6000819050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600381106119f7576119f6611875565b5b50565b611a03816117ac565b8114611a0e57600080fd5b50565b611a1a816117be565b8114611a2557600080fd5b50565b611a31816117ca565b8114611a3c57600080fd5b50565b611a4881611807565b8114611a5357600080fd5b50565b611a5f81611811565b8114611a6a57600080fd5b5056fea26469706673582212203627df57137167117a8b6a86b57f3b67720bf4ac4c891270e8cd48d16778427264736f6c63430008070033",
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

// GetCompletedKey is a free data retrieval call binding the contract method 0x46c368c2.
//
// Solidity: function getCompletedKey((address,address,address,address,uint256,uint256,uint256) order) pure returns(bytes32)
func (_Withdrawal *WithdrawalCaller) GetCompletedKey(opts *bind.CallOpts, order WithdrawalOrder) ([32]byte, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "getCompletedKey", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCompletedKey is a free data retrieval call binding the contract method 0x46c368c2.
//
// Solidity: function getCompletedKey((address,address,address,address,uint256,uint256,uint256) order) pure returns(bytes32)
func (_Withdrawal *WithdrawalSession) GetCompletedKey(order WithdrawalOrder) ([32]byte, error) {
	return _Withdrawal.Contract.GetCompletedKey(&_Withdrawal.CallOpts, order)
}

// GetCompletedKey is a free data retrieval call binding the contract method 0x46c368c2.
//
// Solidity: function getCompletedKey((address,address,address,address,uint256,uint256,uint256) order) pure returns(bytes32)
func (_Withdrawal *WithdrawalCallerSession) GetCompletedKey(order WithdrawalOrder) ([32]byte, error) {
	return _Withdrawal.Contract.GetCompletedKey(&_Withdrawal.CallOpts, order)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 key) view returns(uint8)
func (_Withdrawal *WithdrawalCaller) GetStatus(opts *bind.CallOpts, key [32]byte) (uint8, error) {
	var out []interface{}
	err := _Withdrawal.contract.Call(opts, &out, "getStatus", key)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 key) view returns(uint8)
func (_Withdrawal *WithdrawalSession) GetStatus(key [32]byte) (uint8, error) {
	return _Withdrawal.Contract.GetStatus(&_Withdrawal.CallOpts, key)
}

// GetStatus is a free data retrieval call binding the contract method 0x5de28ae0.
//
// Solidity: function getStatus(bytes32 key) view returns(uint8)
func (_Withdrawal *WithdrawalCallerSession) GetStatus(key [32]byte) (uint8, error) {
	return _Withdrawal.Contract.GetStatus(&_Withdrawal.CallOpts, key)
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
	Signer common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelEvent is a free log retrieval operation binding the contract event 0xef4a2fffd8d48cda4f0eb1d1ce2a38ea8b52f8751c3897c5eccce16f1d18f234.
//
// Solidity: event CancelEvent(address signer, address to, address token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) FilterCancelEvent(opts *bind.FilterOpts) (*WithdrawalCancelEventIterator, error) {

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "CancelEvent")
	if err != nil {
		return nil, err
	}
	return &WithdrawalCancelEventIterator{contract: _Withdrawal.contract, event: "CancelEvent", logs: logs, sub: sub}, nil
}

// WatchCancelEvent is a free log subscription operation binding the contract event 0xef4a2fffd8d48cda4f0eb1d1ce2a38ea8b52f8751c3897c5eccce16f1d18f234.
//
// Solidity: event CancelEvent(address signer, address to, address token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) WatchCancelEvent(opts *bind.WatchOpts, sink chan<- *WithdrawalCancelEvent) (event.Subscription, error) {

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "CancelEvent")
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

// ParseCancelEvent is a log parse operation binding the contract event 0xef4a2fffd8d48cda4f0eb1d1ce2a38ea8b52f8751c3897c5eccce16f1d18f234.
//
// Solidity: event CancelEvent(address signer, address to, address token, uint256 amount)
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
	Signer common.Address
	To     common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalEvent is a free log retrieval operation binding the contract event 0xac1b5146d64cce65efba97457b49a2d72f9be2601587498d2be35fc679d229e1.
//
// Solidity: event WithdrawalEvent(address signer, address to, address token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) FilterWithdrawalEvent(opts *bind.FilterOpts) (*WithdrawalWithdrawalEventIterator, error) {

	logs, sub, err := _Withdrawal.contract.FilterLogs(opts, "WithdrawalEvent")
	if err != nil {
		return nil, err
	}
	return &WithdrawalWithdrawalEventIterator{contract: _Withdrawal.contract, event: "WithdrawalEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawalEvent is a free log subscription operation binding the contract event 0xac1b5146d64cce65efba97457b49a2d72f9be2601587498d2be35fc679d229e1.
//
// Solidity: event WithdrawalEvent(address signer, address to, address token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) WatchWithdrawalEvent(opts *bind.WatchOpts, sink chan<- *WithdrawalWithdrawalEvent) (event.Subscription, error) {

	logs, sub, err := _Withdrawal.contract.WatchLogs(opts, "WithdrawalEvent")
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

// ParseWithdrawalEvent is a log parse operation binding the contract event 0xac1b5146d64cce65efba97457b49a2d72f9be2601587498d2be35fc679d229e1.
//
// Solidity: event WithdrawalEvent(address signer, address to, address token, uint256 amount)
func (_Withdrawal *WithdrawalFilterer) ParseWithdrawalEvent(log types.Log) (*WithdrawalWithdrawalEvent, error) {
	event := new(WithdrawalWithdrawalEvent)
	if err := _Withdrawal.contract.UnpackLog(event, "WithdrawalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
