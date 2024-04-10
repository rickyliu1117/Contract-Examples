# Token交换

[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/)

不同的ERC20代币合约彼此独立，需要有中间的合约才能实现价值兑换，这样的中间合约会提供一个兑换池子，并给两种代币定一个兑换比例，帮助想将一个代币转换为另一个代币的用户进行兑换，exchange合约提供了一个简单的ERC20代币之间的兑换方式，部署后设置一个固定的转换比例，给合约转入对应的两种20币，以作为一个交易池，用户可借助此交易池进行20币的交换。


## 前提条件

- 掌握最基本的[ERC20标准](https://eips.ethereum.org/EIPS/eip-20)。
- 合约owner、交易的sender的概念。
- 在使用智能合约之前，必须对以太坊和Solidity有基本的了解。

## 合约概述

本应用提供了两个合约：a-token合约与exchange合约。其中a-token合约是标准ERC20合约，作为前置合约使用，表示生成某个ERC20代币的合约。exchange合约是实现交换token功能的兑换池合约。exchange合约提供交易方法用来交互部署的两套a-token合约的ERC20代币。
在此场景中，包含两个角色：exchange合约管理员、token交换者：

- exchange合约管理员：部署exchange合约，调用setExchangeRate方法设置兑换比率，往兑换池中转入两种token，即token0和token1，调用updateExchangePool方法更新合约内交易池初始值。

- token交换者：将自己持有的token0或者token1转给exchange合约，再调用exchange合约中swap方法换取另一种token。

## 部署说明

通过以下命令拉取 Beginner-Smart-Contracts 仓库下的应用合约代码，进入 token-exchange 目录：

```
$ git clone git@github.com:BSN-DDC/Beginner-Smart-Contracts.git
```

Token交换合约的部署过程，请按以下步骤操作：

1. 任一用户部署两个ERC20合约。此例中部署a-token合约生成token0，将a-token合约更名为b-token合约并部署，生成token1。这两种token用于兑换。
2. exchange合约管理员分别调用这两个合约中的mint()方法生成并持有不同数量的token0和token1。
3. exchange合约管理员部署exchange合约。在部署时，需要填入a-token合约和b-token合约的地址。
4. exchange合约管理员调用ERC20中的transfer()方法将持有的token0和token1转给token-exchange合约的地址。
5. exchange合约管理员调用exchange合约中的setExchangeRate()方法，输入几个token0兑换几个token1的比率（此信息可通过queryExchangeRate()方法查看）;再调用updateExchangePool()方法,将合约当前拥有的ERC20币数值，更新到合约内变量里。

## 主要功能

### token交换方将token0转给exchange合约地址

token交换方持有一定数量的token0，希望兑换为token1。token交换方将token0转给exchange合约地址。

AToken.transfer(address to, uint256 amount)

其中AToken是token0合约实例，to为exchange合约地址，amount是要发起交互的token0数量。

### 将token0换成token1

token交换方调用exchange合约中的swap()方法，将自己刚转入的token0兑换为token1并给指定地址。参数to为转出的指定地址。

Exchange.swap(address to)

其中Exchange是token-exchange合约实例，参数to为转出的指定地址。
