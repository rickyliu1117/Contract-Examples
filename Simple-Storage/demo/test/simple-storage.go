package test

import (
	"errors"
	"fmt"

	"storage/eth"
	"storage/log"
	storage "storage/abi/type3/storage"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const NodeUrl = "http://52.83.148.57:20001"
const NodeWsUrl = "ws://52.83.148.57:8551"
const Address = "0xDDA883936Ac4044C7EdCE7aF5d5a61F7e594dFE3"
const PrivateKey = "79c57b4d61fbcb281312ddb4d1fde89e3479965b21d8c2804c74b4f5db2d47f9"

var key1 = "key1"
var value1 = "bb"
var value2 = "cc"

func TestStoreData(t *testing.T) {
	log.InitLog()
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Logger.Error(err)
	}
	instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
	if err != nil {
		log.Logger.Error(err)
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Logger.Error(err)
	}
	tx, err := instance.StoreData(auth, key1, value1)
	if err != nil {
		log.Logger.Error(err)
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}

func TestDeleteData(t *testing.T) {
	log.InitLog()
    cli, err := ethclient.Dial(NodeUrl)
    if err != nil {
        log.Logger.Error(err)
    }
    instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
    if err != nil {
        log.Logger.Error(err)
    }
    auth, err := eth.GenAuth(cli, PrivateKey)
    if err != nil {
        log.Logger.Error(err)
    }
    tx, err := instance.StoreData(auth, key1, "")
    if err != nil {
        log.Logger.Error(err)
    }
    fmt.Println("tx Hash:", tx.Hash().String())
}

func TestQueryData(t *testing.T) {
	log.InitLog()
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Logger.Error(err)
	}
	instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
	if err != nil {
		log.Logger.Error(err)
	}
	value, err := instance.QueryData(nil, key1)
	if err != nil {
		log.Logger.Error(err)
	}
	fmt.Println("value:", value)
}

func TestUpdateData(t *testing.T) {
	log.InitLog()
    cli, err := ethclient.Dial(NodeUrl)
    if err != nil {
        log.Logger.Error(err)
    }
    instance, err := storage.NewStorage(common.HexToAddress(Address), cli)
    if err != nil {
        log.Logger.Error(err)
    }
	value, err := instance.QueryData(nil, key1)
    if err != nil {
        log.Logger.Error(err)
    }
	if value == "" {
		log.Logger.Error(errors.New("key1 has no corresponding value in the contract and will not be an update operation"))
	}
	if value == value2 {
		log.Logger.Error(errors.New("the value2 to be updated is the same as the existing value, no update is required"))
	}
    auth, err := eth.GenAuth(cli, PrivateKey)
    if err != nil {
        log.Logger.Error(err)
    }
    tx, err := instance.StoreData(auth, key1, value2)
    if err != nil {
        log.Logger.Error(err)
    }
    fmt.Println("tx Hash:", tx.Hash().String())
}

func TestCaptureStoreDataEvent(t *testing.T) {
	log.InitLog()
	var storageStoreData = make(chan *storage.StorageStoreData, 100)
	wsCli, err := ethclient.Dial(NodeWsUrl)
	if err != nil {
		log.Logger.Error(err)
	}
	wsInstance, err := storage.NewStorage(common.HexToAddress(Address), wsCli)
	if err != nil {
		log.Logger.Error(err)
	}
	go func() {
	hear:
		transfer, err := wsInstance.StorageFilterer.WatchStoreData(nil, storageStoreData)
		if err != nil {
			log.Logger.Error(err)
			goto hear
		}
		defer transfer.Unsubscribe()
		select {
		case err := <-transfer.Err():
			log.Logger.Warn(err)
			goto hear
		}
	}()
	for {
		select {
		case event := <-storageStoreData:
			fmt.Println("StoreData key:", event.Key)
			fmt.Println("StoreData value:", event.Value)
		}
	}
}