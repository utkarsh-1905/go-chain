package main

import (
	"log"
	"net/http"

	"github.com/utkarsh-1905/go-chain/websocket"
)

//this file is the node(host) of the blockchain

func pubSub(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("pubKey") == "" {
		w.Write([]byte("pubKey is required"))
		return
	}
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Print("upgrade failed: ", err)
		return
	}

	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
		ID:   r.URL.Query().Get("pubKey"),
	}
	pool.Register <- client
	client.Read()
}

func handleConnection() {
	txpool := websocket.NewPool()
	go txpool.Start()

	http.HandleFunc("/tx", func(w http.ResponseWriter, r *http.Request) {
		pubSub(txpool, w, r)
	})
}

func main() {
	handleConnection()
	http.ListenAndServe(":8080", nil)

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

// func ShowStakePool(w http.ResponseWriter, r *http.Request) {

// }

// func AddToPool(w http.ResponseWriter, r *http.Request) {
// 	var data miner.StakeData
// 	_ = json.NewDecoder(r.Body).Decode(&data)
// 	fmt.Println(data)
// }
