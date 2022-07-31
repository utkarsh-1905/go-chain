package main

import (
	"fmt"
	"io"
	"net"

	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/pos"
)

//this file is the node(host) of the blockchain

func main() {

	server, err := net.Listen("tcp", ":9000")
	helpers.HandleErr(err)

	defer server.Close()

	for {
		conn, err := server.Accept()
		helpers.HandleErr(err)
		go showStakePool(conn)
	}

	// blockchain.Genesis()
	// transactions.CreateTransaction("utkarsh", "utkarsh", 10, "hello")
	// b1 := blockchain.CreateBlock("utkarsh")
	// b1json, _ := json.Marshal(b1)
	// fmt.Println(string(b1json))
	// fmt.Println("----------------------------------------------------")
	// transactions.CreateTransaction("harish", "tushar", 7, "helu")
	// b2 := blockchain.CreateBlock("utkarsh")
	// b2json, _ := json.Marshal(b2)
	// fmt.Println(string(b2json))
	// fmt.Println("----------------------------------------------------")
}

func showStakePool(conn net.Conn) {
	var total int
	total += <-pos.StakePool
	io.WriteString(conn, fmt.Sprint(total))
	defer conn.Close()
}
