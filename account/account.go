package account

import (
	"context"

	"github.com/utkarsh-1905/go-chain/helpers"
	"github.com/utkarsh-1905/go-chain/transactions"
	"go.mongodb.org/mongo-driver/bson"
)

var Accounts = make(map[string]int)

func Initialize(pubKey string) {
	if Accounts[pubKey] == 0 {
		mongo := helpers.ConnectWithDatabase()
		acc := mongo.Database("go-chain").Collection("accounts")
		Accounts[pubKey] = 0
		ctx := context.TODO()
		acc.InsertOne(ctx, bson.M{"pubkey": pubKey, "balance": 0})
	}
}

func Transfer(from string, to string, value int) {
	Initialize(from)
	Initialize(to)
	Increment(to, value)
	Decrement(from, value)
}

func Increment(pubKey string, value int) {
	mongo := helpers.ConnectWithDatabase()
	acc := mongo.Database("go-chain").Collection("accounts")
	ctx := context.TODO()
	Accounts[pubKey] += value
	acc.UpdateOne(ctx, bson.M{"pubkey": pubKey}, bson.M{"$inc": bson.M{"balance": value}})
}

func Decrement(pubKey string, value int) {
	Increment(pubKey, -value)
}

func GetBalance(pubKey string) int {
	Initialize(pubKey)
	mongo := helpers.ConnectWithDatabase()
	acc := mongo.Database("go-chain").Collection("accounts")
	ctx := context.TODO()
	var res struct {
		pubKey  string `bson:"pubkey"`
		balance int    `bson:"balance"`
	}
	_ = acc.FindOne(ctx, bson.M{"pubkey": pubKey}).Decode(&res)
	return res.balance
}

func Update(tx transactions.Transaction) {
	from := tx.From
	to := tx.To
	amount := tx.Value
	Transfer(from, to, amount)
}
