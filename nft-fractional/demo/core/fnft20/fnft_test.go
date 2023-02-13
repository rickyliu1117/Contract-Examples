package fnft20

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

	//contract deployed success, contract address: 0x4387a667ec5fdab9a7cbf972a9973e88a3044338, tx hash: 0xad2426bd77aab4551f58ca9dd8c1f602c323acfea6cb81e62c75e2b16f00099e
	//address: 0x4387A667ec5fDab9A7cBF972a9973E88A3044338
	//txId 0xad2426bd77aab4551f58ca9dd8c1f602c323acfea6cb81e62c75e2b16f00099e

}

func TestFractional(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(1)
	amount := new(big.Int).SetInt64(10)
	tx, err := cli.Fractional(common.HexToAddress("0x89829B33e9B1EBA4EF78f3faCe52E2e0Caed65cA"), tokenId, amount)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())

}

func TestBalanceOf(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}

	amount, err := cli.Session().BalanceOf(defChain.KeyAddress)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("BalanceOf :", amount.String())
}

func TestBalanceOfTo(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}

	amount, err := cli.Session().BalanceOf(toChain.KeyAddress)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("BalanceOf :", amount.String())
}

func TestTotalSupply(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}

	amount, err := cli.Session().TotalSupply()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TotalSupply :", amount.String())
}

func TestTransfer(t *testing.T) {
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	amount := new(big.Int).SetInt64(10)
	tx, err := cli.Session().Transfer(toChain.KeyAddress, amount)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
}

func TestCompose(t *testing.T) {
	cli, err := NewMyFNFTClient(toChain)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := cli.Compose(toChain.KeyAddress)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
}
