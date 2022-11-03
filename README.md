# ethereum-wallet
ethereum wallet package for creating and generating wallet, transferring ETH, getting wallet balance and crawling blocks to find wallet transactions


### Nodes 
using infura nodes in this repository 
```
node := ethereumWallet.Node{
		Http: "https://goerli.infura.io/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
		Ws:   "wss://goerli.infura.io/ws/v3/89aae5ec52f9450ebe4fc58cbb8138fd",
}
```

### Main methods
- generating ethereum wallet
```
w := GenerateEthereumWallet(node)
w.Address // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```
- creating ethereum wallet from private key
```
w := CreateEthereumWallet(node,privateKeyHex)
w.Address // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```
- getting wallet balance
```
balanceInWei,err := w.Balance()
balanceInWei // int64 
```