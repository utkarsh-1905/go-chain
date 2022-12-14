package block

import (
	"github.com/utkarsh-1905/go-chain/transactions"
)

type Block struct {
	BlockNumber          int                        `json:"blockNumber"`
	ParentHash           string                     `json:"parentHash"`
	CurrHash             string                     `json:"currHash"`
	Timestamp            int64                      `json:"timestamp"`
	NumberOfTransactions int                        `json:"numberOfTransactions"`
	Miner                string                     `json:"miner"`
	Nonce                int                        `json:"nonce"`
	Transactions         []transactions.Transaction `json:"transactions"`
}
