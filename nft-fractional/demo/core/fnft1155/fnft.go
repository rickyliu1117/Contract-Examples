package fnft1155

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"nft-fractional-sdk-go/client"
	"nft-fractional-sdk-go/contracts/fnft1155"
	"nft-fractional-sdk-go/core/chain"
)

func NewChainInfo() *chain.ChainInfo {
	return &chain.ChainInfo{
		Name:            "Test Chain",
		Url:             "http://10.0.51.174:20001",
		KeyHex:          "0x4dff7ea507a09bcce8619f78b8e4a165a28ea467830a5c128c4c32544cfbda02", //"0xb03108353f9beb3f07a81af79a496874b4e2b01edc9e12ed92a1a4d3cdd39e15",
		KeyAddress:      common.HexToAddress("0x30CC0ad51182B108Af7a54ffAdb39EF1f07abC16"),
		ContractAddress: common.HexToAddress("0x76EA6671F1AeEfcAD51972bb544Ebbd45203b013"),
	}
}

func NewChainInfo2() *chain.ChainInfo {
	return &chain.ChainInfo{
		Name:            "Test Chain",
		Url:             "http://10.0.51.174:20001",
		KeyHex:          "0x1e36c3a7b1594b964aab7796f0f86facd1a0109b1749169e63347d4e425a2f9f", //"0xb03108353f9beb3f07a81af79a496874b4e2b01edc9e12ed92a1a4d3cdd39e15",
		KeyAddress:      common.HexToAddress("0x19f90351A0D3D15CA835E15aE8385e8f8d30b561"),
		ContractAddress: common.HexToAddress("0x76EA6671F1AeEfcAD51972bb544Ebbd45203b013"),
	}
}

var defChain = NewChainInfo()

var toChain = NewChainInfo2()

type MyFNFTClient struct {
	*chain.ChainClient

	session *fnft1155.TokenFractional1155Session
	abi     *abi.ABI
}

func MyFNFTClientSesson(cli *client.ETHClient, addr common.Address) (*fnft1155.TokenFractional1155Session, error) {
	session, err := fnft1155.NewTokenFractional1155(addr, cli.Client())
	if err != nil {
		return nil, err
	}
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	return &fnft1155.TokenFractional1155Session{Contract: session, CallOpts: callOpts, TransactOpts: *transactOpts}, err
}

func NewMyFNFTClient(chainInfo *chain.ChainInfo) (*MyFNFTClient, error) {

	cli, err := chain.NewChainClient(chainInfo)
	if err != nil {
		return nil, err
	}

	session, err := MyFNFTClientSesson(cli.EthCli(), chainInfo.ContractAddress)
	if err != nil {
		return nil, err
	}

	s := &MyFNFTClient{
		ChainClient: cli,
		session:     session,
	}

	return s, nil
}

func (c *MyFNFTClient) Session() *fnft1155.TokenFractional1155Session {
	return c.session
}

func (c *MyFNFTClient) Fractional(erc721address common.Address, erc721tokenId *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return c.session.Fractional(erc721address, erc721tokenId, amount, data)
}

func (c *MyFNFTClient) Compose(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return c.session.Compose(to, tokenId)
}
