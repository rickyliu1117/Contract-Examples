# 分布式数字身份DID
[![Smart Contract](https://badgen.net/badge/smart-contract/Solidity/orange)](https://soliditylang.org/) 

数据身份被认为是真实身份信息浓缩为数字代码，形成可通过网络、相关设备等查询和识别的公共秘钥。随着互联网的出现和普及，一般认为数字身份演进经历了三个阶段，分别是:中心化身份、联盟身份、分布式身份。理想情况下，在分布式身份阶段，用户可以完全掌控自己的信息。
在万维网联盟（World Wide Web Consortium, W3C）发布的分布式数字身份规范中，将DID（Decentralized Identifiers）定义为一种新的全球唯一标识符，能够实现可验证的、去中心化的数字身份。这种标识符不仅可以用于人，也可以用于万事万物，包括一辆车、一只动物，甚至是一台机器。[点击查看更多DID信息](https://www.w3.org/TR/2022/REC-did-core-20220719/#abstract)。

## 使用说明 
* DID达不到实名认证的效果，所以如有此需求，可在业务系统内实现实名认证，然后将系统用户与其DID进行绑定即可。
* 即使同一个Sender，调用多次“创建DID”也会为其生成多个DID，所以如需保证每个用户只有一个DID，可在业务系统内判断是否已绑定DID即可。

## 前提条件

在使用智能合约之前，必须对以太坊和Solidity有基本的了解。
有关智能合约入门的全面指南，请参阅我们的[初学者指南](https://github.com/BSN-DDC/docs/blob/main/BSN-DDC%E7%BD%91%E7%BB%9C%E9%83%A8%E7%BD%B2Solidity%E5%90%88%E7%BA%A6%E5%BF%AB%E9%80%9F%E4%B8%8A%E6%89%8B%E6%8C%87%E5%8D%97.pdf)。

## 合约概述
本合约需用户先行部署，然后可直接调用。合约内提供三个DID基础功能的方法：

* “创建DID”的方法是公共方法，会为sender生成一个DID标识符和描述文件，标识符和描述文件是一对一的关系，在链上以Key-Value的形式进行存储。
* “查询描述文件”的方法是公共方法，任何人都可通过DID标识符查询链上对应的描述文件。获得描述文件便可知晓该DID的创建日期、公钥信息等，从而进一步去做签名验证以保证传输数据的完整性和真实性。
* “更新描述文件”的方法是公共方法，如果DID对应的主私钥丢失或泄漏，可以本地生成一对新的公私钥，修改描述文件内的主公钥和签名值，将新的描述文件发送到链上进行更新。合约内部会对Sender进行验证，如果与“创建DID”时的Sender一致才会进行逻辑处理，否则便会提醒更新失败。


## 合约方法

编号  |  方法声明   |  方法描述 
----|----|----
1 | *createDid(string memory did, string memory didDocument) public returns (string memory msgcode, string memory msg)* |      创建DID              
2 | *getDocument(string memory did) public view returns (string memory msgcode, string memory msg, string memory didDocument)* |      查询描述文件
3 | *updateDocument(string memory did, string memory didDocument) public returns (string memory msgcode, string memory msg)* |      更新描述文件


## Go语言调用示例

### 1）创建DID

```
func TestCreateDid(t *testing.T) {
	didDocument, did, err := utils.CreateDid()
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("did document: %s\n", didDocument)
	log.Infof("did : %s\n", did)

	cli, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := DidContract.NewDidContract(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.CreateDid(auth, did, didDocument)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("tx Hash: %s\n", tx.Hash().String())
}

```

### 2）查询描述文件

```
func TestGetDocument(t *testing.T) {
	did := "did:bsn:2VwpU9b8cktekvGhPeD89gSt3EyR"
	cli, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := DidContract.NewDidContract(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	document, err := instance.GetDocument(&bind.CallOpts{}, did)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("did document: %s\n", document.DidDocument)
}
```

### 3）更新描述文件

```
func TestUpdateDidAuth(t *testing.T) {
	did := "did:bsn:2VwpU9b8cktekvGhPeD89gSt3EyR"
	privateKey := "09388158cf84cc1b4ada96cac2e912d3651cc428d3b179c6758ffc58e3eadc00"

	cli, err := ethclient.Dial(Rpc)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	instance, err := DidContract.NewDidContract(Address, cli)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	result, err := instance.GetDocument(nil, did)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	var didDocument utils.Document

	err = json.Unmarshal([]byte(result.DidDocument), &didDocument)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	err = didDocument.ResetDidAuthTest(privateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	didDocumentBytes, err := json.Marshal(didDocument)
	if err != nil {
		log.Errorf(err.Error())
		return
	}

	auth, err := eth.GenAuth(cli, PrivateKey)
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	tx, err := instance.UpdateDocument(auth, did, string(didDocumentBytes))
	if err != nil {
		log.Errorf(err.Error())
		return
	}
	log.Infof("tx Hash: %s\n", tx.Hash().String())
}

```