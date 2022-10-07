package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//this file is the node(host) of the blockchain

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {

	// router := mux.NewRouter()

	var Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		ws.WriteMessage(1, []byte("Hello"))
		reader(ws)
	})

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

func ShowStakePool(w http.ResponseWriter, r *http.Request) {

}

// func AddToPool(w http.ResponseWriter, r *http.Request) {
// 	var data miner.StakeData
// 	_ = json.NewDecoder(r.Body).Decode(&data)
// 	fmt.Println(data)
// }
