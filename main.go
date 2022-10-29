package main

import (
	"log"

	gnats "github.com/utkarsh-1905/go-chain/nats"
)

func main() {
	nc := gnats.Connect()
	err := nc.Publish("block", []byte("block created"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
}
