// library project main.go
package main

import (
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"library/Chain3Go"
	//	"library/Chain3Go/utils"
	"library/Chain3Go/requestData"
	"library/Chain3Go/types"
	//	"github.com/innowells/moac-lib/common/hexutil"
	"sync"
)

var group sync.WaitGroup

func main() {

	//	jsonStr := `{"address":"d58592114ebd97525856929d5c662b72d58b767b","crypto":{"cipher":"aes-128-ctr","ciphertext":"db96d030406419a3ca0d6e6901b3b688ad4c6f34376f048c8ed56f39d1b37169","cipherparams":{"iv":"9e6aee2025bd919866da952d3df40566"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f22f59cac654619c29efba26fef1c1b894691b0e2e6a2bf903104dae96d824c4"},"mac":"af807312730a542925badcb3830853a43324c8b19b27cbefff7e72f150e6576f"},"id":"d4bfe897-6bcb-4f5d-9166-cc9a4b5ba488","version":3}`
	//	Chain3Go.GetPrivateKey(jsonStr, "Wlr7523286")

	var rpcClient *Chain3Go.RpcClient
	//via : 0xcEAC4CC8524a5f8afa2cCA6cbde7270F1942b2d5
	rpcClient = Chain3Go.NewRpcClient("http://192.168.1.23:8545", 101) //子链
	//	fmt.Println(rpcClient.Mc().MC_getTransactionCount("0xd58592114ebd97525856929d5c662b72d58b767b", "latest"))

	//	rpcClient = Chain3Go.NewRpcClient("http://35.236.108.149:8545", 101) //主链

	//	rpcClient = Chain3Go.NewRpcClient("http://47.107.99.26:8545", 101)

	//	rpcClient = Chain3Go.NewRpcClient("http://192.168.1.15:8545", 99) //主链

	//	rpcClient = Chain3Go.NewRpcClient("http://106.14.225.49:8545", 99) //主链

	//	rpcClient = Chain3Go.NewRpcClient("http://192.168.1.34:8545", 99) //主链

	//	fmt.Println(rpcClient.Mc().CHAIN3_clientVersion())
	//	fmt.Println(rpcClient.Mc().CHAIN3_sha3("123"))
	//	fmt.Println(rpcClient.Mc().NET_version())
	//	fmt.Println(rpcClient.Mc().NET_listening())
	//	fmt.Println(rpcClient.Mc().NET_peerCount())
	//	fmt.Println(rpcClient.Mc().MC_protocolVersion())
	//	fmt.Println(rpcClient.Mc().MC_syncing())
	//	fmt.Println(rpcClient.Mc().MC_coinbase())
	//	fmt.Println(rpcClient.Mc().MC_mining())
	//	fmt.Println(rpcClient.Mc().MC_hashrate())
	//	fmt.Println(rpcClient.Mc().MC_gasPrice())
	//	fmt.Println(rpcClient.Mc().MC_accounts())
	//	fmt.Println(rpcClient.Mc().MC_blockNumber())
	//	fmt.Println(rpcClient.Mc().MC_getBalance("0x955687b78de0a7336d144138dabc4de419d5bc41", "latest"))
	//	fmt.Println(rpcClient.Mc().MC_getStorageAt("0xd58592114ebd97525856929d5c662b72d58b767b", "0x1", "latest"))
	//	fmt.Println(rpcClient.Mc().MC_getTransactionCount("0xd58592114ebd97525856929d5c662b72d58b767b", "latest"))
	//	fmt.Println(rpcClient.Mc().MC_getBlockTransactionCountByHash("0x0f2b4da92a7ffd6e47d71a59022c7c8dddea3d38ab01e26663188122014a7fde"))
	//	fmt.Println(rpcClient.Mc().MC_getBlockTransactionCountByNumber("0x112842"))
	//	fmt.Println(rpcClient.Mc().MC_getUncleCountByBlockHash("0x0f2b4da92a7ffd6e47d71a59022c7c8dddea3d38ab01e26663188122014a7fde"))
	//	fmt.Println(rpcClient.Mc().MC_getUncleCountByBlockNumber("0x112842"))
	//	fmt.Println(rpcClient.Mc().MC_getCode("0xd58592114ebd97525856929d5c662b72d58b767b", "latest"))
	//	fmt.Println(rpcClient.Mc().MC_sign("0xd58592114ebd97525856929d5c662b72d58b767b", "0x12345678"))

	//	fmt.Println(rpcClient.Mc().PERSONAL_unlockAccount("0xd58592114ebd97525856929d5c662b72d58b767b", "Wlr7523286"))
	//	trData := new(requestData.TransactionParameters)
	//	trData.From = "0xd58592114ebd97525856929d5c662b72d58b767b"
	//	trData.To = "0x50c1fc4ce9ae1012a942fb411ee69b94623c69f8"
	//	trData.Gas = 50000
	//	trData.GasPrice = 20000000000
	//	trData.Value = 10000000000000000
	//	trData.Data = "0x"
	//	fmt.Println(rpcClient.Mc().MC_sendTransaction(trData))
	//	fmt.Println(rpcClient.Mc().PERSONAL_lockAccount("0xd58592114ebd97525856929d5c662b72d58b767b"))

	//	fmt.Println(rpcClient.Mc().PERSONAL_unlockAccount("0xe3153Ca7b163cE80150cC02b16A58a7D197EfF52", "123456"))
	//	fmt.Println(rpcClient.Mc().PERSONAL_unlockAccount("0x7419C28e9283369FF7bEb8B3E9d1B8F622A9DCB2", "Wlr7523286"))
	//	trData := new(requestData.TransactionParameters)
	//	trData.From = "0xd58592114ebd97525856929d5c662b72d58b767b"
	//	trData.To = "0x60a602e92a86A44c89873f92BF15258b6b114E34"
	//	trData.Gas = 50000
	//	trData.GasPrice = 20000000000
	//	trData.Value = 0
	//	trData.Data = "0xa9059cbb00000000000000000000000050c1Fc4Ce9ae1012a942FB411ee69B94623C69F80000000000000000000000000000000000000000000000000000000000002710"
	//	fmt.Println(rpcClient.Mc().MC_sendTransaction(trData))
	//	fmt.Println(rpcClient.Mc().PERSONAL_lockAccount("0xd58592114ebd97525856929d5c662b72d58b767b"))

	//	txData := new(requestData.TransactionParameters)
	//	txData.To = "0x60a602e92a86A44c89873f92BF15258b6b114E34"
	//	//	txData.Data = "0x06fdde03"
	//	txData.Data = "0x70a0823100000000000000000000000050c1Fc4Ce9ae1012a942FB411ee69B94623C69F8"
	//	fmt.Println(rpcClient.Mc().MC_call(txData, "latest"))

	//	txData := new(requestData.TransactionParameters)
	//	txData.From = "0xd58592114ebd97525856929d5c662b72d58b767b"
	//	txData.To = "0x50c1fc4ce9ae1012a942fb411ee69b94623c69f8"
	//	txData.Gas = 50000
	//	txData.GasPrice = 20000000000
	//	txData.Value = 10000000000000000
	//	txData.Data = "0x"
	//	fmt.Println(rpcClient.Mc().MC_estimateGas(txData))

	//	fmt.Println(rpcClient.Mc().MC_getBlockByHash("0xdebef8cd4882c4ab62b93357833ca12557f65d7bc15cb999ecb77367b71e926e", true))
	//	fmt.Println(rpcClient.Mc().MC_getBlockByNumber(608801, true))
	//	fmt.Println(rpcClient.Mc().MC_getTransactionByHash("0xd5f6af416db45577bd95829d1353229da37231a8fe18eabbcb1f210e64d246c4"))
	//	fmt.Println(rpcClient.Mc().MC_getTransactionByBlockHashAndIndex("0xdebef8cd4882c4ab62b93357833ca12557f65d7bc15cb999ecb77367b71e926e", 1))
	//	fmt.Println(rpcClient.Mc().MC_getTransactionByBlockNumberAndIndex(1137688, 1))
	//	fmt.Println(rpcClient.Mc().MC_getTransactionReceipt("0xb5e813306265e80d4af39c09cf06039d75a892ddca8931479a5f6afa87a142a2"))
	//	fmt.Println(rpcClient.Mc().MC_getUncleByBlockHashAndIndex("0xdebef8cd4882c4ab62b93357833ca12557f65d7bc15cb999ecb77367b71e926e", 0))
	//	fmt.Println(rpcClient.Mc().MC_getUncleByBlockNumberAndIndex("115c18", 0))
	//	fmt.Println(rpcClient.Mc().MC_newFilter("", "", "", []interface{}{"", ""}))
	//	fmt.Println(rpcClient.Mc().MC_newBlockFilter())
	//	fmt.Println(rpcClient.Mc().MC_newPendingTransactionFilter())
	//	fmt.Println(rpcClient.Mc().MC_uninstallFilter("0x8ae65655e5a094e9cef2fb2774b795b2"))
	//	fmt.Println(rpcClient.Mc().MC_getFilterChanges("0xb17907cfe1c307a187e3a99ae6ed1562"))
	//	fmt.Println(rpcClient.Mc().MC_getFilterLogs("0x1d82393d72c9537c8275f85650e1dada"))
	//	fmt.Println(rpcClient.Mc().MC_getLogs([]string{"0x1d82393d72c9537c8275f85650e1dada"}))
	//	fmt.Println(rpcClient.Mc().MC_getWork())

	//	fmt.Println(rpcClient.Mc().VNODE_address())
	//	fmt.Println(rpcClient.Mc().VNODE_scsService())
	//	fmt.Println(rpcClient.Mc().VNODE_serviceCfg())
	//	fmt.Println(rpcClient.Mc().VNODE_showToPublic())
	//	fmt.Println(rpcClient.Mc().VNODE_vnodeIP())
	//子链call	fmt.Println(rpcClient.Mc().VNODE_directCall())

	//	placeholderStr := "0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000002e"
	//remove:0x80599e4b
	//read:0x616ffe83
	//write:0xebaac771
	//	txParams := new(requestData.TransactionParameters)
	//	txParams.From = "0xd58592114ebd97525856929d5c662b72d58b767b"
	//	txParams.To = "0x6abeb836a97312aae6ff93528640f75ba2cd0732"
	//	txParams.Data = "0xa87d942c"
	//	fmt.Println(rpcClient.Mc().SCS_directCall(txParams))

	//	fmt.Println(rpcClient.Mc().SCS_getBlock("0x8d3a5eC8EE86262790bed060f8751851173F1112", "0x186a0"))
	//	fmt.Println(rpcClient.Mc().SCS_getBlockNumber("0x6abeb836a97312aae6ff93528640f75ba2cd0732"))
	//	fmt.Println(rpcClient.Mc().SCS_getDappState("0x8d3a5eC8EE86262790bed060f8751851173F1112"))
	//	fmt.Println(rpcClient.Mc().SCS_getMicroChainList())
	//	fmt.Println(rpcClient.Mc().SCS_getNonce("0xA0b2a81afFd03522aF38664730FFD2300A359CC4", "0xAC99cC5b00f41FC772AA0ebE33Ac03D94Bc31444"))
	//	fmt.Println(rpcClient.Mc().SCS_getSCSId())

	//	fmt.Println(rpcClient.Mc().SCS_getTransactionReceipt("0xA0b2a81afFd03522aF38664730FFD2300A359CC4", "0x674f5a346252186bc3c1e4fa1c580fe7bf53c2edb4309f8341ea2baa99ff8bb4"))

	//	fmt.Println(rpcClient.Mc().PERSONAL_newAccount("Wlr7523286"))
	//	fmt.Println(rpcClient.Mc().PERSONAL_sign("0x01", "0x5b40d3f6c19a0e9567fac86c334da31277ce1648", "Wlr7523286"))
	//	fmt.Println(rpcClient.Mc().PERSONAL_ecRecover("0x01", "0x718e1f7e16e10fd96a2fb11dee2b2e8115f0d3640344d523fd160869a43940ad540ad2acf4a53efd716a54434f5d928fc6254fc01a6c5e5990967dabddc398071b"))

	//	trData := new(requestData.TransactionParameters)
	//	trData.From = "0xd58592114ebd97525856929d5c662b72d58b767b"
	//	trData.To = "0x50c1fc4ce9ae1012a942fb411ee69b94623c69f8"
	//	trData.Gas = 50000
	//	trData.GasPrice = 20000000000
	//	trData.Value = 10000000000000000
	//	trData.Data = "0x"
	//	fmt.Println(rpcClient.Mc().PERSONAL_sendTransaction(trData, "Wlr7523286"))

	//	fmt.Println(Chain3Go.CreateKeystoreAddress("Wlr7523286"))

	//转账
	signStr, err := Chain3Go.TxSign(
		`{"address":"955687b78de0a7336d144138dabc4de419d5bc41","crypto":{"cipher":"aes-128-ctr","ciphertext":"d7e07bb3cd77e6ac9914d495732d3aaf7cbd51425c384663fff81b8f16aa5b70","cipherparams":{"iv":"9f02ff24f52933e9617502e8b539ca8f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"3b0a4ebe07c7d6e4384cf87e8b87bdb3312f1bf1e1098699c527dc6eb54dae52"},"mac":"746a2b5ec1f8588e509dec17fce4e5b20d38a9e0bf74aa147fed6df0073b881c"},"id":"820bdcf9-5c24-4b11-b04d-3e5a4d8a5696","version":3}`,
		"Wlr7523286",
		"0x955687b78de0a7336d144138dabc4de419d5bc41",
		"0xea1c67e31f56b878af06de6b6a7eee1ac6e9836a",
		big.NewInt(9000000000000000000),
		big.NewInt(500000),
		big.NewInt(20000000000),
		0,
		nil,
	)
	fmt.Println(signStr)
	if err == nil {
		fmt.Println(rpcClient.Mc().MC_sendRawTransaction(signStr))
	}

	//	bytes, _ := hexutil.Decode("0xa9059cbb00000000000000000000000054f3fb40df3cf0d234839ae237e8c500fabd71a9000000000000000000000000000000000000000000000000112210f4b16c1d00")

	//	//转账
	//	signStr, err := Chain3Go.TxSign(
	//		`{"address":"7312f4b8a4457a36827f185325fd6b66a3f8bb8b","crypto":{"cipher":"aes-128-ctr","ciphertext":"ab6e3834958a5f2969d767a21915d96466733e2aa5d873d1b778ad82a3c0874c","cipherparams":{"iv":"bcf3bebd97bdd5095d53cc0aa0b564fa"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"67c206ffb7ba8c39d55157d7bd7799dbd0734dcb7ca5a06ca8c492f78f381240"},"mac":"b07364780bf92df5fcdaef17127f3842ccba34ac0264278b5e741a916df59d0e"},"id":"7ceb8ed1-14e8-4c09-a47d-5dcd612d84b3","version":3}`,
	//		"test",
	//		"0x7312F4B8A4457a36827f185325Fd6B66a3f8BB8B",
	//		"0xf2f4eec6c2adfcf780aae828de0b25f86506ffae",
	//		big.NewInt(0),
	//		big.NewInt(4000100),
	//		big.NewInt(30000000000),
	//		0,
	//		bytes,
	//	)
	//	if err == nil {
	//		//		fmt.Println(rpcClient.Mc().MC_sendRawTransaction(signStr))
	//		fmt.Println(signStr)
	//	} else {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(signStr, err)

	//	group.Add(1)
	//	for k := 0; k < 1; k++ {
	//		go txCBE()
	//	}
	//	group.Wait()
}

func txCBE() {

	var rpcClient *Chain3Go.RpcClient

	rpcClient = Chain3Go.NewRpcClient("http://192.168.1.15:8545", 99)

	for i := 0; i < 1; i++ {
		jsonStr, address, err := Chain3Go.CreateKeystoreAddress("Wlr7523286")

		data := "0xa9059cbb000000000000000000000000" + address[2:] + txNumber()

		timeStr := time.Now().Add(-8 * time.Hour).String()[:26]
		timeStr = strings.Replace(timeStr, " ", "T", -1)
		timeStr = strings.Replace(timeStr, ":", "-", -1)
		fileNameStr := "UTC--" + timeStr + "000Z--" + strings.ToLower(address[2:])
		ioutil.WriteFile("./keystore/"+fileNameStr, []byte(jsonStr), 0644)

		if err == nil {
			//转账
			trData := new(requestData.TransactionParameters)
			trData.From = "0xd58592114ebd97525856929d5c662b72d58b767b"
			trData.To = "0xd98792127cb7a0953669f2986af6fcaa37e40cd0"
			trData.Gas = 90000
			trData.GasPrice = 5000000000
			trData.Value = 0
			trData.Data = types.ComplexString(data)
			fmt.Println(rpcClient.Mc().PERSONAL_sendTransaction(trData, "Wlr7523286"))
		}
	}
	group.Done()
}

func txNumber() string {

	var number uint64
	//以当前的系统时间作为种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		number = rand.Uint64() % 1000000
		if number >= 500000 {
			break
		}
	}

	tmpStr := strconv.FormatUint(number, 16)
	return pieceTogether(tmpStr, "0", 64, true)
}

//拼接字符串 flag:true(前补) flag:false(后补)
func pieceTogether(str, subStr string, number uint64, flag bool) string {

	var reParam string = str
	for i := len(str); uint64(i) < number; i++ {
		if flag {
			reParam = subStr + reParam[:]
		} else {
			reParam = reParam[:] + subStr
		}
	}

	return reParam
}
