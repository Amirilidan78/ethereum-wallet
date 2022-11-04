package main

import (
	"fmt"
	ethereumWallet "github.com/Amirilidan78/ethereum-wallet"
	"math/big"
)

const PrivateKey = "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec"
const GOERLI_HTTP_NODE = "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd"
const GOERLI_WS_NODE = "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd"

func main() {

	toAddress := "0x75c07e7207Bb00Cf60c77f2506D7CE2B8d18bf0f"

	node := ethereumWallet.Node{
		Http: GOERLI_HTTP_NODE,
		Ws:   GOERLI_WS_NODE,
	}

	w := ethereumWallet.CreateEthereumWallet(node, PrivateKey)

	txId, err := w.Transfer(toAddress, big.NewInt(100000000000000)) // 0.0001 eth

	fmt.Println(txId)
	fmt.Println(err)

	//
	//c := ethereumWallet.Crawler{
	//	Node:      node,
	//	Addresses: []string{w.Address},
	//}
	//fmt.Println(c.ScanBlocks(3))

}
