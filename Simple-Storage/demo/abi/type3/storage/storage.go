// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"StoreData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"queryData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610496806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634ece5b4c1461003b5780639027696b14610050575b600080fd5b61004e6100493660046102ee565b610079565b005b61006361005e3660046102b3565b6100e9565b6040516100709190610397565b60405180910390f35b8060008360405161008a919061037b565b908152602001604051809103902090805190602001906100ab929190610199565b507f5db62cb3c537e53fcb4f478359589712e4ecfe4434ab6aa7acd408c3be0eb1f982826040516100dd9291906103b1565b60405180910390a15050565b60606000826040516100fb919061037b565b908152602001604051809103902080546101149061040f565b80601f01602080910402602001604051908101604052809291908181526020018280546101409061040f565b801561018d5780601f106101625761010080835404028352916020019161018d565b820191906000526020600020905b81548152906001019060200180831161017057829003601f168201915b50505050509050919050565b8280546101a59061040f565b90600052602060002090601f0160209004810192826101c7576000855561020d565b82601f106101e057805160ff191683800117855561020d565b8280016001018555821561020d579182015b8281111561020d5782518255916020019190600101906101f2565b5061021992915061021d565b5090565b5b80821115610219576000815560010161021e565b600082601f830112610242578081fd5b813567ffffffffffffffff8082111561025d5761025d61044a565b604051601f8301601f1916810160200182811182821017156102815761028161044a565b604052828152848301602001861015610298578384fd5b82602086016020830137918201602001929092529392505050565b6000602082840312156102c4578081fd5b813567ffffffffffffffff8111156102da578182fd5b6102e684828501610232565b949350505050565b60008060408385031215610300578081fd5b823567ffffffffffffffff80821115610317578283fd5b61032386838701610232565b93506020850135915080821115610338578283fd5b5061034585828601610232565b9150509250929050565b600081518084526103678160208601602086016103df565b601f01601f19169290920160200192915050565b6000825161038d8184602087016103df565b9190910192915050565b6000602082526103aa602083018461034f565b9392505050565b6000604082526103c4604083018561034f565b82810360208401526103d6818561034f565b95945050505050565b60005b838110156103fa5781810151838201526020016103e2565b83811115610409576000848401525b50505050565b60028104600182168061042357607f821691505b6020821081141561044457634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea26469706673582212202573f89da75d182024c865cd3206152d95651898ec9f81c020724e9e3565314d64736f6c63430008000033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// QueryData is a free data retrieval call binding the contract method 0x9027696b.
//
// Solidity: function queryData(string key) view returns(string)
func (_Storage *StorageCaller) QueryData(opts *bind.CallOpts, key string) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "queryData", key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryData is a free data retrieval call binding the contract method 0x9027696b.
//
// Solidity: function queryData(string key) view returns(string)
func (_Storage *StorageSession) QueryData(key string) (string, error) {
	return _Storage.Contract.QueryData(&_Storage.CallOpts, key)
}

// QueryData is a free data retrieval call binding the contract method 0x9027696b.
//
// Solidity: function queryData(string key) view returns(string)
func (_Storage *StorageCallerSession) QueryData(key string) (string, error) {
	return _Storage.Contract.QueryData(&_Storage.CallOpts, key)
}

// StoreData is a paid mutator transaction binding the contract method 0x4ece5b4c.
//
// Solidity: function storeData(string key, string value) returns()
func (_Storage *StorageTransactor) StoreData(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "storeData", key, value)
}

// StoreData is a paid mutator transaction binding the contract method 0x4ece5b4c.
//
// Solidity: function storeData(string key, string value) returns()
func (_Storage *StorageSession) StoreData(key string, value string) (*types.Transaction, error) {
	return _Storage.Contract.StoreData(&_Storage.TransactOpts, key, value)
}

// StoreData is a paid mutator transaction binding the contract method 0x4ece5b4c.
//
// Solidity: function storeData(string key, string value) returns()
func (_Storage *StorageTransactorSession) StoreData(key string, value string) (*types.Transaction, error) {
	return _Storage.Contract.StoreData(&_Storage.TransactOpts, key, value)
}

// StorageStoreDataIterator is returned from FilterStoreData and is used to iterate over the raw logs and unpacked data for StoreData events raised by the Storage contract.
type StorageStoreDataIterator struct {
	Event *StorageStoreData // Event containing the contract specifics and raw log

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
func (it *StorageStoreDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageStoreData)
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
		it.Event = new(StorageStoreData)
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
func (it *StorageStoreDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageStoreDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageStoreData represents a StoreData event raised by the Storage contract.
type StorageStoreData struct {
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStoreData is a free log retrieval operation binding the contract event 0x5db62cb3c537e53fcb4f478359589712e4ecfe4434ab6aa7acd408c3be0eb1f9.
//
// Solidity: event StoreData(string key, string value)
func (_Storage *StorageFilterer) FilterStoreData(opts *bind.FilterOpts) (*StorageStoreDataIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "StoreData")
	if err != nil {
		return nil, err
	}
	return &StorageStoreDataIterator{contract: _Storage.contract, event: "StoreData", logs: logs, sub: sub}, nil
}

// WatchStoreData is a free log subscription operation binding the contract event 0x5db62cb3c537e53fcb4f478359589712e4ecfe4434ab6aa7acd408c3be0eb1f9.
//
// Solidity: event StoreData(string key, string value)
func (_Storage *StorageFilterer) WatchStoreData(opts *bind.WatchOpts, sink chan<- *StorageStoreData) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "StoreData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageStoreData)
				if err := _Storage.contract.UnpackLog(event, "StoreData", log); err != nil {
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

// ParseStoreData is a log parse operation binding the contract event 0x5db62cb3c537e53fcb4f478359589712e4ecfe4434ab6aa7acd408c3be0eb1f9.
//
// Solidity: event StoreData(string key, string value)
func (_Storage *StorageFilterer) ParseStoreData(log types.Log) (*StorageStoreData, error) {
	event := new(StorageStoreData)
	if err := _Storage.contract.UnpackLog(event, "StoreData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
