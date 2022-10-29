package gnats

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func Connect() *nats.Conn {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal("Could not connect to NATS server")
		log.Fatal(err)
	}
	fmt.Println("Connected to NATS server")
	return nc
}
