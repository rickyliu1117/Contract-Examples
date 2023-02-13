package chain

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"nft-fractional-sdk-go/client"
	"time"
)

type ChainInfo struct {
	Name            string
	Url             string
	KeyHex          string
	KeyAddress      common.Address
	ContractAddress common.Address
}

//94e3e21f6e7221becb289fffc7b1d0aa66a7dbd77f035be6a1e96737990852e8
func NewChainInfo() *ChainInfo {
	return &ChainInfo{
		Name:            "Test Chain",
		Url:             "http://10.0.51.174:20001",
		KeyHex:          "0x94e3e21f6e7221becb289fffc7b1d0aa66a7dbd77f035be6a1e96737990852e8", //"0xb03108353f9beb3f07a81af79a496874b4e2b01edc9e12ed92a1a4d3cdd39e15",
		KeyAddress:      common.HexToAddress("0x30CC0ad51182B108Af7a54ffAdb39EF1f07abC16"),
		ContractAddress: common.HexToAddress("0xc0b3382E3D88ef66941278A9A58A49b442a5f8d0"),
	}
}

func NewETHCli(chain *ChainInfo) (*client.ETHClient, error) {

	return client.NewETHCliByURL(chain.Url, chain.KeyHex)
}

const defTimeout = 20 * time.Second

type ChainClient struct {
	cli *client.ETHClient

	timeout time.Duration
}

func (s *ChainClient) EthCli() *client.ETHClient {
	return s.cli
}

func NewChainClient(chain *ChainInfo) (*ChainClient, error) {

	cli, err := NewETHCli(chain)
	if err != nil {
		return nil, err
	}

	s := &ChainClient{
		cli:     cli,
		timeout: defTimeout,
	}

	return s, nil
}

func (s *ChainClient) WaitTx(txId string) error {
	tick := time.NewTicker(s.timeout)
	for {
		select {
		case <-tick.C:
			return errors.New("tx pending , wait time out")
		default:

			time.Sleep(time.Second)
			_, p, err := s.EthCli().Client().TransactionByHash(context.Background(), common.HexToHash(txId))
			if err != nil {
				return err
			}
			if !p {
				return nil
			}

		}
	}
}
