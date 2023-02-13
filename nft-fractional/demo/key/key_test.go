package key

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
	"nft-fractional-sdk-go/common"
	"testing"
)

func TestName(t *testing.T) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		t.Fatal(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)

	fmt.Println("助记词:", mnemonic)

	seed := bip39.NewSeed(mnemonic, "")

	keyBytes := crypto.Keccak256(seed)

	key, err := crypto.ToECDSA(keyBytes)

	if err != nil {
		t.Fatal(err)
	}

	pubK := &key.PublicKey
	prk := crypto.FromECDSA(key)

	fmt.Println("私钥：", hexutil.Encode(prk))
	fmt.Println("公钥：", hexutil.Encode(crypto.FromECDSAPub(pubK)))
	fmt.Println("地址：", crypto.PubkeyToAddress(*pubK).String())

	keyJson, err := EncryptKey(prk, "abc123")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(keyJson))

	kk, err := DecryptKey(keyJson, "abc123")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(hexutil.Encode(kk))
}

func TestGenKey(t *testing.T) {

	key, _ := crypto.GenerateKey()

	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)

	keyB := crypto.FromECDSA(key)

	keyHex := hex.EncodeToString(keyB)

	//fmt.Println(keyHex)
	common.WriteFile([]byte(keyHex), fmt.Sprintf("./account/%s.key", address.String()), true)

}

func TestString(t *testing.T) {
	keyJson, err := EncryptKey([]byte("shiyuetang"), "abc123")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(keyJson))
}
