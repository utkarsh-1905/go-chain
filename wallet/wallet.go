package wallet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/utkarsh-1905/go-chain/account"
	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
)

type Wallet struct {
	Balance      int                        `json:"balance"`
	PrivateKey   string                     `json:"privateKey"`
	PublicKey    string                     `json:"publicKey"`
	Transactions []transactions.Transaction `json:"transactions"`
}

func GenerateWallet(name string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey := hexutil.Encode(privateKeyBytes)[2:]
	pubKey := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	fmt.Println("Private Key:", privKey)
	fmt.Println("Public Key:", pubKey)
	wallet := Wallet{
		Balance:    0,
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
	account.Initialize(string(pubKey))
	wl, _ := json.MarshalIndent(wallet, "", "\t")
	_ = ioutil.WriteFile(name+".json", wl, 0644)
}

func GetPubKey(name string) string {
	content, err := ioutil.ReadFile(name + ".json")
	helpers.HandleErr(err)
	var wallet Wallet
	_ = json.Unmarshal([]byte(content), &wallet)
	return string(wallet.PublicKey)
}

func GetBalance(name string) int {
	pubKey := GetPubKey(name)
	balance := account.GetBalance(pubKey)
	return balance
}

// func Sign(name string) string {
// 	content, err := ioutil.ReadFile(name + ".json")
// 	helpers.HandleErr(err)
// 	var wallet Wallet
// 	_ = json.Unmarshal([]byte(content), &wallet)
// 	privKey := wallet.PrivateKey
// }
