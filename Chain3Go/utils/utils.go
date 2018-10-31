// utils project utils.go
package utils

import (
	"bytes"
	"encoding/json"
	"math/big"
	"strconv"

	"library/Chain3Go/accounts/keystore"
	"library/Chain3Go/utils"

	"github.com/tonnerre/golang-go.crypto/sha3"

	"github.com/innowells/moac-lib/common"
	"github.com/innowells/moac-lib/rlp"
	"github.com/innowells/moac-lib/types"
)

type JsonRpc struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

//转换成json字符串
func (jRpc *JsonRpc) AsJsonString() string {

	resultBytes, err := json.Marshal(jRpc)
	if err != nil {
		return ""
	}
	return string(resultBytes)
}

func bytesToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {
		s := strconv.FormatInt(int64(b&0xFF), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	return "0x" + buffer.String()
}

func Sha3Hash(dataStr string) string {

	d := sha3.NewKeccak256()
	d.Write([]byte(dataStr))

	return bytesToHex(d.Sum(nil))[:10]
}

//根据keystore字符串获取私钥
func GetPrivateKey(jsonStr string, password string) (*keystore.Key, error) {

	var jsonBytes []byte
	jsonBytes = []byte(jsonStr)

	var err error
	var storeKey *keystore.Key
	storeKey, err = keystore.DecryptKey(jsonBytes, password)

	return storeKey, err
}

//创建地址返回keystore字符串
func CreateKeystoreAddress(tradePassword string) (string, error) {
	//创建地址

	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP
	key, err := keystore.CreateKey()
	if err == nil {
		keyjson, err := keystore.EncryptKey(key, tradePassword, scryptN, scryptP)
		return string(keyjson), err
	}
	return "", err
}

//交易签名
func TxSign(keystoreStr, keystorePassword, from, to string, amount, gas, gasPrice *big.Int, fromNonce uint64) (string, error) {

	via := common.HexToAddress("0x868183417b201a31287cdc60058c7626ee53fea6")

	tx := types.NewTransaction(fromNonce, common.HexToAddress(to), amount, gas, gasPrice, 0, &via, nil)
	priKey, _ := utils.GetPrivateKey(keystoreStr, keystorePassword) //获取私钥

	tx, err = types.SignTx(tx, types.NewPanguSigner(big.NewInt(101)), priKey.PrivateKey) //根据私钥进行签名
	if err != nil {
		return "", err
	}

	enc, err := rlp.EncodeToBytes(&tx.TxData) //解析获取交易Hex

	return fmt.Sprintf("0x%x", enc), err
}
