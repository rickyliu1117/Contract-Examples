package key

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pborman/uuid"
	"math/big"
)

//根据助记词导入秘钥，然后输入一个密码加密，输出到文件
// AES 加密

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	version          = 3
)

func EncryptKey(data []byte, auth string) (keyJson []byte, err error) {
	d := new(big.Int).SetBytes(data)
	keyBytes := math.PaddedBigBytes(d, 32)
	cryptoStruct, err := keystore.EncryptDataV3(keyBytes, []byte(auth), veryLightScryptN, veryLightScryptP)
	if err != nil {
		return
	}

	encryptedKeyJSONV3 := encryptedKeyJSONV3{
		cryptoStruct,
		uuid.New(),
		version,
	}

	keyJson, err = json.Marshal(&encryptedKeyJSONV3)
	if err != nil {
		return
	}
	var str bytes.Buffer
	_ = json.Indent(&str, keyJson, "", "    ")

	keyJson = str.Bytes()
	return
}

func DecryptKey(keyjson []byte, auth string) (key []byte, err error) {
	encryptedKeyJSONV3 := &encryptedKeyJSONV3{}

	err = json.Unmarshal(keyjson, encryptedKeyJSONV3)

	if err != nil {
		return
	}

	return keystore.DecryptDataV3(encryptedKeyJSONV3.Crypto, auth)

}

type encryptedKeyJSONV3 struct {
	//Address string     `json:"address"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
	Id      string              `json:"id"`
	Version int                 `json:"version"`
}
