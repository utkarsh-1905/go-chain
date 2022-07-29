package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/utkarsh-1905/go-chain/blockchain"
	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
)

//this file is the node(host) of the blockchain

func main() {

	server, err := net.Listen("tcp", ":9000")
	helpers.HandleErr(err)

	defer server.Close()

	// blockchain.Genesis()
	transactions.CreateTransaction("utkarsh", "utkarsh", 10, "hello")
	b1 := blockchain.CreateBlock("utkarsh")
	b1json, _ := json.Marshal(b1)
	fmt.Println(string(b1json))
	// fmt.Println("----------------------------------------------------")
	// transactions.CreateTransaction("harish", "tushar", 7, "helu")
	// b2 := blockchain.CreateBlock("utkarsh")
	// b2json, _ := json.Marshal(b2)
	// fmt.Println(string(b2json))
	// fmt.Println("----------------------------------------------------")
}
