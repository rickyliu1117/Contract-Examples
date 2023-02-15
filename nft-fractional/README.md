# NFT 碎片化
[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) [![License](https://badgen.net/badge/license/MIT/blue)](https://typescriptlang.org)

## 前提条件

在使用智能合约之前，必须对以太坊和Solidity有基本的了解。
有关智能合约入门的全面指南，请参阅我们的[初学者指南](https://github.com/BSN-DDC/docs/blob/main/BSN-DDC%E7%BD%91%E7%BB%9C%E9%83%A8%E7%BD%B2Solidity%E5%90%88%E7%BA%A6%E5%BF%AB%E9%80%9F%E4%B8%8A%E6%89%8B%E6%8C%87%E5%8D%97.pdf)。

## 什么是NFT碎片化

通过智能合约，基于ERC-721的NFT作品被分割成多个ERC-20代币（ERC-1155 Token），其所有权也被一同拆分，易于流转与交易。
简单来说，NFT碎片化的过程相当于「资产再分配」，但这一过程改变资产标准（由ERC-721转变为ERC-20/ERC-1155）。

## 为什么要将NFT碎片化

流动性不足已经成为NFT市场向前发展过程中亟需解决的问题。

目前，热门NFT收藏品的价格高企，想要快速找到合适的买家并不容易，而希望拥有这些NFT的普通投资者则因价格过高难以入场。最终，卖家无法脱手，潜在买家无法入手，双方买卖需求不能满足。

通过智能合约，NFT所有权将被分割， 散户投资者可以共同拥有一件NFT作品，其进入市场的门槛降低，而这将为NFT二级市场注入更多流动性。

同时，艺术家以及NFT创作者可以在不完全出售作品的情况下，轻松将作品的部分所有权代币化，获得现金流。

## 一些NFT碎片化平台

### NIFTEX

宣布支持基于ERC-721、ERC-777 和 ERC-1155标准的NFT作品，同时决定搭建自己的交易平台。目前，Niftex平台有近100种NTF碎片化资产，其中包括Cryptopunks、Axie Infinity等项目。

当使用NIFTEX拆分NFT时，用户可以选择Layer 2网络,Matic或以太坊，并从OpenSea 复制NFT资产的url 地址，继而启动碎片化功能。

### Unic.ly

在提供NFT碎片化功能之上，Unicly引入AMM（自动做市商）及流动性挖矿。这意味着，用户不仅可以将NFT拆分成ERC-20代币，还可以在平台进行交易，并通过提供流动性获得协议治理代币UNIC。

在Unicly平台，用户可以将任意数量的NFT（支持基于ERC-721、ERC-1155协议的资产）锁入智能合约，继而创建自己的uToken。代币的名称、总供应量、概述以及投票解锁NFT所需设置的百分比均由用户自己决定。

### Fractional
在Fractional平台，用户可以创建NFT Vault（支持多个NFT），由金库来保管自己的NFT，随后则可发行与之对应的ERC-20代币。如果用户想要为碎片化代币创建流动性池，需要去SushiSwap和UniSwap等第三方平台操作，Fractional自身仅支持交易。

> 8月18日，NFT 碎片化协议 Fractional 宣布更名为 Tessera

## NFT碎片化合约设计

可以通过ERC-20、ERC-1155 对NFT进行碎片化，包含NFT碎片化方法、ERC-20/ERC-1155标准、NFT组合、NFT查询等方法。

> 对于合约的权限，以及合约在生成，授权过程中的NFT Owner 要描述清楚  
> 合约没有管理权限，不能主动销毁


### 分割为`ERC-20`的碎片化

将单个ERC-721 NFT 授权给该合约，生成ERC-20 代币，继承ERC-20 标准，一个合约只能支持一个NFT的碎片化，NFT重新组合后，该合约重置，可以接受新的NFT碎片化操作。


#### NFT碎片化
1. NFT的所有者将NFT授权给F-NFT合约
2. NFT的所有者调用初始化方法设置符号，碎片（代币）总数等，对已经授权的NFT进行碎片化
3. 合约将授权的NFT转移给自己，按照输入参数铸造代币给NFT所有者

#### NFT碎片流转
继承ERC-20的标准方法

#### NFT组合
拥有所有代币的账户可以将NFT转移走，并将合约重置。

#### NFT查询
查询该合约关联的NFT的合约地址、TokenId等信息。



### 分割为`ERC-1155`的碎片化
将ERC-721 NFT 授权给该合约，生成ERC-1155 Token，继承ERC-1155 标准，支持将多个NFT在一个合约中碎片化管理。

#### NFT碎片化
1. NFT的所有者将NFT授权给F-NFT合约
2. NFT的所有者调用Mint，生成ERC-1155 Token，对已经授权的NFT进行碎片化
3. 合约将授权的NFT转移给自己，按照输入参数生成对应的1155Token给NFT所有者

#### NFT碎片流转
继承ERC-1155的标准方法

#### NFT组合
拥有某个ERC-1155 Token的所有份数的账户可以将该NFT转移走，并销毁对应Token。

#### NFT查询
查询对应Token关联的NFT的合约地址、TokenId等信息。








