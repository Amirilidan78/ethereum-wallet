package main

import (
	"fmt"
	ethereumWallet "github.com/Amirilidan78/ethereum-wallet"
)

const PrivateKey = "a24031202755246def61140ae1bce297d0c4886b2faea5ce79001748ef97e8ec"
const GOERLI_HTTP_NODE = "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd"
const GOERLI_WS_NODE = "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd"

func main() {

	w := ethereumWallet.CreateEthereumWallet(ethereumWallet.Node{
		Http: GOERLI_HTTP_NODE,
		Ws:   GOERLI_WS_NODE,
	}, PrivateKey)

	fmt.Println(w.Balance())

}
