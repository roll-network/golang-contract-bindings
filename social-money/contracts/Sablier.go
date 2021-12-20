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
)

// SablierMetaData contains all meta data concerning the Sablier contract.
var SablierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"senderBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recipientBalance\",\"type\":\"uint256\"}],\"name\":\"CancelStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stopTime\",\"type\":\"uint256\"}],\"name\":\"CreateStream\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFromStream\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"cancelStream\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stopTime\",\"type\":\"uint256\"}],\"name\":\"createStream\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"deltaOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delta\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mantissa\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"getStream\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratePerSecond\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"}],\"name\":\"isEntity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextStreamId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"streamId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromStream\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100935760003560e01c8063894e9a0d11610066578063894e9a0d146101aa578063a82ccd4d1461025f578063cc1b4bf6146102a1578063d7a42b5414610337578063ddca3f431461037b57610093565b80631e99d569146100985780633656eec2146100b65780636db9241b146101185780637a9b2c6c1461015c575b600080fd5b6100a0610399565b6040518082815260200191505060405180910390f35b610102600480360360408110156100cc57600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061039f565b6040518082815260200191505060405180910390f35b6101446004803603602081101561012e57600080fd5b8101908080359060200190929190505050610830565b60405180821515815260200191505060405180910390f35b6101926004803603604081101561017257600080fd5b810190808035906020019092919080359060200190929190505050610fbc565b60405180821515815260200191505060405180910390f35b6101d6600480360360208110156101c057600080fd5b810190808035906020019092919050505061173e565b604051808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018381526020018281526020019850505050505050505060405180910390f35b61028b6004803603602081101561027557600080fd5b8101908080359060200190929190505050611915565b6040518082815260200191505060405180910390f35b610321600480360360a08110156102b757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050611b65565b6040518082815260200191505060405180910390f35b6103636004803603602081101561034d57600080fd5b810190808035906020019092919050505061249b565b60405180821515815260200191505060405180910390f35b6103836124c8565b6040518082815260200191505060405180910390f35b60035481565b6000826004600082815260200190815260200160002060070160149054906101000a900460ff16610438576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73747265616d20646f6573206e6f74206578697374000000000000000000000081525060200191505060405180910390fd5b6000600460008681526020019081526020016000206040518061012001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160149054906101000a900460ff16151515158152505090506105b26125a9565b60006105bd87611915565b90506105cd8184602001516124d4565b8360000184602001828152508260038111156105e557fe5b60038111156105f057fe5b81525050506000600381111561060257fe5b8260000151600381111561061257fe5b14610668576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061268e6023913960400191505060405180910390fd5b8260400151836000015111156107355761068a83600001518460400151612521565b8360000184604001828152508260038111156106a257fe5b60038111156106ad57fe5b8152505050600060038111156106bf57fe5b826000015160038111156106cf57fe5b146106d657fe5b6106e882602001518360400151612521565b83600001846020018281525082600381111561070057fe5b600381111561070b57fe5b81525050506000600381111561071d57fe5b8260000151600381111561072d57fe5b1461073457fe5b5b8260a0015173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561077c5781602001519450505050610829565b8260c0015173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415610821576107c683604001518360200151612521565b8360000184606001828152508260038111156107de57fe5b60038111156107e957fe5b8152505050600060038111156107fb57fe5b8260000151600381111561080b57fe5b1461081257fe5b81606001519450505050610829565b600094505050505b5092915050565b6000600260005414156108ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0081525060200191505060405180910390fd5b6002600081905550816004600082815260200190815260200160002060070160149054906101000a900460ff1661094a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73747265616d20646f6573206e6f74206578697374000000000000000000000081525060200191505060405180910390fd5b826004600082815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610a1c57506004600082815260200190815260200160002060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610a71576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806126b16037913960400191505060405180910390fd5b6000600460008681526020019081526020016000206040518061012001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160149054906101000a900460ff16151515158152505090506000610bf3868360c0015161039f565b90506000610c05878460a0015161039f565b90506004600088815260200190815260200160002060008082016000905560018201600090556002820160009055600382016000905560048201600090556005820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556006820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556007820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556007820160146101000a81549060ff0219169055505060008360e0015190506000821115610e03578073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8560a00151846040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610d5557600080fd5b505af1158015610d69573d6000803e3d6000fd5b505050506040513d6020811015610d7f57600080fd5b8101908080519060200190929190505050610e02576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f726563697069656e7420746f6b656e207472616e73666572206661696c75726581525060200191505060405180910390fd5b5b6000831115610f2f578073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8560c00151856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610e8157600080fd5b505af1158015610e95573d6000803e3d6000fd5b505050506040513d6020811015610eab57600080fd5b8101908080519060200190929190505050610f2e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f73656e64657220746f6b656e207472616e73666572206661696c75726500000081525060200191505060405180910390fd5b5b8360a0015173ffffffffffffffffffffffffffffffffffffffff168460c0015173ffffffffffffffffffffffffffffffffffffffff16897fca3e6079b726e7728802a0537949e2d1c7762304fa641fb06eb56daf2ba8c6b98686604051808381526020018281526020019250505060405180910390a4600196505050505050506001600081905550919050565b600060026000541415611037576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0081525060200191505060405180910390fd5b6002600081905550826004600082815260200190815260200160002060070160149054906101000a900460ff166110d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73747265616d20646f6573206e6f74206578697374000000000000000000000081525060200191505060405180910390fd5b836004600082815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806111a857506004600082815260200190815260200160002060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6111fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001806126b16037913960400191505060405180910390fd5b60008411611273576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f616d6f756e74206973207a65726f00000000000000000000000000000000000081525060200191505060405180910390fd5b6000600460008781526020019081526020016000206040518061012001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160149054906101000a900460ff16151515158152505090506113ed6125dc565b60006113fd888460a0015161039f565b905086811015611458576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806126276024913960400191505060405180910390fd5b611466836040015188612521565b83600001600460008c81526020019081526020016000206002016000839190505582600381111561149357fe5b600381111561149e57fe5b8152505050600060038111156114b057fe5b826000015160038111156114c057fe5b146114c757fe5b6000600460008a81526020019081526020016000206002015414156115ae576004600089815260200190815260200160002060008082016000905560018201600090556002820160009055600382016000905560048201600090556005820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556006820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556007820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556007820160146101000a81549060ff021916905550505b8260e0015173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8460a00151896040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561162757600080fd5b505af115801561163b573d6000803e3d6000fd5b505050506040513d602081101561165157600080fd5b81019080805190602001909291905050506116d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f746f6b656e207472616e73666572206661696c7572650000000000000000000081525060200191505060405180910390fd5b8260a0015173ffffffffffffffffffffffffffffffffffffffff16887f36c3ab437e6a424ed25dc4bfdeb62706aa06558660fab2dab229d2555adaf89c896040518082815260200191505060405180910390a3600195505050505050600160008190555092915050565b600080600080600080600080886004600082815260200190815260200160002060070160149054906101000a900460ff166117e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73747265616d20646f6573206e6f74206578697374000000000000000000000081525060200191505060405180910390fd5b600460008b815260200190815260200160002060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169850600460008b815260200190815260200160002060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169750600460008b8152602001908152602001600020600001549650600460008b815260200190815260200160002060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169550600460008b8152602001908152602001600020600301549450600460008b8152602001908152602001600020600401549350600460008b8152602001908152602001600020600201549250600460008b815260200190815260200160002060010154915050919395975091939597565b6000816004600082815260200190815260200160002060070160149054906101000a900460ff166119ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73747265616d20646f6573206e6f74206578697374000000000000000000000081525060200191505060405180910390fd5b6000600460008581526020019081526020016000206040518061012001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160149054906101000a900460ff161515151581525050905080606001514211611b35576000925050611b5f565b8060800151421015611b505780606001514203925050611b5f565b80606001518160800151039250505b50919050565b60008073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415611c09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f73747265616d20746f20746865207a65726f206164647265737300000000000081525060200191505060405180910390fd5b3073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415611cab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f73747265616d20746f2074686520636f6e747261637420697473656c6600000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff161415611d4d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f73747265616d20746f207468652063616c6c657200000000000000000000000081525060200191505060405180910390fd5b60008511611dc3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6465706f736974206973207a65726f000000000000000000000000000000000081525060200191505060405180910390fd5b42831015611e1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061264b6021913960400191505060405180910390fd5b828211611e91576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f73746f702074696d65206265666f7265207468652073746172742074696d650081525060200191505060405180910390fd5b611e996125fa565b611ea38385612521565b826000018360200182815250826003811115611ebb57fe5b6003811115611ec657fe5b815250505060006003811115611ed857fe5b81600001516003811115611ee857fe5b14611eef57fe5b8060200151861015611f69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6465706f73697420736d616c6c6572207468616e2074696d652064656c74610081525060200191505060405180910390fd5b600081602001518781611f7857fe5b0614611fcf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061266c6022913960400191505060405180910390fd5b611fdd868260200151612549565b826000018360400182815250826003811115611ff557fe5b600381111561200057fe5b81525050506000600381111561201257fe5b8160000151600381111561202257fe5b1461202957fe5b60006003549050604051806101200160405280888152602001836040015181526020018881526020018681526020018581526020018973ffffffffffffffffffffffffffffffffffffffff1681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020016001151581525060046000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101008201518160070160146101000a81548160ff021916908315150217905550905050612204600354600161257a565b8360000160036000839190505582600381111561221d57fe5b600381111561222857fe5b81525050506000600381111561223a57fe5b8260000151600381111561224a57fe5b146122bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f6e6578742073747265616d2069642063616c63756c6174696f6e206572726f7281525060200191505060405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff166323b872dd33308a6040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561234c57600080fd5b505af1158015612360573d6000803e3d6000fd5b505050506040513d602081101561237657600080fd5b81019080805190602001909291905050506123f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f746f6b656e207472616e73666572206661696c7572650000000000000000000081525060200191505060405180910390fd5b8773ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16827f7b01d409597969366dc268d7f957a990d1ca3d3449baf8fb45db67351aecfe788a8a8a8a604051808581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200194505050505060405180910390a4809250505095945050505050565b60006004600083815260200190815260200160002060070160149054906101000a900460ff169050919050565b60028060000154905081565b60008060008414156124ec576000809150915061251a565b60008385029050838582816124fd57fe5b04146125115760026000925092505061251a565b60008192509250505b9250929050565b60008083831161253957600083850391509150612542565b60036000915091505b9250929050565b6000806000831415612562576001600091509150612573565b600083858161256d57fe5b04915091505b9250929050565b600080600083850190508481106125985760008192509250506125a2565b6002600092509250505b9250929050565b6040518060800160405280600060038111156125c157fe5b81526020016000815260200160008152602001600081525090565b6040518060200160405280600060038111156125f457fe5b81525090565b60405180606001604052806000600381111561261257fe5b81526020016000815260200160008152509056fe616d6f756e7420657863656564732074686520617661696c61626c652062616c616e636573746172742074696d65206265666f726520626c6f636b2e74696d657374616d706465706f736974206e6f74206d756c7469706c65206f662074696d652064656c7461726563697069656e742062616c616e63652063616c63756c6174696f6e206572726f7263616c6c6572206973206e6f74207468652073656e646572206f722074686520726563697069656e74206f66207468652073747265616da26469706673582212204e2cbcadcc03dac06e56d6d484540e19013a07984e308b4d735c9d2b60d1b61a64736f6c63430007060033",
}

// SablierABI is the input ABI used to generate the binding from.
// Deprecated: Use SablierMetaData.ABI instead.
var SablierABI = SablierMetaData.ABI

// SablierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SablierMetaData.Bin instead.
var SablierBin = SablierMetaData.Bin

// DeploySablier deploys a new Ethereum contract, binding an instance of Sablier to it.
func DeploySablier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sablier, error) {
	parsed, err := SablierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SablierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sablier{SablierCaller: SablierCaller{contract: contract}, SablierTransactor: SablierTransactor{contract: contract}, SablierFilterer: SablierFilterer{contract: contract}}, nil
}

// Sablier is an auto generated Go binding around an Ethereum contract.
type Sablier struct {
	SablierCaller     // Read-only binding to the contract
	SablierTransactor // Write-only binding to the contract
	SablierFilterer   // Log filterer for contract events
}

// SablierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SablierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SablierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SablierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SablierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SablierSession struct {
	Contract     *Sablier          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SablierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SablierCallerSession struct {
	Contract *SablierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SablierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SablierTransactorSession struct {
	Contract     *SablierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SablierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SablierRaw struct {
	Contract *Sablier // Generic contract binding to access the raw methods on
}

// SablierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SablierCallerRaw struct {
	Contract *SablierCaller // Generic read-only contract binding to access the raw methods on
}

// SablierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SablierTransactorRaw struct {
	Contract *SablierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSablier creates a new instance of Sablier, bound to a specific deployed contract.
func NewSablier(address common.Address, backend bind.ContractBackend) (*Sablier, error) {
	contract, err := bindSablier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sablier{SablierCaller: SablierCaller{contract: contract}, SablierTransactor: SablierTransactor{contract: contract}, SablierFilterer: SablierFilterer{contract: contract}}, nil
}

// NewSablierCaller creates a new read-only instance of Sablier, bound to a specific deployed contract.
func NewSablierCaller(address common.Address, caller bind.ContractCaller) (*SablierCaller, error) {
	contract, err := bindSablier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SablierCaller{contract: contract}, nil
}

// NewSablierTransactor creates a new write-only instance of Sablier, bound to a specific deployed contract.
func NewSablierTransactor(address common.Address, transactor bind.ContractTransactor) (*SablierTransactor, error) {
	contract, err := bindSablier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SablierTransactor{contract: contract}, nil
}

// NewSablierFilterer creates a new log filterer instance of Sablier, bound to a specific deployed contract.
func NewSablierFilterer(address common.Address, filterer bind.ContractFilterer) (*SablierFilterer, error) {
	contract, err := bindSablier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SablierFilterer{contract: contract}, nil
}

// bindSablier binds a generic wrapper to an already deployed contract.
func bindSablier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SablierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sablier *SablierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sablier.Contract.SablierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sablier *SablierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sablier.Contract.SablierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sablier *SablierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sablier.Contract.SablierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sablier *SablierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sablier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sablier *SablierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sablier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sablier *SablierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sablier.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x3656eec2.
//
// Solidity: function balanceOf(uint256 streamId, address who) view returns(uint256 balance)
func (_Sablier *SablierCaller) BalanceOf(opts *bind.CallOpts, streamId *big.Int, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Sablier.contract.Call(opts, &out, "balanceOf", streamId, who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x3656eec2.
//
// Solidity: function balanceOf(uint256 streamId, address who) view returns(uint256 balance)
func (_Sablier *SablierSession) BalanceOf(streamId *big.Int, who common.Address) (*big.Int, error) {
	return _Sablier.Contract.BalanceOf(&_Sablier.CallOpts, streamId, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x3656eec2.
//
// Solidity: function balanceOf(uint256 streamId, address who) view returns(uint256 balance)
func (_Sablier *SablierCallerSession) BalanceOf(streamId *big.Int, who common.Address) (*big.Int, error) {
	return _Sablier.Contract.BalanceOf(&_Sablier.CallOpts, streamId, who)
}

// DeltaOf is a free data retrieval call binding the contract method 0xa82ccd4d.
//
// Solidity: function deltaOf(uint256 streamId) view returns(uint256 delta)
func (_Sablier *SablierCaller) DeltaOf(opts *bind.CallOpts, streamId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Sablier.contract.Call(opts, &out, "deltaOf", streamId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeltaOf is a free data retrieval call binding the contract method 0xa82ccd4d.
//
// Solidity: function deltaOf(uint256 streamId) view returns(uint256 delta)
func (_Sablier *SablierSession) DeltaOf(streamId *big.Int) (*big.Int, error) {
	return _Sablier.Contract.DeltaOf(&_Sablier.CallOpts, streamId)
}

// DeltaOf is a free data retrieval call binding the contract method 0xa82ccd4d.
//
// Solidity: function deltaOf(uint256 streamId) view returns(uint256 delta)
func (_Sablier *SablierCallerSession) DeltaOf(streamId *big.Int) (*big.Int, error) {
	return _Sablier.Contract.DeltaOf(&_Sablier.CallOpts, streamId)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256 mantissa)
func (_Sablier *SablierCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sablier.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256 mantissa)
func (_Sablier *SablierSession) Fee() (*big.Int, error) {
	return _Sablier.Contract.Fee(&_Sablier.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256 mantissa)
func (_Sablier *SablierCallerSession) Fee() (*big.Int, error) {
	return _Sablier.Contract.Fee(&_Sablier.CallOpts)
}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns(address sender, address recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime, uint256 remainingBalance, uint256 ratePerSecond)
func (_Sablier *SablierCaller) GetStream(opts *bind.CallOpts, streamId *big.Int) (struct {
	Sender           common.Address
	Recipient        common.Address
	Deposit          *big.Int
	TokenAddress     common.Address
	StartTime        *big.Int
	StopTime         *big.Int
	RemainingBalance *big.Int
	RatePerSecond    *big.Int
}, error) {
	var out []interface{}
	err := _Sablier.contract.Call(opts, &out, "getStream", streamId)

	outstruct := new(struct {
		Sender           common.Address
		Recipient        common.Address
		Deposit          *big.Int
		TokenAddress     common.Address
		StartTime        *big.Int
		StopTime         *big.Int
		RemainingBalance *big.Int
		RatePerSecond    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokenAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StopTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RemainingBalance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RatePerSecond = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns(address sender, address recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime, uint256 remainingBalance, uint256 ratePerSecond)
func (_Sablier *SablierSession) GetStream(streamId *big.Int) (struct {
	Sender           common.Address
	Recipient        common.Address
	Deposit          *big.Int
	TokenAddress     common.Address
	StartTime        *big.Int
	StopTime         *big.Int
	RemainingBalance *big.Int
	RatePerSecond    *big.Int
}, error) {
	return _Sablier.Contract.GetStream(&_Sablier.CallOpts, streamId)
}

// GetStream is a free data retrieval call binding the contract method 0x894e9a0d.
//
// Solidity: function getStream(uint256 streamId) view returns(address sender, address recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime, uint256 remainingBalance, uint256 ratePerSecond)
func (_Sablier *SablierCallerSession) GetStream(streamId *big.Int) (struct {
	Sender           common.Address
	Recipient        common.Address
	Deposit          *big.Int
	TokenAddress     common.Address
	StartTime        *big.Int
	StopTime         *big.Int
	RemainingBalance *big.Int
	RatePerSecond    *big.Int
}, error) {
	return _Sablier.Contract.GetStream(&_Sablier.CallOpts, streamId)
}

// IsEntity is a free data retrieval call binding the contract method 0xd7a42b54.
//
// Solidity: function isEntity(uint256 streamId) view returns(bool)
func (_Sablier *SablierCaller) IsEntity(opts *bind.CallOpts, streamId *big.Int) (bool, error) {
	var out []interface{}
	err := _Sablier.contract.Call(opts, &out, "isEntity", streamId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEntity is a free data retrieval call binding the contract method 0xd7a42b54.
//
// Solidity: function isEntity(uint256 streamId) view returns(bool)
func (_Sablier *SablierSession) IsEntity(streamId *big.Int) (bool, error) {
	return _Sablier.Contract.IsEntity(&_Sablier.CallOpts, streamId)
}

// IsEntity is a free data retrieval call binding the contract method 0xd7a42b54.
//
// Solidity: function isEntity(uint256 streamId) view returns(bool)
func (_Sablier *SablierCallerSession) IsEntity(streamId *big.Int) (bool, error) {
	return _Sablier.Contract.IsEntity(&_Sablier.CallOpts, streamId)
}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_Sablier *SablierCaller) NextStreamId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sablier.contract.Call(opts, &out, "nextStreamId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_Sablier *SablierSession) NextStreamId() (*big.Int, error) {
	return _Sablier.Contract.NextStreamId(&_Sablier.CallOpts)
}

// NextStreamId is a free data retrieval call binding the contract method 0x1e99d569.
//
// Solidity: function nextStreamId() view returns(uint256)
func (_Sablier *SablierCallerSession) NextStreamId() (*big.Int, error) {
	return _Sablier.Contract.NextStreamId(&_Sablier.CallOpts)
}

// CancelStream is a paid mutator transaction binding the contract method 0x6db9241b.
//
// Solidity: function cancelStream(uint256 streamId) returns(bool)
func (_Sablier *SablierTransactor) CancelStream(opts *bind.TransactOpts, streamId *big.Int) (*types.Transaction, error) {
	return _Sablier.contract.Transact(opts, "cancelStream", streamId)
}

// CancelStream is a paid mutator transaction binding the contract method 0x6db9241b.
//
// Solidity: function cancelStream(uint256 streamId) returns(bool)
func (_Sablier *SablierSession) CancelStream(streamId *big.Int) (*types.Transaction, error) {
	return _Sablier.Contract.CancelStream(&_Sablier.TransactOpts, streamId)
}

// CancelStream is a paid mutator transaction binding the contract method 0x6db9241b.
//
// Solidity: function cancelStream(uint256 streamId) returns(bool)
func (_Sablier *SablierTransactorSession) CancelStream(streamId *big.Int) (*types.Transaction, error) {
	return _Sablier.Contract.CancelStream(&_Sablier.TransactOpts, streamId)
}

// CreateStream is a paid mutator transaction binding the contract method 0xcc1b4bf6.
//
// Solidity: function createStream(address recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime) returns(uint256)
func (_Sablier *SablierTransactor) CreateStream(opts *bind.TransactOpts, recipient common.Address, deposit *big.Int, tokenAddress common.Address, startTime *big.Int, stopTime *big.Int) (*types.Transaction, error) {
	return _Sablier.contract.Transact(opts, "createStream", recipient, deposit, tokenAddress, startTime, stopTime)
}

// CreateStream is a paid mutator transaction binding the contract method 0xcc1b4bf6.
//
// Solidity: function createStream(address recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime) returns(uint256)
func (_Sablier *SablierSession) CreateStream(recipient common.Address, deposit *big.Int, tokenAddress common.Address, startTime *big.Int, stopTime *big.Int) (*types.Transaction, error) {
	return _Sablier.Contract.CreateStream(&_Sablier.TransactOpts, recipient, deposit, tokenAddress, startTime, stopTime)
}

// CreateStream is a paid mutator transaction binding the contract method 0xcc1b4bf6.
//
// Solidity: function createStream(address recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime) returns(uint256)
func (_Sablier *SablierTransactorSession) CreateStream(recipient common.Address, deposit *big.Int, tokenAddress common.Address, startTime *big.Int, stopTime *big.Int) (*types.Transaction, error) {
	return _Sablier.Contract.CreateStream(&_Sablier.TransactOpts, recipient, deposit, tokenAddress, startTime, stopTime)
}

// WithdrawFromStream is a paid mutator transaction binding the contract method 0x7a9b2c6c.
//
// Solidity: function withdrawFromStream(uint256 streamId, uint256 amount) returns(bool)
func (_Sablier *SablierTransactor) WithdrawFromStream(opts *bind.TransactOpts, streamId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sablier.contract.Transact(opts, "withdrawFromStream", streamId, amount)
}

// WithdrawFromStream is a paid mutator transaction binding the contract method 0x7a9b2c6c.
//
// Solidity: function withdrawFromStream(uint256 streamId, uint256 amount) returns(bool)
func (_Sablier *SablierSession) WithdrawFromStream(streamId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sablier.Contract.WithdrawFromStream(&_Sablier.TransactOpts, streamId, amount)
}

// WithdrawFromStream is a paid mutator transaction binding the contract method 0x7a9b2c6c.
//
// Solidity: function withdrawFromStream(uint256 streamId, uint256 amount) returns(bool)
func (_Sablier *SablierTransactorSession) WithdrawFromStream(streamId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Sablier.Contract.WithdrawFromStream(&_Sablier.TransactOpts, streamId, amount)
}

// SablierCancelStreamIterator is returned from FilterCancelStream and is used to iterate over the raw logs and unpacked data for CancelStream events raised by the Sablier contract.
type SablierCancelStreamIterator struct {
	Event *SablierCancelStream // Event containing the contract specifics and raw log

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
func (it *SablierCancelStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierCancelStream)
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
		it.Event = new(SablierCancelStream)
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
func (it *SablierCancelStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierCancelStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierCancelStream represents a CancelStream event raised by the Sablier contract.
type SablierCancelStream struct {
	StreamId         *big.Int
	Sender           common.Address
	Recipient        common.Address
	SenderBalance    *big.Int
	RecipientBalance *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancelStream is a free log retrieval operation binding the contract event 0xca3e6079b726e7728802a0537949e2d1c7762304fa641fb06eb56daf2ba8c6b9.
//
// Solidity: event CancelStream(uint256 indexed streamId, address indexed sender, address indexed recipient, uint256 senderBalance, uint256 recipientBalance)
func (_Sablier *SablierFilterer) FilterCancelStream(opts *bind.FilterOpts, streamId []*big.Int, sender []common.Address, recipient []common.Address) (*SablierCancelStreamIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Sablier.contract.FilterLogs(opts, "CancelStream", streamIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SablierCancelStreamIterator{contract: _Sablier.contract, event: "CancelStream", logs: logs, sub: sub}, nil
}

// WatchCancelStream is a free log subscription operation binding the contract event 0xca3e6079b726e7728802a0537949e2d1c7762304fa641fb06eb56daf2ba8c6b9.
//
// Solidity: event CancelStream(uint256 indexed streamId, address indexed sender, address indexed recipient, uint256 senderBalance, uint256 recipientBalance)
func (_Sablier *SablierFilterer) WatchCancelStream(opts *bind.WatchOpts, sink chan<- *SablierCancelStream, streamId []*big.Int, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Sablier.contract.WatchLogs(opts, "CancelStream", streamIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierCancelStream)
				if err := _Sablier.contract.UnpackLog(event, "CancelStream", log); err != nil {
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

// ParseCancelStream is a log parse operation binding the contract event 0xca3e6079b726e7728802a0537949e2d1c7762304fa641fb06eb56daf2ba8c6b9.
//
// Solidity: event CancelStream(uint256 indexed streamId, address indexed sender, address indexed recipient, uint256 senderBalance, uint256 recipientBalance)
func (_Sablier *SablierFilterer) ParseCancelStream(log types.Log) (*SablierCancelStream, error) {
	event := new(SablierCancelStream)
	if err := _Sablier.contract.UnpackLog(event, "CancelStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierCreateStreamIterator is returned from FilterCreateStream and is used to iterate over the raw logs and unpacked data for CreateStream events raised by the Sablier contract.
type SablierCreateStreamIterator struct {
	Event *SablierCreateStream // Event containing the contract specifics and raw log

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
func (it *SablierCreateStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierCreateStream)
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
		it.Event = new(SablierCreateStream)
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
func (it *SablierCreateStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierCreateStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierCreateStream represents a CreateStream event raised by the Sablier contract.
type SablierCreateStream struct {
	StreamId     *big.Int
	Sender       common.Address
	Recipient    common.Address
	Deposit      *big.Int
	TokenAddress common.Address
	StartTime    *big.Int
	StopTime     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateStream is a free log retrieval operation binding the contract event 0x7b01d409597969366dc268d7f957a990d1ca3d3449baf8fb45db67351aecfe78.
//
// Solidity: event CreateStream(uint256 indexed streamId, address indexed sender, address indexed recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime)
func (_Sablier *SablierFilterer) FilterCreateStream(opts *bind.FilterOpts, streamId []*big.Int, sender []common.Address, recipient []common.Address) (*SablierCreateStreamIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Sablier.contract.FilterLogs(opts, "CreateStream", streamIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SablierCreateStreamIterator{contract: _Sablier.contract, event: "CreateStream", logs: logs, sub: sub}, nil
}

// WatchCreateStream is a free log subscription operation binding the contract event 0x7b01d409597969366dc268d7f957a990d1ca3d3449baf8fb45db67351aecfe78.
//
// Solidity: event CreateStream(uint256 indexed streamId, address indexed sender, address indexed recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime)
func (_Sablier *SablierFilterer) WatchCreateStream(opts *bind.WatchOpts, sink chan<- *SablierCreateStream, streamId []*big.Int, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Sablier.contract.WatchLogs(opts, "CreateStream", streamIdRule, senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierCreateStream)
				if err := _Sablier.contract.UnpackLog(event, "CreateStream", log); err != nil {
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

// ParseCreateStream is a log parse operation binding the contract event 0x7b01d409597969366dc268d7f957a990d1ca3d3449baf8fb45db67351aecfe78.
//
// Solidity: event CreateStream(uint256 indexed streamId, address indexed sender, address indexed recipient, uint256 deposit, address tokenAddress, uint256 startTime, uint256 stopTime)
func (_Sablier *SablierFilterer) ParseCreateStream(log types.Log) (*SablierCreateStream, error) {
	event := new(SablierCreateStream)
	if err := _Sablier.contract.UnpackLog(event, "CreateStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SablierWithdrawFromStreamIterator is returned from FilterWithdrawFromStream and is used to iterate over the raw logs and unpacked data for WithdrawFromStream events raised by the Sablier contract.
type SablierWithdrawFromStreamIterator struct {
	Event *SablierWithdrawFromStream // Event containing the contract specifics and raw log

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
func (it *SablierWithdrawFromStreamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SablierWithdrawFromStream)
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
		it.Event = new(SablierWithdrawFromStream)
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
func (it *SablierWithdrawFromStreamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SablierWithdrawFromStreamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SablierWithdrawFromStream represents a WithdrawFromStream event raised by the Sablier contract.
type SablierWithdrawFromStream struct {
	StreamId  *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFromStream is a free log retrieval operation binding the contract event 0x36c3ab437e6a424ed25dc4bfdeb62706aa06558660fab2dab229d2555adaf89c.
//
// Solidity: event WithdrawFromStream(uint256 indexed streamId, address indexed recipient, uint256 amount)
func (_Sablier *SablierFilterer) FilterWithdrawFromStream(opts *bind.FilterOpts, streamId []*big.Int, recipient []common.Address) (*SablierWithdrawFromStreamIterator, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Sablier.contract.FilterLogs(opts, "WithdrawFromStream", streamIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SablierWithdrawFromStreamIterator{contract: _Sablier.contract, event: "WithdrawFromStream", logs: logs, sub: sub}, nil
}

// WatchWithdrawFromStream is a free log subscription operation binding the contract event 0x36c3ab437e6a424ed25dc4bfdeb62706aa06558660fab2dab229d2555adaf89c.
//
// Solidity: event WithdrawFromStream(uint256 indexed streamId, address indexed recipient, uint256 amount)
func (_Sablier *SablierFilterer) WatchWithdrawFromStream(opts *bind.WatchOpts, sink chan<- *SablierWithdrawFromStream, streamId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var streamIdRule []interface{}
	for _, streamIdItem := range streamId {
		streamIdRule = append(streamIdRule, streamIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Sablier.contract.WatchLogs(opts, "WithdrawFromStream", streamIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SablierWithdrawFromStream)
				if err := _Sablier.contract.UnpackLog(event, "WithdrawFromStream", log); err != nil {
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

// ParseWithdrawFromStream is a log parse operation binding the contract event 0x36c3ab437e6a424ed25dc4bfdeb62706aa06558660fab2dab229d2555adaf89c.
//
// Solidity: event WithdrawFromStream(uint256 indexed streamId, address indexed recipient, uint256 amount)
func (_Sablier *SablierFilterer) ParseWithdrawFromStream(log types.Log) (*SablierWithdrawFromStream, error) {
	event := new(SablierWithdrawFromStream)
	if err := _Sablier.contract.UnpackLog(event, "WithdrawFromStream", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
