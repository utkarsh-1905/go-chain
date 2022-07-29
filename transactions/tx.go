package transactions

type Transaction struct {
	To    string `json:"to"`
	From  string `json:"from"`
	Value int    `json:"value"`
	Data  string `json:"data"`
}

var Pool = make([]Transaction, 0)

//create transaction or transfer tokens, and send it to Pool
//If transaction mined, the wallet file will show the updated balance

func CreateTransaction(to string, from string, value int, data string) Transaction {
	tx := Transaction{to, from, value, data}
	Pool = append(Pool, tx)
	return tx
}
