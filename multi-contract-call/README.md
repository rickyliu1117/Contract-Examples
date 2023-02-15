# 多合约调用结构

[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) [![License](https://badgen.net/badge/license/MIT/blue)](https://typescriptlang.org)

​		一个应用如果只使用了一个合约来实现业务逻辑肯定是不合理的，因为数据承载会超出区块的限制，所以应该在初始设计阶段就应该考虑多合约的结构。这里介绍了两种多合约调用方式以供参考。

- 合约引用
- 合约call和delegatecall

## 前提条件

在使用智能合约之前，必须对以太坊和Solidity有基本的了解。
有关智能合约入门的全面指南，请参阅我们的[初学者指南](https://github.com/BSN-DDC/docs/blob/main/BSN-DDC%E7%BD%91%E7%BB%9C%E9%83%A8%E7%BD%B2Solidity%E5%90%88%E7%BA%A6%E5%BF%AB%E9%80%9F%E4%B8%8A%E6%89%8B%E6%8C%87%E5%8D%97.pdf)。


## 使用说明 

​		用户下载示例合约，通过remix连接并进行简单调用验证后可自行对示例合约进行修改。或者通过abigen编译工具编译出的golang语言的文件进行调用。编译命令示例：

`./abigen --bin ./Test.bin --abi ./Test.abi --pkg Test --out Test.go`

## 合约概述

​		如下面的两个合约代码所示，Logic合约通过引用Data合约后实现了Data合约里的方法调用，当给定Logic合约里的`setX`方法第一个入参为Data合约的合约地址，第二个入参为数值后调用，会发现Data合约里的数值也相应的产生了变化。

​	`setX`和`setXFromAddress`两个方法实现逻辑一模一样只是写法不同。

```
// data.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Data {
    uint public x;
    /** 
     * @dev Set a value.
     * @param _x value
     */
    function setX(uint _x) public returns (uint) {
        x = _x;
        return x;
    }
}
```

```
// logic.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./data.sol";

contract Logic {
    /** 
     * @dev Set a value.
     * @param _data Contract address
     * @param _x value
     */
    function setX(Data _data, uint _x) public {
        _data.setX(_x);
    }

    /** 
     * @dev Set a value.
     * @param _addr Contract address
     * @param _x value
     */
    function setXFromAddress(address _addr, uint _x) public {
        Data data = Data(_addr);
        data.setX(_x);
    }
}

```

## Go语言调用示例

- ***setX方法调用***

```go
func TestSetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	instance, err := Logic.NewLogic(Address, cli)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	tx, err := instance.SetX(auth, dataContractAddr, x)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}
```

- ***setXFromAddress方法调用***

```go
func TestSetXFromAddress(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	instance, err := Logic.NewLogic(Address, cli)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	tx, err := instance.SetXFromAddress(auth, dataContractAddr, x)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}
```

## 合约call和delegatecall

### 使用说明 

​		用户下载示例合约，通过remix连接并进行简单调用验证后可自行对示例合约进行修改。或者通过abigen编译工具编译出的golang语言的文件进行调用。编译命令示例：

`./abigen --bin ./Test.bin --abi ./Test.abi --pkg Test --out Test.go`

### 合约概述

​		如下面的两个合约代码所示，

​        Caller合约中的`callSetX`实现了call方法并指向了某个合约的setX方法调用。当调用Caller合约中的`callSetX`方法第一个入参为Receiver合约地址，第二个入参为数值时，根据call方法的特性会转而调用Receiver合约中的setX方法并将数值加一后的结果返回，单独查询Recevier合约的x会发现其数值**有变动**。

​		Caller合约中的`delegatecallSetX`实现了delegatecall方法并指向了某个合约的setX方法调用。当调用Caller合约中的`delegatecallSetX`方法第一个入参为Receiver合约地址，第二个入参为数值时，根据delegatecall方法的特性会转而调用Receiver合约中的setX方法并将数值加一后的结果返回，单独查询Recevier合约的x会发现其数值**没有变动**但Caller合约里的x其数值已经**发生变动**。

​		**特别注意**：实现delegatecall的合约Caller存储布局必须与合约Receiver相同。



```
// receiver.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Receiver {
    uint public x;
    
    /** 
     * @dev Set a value.
     * @param _x value
     */
    function setX(uint _x) public returns (uint) {
        x = _x + 1;
        return x;
    }
}

```

```
// caller.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract Caller {
	uint public x;
	
	// event for EVM logging
    event Response(bool success, bytes data);

    /**
     * @dev Call the setX method in the contract for the given contract address.
     * @param _addr contract address
     * @param _x value
     */
    function callSetX(address  _addr, uint _x) public {
        (bool success, bytes memory data) = _addr.call(
            abi.encodeWithSignature("setX(uint256)", _x)
        );

        emit Response(success, data);
    }

    /**
     * @dev Delegatecall the setX method in the contract for the given contract address.
     * @param _addr contract address
     * @param _x value    
     */
    function delegatecallSetX(address  _addr, uint _x) public {
        (bool success, bytes memory data) = _addr.delegatecall(
            abi.encodeWithSignature("setX(uint256)", _x)
        );

        emit Response(success, data);
    }

}
```

### Go语言调用示例

- ***callSetX方法调用***

```go
func TestCallSetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	instance, err := Caller.NewCaller(Address, cli)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	tx, err := instance.CallSetX(auth, contractAddr, x)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}
```

- ***delegatecallSetX方法调用***

```go
func TestDelegatecallSetX(t *testing.T) {
	cli, err := ethclient.Dial(NodeUrl)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	instance, err := Caller.NewCaller(Address, cli)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	tx, err := instance.DelegatecallSetX(auth, contractAddr, x)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	fmt.Println("tx Hash:", tx.Hash().String())
}
```

- ***捕获解析事件示例***

```go
func TestEvent(t *testing.T) {
	var callerResponse = make(chan *Caller.CallerResponse, 100)
	wsCli, err := ethclient.Dial(NodeWsUrl)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	wsInstance, err := Caller.NewCaller(Address, wsCli)
	if err != nil {
		log.Logger.Error(err)
        return
	}
	go func() {
	hear:
		transfer, err := wsInstance.CallerFilterer.WatchResponse(nil, callerResponse)
		if err != nil {
			log.Logger.Error(err)
			goto hear
		}
		defer transfer.Unsubscribe()
		select {
		case errs := <-transfer.Err():
			log.Logger.Warn(errs)
			goto hear
		}
	}()
	for {
		select {
		case event := <-callerResponse:
			fmt.Println("success:", event.Success)
			fmt.Println("data:", event.Data)
		}
	}
}
```

