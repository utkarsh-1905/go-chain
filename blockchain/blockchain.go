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
	Blockchain = append(Blockchain, block.Block{BlockNumber: 1, ParentHash: "", CurrHash: "", Timestamp: time.Now(), NumberOfTransactions: 0, Miner: "", Nonce: 0, Transactions: tx})
	return Blockchain
}

func GetLatestBlock() block.Block {
	return Blockchain[len(Blockchain)-1]
}

func CreateBlock() block.Block {
	var chash string
	txs, _ := json.Marshal(transactions.Pool)
	fmt.Println(string(txs))
	h := sha256.New()
	//from google
	h.Write([]byte(string(txs)))
	chash = fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(chash)
	block := block.Block{
		BlockNumber:          (GetLatestBlock().BlockNumber) + 1,
		ParentHash:           GetLatestBlock().ParentHash,
		CurrHash:             chash,
		Timestamp:            time.Now(),
		NumberOfTransactions: len(transactions.Pool),
		Miner:                "utkarsh",
		Nonce:                10,
		Transactions:         transactions.Pool,
	}
	Blockchain = append(Blockchain, block)
	transactions.Pool = nil
	return block
}
