package fnft1155

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestDeploy(t *testing.T) {
	address, txId, err := Deploy(defChain)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("address:", address.String())
	fmt.Println("txId", txId)

	//contract deployed success, contract address: 0x76ea6671f1aeefcad51972bb544ebbd45203b013, tx hash: 0x7259f9dc69279516817ef3d0bca479651b6819c2c86b5ebf82c37519e4b8b925
	//address: 0x76EA6671F1AeEfcAD51972bb544Ebbd45203b013
	//txId 0x7259f9dc69279516817ef3d0bca479651b6819c2c86b5ebf82c37519e4b8b925
}

func TestFractional(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(1)
	amount := new(big.Int).SetInt64(10)
	tx, err := cli.Fractional(common.HexToAddress("0x2A89Dc26a158081073025640616496F5A3F031b1"), tokenId, amount, []byte("test"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
	cli.WaitTx(tx.Hash().String())

}

func TestBalanceOf(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(0)
	balance, err := cli.Session().BalanceOf(toChain.KeyAddress, tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance:", balance.String())
}

func TestTokenOf(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(0)

	erc721Addr, erc721TokenId, err := cli.Session().TokenOf(tokenId)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("address", erc721Addr.String())
	fmt.Println("tokenId", erc721TokenId.String())
}

func TestName(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(0)

	amount, err := cli.Session().TotalSupply(tokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("amount", amount.String())
}

func TestSafeTransferFrom(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(0)
	amount := new(big.Int).SetInt64(10)
	tx, err := cli.Session().SafeTransferFrom(defChain.KeyAddress, toChain.KeyAddress, tokenId, amount, []byte("test"))

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
	cli.WaitTx(tx.Hash().String())

}

func TestCompose(t *testing.T) {
	cli, err := NewMyFNFTClient(toChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(0)
	tx, err := cli.Compose(toChain.KeyAddress, tokenId)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
	cli.WaitTx(tx.Hash().String())

}
