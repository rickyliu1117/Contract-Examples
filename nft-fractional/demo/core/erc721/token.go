package erc721

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"nft-fractional-sdk-go/client"
	"nft-fractional-sdk-go/contracts/erc721"
	"nft-fractional-sdk-go/core/chain"
)

func NewChainInfo() *chain.ChainInfo {
	return &chain.ChainInfo{
		Name:            "Test Chain",
		Url:             "http://10.0.51.174:20001",
		KeyHex:          "0x4dff7ea507a09bcce8619f78b8e4a165a28ea467830a5c128c4c32544cfbda02", //"0xb03108353f9beb3f07a81af79a496874b4e2b01edc9e12ed92a1a4d3cdd39e15",
		KeyAddress:      common.HexToAddress("0x30CC0ad51182B108Af7a54ffAdb39EF1f07abC16"),
		ContractAddress: common.HexToAddress("0x2A89Dc26a158081073025640616496F5A3F031b1"),
	}
}

var defChain = NewChainInfo()

type MyTokenClient struct {
	*chain.ChainClient

	session *erc721.MyERC721Session
	abi     *abi.ABI
}

func MyTokenClientSesson(cli *client.ETHClient, addr common.Address) (*erc721.MyERC721Session, error) {
	session, err := erc721.NewMyERC721(addr, cli.Client())
	if err != nil {
		return nil, err
	}
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	return &erc721.MyERC721Session{Contract: session, CallOpts: callOpts, TransactOpts: *transactOpts}, err
}

func NewMyTokenClient(chainInfo *chain.ChainInfo) (*MyTokenClient, error) {

	cli, err := chain.NewChainClient(chainInfo)
	if err != nil {
		return nil, err
	}

	session, err := MyTokenClientSesson(cli.EthCli(), chainInfo.ContractAddress)
	if err != nil {
		return nil, err
	}
	parsed, err := erc721.MyERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	s := &MyTokenClient{
		ChainClient: cli,
		session:     session,
		abi:         parsed,
	}

	return s, nil
}

const TransferEventName = "Transfer"

func (s *MyTokenClient) GetTokenId(txId common.Hash) (*big.Int, error) {
	err := s.WaitTx(txId.String())
	if err != nil {
		return nil, err
	}
	r, err := s.EthCli().Client().TransactionReceipt(context.Background(), txId)
	if err != nil {
		return nil, err
	}
	event := s.abi.Events[TransferEventName]
	for _, log := range r.Logs {
		if log.Topics[0] == event.ID {
			transfer, err := s.session.Contract.ParseTransfer(*log)
			if err != nil {
				return nil, err
			} else {
				return transfer.TokenId, nil
			}
		}
	}
	return nil, err
}
