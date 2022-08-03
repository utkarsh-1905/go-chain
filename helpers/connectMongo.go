package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectWithDatabase() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Mongo, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	HandleErr(err)
	defer func() {
		if err = Mongo.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return Mongo
}
