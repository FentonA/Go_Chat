package main

import (
	"context"
	"fmt"
	chat "go-chat/chatapi"
	"log"
	"net/http"

	"github.com/m1gwings/livegollection"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	coll, err := chat.NewChat()
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating new chat collection: %v", err))
	}
	liveGoll := livegollection.NewLiveGollection[primitive.ObjectID, *chat.Messages](context.TODO(), coll, log.Default())

	http.HandleFunc("/livegollection", liveGoll.Join)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
