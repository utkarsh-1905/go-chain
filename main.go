package main

import (
	"fmt"

	"github.com/utkarsh-1905/go-chain/blockchain"
	"github.com/utkarsh-1905/go-chain/transactions"
)

func main() {
	blockchain.Genesis()
	transactions.CreateTransaction("utkarsh", "utkarsh", 10, "hello")
	blockchain.CreateBlock()
	transactions.CreateTransaction("harish", "tushar", 7, "helu")
	blockchain.CreateBlock()
	bc := blockchain.Blockchain
	fmt.Println(bc)
}
