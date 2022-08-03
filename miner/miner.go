package miner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/wallet"
)

type StakeData struct {
	Amount   int    `json:"amount"`
	PublicId string `json:"public_id"`
}

func Stake(amount int, name string) {
	content := helpers.ReadFileWithName(name)
	var wallet wallet.Wallet
	_ = json.Unmarshal([]byte(content), &wallet)
	wallet.Balance = 100
	if amount < wallet.Balance {
		body, _ := json.Marshal(StakeData{Amount: amount, PublicId: string(wallet.PublicKey)})
		http.Post("http://localhost:9000/stake", "application/json", bytes.NewBuffer(body))
		wallet.Balance -= amount
		wl, _ := json.MarshalIndent(wallet, "", "\t")
		_ = ioutil.WriteFile(name+".json", wl, 0644)
	} else {
		fmt.Println("Insufficient balance")
		wl, _ := json.MarshalIndent(wallet, "", "\t")
		_ = ioutil.WriteFile(name+".json", wl, 0644)

	}
}
