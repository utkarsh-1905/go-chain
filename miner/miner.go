package miner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/pos"
	"github.com/utkarsh-1905/go-chain/wallet"
)

func Stake(amount int, name string) {
	content := helpers.ReadFileWithName(name)
	var wallet wallet.Wallet
	_ = json.Unmarshal([]byte(content), &wallet)
	if amount < wallet.Balance {
		pos.StakePool <- amount
		wallet.Balance -= amount
		wl, _ := json.MarshalIndent(wallet, "", "\t")
		_ = ioutil.WriteFile(name+".json", wl, 0644)
	} else {
		fmt.Println("Insufficient balance")
	}
}
