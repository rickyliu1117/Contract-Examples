# nft-fractional-sdk-go


## Fractional NFT to ERC-20

### Deploy

Deploy the contract

```
	cli, err := chain.NewETHCli(defChain)
	if err != nil {
		t.Fatal()
	}

	address, txId, err := cli.Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {

		address, tx, _, err := fnft20.DeployTokenFractional20(auth, backend, "Token Fractional", "FNFT")

		return address, tx, err
	})

	if err != nil {
		t.Fatal()
	}

	fmt.Println("address:", address.String())
	fmt.Println("txId", txId)

```

### Fractionalize the NFT

The NFT needs to be authorized to the F-NFT contract before the fractionalization.

```
	cli, err := NewMyFNFTClient(defChain)
	if err != nil {
		t.Fatal(err)
	}
	tokenId := new(big.Int).SetInt64(1)
	amount := new(big.Int).SetInt64(10)
	tx, err := cli.session.Fractional(common.HexToAddress("0x89829B33e9B1EBA4EF78f3faCe52E2e0Caed65cA"), tokenId, amount)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
```

### Compose the NFT

```
	cli, err := NewMyFNFTClient(toChain)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := cli.session.Compose(toChain.KeyAddress)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("txId", tx.Hash().String())
```
