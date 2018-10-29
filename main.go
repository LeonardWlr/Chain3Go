// library project main.go
package main

import (
	//	"encoding/json"
	"fmt"

	"library/Lchain3"
	//	"library/Lchain3/requestData"
	//	"library/Lchain3/accounts/keystore"
)

func main() {

	var rpcClient *Lchain3.RpcClient
	//	rpcClient = Lchain3.NewRpcClient("http://192.168.3.35:50066")
	rpcClient = Lchain3.NewRpcClient("http://192.168.3.9:8545")

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
	//	fmt.Println(rpcClient.Mc().MC_getBalance("0xd58592114ebd97525856929d5c662b72d58b767b", "latest"))
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

	fmt.Println(rpcClient.Mc().PERSONAL_unlockAccount("0xd58592114ebd97525856929d5c662b72d58b767b", "Wlr7523286"))
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

	//	signTxStr := "0x123456789"
	//	fmt.Println(rpcClient.Mc().MC_sendRawTransaction(signTxStr))

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
	//	fmt.Println(rpcClient.Mc().MC_getTransactionReceipt("0xd5f6af416db45577bd95829d1353229da37231a8fe18eabbcb1f210e64d246c4"))
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

	//	fmt.Println(rpcClient.Mc().SCS_getBlock("0x8d3a5eC8EE86262790bed060f8751851173F1112", "0x186a0"))
	//	fmt.Println(rpcClient.Mc().SCS_getBlockNumber("0x8d3a5eC8EE86262790bed060f8751851173F1112"))
	//	fmt.Println(rpcClient.Mc().SCS_getDappState("0x8d3a5eC8EE86262790bed060f8751851173F1112"))
	//	fmt.Println(rpcClient.Mc().SCS_getMicroChainList())
	//	fmt.Println(rpcClient.Mc().SCS_getNonce("0x8d3a5eC8EE86262790bed060f8751851173F1112", "0x180d14bc9db75d64213b64fccbe5b7338baf777d"))
	//	fmt.Println(rpcClient.Mc().SCS_getSCSId())

	//未测试	fmt.Println(rpcClient.Mc().SCS_getTransactionReceipt("",""))

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

}
