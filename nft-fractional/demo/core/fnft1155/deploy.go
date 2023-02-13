package fnft1155

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"nft-fractional-sdk-go/contracts/fnft1155"
	"nft-fractional-sdk-go/core/chain"
)

func Deploy(chainInfo *chain.ChainInfo) (common.Address, string, error) {
	cli, err := chain.NewETHCli(defChain)
	if err != nil {
		return [20]byte{}, "", err
	}

	address, txId, err := cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {

		address, tx, _, err := fnft1155.DeployTokenFractional1155(auth, backend)

		return address, tx, err
	})

	if err != nil {
		return [20]byte{}, "", err
	}

	return address, txId, nil
}
