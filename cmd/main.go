package main

import (
	"context"
	"fmt"
	ethereumWallet "github.com/Amirilidan78/ethereum-wallet"
	"github.com/Amirilidan78/ethereum-wallet/geth"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const fromAddress = "0x5A2187B6d76a09F649CDC5d69F182697cFBA126B"
const toAddress = "0x75c07e7207Bb00Cf60c77f2506D7CE2B8d18bf0f"
const PrivateKey = "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec"
const HTTP_NODE = "https://sepolia.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd"
const WS_NODE = "wss://sepolia.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd"

func main() {

	crawl()
}

func nonce() {
	client, _ := geth.GetGETHClient(HTTP_NODE)

	n, _ := client.NonceAt(context.Background(), common.HexToAddress(fromAddress), nil)
	pn, _ := client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))

	fmt.Println("===============")
	fmt.Println("===Nonce==")
	fmt.Println(n)
	fmt.Println("===Pending Nonce==")
	fmt.Println(pn)
	fmt.Println("===============")

}

func crawl() {

	c := ethereumWallet.Crawler{
		Node: ethereumWallet.Node{
			Http: HTTP_NODE,
			Ws:   WS_NODE,
		},
		Addresses: []string{fromAddress},
	}

	res, err := c.ScanBlocks(10)

	fmt.Println(res)
	fmt.Println(err)

}

func transfer() {

	node := ethereumWallet.Node{
		Http: HTTP_NODE,
		Ws:   WS_NODE,
	}

	w := ethereumWallet.CreateEthereumWallet(node, PrivateKey)

	txId, err := w.Transfer(toAddress, big.NewInt(100000000000000)) // 0.0001 eth

	fmt.Println(txId)
	fmt.Println(err)

}
