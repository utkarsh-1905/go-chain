package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/utkarsh-1905/go-chain/block"
	"github.com/utkarsh-1905/go-chain/transactions"
)

var Blockchain = make([]block.Block, 0)

func Genesis() []block.Block {
	var tx []transactions.Transaction
	tx = append(tx, transactions.Transaction{To: "", From: "", Value: 1000, Data: ""})
	Blockchain = append(Blockchain, block.Block{BlockNumber: 1, ParentHash: "", CurrHash: "", Timestamp: time.Now().Unix(), NumberOfTransactions: 0, Miner: "", Nonce: 0, Transactions: tx})
	return Blockchain
}

func GetLatestBlock() block.Block {
	return Blockchain[len(Blockchain)-1]
}

func CreateBlock(mnr string) block.Block {
	var chash string
	t := int64(time.Now().Unix())
	block := block.Block{
		BlockNumber:          (GetLatestBlock().BlockNumber) + 1,
		ParentHash:           GetLatestBlock().CurrHash,
		Timestamp:            t,
		NumberOfTransactions: len(transactions.Pool),
		Miner:                mnr,
		Nonce:                10,
		Transactions:         transactions.Pool,
	}
	blockData, _ := json.Marshal(block)
	h := sha256.New()
	h.Write([]byte(string(blockData)))
	chash = fmt.Sprintf("%x", h.Sum(nil)) //thanks copilot
	block.CurrHash = chash
	Blockchain = append(Blockchain, block)
	transactions.Pool = nil
	return block
}
