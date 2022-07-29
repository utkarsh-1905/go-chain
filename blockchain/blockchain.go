package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/utkarsh-1905/go-chain/block"
	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
)

func Genesis() []block.Block {
	var tx []transactions.Transaction
	tx = append(tx, transactions.Transaction{
		To:    "MzMxNTQwOTM1MDY4NDQ5ODg5NTg4MjEwMDE0NzQ1MjkwMDA3NjI3MDQxNjgxOTc0NzM1MDM4MjAxOTU4OTg5MDMyNzkyNTU3OTMwMTA=",
		From:  "0",
		Value: 1000,
		Data:  "Genesis Block, 1000 tokens created",
	})
	var blockchain = make([]block.Block, 0)
	block := block.Block{
		BlockNumber:          1,
		ParentHash:           "",
		Timestamp:            time.Now().Unix(),
		NumberOfTransactions: len(tx),
		Miner:                "utkarsh",
		Nonce:                10,
		Transactions:         tx,
	}
	blockData, _ := json.Marshal(block)
	h := sha256.New()
	h.Write([]byte(string(blockData)))
	chash := fmt.Sprintf("%x", h.Sum(nil)) //thanks copilot
	block.CurrHash = chash
	blockchain = append(blockchain, block)
	bchain, _ := json.MarshalIndent(blockchain, "", "\t")
	_ = ioutil.WriteFile("blockchain.json", bchain, 0644)
	return blockchain
}

func GetLatestBlock() block.Block {
	blockchain := helpers.ReadAndUnmarshallBlockchain()
	return blockchain[len(blockchain)-1]
}

func CreateBlock(mnr string) block.Block {
	// blockchain := helpers.ReadAndUnmarshallBlockchain()
	var chash string
	block := block.Block{
		BlockNumber:          (GetLatestBlock().BlockNumber) + 1,
		ParentHash:           GetLatestBlock().CurrHash,
		Timestamp:            time.Now().Unix(),
		NumberOfTransactions: len(transactions.Pool()),
		Miner:                mnr,
		Nonce:                10,
		Transactions:         transactions.Pool(),
	}
	blockData, _ := json.Marshal(block)
	h := sha256.New()
	h.Write([]byte(string(blockData)))
	chash = fmt.Sprintf("%x", h.Sum(nil)) //thanks copilot
	block.CurrHash = chash
	// blockchain = append(blockchain, block)
	// transactions.Pool = nil
	// bchain, _ := json.MarshalIndent(blockchain, "", "\t")
	// _ = ioutil.WriteFile("blockchain.json", bchain, 0644)
	return block
}
