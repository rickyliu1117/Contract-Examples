package client

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/params"
	"strconv"
	"strings"

	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func NewETHClient(cli *ethclient.Client, pk *ecdsa.PrivateKey) (*ETHClient, error) {
	contractCli := &ETHClient{
		client: cli,
		key:    pk,
	}

	chainId, err := cli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	contractCli.chainId = chainId

	return contractCli, nil
}

type ETHClient struct {
	client       *ethclient.Client
	session      interface{}
	key          *ecdsa.PrivateKey
	chainId      *big.Int
	transactOpts *bind.TransactOpts

	GasPrice string
	GasLimit string
}

type DeployContract func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error)

func (cli *ETHClient) GetTrandOpts() *bind.TransactOpts {
	if cli.transactOpts == nil {
		cli.transactOpts = NewKeyedTransactor(cli.key, cli.chainId, cli.GasPrice, cli.GasLimit)
		gasPrice, err := cli.ChainGasPrice()
		if err == nil {
			cli.transactOpts.GasPrice = gasPrice
		}
	}
	return cli.transactOpts

}

func (cli *ETHClient) Deploy(deploy DeployContract) (addr common.Address, txId string, err error) {

	addr, tx, err := deploy(cli.GetTrandOpts(), cli.client)
	if err != nil {

		fmt.Println("deploy contract has err :", err.Error())
		return
	}

	fmt.Println(fmt.Sprintf("contract deployed success, contract address: %s, tx hash: %s",
		strings.ToLower(addr.String()), tx.Hash().String()))

	return addr, tx.Hash().String(), err
}

type NewSession func(auth *bind.TransactOpts, opts bind.CallOpts, client *ethclient.Client) (interface{}, error)

func (cli *ETHClient) Connection(session NewSession) error {
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	var err error
	cli.session, err = session(transactOpts, callOpts, cli.client)

	return err

}

func (cli *ETHClient) Client() *ethclient.Client {
	return cli.client
}

func (cli *ETHClient) Session() interface{} {
	return cli.session
}

func (cli *ETHClient) ChainGasPrice() (*big.Int, error) {
	if cli.GasPrice != "" {
		gp, ok := new(big.Int).SetString(cli.GasPrice, 10)
		if ok {
			return gp, nil
		}
	}

	return cli.client.SuggestGasPrice(context.Background())

}

//1900000000
var gasPrice = NewGwei(1)

func NewGwei(n int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(n), big.NewInt(params.GWei))
}

func NewEther(n int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(n), big.NewInt(params.Ether))
}

//var gasPrice = new(big.Int).SetInt64(1900000000)

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
func NewKeyedTransactor(key *ecdsa.PrivateKey, chainId *big.Int, gasPrice string, gaslimit string) *bind.TransactOpts {
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	to := &bind.TransactOpts{
		From:     keyAddr,
		GasLimit: 4000000,
		//GasPrice: NewGwei(1),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.NewLondonSigner(chainId)
			if address != keyAddr {
				return nil, errors.New("not authorized to sign this account")
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

	//if gasPrice != "" {
	//	gp, ok := new(big.Int).SetString(gasPrice, 10)
	//	if ok {
	//		to.GasPrice = gp
	//	}
	//}

	if gaslimit != "" {
		gl, err := strconv.ParseUint(gaslimit, 10, 64)
		if err != nil {
			to.GasLimit = gl
		}

	}

	return to
}

func (cli *ETHClient) OwnerAddress() common.Address {
	return crypto.PubkeyToAddress(cli.key.PublicKey)
}

func (cli *ETHClient) Nonce() uint64 {
	nonce, _ := cli.client.NonceAt(context.Background(), cli.OwnerAddress(), nil)
	return nonce
}

// Transfer 向 `to`转账 `amount` GAS
func (cli *ETHClient) Transfer(to common.Address, amount *big.Int) (common.Hash, error) {
	tranOpt := cli.GetTrandOpts()

	data := &types.LegacyTx{
		Nonce:    cli.Nonce(),
		To:       &to,
		Value:    amount,
		Gas:      21000,
		GasPrice: gasPrice,
		Data:     nil,
	}

	tx := types.NewTx(data)

	signTx, err := tranOpt.Signer(cli.OwnerAddress(), tx)
	if err != nil {
		return common.Hash{}, err
	}
	err = cli.Client().SendTransaction(context.Background(), signTx)
	if err != nil {
		return common.Hash{}, err
	}
	return signTx.Hash(), nil
}
