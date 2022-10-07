package transactions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/gorilla/websocket"
)

type Transaction struct {
	To    string `json:"to"`
	From  string `json:"from"`
	Value int    `json:"value"`
	Data  string `json:"data"`
}

func Pool() []Transaction {
	content, err := ioutil.ReadFile("mempool.json")
	if err != nil {
		fmt.Println("No transactions in mempool")
		return nil
	}
	var mempool []Transaction
	_ = json.Unmarshal([]byte(content), &mempool)
	return mempool
}

//create transaction or transfer tokens, and send it to Pool
//If transaction mined, the wallet file will show the updated balance

func CreateTransaction(to string, from string, value int, data string) Transaction {
	tx := Transaction{to, from, value, data}
	content, err := ioutil.ReadFile("mempool.json")
	if err != nil {
		fmt.Println("Creating new mempool")
		var mempool []Transaction
		mempool = append(mempool, tx)
		wl, _ := json.MarshalIndent(mempool, "", "\t")
		_ = ioutil.WriteFile("mempool.json", wl, 0644)
	} else {
		var mempool []Transaction
		_ = json.Unmarshal([]byte(content), &mempool)
		mempool = append(mempool, tx)
		wl, _ := json.MarshalIndent(mempool, "", "\t")
		_ = ioutil.WriteFile("mempool.json", wl, 0644)
	}
	fmt.Println("Transaction created")
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/tx"}
	conn, _, wsserr := websocket.DefaultDialer.Dial(u.String(), nil)
	if wsserr != nil {
		panic(wsserr)
	}
	werr := conn.WriteJSON(tx)
	if werr != nil {
		panic(werr)
	}
	defer conn.Close()
	return tx
}
