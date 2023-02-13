package chain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"nft-fractional-sdk-go/client"
	"testing"
)

func TestTranfer(t *testing.T) {

	cli, err := NewETHCli(NewChainInfo())
	if err != nil {
		t.Fatal(err)
	}
	tx, err := cli.Transfer(common.HexToAddress("0x8e4aaBE269F30921eDEACA3239805e86E79Ce020"), client.NewEther(10))

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("txId", tx.String())
}

func TestBalance(t *testing.T) {
	cli, err := NewETHCli(NewChainInfo())
	if err != nil {
		t.Fatal(err)
	}

	b, err := cli.Client().BalanceAt(context.Background(), common.HexToAddress("0x8e4aaBE269F30921eDEACA3239805e86E79Ce020"), nil)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(b.String())

}
