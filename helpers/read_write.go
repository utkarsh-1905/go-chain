package helpers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/utkarsh-1905/go-chain/block"
)

func ReadAndUnmarshallBlockchain() []block.Block {
	content, err := ioutil.ReadFile("blockchain.json")
	HandleErr(err)
	var blockchain []block.Block
	_ = json.Unmarshal(content, &blockchain)
	return blockchain
}
