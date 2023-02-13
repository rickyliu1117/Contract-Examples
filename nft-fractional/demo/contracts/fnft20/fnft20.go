// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fnft20

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

// TokenFractional20MetaData contains all meta data concerning the TokenFractional20 contract.
var TokenFractional20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Compose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Fractional\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"compose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fractional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200165938038062001659833981016040819052620000349162000193565b818160036200004483826200028c565b5060046200005382826200028c565b505050620000706200006a6200007860201b60201c565b6200007c565b505062000358565b3390565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000f657600080fd5b81516001600160401b0380821115620001135762000113620000ce565b604051601f8301601f19908116603f011681019082821181831017156200013e576200013e620000ce565b816040528381526020925086838588010111156200015b57600080fd5b600091505b838210156200017f578582018301518183018401529082019062000160565b600093810190920192909252949350505050565b60008060408385031215620001a757600080fd5b82516001600160401b0380821115620001bf57600080fd5b620001cd86838701620000e4565b93506020850151915080821115620001e457600080fd5b50620001f385828601620000e4565b9150509250929050565b600181811c908216806200021257607f821691505b6020821081036200023357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028757600081815260208120601f850160051c81016020861015620002625750805b601f850160051c820191505b8181101562000283578281556001016200026e565b5050505b505050565b81516001600160401b03811115620002a857620002a8620000ce565b620002c081620002b98454620001fd565b8462000239565b602080601f831160018114620002f85760008415620002df5750858301515b600019600386901b1c1916600185901b17855562000283565b600085815260208120601f198616915b82811015620003295788860151825594840194600190910190840162000308565b5085821015620003485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6112f180620003686000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c8063715018a6116100a2578063a457c2d711610071578063a457c2d714610214578063a9059cbb14610227578063dd62ed3e1461023a578063f2fde38b1461024d578063fc0c546a1461026057600080fd5b8063715018a6146101d657806371772471146101de5780638da5cb5b146101f157806395d89b411461020c57600080fd5b8063313ce567116100de578063313ce5671461017657806339509351146101855780633d636c831461019857806370a08231146101ad57600080fd5b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015157806323b872dd14610163575b600080fd5b610118610287565b60405161012591906110c8565b60405180910390f35b61014161013c36600461112b565b610319565b6040519015158152602001610125565b6002545b604051908152602001610125565b610141610171366004611157565b610333565b60405160128152602001610125565b61014161019336600461112b565b610357565b6101ab6101a6366004611198565b610379565b005b6101556101bb366004611198565b6001600160a01b031660009081526020819052604090205490565b6101ab610611565b6101ab6101ec3660046111bc565b610625565b6005546040516001600160a01b039091168152602001610125565b61011861095c565b61014161022236600461112b565b61096b565b61014161023536600461112b565b6109e6565b6101556102483660046111f1565b6109f4565b6101ab61025b366004611198565b610a1f565b610268610a98565b604080516001600160a01b039093168352602083019190915201610125565b6060600380546102969061122a565b80601f01602080910402602001604051908101604052809291908181526020018280546102c29061122a565b801561030f5780601f106102e45761010080835404028352916020019161030f565b820191906000526020600020905b8154815290600101906020018083116102f257829003601f168201915b5050505050905090565b600033610327818585610af9565b60019150505b92915050565b600033610341858285610c15565b61034c858585610c8f565b506001949350505050565b60003361032781858561036a83836109f4565b6103749190611264565b610af9565b6001600160a01b0381166103d45760405162461bcd60e51b815260206004820152601a60248201527f746f2063616e206e6f74206265207a65726f206164647265737300000000000060448201526064015b60405180910390fd5b6000339050306001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610417573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043b9190611285565b6040516370a0823160e01b81526001600160a01b038316600482015230906370a0823190602401602060405180830381865afa15801561047f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a39190611285565b146104f05760405162461bcd60e51b815260206004820152601960248201527f4e6f74206f776e696e6720616c6c2074686520746f6b656e730000000000000060448201526064016103cb565b6006546007546040516323b872dd60e01b81523060048201526001600160a01b03858116602483015260448201929092529116906323b872dd90606401600060405180830381600087803b15801561054757600080fd5b505af115801561055b573d6000803e3d6000fd5b505050506105ca81306001600160a01b03166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c59190611285565b610e33565b6008805460ff191690556040516001600160a01b0380841691908316907f577991118e5e746b33f8f0107b7931b64a47547ddeb8c70633eabd9ba6cfdc2f90600090a35050565b610619610f5d565b6106236000610fb7565b565b61062d610f5d565b60085460ff16156106805760405162461bcd60e51b815260206004820152601d60248201527f636f6e7472616374206f6e6c79206672616374696f6e616c206f6e636500000060448201526064016103cb565b600081116106d05760405162461bcd60e51b815260206004820152601d60248201527f616d6f756e74206d7573742067726561746572207468616e207a65726f00000060448201526064016103cb565b600680546001600160a01b0319166001600160a01b038516179055336006546040516331a9108f60e11b8152600481018590526001600160a01b039283169290911690636352211e90602401602060405180830381865afa158015610739573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075d919061129e565b6001600160a01b0316146107b35760405162461bcd60e51b815260206004820152601860248201527f63616c6c6572206d757374206265204e4654206f776e6572000000000000000060448201526064016103cb565b6006546001600160a01b03166323b872dd336040516001600160e01b031960e084901b1681526001600160a01b03909116600482015230602482015260448101859052606401600060405180830381600087803b15801561081357600080fd5b505af1158015610827573d6000803e3d6000fd5b50506006546040516331a9108f60e11b8152600481018690523093506001600160a01b039091169150636352211e90602401602060405180830381865afa158015610876573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089a919061129e565b6001600160a01b0316146108df5760405162461bcd60e51b815260206004820152600c60248201526b18da1958dac819985a5b195960a21b60448201526064016103cb565b60078290556108ee3382611009565b6008805460ff191660011790556109023390565b6001600160a01b0316836001600160a01b03167fd7a388001e99d27eeb3c1cc2c4c55687c27f84e3c1b4c729e472f8079ec0a063848460405161094f929190918252602082015260400190565b60405180910390a3505050565b6060600480546102969061122a565b6000338161097982866109f4565b9050838110156109d95760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016103cb565b61034c8286868403610af9565b600033610327818585610c8f565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b610a27610f5d565b6001600160a01b038116610a8c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103cb565b610a9581610fb7565b50565b600854600090819060ff16610ae35760405162461bcd60e51b815260206004820152601160248201527013919508191bc81b9bdd081a5b9a5d1959607a1b60448201526064016103cb565b50506006546007546001600160a01b0390911691565b6001600160a01b038316610b5b5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103cb565b6001600160a01b038216610bbc5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103cb565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910161094f565b6000610c2184846109f4565b90506000198114610c895781811015610c7c5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103cb565b610c898484848403610af9565b50505050565b6001600160a01b038316610cf35760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103cb565b6001600160a01b038216610d555760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103cb565b6001600160a01b03831660009081526020819052604090205481811015610dcd5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103cb565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610c89565b6001600160a01b038216610e935760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103cb565b6001600160a01b03821660009081526020819052604090205481811015610f075760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103cb565b6001600160a01b0383166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910161094f565b6005546001600160a01b031633146106235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103cb565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03821661105f5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103cb565b80600260008282546110719190611264565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b818110156110f5578581018301518582016040015282016110d9565b506000604082860101526040601f19601f8301168501019250505092915050565b6001600160a01b0381168114610a9557600080fd5b6000806040838503121561113e57600080fd5b823561114981611116565b946020939093013593505050565b60008060006060848603121561116c57600080fd5b833561117781611116565b9250602084013561118781611116565b929592945050506040919091013590565b6000602082840312156111aa57600080fd5b81356111b581611116565b9392505050565b6000806000606084860312156111d157600080fd5b83356111dc81611116565b95602085013595506040909401359392505050565b6000806040838503121561120457600080fd5b823561120f81611116565b9150602083013561121f81611116565b809150509250929050565b600181811c9082168061123e57607f821691505b60208210810361125e57634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561032d57634e487b7160e01b600052601160045260246000fd5b60006020828403121561129757600080fd5b5051919050565b6000602082840312156112b057600080fd5b81516111b58161111656fea2646970667358221220335da6eefdb1f3abe9865f78445674323886b614e0b67567bfbc13080fdfefa864736f6c63430008110033",
}

// TokenFractional20ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenFractional20MetaData.ABI instead.
var TokenFractional20ABI = TokenFractional20MetaData.ABI

// TokenFractional20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenFractional20MetaData.Bin instead.
var TokenFractional20Bin = TokenFractional20MetaData.Bin

// DeployTokenFractional20 deploys a new Ethereum contract, binding an instance of TokenFractional20 to it.
func DeployTokenFractional20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *TokenFractional20, error) {
	parsed, err := TokenFractional20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenFractional20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenFractional20{TokenFractional20Caller: TokenFractional20Caller{contract: contract}, TokenFractional20Transactor: TokenFractional20Transactor{contract: contract}, TokenFractional20Filterer: TokenFractional20Filterer{contract: contract}}, nil
}

// TokenFractional20 is an auto generated Go binding around an Ethereum contract.
type TokenFractional20 struct {
	TokenFractional20Caller     // Read-only binding to the contract
	TokenFractional20Transactor // Write-only binding to the contract
	TokenFractional20Filterer   // Log filterer for contract events
}

// TokenFractional20Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFractional20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFractional20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFractional20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFractional20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFractional20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFractional20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFractional20Session struct {
	Contract     *TokenFractional20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TokenFractional20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFractional20CallerSession struct {
	Contract *TokenFractional20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TokenFractional20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFractional20TransactorSession struct {
	Contract     *TokenFractional20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TokenFractional20Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFractional20Raw struct {
	Contract *TokenFractional20 // Generic contract binding to access the raw methods on
}

// TokenFractional20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFractional20CallerRaw struct {
	Contract *TokenFractional20Caller // Generic read-only contract binding to access the raw methods on
}

// TokenFractional20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFractional20TransactorRaw struct {
	Contract *TokenFractional20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFractional20 creates a new instance of TokenFractional20, bound to a specific deployed contract.
func NewTokenFractional20(address common.Address, backend bind.ContractBackend) (*TokenFractional20, error) {
	contract, err := bindTokenFractional20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20{TokenFractional20Caller: TokenFractional20Caller{contract: contract}, TokenFractional20Transactor: TokenFractional20Transactor{contract: contract}, TokenFractional20Filterer: TokenFractional20Filterer{contract: contract}}, nil
}

// NewTokenFractional20Caller creates a new read-only instance of TokenFractional20, bound to a specific deployed contract.
func NewTokenFractional20Caller(address common.Address, caller bind.ContractCaller) (*TokenFractional20Caller, error) {
	contract, err := bindTokenFractional20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20Caller{contract: contract}, nil
}

// NewTokenFractional20Transactor creates a new write-only instance of TokenFractional20, bound to a specific deployed contract.
func NewTokenFractional20Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenFractional20Transactor, error) {
	contract, err := bindTokenFractional20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20Transactor{contract: contract}, nil
}

// NewTokenFractional20Filterer creates a new log filterer instance of TokenFractional20, bound to a specific deployed contract.
func NewTokenFractional20Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenFractional20Filterer, error) {
	contract, err := bindTokenFractional20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20Filterer{contract: contract}, nil
}

// bindTokenFractional20 binds a generic wrapper to an already deployed contract.
func bindTokenFractional20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFractional20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFractional20 *TokenFractional20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFractional20.Contract.TokenFractional20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFractional20 *TokenFractional20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFractional20.Contract.TokenFractional20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFractional20 *TokenFractional20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFractional20.Contract.TokenFractional20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFractional20 *TokenFractional20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenFractional20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFractional20 *TokenFractional20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFractional20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFractional20 *TokenFractional20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFractional20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenFractional20 *TokenFractional20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenFractional20 *TokenFractional20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenFractional20.Contract.Allowance(&_TokenFractional20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TokenFractional20 *TokenFractional20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TokenFractional20.Contract.Allowance(&_TokenFractional20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenFractional20 *TokenFractional20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenFractional20 *TokenFractional20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenFractional20.Contract.BalanceOf(&_TokenFractional20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TokenFractional20 *TokenFractional20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TokenFractional20.Contract.BalanceOf(&_TokenFractional20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFractional20 *TokenFractional20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFractional20 *TokenFractional20Session) Decimals() (uint8, error) {
	return _TokenFractional20.Contract.Decimals(&_TokenFractional20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TokenFractional20 *TokenFractional20CallerSession) Decimals() (uint8, error) {
	return _TokenFractional20.Contract.Decimals(&_TokenFractional20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFractional20 *TokenFractional20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFractional20 *TokenFractional20Session) Name() (string, error) {
	return _TokenFractional20.Contract.Name(&_TokenFractional20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenFractional20 *TokenFractional20CallerSession) Name() (string, error) {
	return _TokenFractional20.Contract.Name(&_TokenFractional20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFractional20 *TokenFractional20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFractional20 *TokenFractional20Session) Owner() (common.Address, error) {
	return _TokenFractional20.Contract.Owner(&_TokenFractional20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenFractional20 *TokenFractional20CallerSession) Owner() (common.Address, error) {
	return _TokenFractional20.Contract.Owner(&_TokenFractional20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFractional20 *TokenFractional20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFractional20 *TokenFractional20Session) Symbol() (string, error) {
	return _TokenFractional20.Contract.Symbol(&_TokenFractional20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenFractional20 *TokenFractional20CallerSession) Symbol() (string, error) {
	return _TokenFractional20.Contract.Symbol(&_TokenFractional20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address, uint256)
func (_TokenFractional20 *TokenFractional20Caller) Token(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address, uint256)
func (_TokenFractional20 *TokenFractional20Session) Token() (common.Address, *big.Int, error) {
	return _TokenFractional20.Contract.Token(&_TokenFractional20.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address, uint256)
func (_TokenFractional20 *TokenFractional20CallerSession) Token() (common.Address, *big.Int, error) {
	return _TokenFractional20.Contract.Token(&_TokenFractional20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFractional20 *TokenFractional20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenFractional20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFractional20 *TokenFractional20Session) TotalSupply() (*big.Int, error) {
	return _TokenFractional20.Contract.TotalSupply(&_TokenFractional20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenFractional20 *TokenFractional20CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenFractional20.Contract.TotalSupply(&_TokenFractional20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Approve(&_TokenFractional20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Approve(&_TokenFractional20.TransactOpts, spender, amount)
}

// Compose is a paid mutator transaction binding the contract method 0x3d636c83.
//
// Solidity: function compose(address to) returns()
func (_TokenFractional20 *TokenFractional20Transactor) Compose(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "compose", to)
}

// Compose is a paid mutator transaction binding the contract method 0x3d636c83.
//
// Solidity: function compose(address to) returns()
func (_TokenFractional20 *TokenFractional20Session) Compose(to common.Address) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Compose(&_TokenFractional20.TransactOpts, to)
}

// Compose is a paid mutator transaction binding the contract method 0x3d636c83.
//
// Solidity: function compose(address to) returns()
func (_TokenFractional20 *TokenFractional20TransactorSession) Compose(to common.Address) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Compose(&_TokenFractional20.TransactOpts, to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenFractional20 *TokenFractional20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenFractional20 *TokenFractional20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.DecreaseAllowance(&_TokenFractional20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TokenFractional20 *TokenFractional20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.DecreaseAllowance(&_TokenFractional20.TransactOpts, spender, subtractedValue)
}

// Fractional is a paid mutator transaction binding the contract method 0x71772471.
//
// Solidity: function fractional(address erc721, uint256 tokenId, uint256 amount) returns()
func (_TokenFractional20 *TokenFractional20Transactor) Fractional(opts *bind.TransactOpts, erc721 common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "fractional", erc721, tokenId, amount)
}

// Fractional is a paid mutator transaction binding the contract method 0x71772471.
//
// Solidity: function fractional(address erc721, uint256 tokenId, uint256 amount) returns()
func (_TokenFractional20 *TokenFractional20Session) Fractional(erc721 common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Fractional(&_TokenFractional20.TransactOpts, erc721, tokenId, amount)
}

// Fractional is a paid mutator transaction binding the contract method 0x71772471.
//
// Solidity: function fractional(address erc721, uint256 tokenId, uint256 amount) returns()
func (_TokenFractional20 *TokenFractional20TransactorSession) Fractional(erc721 common.Address, tokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Fractional(&_TokenFractional20.TransactOpts, erc721, tokenId, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenFractional20 *TokenFractional20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenFractional20 *TokenFractional20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.IncreaseAllowance(&_TokenFractional20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TokenFractional20 *TokenFractional20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.IncreaseAllowance(&_TokenFractional20.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFractional20 *TokenFractional20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFractional20 *TokenFractional20Session) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFractional20.Contract.RenounceOwnership(&_TokenFractional20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenFractional20 *TokenFractional20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenFractional20.Contract.RenounceOwnership(&_TokenFractional20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Transfer(&_TokenFractional20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.Transfer(&_TokenFractional20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.TransferFrom(&_TokenFractional20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TokenFractional20 *TokenFractional20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenFractional20.Contract.TransferFrom(&_TokenFractional20.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFractional20 *TokenFractional20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenFractional20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFractional20 *TokenFractional20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFractional20.Contract.TransferOwnership(&_TokenFractional20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenFractional20 *TokenFractional20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenFractional20.Contract.TransferOwnership(&_TokenFractional20.TransactOpts, newOwner)
}

// TokenFractional20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenFractional20 contract.
type TokenFractional20ApprovalIterator struct {
	Event *TokenFractional20Approval // Event containing the contract specifics and raw log

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
func (it *TokenFractional20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional20Approval)
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
		it.Event = new(TokenFractional20Approval)
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
func (it *TokenFractional20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional20Approval represents a Approval event raised by the TokenFractional20 contract.
type TokenFractional20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenFractional20 *TokenFractional20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenFractional20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFractional20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20ApprovalIterator{contract: _TokenFractional20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TokenFractional20 *TokenFractional20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenFractional20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFractional20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional20Approval)
				if err := _TokenFractional20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_TokenFractional20 *TokenFractional20Filterer) ParseApproval(log types.Log) (*TokenFractional20Approval, error) {
	event := new(TokenFractional20Approval)
	if err := _TokenFractional20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional20ComposeIterator is returned from FilterCompose and is used to iterate over the raw logs and unpacked data for Compose events raised by the TokenFractional20 contract.
type TokenFractional20ComposeIterator struct {
	Event *TokenFractional20Compose // Event containing the contract specifics and raw log

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
func (it *TokenFractional20ComposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional20Compose)
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
		it.Event = new(TokenFractional20Compose)
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
func (it *TokenFractional20ComposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional20ComposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional20Compose represents a Compose event raised by the TokenFractional20 contract.
type TokenFractional20Compose struct {
	Sender common.Address
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCompose is a free log retrieval operation binding the contract event 0x577991118e5e746b33f8f0107b7931b64a47547ddeb8c70633eabd9ba6cfdc2f.
//
// Solidity: event Compose(address indexed sender, address indexed to)
func (_TokenFractional20 *TokenFractional20Filterer) FilterCompose(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*TokenFractional20ComposeIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional20.contract.FilterLogs(opts, "Compose", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20ComposeIterator{contract: _TokenFractional20.contract, event: "Compose", logs: logs, sub: sub}, nil
}

// WatchCompose is a free log subscription operation binding the contract event 0x577991118e5e746b33f8f0107b7931b64a47547ddeb8c70633eabd9ba6cfdc2f.
//
// Solidity: event Compose(address indexed sender, address indexed to)
func (_TokenFractional20 *TokenFractional20Filterer) WatchCompose(opts *bind.WatchOpts, sink chan<- *TokenFractional20Compose, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional20.contract.WatchLogs(opts, "Compose", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional20Compose)
				if err := _TokenFractional20.contract.UnpackLog(event, "Compose", log); err != nil {
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

// ParseCompose is a log parse operation binding the contract event 0x577991118e5e746b33f8f0107b7931b64a47547ddeb8c70633eabd9ba6cfdc2f.
//
// Solidity: event Compose(address indexed sender, address indexed to)
func (_TokenFractional20 *TokenFractional20Filterer) ParseCompose(log types.Log) (*TokenFractional20Compose, error) {
	event := new(TokenFractional20Compose)
	if err := _TokenFractional20.contract.UnpackLog(event, "Compose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional20FractionalIterator is returned from FilterFractional and is used to iterate over the raw logs and unpacked data for Fractional events raised by the TokenFractional20 contract.
type TokenFractional20FractionalIterator struct {
	Event *TokenFractional20Fractional // Event containing the contract specifics and raw log

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
func (it *TokenFractional20FractionalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional20Fractional)
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
		it.Event = new(TokenFractional20Fractional)
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
func (it *TokenFractional20FractionalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional20FractionalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional20Fractional represents a Fractional event raised by the TokenFractional20 contract.
type TokenFractional20Fractional struct {
	Erc721  common.Address
	Spender common.Address
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFractional is a free log retrieval operation binding the contract event 0xd7a388001e99d27eeb3c1cc2c4c55687c27f84e3c1b4c729e472f8079ec0a063.
//
// Solidity: event Fractional(address indexed erc721, address indexed spender, uint256 tokenId, uint256 amount)
func (_TokenFractional20 *TokenFractional20Filterer) FilterFractional(opts *bind.FilterOpts, erc721 []common.Address, spender []common.Address) (*TokenFractional20FractionalIterator, error) {

	var erc721Rule []interface{}
	for _, erc721Item := range erc721 {
		erc721Rule = append(erc721Rule, erc721Item)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFractional20.contract.FilterLogs(opts, "Fractional", erc721Rule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20FractionalIterator{contract: _TokenFractional20.contract, event: "Fractional", logs: logs, sub: sub}, nil
}

// WatchFractional is a free log subscription operation binding the contract event 0xd7a388001e99d27eeb3c1cc2c4c55687c27f84e3c1b4c729e472f8079ec0a063.
//
// Solidity: event Fractional(address indexed erc721, address indexed spender, uint256 tokenId, uint256 amount)
func (_TokenFractional20 *TokenFractional20Filterer) WatchFractional(opts *bind.WatchOpts, sink chan<- *TokenFractional20Fractional, erc721 []common.Address, spender []common.Address) (event.Subscription, error) {

	var erc721Rule []interface{}
	for _, erc721Item := range erc721 {
		erc721Rule = append(erc721Rule, erc721Item)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TokenFractional20.contract.WatchLogs(opts, "Fractional", erc721Rule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional20Fractional)
				if err := _TokenFractional20.contract.UnpackLog(event, "Fractional", log); err != nil {
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

// ParseFractional is a log parse operation binding the contract event 0xd7a388001e99d27eeb3c1cc2c4c55687c27f84e3c1b4c729e472f8079ec0a063.
//
// Solidity: event Fractional(address indexed erc721, address indexed spender, uint256 tokenId, uint256 amount)
func (_TokenFractional20 *TokenFractional20Filterer) ParseFractional(log types.Log) (*TokenFractional20Fractional, error) {
	event := new(TokenFractional20Fractional)
	if err := _TokenFractional20.contract.UnpackLog(event, "Fractional", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenFractional20 contract.
type TokenFractional20OwnershipTransferredIterator struct {
	Event *TokenFractional20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenFractional20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional20OwnershipTransferred)
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
		it.Event = new(TokenFractional20OwnershipTransferred)
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
func (it *TokenFractional20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional20OwnershipTransferred represents a OwnershipTransferred event raised by the TokenFractional20 contract.
type TokenFractional20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFractional20 *TokenFractional20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenFractional20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFractional20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20OwnershipTransferredIterator{contract: _TokenFractional20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenFractional20 *TokenFractional20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenFractional20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenFractional20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional20OwnershipTransferred)
				if err := _TokenFractional20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenFractional20 *TokenFractional20Filterer) ParseOwnershipTransferred(log types.Log) (*TokenFractional20OwnershipTransferred, error) {
	event := new(TokenFractional20OwnershipTransferred)
	if err := _TokenFractional20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenFractional20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenFractional20 contract.
type TokenFractional20TransferIterator struct {
	Event *TokenFractional20Transfer // Event containing the contract specifics and raw log

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
func (it *TokenFractional20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFractional20Transfer)
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
		it.Event = new(TokenFractional20Transfer)
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
func (it *TokenFractional20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFractional20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFractional20Transfer represents a Transfer event raised by the TokenFractional20 contract.
type TokenFractional20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenFractional20 *TokenFractional20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenFractional20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenFractional20TransferIterator{contract: _TokenFractional20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TokenFractional20 *TokenFractional20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenFractional20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenFractional20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFractional20Transfer)
				if err := _TokenFractional20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_TokenFractional20 *TokenFractional20Filterer) ParseTransfer(log types.Log) (*TokenFractional20Transfer, error) {
	event := new(TokenFractional20Transfer)
	if err := _TokenFractional20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
