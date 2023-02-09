# 永久存储合约

永久存储是区块链技术的一大特点，传统 IT 系统内的业务数据存储均依赖于先有存储空间，意味着用户需要先购买存储资源；其次如果用户使用的是云资源也面临着需要长期续费，否则用户的整个云主机都将被卸载，业务数据终将丢失。而 BSN 定义的公共 IT 系统，基于区块链的特性和无币公链、开放联盟链的自有特点，用户只有在部署合约和发送交易的时候承担非常少的 GAS （注意：不同于公有链上的 GAS）即可，数据永久存储于链上，不调用或者调用查询方法都不产生任何费用。

## 1. 使用说明 
* 用户向链上提交的数据都将是公开可查的，所以业务数据如涉及到敏感信息，则需要用户在发送交易之前，在业务系统内先对业务数据进行加密处理，然后将密文数据发送到链上，从而确保原始的敏感数据不被公开，在调用查询方法获取到链上的密文数据后，业务系统内做解密处理得到原始明文数据。

* 用户向链上发送交易产生的 GAS 费用，与发送的数据大小有一定的关系，所以用户可以在发送交易之前，在业务系统内对待发送的数据进行压缩处理，从而达到减少 GAS 费用的目的。

## 2.合约概述
本合约支持用户将业务数据提交到区块链网络环境内进行存储，同时可以对其进行删除、修改、查询。合约内仅提供 storeData、queryData 两个方法，业务系统在调用时传入不同的参数值，可以对业务数据进行不同的操作。

* 数据上链：业务系统内自定义 key 值，将业务数据作为 value，调用 storeData 方法可完成数据上链。
* 删除：业务系统内指定需要删除的业务数据的 key 值，将空字符串作为 value，调用 storeData 方法可完成数据的删除。
* 修改：业务系统内指定需要修改的业务数据的 key 值，将新的业务数据作为 value，调用 storeData 方法可完成数据的修改。
* 查询：业务系统内指定需要查询的业务数据的 key 值，调用 queryData 方法可查到链上存储的业务数据。

## 3.合约方法

编号  |  方法声明   |  方法描述 
----|----|----
1 | *storeData(string memory key, string memory value) public* |      key-value 形式操作业务数据              
2 | *queryData(string memory key) public view returns (string memory)* |      根据 key 获取对应的业务数据     


## 4.Go语言调用示例

### 1）数据上链
```
func TestStoreData(t *testing.T) {
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
```

### 2）删除数据

```
func TestDeleteData(t *testing.T) {
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
```

### 3）修改数据

```
func TestUpdateData(t *testing.T) {
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
```

### 4）查询数据
```
func TestQueryData(t *testing.T) {
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
```

### 5）捕获解析StoreData事件示例
```
func TestCaptureStoreDataEvent(t *testing.T) {
	var storageStoreData = make(chan *storage.StorageStoreData, 100)
	wsCli, err := ethclient.Dial(NodeWsUrl)
    if err != nil {
        log.Logger.Error(err)
    }
	wsInstance, err := storage.NewStorage(common.HexToAddress(Address),wsCli)
	if err != nil {
		log.Logger.Error(err)
	}
	go func() {
	hear:
		transfer, err := wsInstance.StorageFilterer.WatchStoreData(nil, storageStoreData, nil)
		if err != nil {
			log.Logger.Error(err)
			goto hear
		}
		defer transfer.Unsubscribe()
		select {
		case errs := <-transfer.Err():
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
```
