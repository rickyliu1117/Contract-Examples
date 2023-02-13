package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewETHCliByURL(url, key string) (*ETHClient, error) {
	//fmt.Println("URL:", url)
	hex, err := hexutil.Decode(key)
	if err != nil {
		return nil, err
	}
	pk, err := crypto.ToECDSA(hex)
	if err != nil {
		return nil, err
	}
	fmt.Println(crypto.PubkeyToAddress(pk.PublicKey).String())
	ethCli, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	cli, err := NewETHClient(ethCli, pk)
	if err != nil {
		return nil, err
	}

	return cli, nil

}
