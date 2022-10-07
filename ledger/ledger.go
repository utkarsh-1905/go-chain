package ledger

import (
	"context"
	"fmt"

	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
	"go.mongodb.org/mongo-driver/bson"
)

var Ledger = make(map[string]int)

func Initialize(pubKey string) {
	if Ledger[pubKey] == 0 {
		mongo := helpers.ConnectWithDatabase()
		acc := mongo.Database("go-chain").Collection("ledger")
		Ledger[pubKey] = 0
		ctx := context.TODO()
		acc.InsertOne(ctx, bson.M{"pubkey": pubKey, "balance": 0})
	}
}

func Transfer(from string, to string, value int) {
	balance, fr := GetBalance(from)
	if balance >= value && fr == from {
		Initialize(to)
		Increment(to, value)
		Decrement(from, value)
	} else {
		fmt.Println("Insufficient funds")
		return
	}
}

func Increment(pubKey string, value int) {
	mongo := helpers.ConnectWithDatabase()
	acc := mongo.Database("go-chain").Collection("ledger")
	ctx := context.TODO()
	Ledger[pubKey] += value
	acc.UpdateOne(ctx, bson.M{"pubkey": pubKey}, bson.M{"$inc": bson.M{"balance": value}})
}

func Decrement(pubKey string, value int) {
	Increment(pubKey, -value)
}

func GetBalance(pubKey string) (int, string) {
	mongo := helpers.ConnectWithDatabase()
	acc := mongo.Database("go-chain").Collection("ledger")
	ctx := context.TODO()
	var res struct {
		pubKey  string `bson:"pubkey"`
		balance int    `bson:"balance"`
	}
	err := acc.FindOne(ctx, bson.M{"pubkey": pubKey}).Decode(&res)
	if err != nil {
		panic(err)
	} else {
		return res.balance, res.pubKey
	}
}

func Update(tx transactions.Transaction) {
	from := tx.From
	to := tx.To
	amount := tx.Value
	Transfer(from, to, amount)
}
