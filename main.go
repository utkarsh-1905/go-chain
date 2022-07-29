package main

import (
	"encoding/json"
	"fmt"

	"github.com/utkarsh-1905/go-chain/blockchain"
	"github.com/utkarsh-1905/go-chain/transactions"
)

func main() {
	blockchain.Genesis()
	transactions.CreateTransaction("utkarsh", "utkarsh", 10, "hello")
	b1 := blockchain.CreateBlock("utkarsh")
	b1json, _ := json.Marshal(b1)
	fmt.Println(string(b1json))
	fmt.Println("----------------------------------------------------")
	transactions.CreateTransaction("harish", "tushar", 7, "helu")
	b2 := blockchain.CreateBlock("utkarsh")
	b2json, _ := json.Marshal(b2)
	fmt.Println(string(b2json))
	fmt.Println("----------------------------------------------------")
	bc := blockchain.Blockchain
	bcjson, _ := json.Marshal(bc)
	fmt.Println(string(bcjson))
}
