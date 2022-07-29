package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"

	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
)

type Wallet struct {
	Balance      int                        `json:"balance"`
	PrivateKey   []byte                     `json:"privateKey"`
	PublicKey    []byte                     `json:"publicKey"`
	Transactions []transactions.Transaction `json:"transactions"`
}

func GenerateWallet(name string) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	helpers.HandleErr(err)
	pubKey, _ := privateKey.PublicKey.X.MarshalJSON()
	privKey, _ := privateKey.Y.MarshalJSON()
	wallet := Wallet{
		Balance:    0,
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
	wl, _ := json.MarshalIndent(wallet, "", "\t")
	_ = ioutil.WriteFile(name+".json", wl, 0644)
}
