# 元交易
[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) [![License](https://badgen.net/badge/license/MIT/blue)](https://typescriptlang.org)

元交易（Meta Transaction），是一种用户使用DApp、发起交易和调用智能合同而无需支付燃气费的一种方式。

聊元交易之前，首先了解一下什么是交易（Transaction）。一笔以太坊交易由以下内容构成：

- from – 发送者地址
- recipient – 接收者地址（如果为一个外部持有的帐户，交易将传输值。 如果为合约帐户，交易将执行合约代码）
- signature – 发送者的签名。 当通过发送者的私钥签名交易来确保发送者已授权此交易时，生成此签名。
- value – 从发送者向接收者转移 ETH 的金额 （以 WEI 为单位，ETH 的一种面值单位）
- data – 可包括任意数据的可选字段
- gasLimit – 交易可以消耗的 Gas 的最大数量。 Gas 单位代表了计算步骤
- gasPrice – 发送者按单位 gas 支付的费用
- nonce – 发送者已经发送过的交易数

注意其中的 signature 字段，通过它任何人都能够验证这笔交易就是发送者地址签署的。交易会被发送给区块链节点，发送者会支付 gas 费，通过验证的交易才会被节点包含进自己的区块链中，并进行广播。而如果说，这样一笔交易发给某个中间人/节点，让他帮忙来付 gas 费并执行该交易，我们的目的就实现了。

但问题是，简单的将这样一笔交易发给中间人，中间人也并不能帮你支付gas费，因为它是一个普通的交易，它会被验证通过，并认为是发送者来支付gas费。

那我们如何绕过这个限制呢？答案是智能合约。

如果这笔交易发生在智能合约内部，也就是说，在普通的交易内部嵌入一个交易（这个交易就被称作元交易），交易被你的中间人/节点签署，并指定接收者地址为元交易智能合约的地址，因此 gas 费由中间人/节点支付；而元交易智能合约在收到一笔元交易后，会验证元交易的签名信息，确认无误后，你的元交易在元交易智能合约中被执行。

## 前提条件

在使用智能合约之前，必须对以太坊和Solidity有基本的了解。
有关智能合约入门的全面指南，请参阅我们的[初学者指南](https://github.com/BSN-DDC/docs/blob/main/BSN-DDC%E7%BD%91%E7%BB%9C%E9%83%A8%E7%BD%B2Solidity%E5%90%88%E7%BA%A6%E5%BF%AB%E9%80%9F%E4%B8%8A%E6%89%8B%E6%8C%87%E5%8D%97.pdf)。


## 场景说明

Alice 想向 Bob 转账 0 ETH，而由于 Alice 账户上没有任何 ETH，即便是转账 0 ETH，但她仍然需要支付一定数额的 gas 费，因此 Alice 无法直接执行这样一笔交易。

而 Alice 知道 Carol 恰好账号上有足够多的 ETH 去支付 gas 费，于是请求他的帮助。

Carol 让 Alice 签署这笔元交易，并将所有内容发送给他；Carol 收到 Alice 的元交易后，构造出一个发送给元交易智能合约地址的交易，广播给区块链的节点。

区块链节点将验证 Carol 的交易合法性；元交易智能合约 扣除 Carol 的 gas 费作为执行智能合约的费用，并验证该交易中的元交易是否合法（验证是否为 Alice 的签名，nonce 值是否合法等）。

验证合法后，元交易智能合约执行该元交易，从而 Alice 在没有花任何 gas 费的情况下，通过中间人 Carol 执行了交易。

## 合约方法

编号  |  方法声明   |  方法描述 
--------|------------|------------
&nbsp;1 | *function getNonce(address from)* | 获取发送者进行元交易的交易数，from是发送者地址           
&nbsp;2 | *function metaTransferFrom(address from,address to,uint256 amount,uint256 nonce,uint256 deadline,uint8 v,bytes32 r,bytes32 s)* | 发送元交易，from是发送者地址，to是接收者地址，amount是交易代币的数量，nonce是发送者已经发送过的元交易数+1，deadline是截止时间（大于当前时间的时间戳），r,s,v是签名数据


## 使用示例

假如现在要进行一次代币的转账交易，用户A需要向用户B转账一定数量的代币，用户A没有足够的gas，想通过元交易让用户C代为进行，整个过程如下：

```
1. 用户C调用合约中的mint方法，为用户A生成一定数量的代币
2. 用户C（也可以用户A获取，查询不消耗gas）通过getNonce方法获取用户A已经发送过的元交易数，封装交易时使用的nonce值需要在获取的nonce值基础上+1
2. 用户C（也可以用户A获取，查询不消耗gas）通过DOMAIN_SEPARATOR方法获取加密值 
3. 用户A通过链下签名程序和用户A的私钥将交易数据进行签名，将r,s,v签名数据和交易数据发送给用户C 
4. 用户C收到数据后调用合约的metaTransferFrom方法将交易发出，注意调用metaTransferFrom传入的from、to、amount、nonce、deadline要和用户A链下签名时的数据一致
5. 合约执行成功后元交易完成  
```

## 链下生成签名数据

提供链下签名程序包ERC20MetaTxSign生成`r,s,v`签名数据。

当前使用nodejs版本

```
node v14.17.6
npm  v6.14.15
```

进入到项目目录中`cd ERC20MetaTxSign `

在命令行执行`npm install`拉取需要使用的node包

打开`ERC20MetaTxSign.js`文件，修改文件中的以下参数

```javascript
var from = "";				//发送者地址（用户A的钱包地址）
var to = "";				//接收者地址（用户B的钱包地址）
var amount = 0;				//转移代币数量
var nonce = 0;				//nonce值 （用户A的元交易nonce值 + 1，使用合约方法getNonce获取账户A的元交易nonce值）
var deadline = 0;			//截止时间，不能为0，必须为大于当前时间的时间戳
var DOMAIN_SEPARATOR = "";	//使用合约方法DOMAIN_SEPARATOR生成加密值
var privateKey = ""; 		//发送者私钥（用户A的私钥）
```

完成参数修改后，在命令行执行 `node ERC20MetaTxSign.js` 

会得到签名数据的JSON，格式如下：

数据中的`r,s,v`就是元交易需要使用的签名参数

```
{
  r: '0x00',
  s: '0x00',
  _vs: '',
  recoveryParam: 0,
  v: 27,
  yParityAndS: '0x00',
  compact: '0x00'
}
```
