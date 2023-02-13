package erc721

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"nft-fractional-sdk-go/contracts/erc721"
	"nft-fractional-sdk-go/core/chain"
	"testing"
)

func TestDeploy(t *testing.T) {
	cli, err := chain.NewETHCli(defChain)
	if err != nil {
		t.Fatal()
	}

	address, txId, err := cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {

		address, tx, _, err := erc721.DeployMyERC721(auth, backend)

		return address, tx, err
	})

	if err != nil {
		t.Fatal()
	}

	fmt.Println("address:", address.String())
	fmt.Println("txId", txId)

	//contract deployed success, contract address: 0x2a89dc26a158081073025640616496f5a3f031b1, tx hash: 0x1dfe9bae913e25f2ca9a1aa688d2f789ba5d8d31bad1cc9b3f63a4a7bbd8d30f
	//address: 0x2A89Dc26a158081073025640616496F5A3F031b1
	//txId 0x1dfe9bae913e25f2ca9a1aa688d2f789ba5d8d31bad1cc9b3f63a4a7bbd8d30f

}

func TestOwner(t *testing.T) {
	cli, err := NewMyTokenClient(defChain)
	if err != nil {
		t.Fatal(err)
	}

	owner, err := cli.session.Owner()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("owner:", owner.String())
}

func TestMint(t *testing.T) {
	cli, err := NewMyTokenClient(defChain)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := cli.session.SafeMint(common.HexToAddress("0x30CC0ad51182B108Af7a54ffAdb39EF1f07abC16"), "this is uri 3")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId:", tx.Hash().String())
	tokenId, err := cli.GetTokenId(tx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("TokenId", tokenId)
}

func TestOwnerOf(t *testing.T) {
	cli, err := NewMyTokenClient(defChain)
	if err != nil {
		t.Fatal(err)
	}

	tokenId := new(big.Int).SetInt64(1)
	owner, err := cli.session.OwnerOf(tokenId)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("owner:", owner.String())
}

func TestApprove(t *testing.T) {
	cli, err := NewMyTokenClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(1)
	tx, err := cli.session.Approve(common.HexToAddress("0x76EA6671F1AeEfcAD51972bb544Ebbd45203b013"), tokenId)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId:", tx.Hash().String())

}

func TestGetApproved(t *testing.T) {
	cli, err := NewMyTokenClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(1)
	approve, err := cli.session.GetApproved(tokenId)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("approve :", approve.String())
}
