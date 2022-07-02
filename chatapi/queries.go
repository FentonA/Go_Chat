package chat

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)


client, err := mongo.Connect(ctx, options.Client().ApplyUri("mongodb+srv://gochat:<password>@cluster0.usyps.mongodb.net/?retryWrites=true&w=majority"))
if != nil{
	log.Fatal(err)
}
ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
err = client.Connect(ctx)
if err != nil{
	log.Fatal(err)
}
defer.client.Disconnect(ctx)

chatDatabase := client.Database("gochat")