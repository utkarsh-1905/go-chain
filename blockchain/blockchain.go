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
	tx = append(tx, transactions.Transaction{To: "", From: "", Value: 1000, Data: ""})
	var blockchain = make([]block.Block, 0)
	blockchain = append(blockchain, block.Block{BlockNumber: 1, ParentHash: "", CurrHash: "", Timestamp: time.Now().Unix(), NumberOfTransactions: 0, Miner: "", Nonce: 0, Transactions: tx})
	bchain, _ := json.MarshalIndent(blockchain, "", "\t")
	_ = ioutil.WriteFile("blockchain.json", bchain, 0644)
	return blockchain
}

func GetLatestBlock() block.Block {
	blockchain := helpers.ReadAndUnmarshallBlockchain()
	return blockchain[len(blockchain)-1]
}

func CreateBlock(mnr string) block.Block {
	blockchain := helpers.ReadAndUnmarshallBlockchain()
	var chash string
	block := block.Block{
		BlockNumber:          (GetLatestBlock().BlockNumber) + 1,
		ParentHash:           GetLatestBlock().CurrHash,
		Timestamp:            time.Now().Unix(),
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
	blockchain = append(blockchain, block)
	transactions.Pool = nil
	bchain, _ := json.MarshalIndent(blockchain, "", "\t")
	_ = ioutil.WriteFile("blockchain.json", bchain, 0644)
	return block
}
