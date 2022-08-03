package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utkarsh-1905/go-chain/miner"
)

//this file is the node(host) of the blockchain

var StakePool int

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/stake", ShowStakePool).Methods("GET")
	router.HandleFunc("/stake", AddToPool).Methods("POST")

	http.ListenAndServe(":9000", router)
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

func ShowStakePool(w http.ResponseWriter, r *http.Request) {

}

func AddToPool(w http.ResponseWriter, r *http.Request) {
	var data miner.StakeData
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
}
