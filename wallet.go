package ethereumWallet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/Amirilidan78/ethereum-wallet/geth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type Node struct {
	Http string
	Ws   string
}

type EthereumWallet struct {
	Node       Node
	Address    string
	PrivateKey string
	PublicKey  string
}

func GenerateEthereumWallet(node Node) *EthereumWallet {

	privateKey, _ := generatePrivateKey()
	privateKeyHex := convertPrivateKeyToHex(privateKey)

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address := getAddressFromPublicKey(publicKey)

	return &EthereumWallet{
		Node:       node,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}
}

func CreateEthereumWallet(node Node, privateKeyHex string) *EthereumWallet {

	privateKey, err := privateKeyFromHex(privateKeyHex)
	if err != nil {
		panic(err)
	}

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address := getAddressFromPublicKey(publicKey)

	return &EthereumWallet{
		Node:       node,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
	}
}

// struct functions

func (ew *EthereumWallet) PrivateKeyRCDSA() (*ecdsa.PrivateKey, error) {
	return privateKeyFromHex(ew.PrivateKey)
}

func (ew *EthereumWallet) PrivateKeyBytes() ([]byte, error) {

	priv, err := ew.PrivateKeyRCDSA()
	if err != nil {
		return []byte{}, err
	}

	return crypto.FromECDSA(priv), nil
}

// private key

func generatePrivateKey() (*ecdsa.PrivateKey, error) {

	return crypto.GenerateKey()
}

func convertPrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

func privateKeyFromHex(hex string) (*ecdsa.PrivateKey, error) {

	return crypto.HexToECDSA(hex)
}

// public key

func getPublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error in getting public key")
	}

	return publicKeyECDSA, nil
}

func convertPublicKeyToHex(publicKey *ecdsa.PublicKey) string {

	privateKeyBytes := crypto.FromECDSAPub(publicKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

// address

func getAddressFromPublicKey(publicKey *ecdsa.PublicKey) string {

	return crypto.PubkeyToAddress(*publicKey).Hex()
}

// balance

func (ew *EthereumWallet) Balance() (int64, error) {

	c, err := geth.GetGETHClient(ew.Node.Http)
	if err != nil {
		return 0, err
	}

	balance, err := c.BalanceAt(context.Background(), common.HexToAddress(ew.Address), nil)
	if err != nil {
		return 0, err
	}

	return balance.Int64(), nil
}

// transaction

func (ew *EthereumWallet) Transfer(toAddressHex string, amountInWei *big.Int) (string, error) {

	privateRCDSA, err := ew.PrivateKeyRCDSA()
	if err != nil {
		return "", fmt.Errorf("RCDSA private key error: %v", err)
	}

	tx, err := createTransactionInput(ew.Node, ew.Address, toAddressHex, amountInWei)
	if err != nil {
		return "", fmt.Errorf("creating tx pb error: %v", err)
	}

	tx, err = signTransaction(ew.Node, tx, privateRCDSA)
	if err != nil {
		return "", fmt.Errorf("signing transaction error: %v", err)
	}

	txId, err := broadcastTransaction(ew.Node, tx)
	if err != nil {
		return "", fmt.Errorf("broadcast transaction error: %v", err)
	}

	return txId, nil
}
