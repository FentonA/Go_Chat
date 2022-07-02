package chat

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx = context.TODO()
)

const dbPath = ""

type Messages struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Sender   string
	SentTime time.Time
	Text     string
}

func InitiateMongoClient() *mongo.Database {
	var client *mongo.Client
	uri := "mongodb+srv://gochat:gochat@cluster0.usyps.mongodb.net/?retryWrites=true&w=majority"
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err := mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
		fmt.Println(client)
	}
	return client.Database("gochat")
}
func (msg *Messages) ID() primitive.ObjectID {
	return msg.Id
}

func NewChat() (mongo.Collection, error) {
	conn := InitiateMongoClient()
	err := conn.CreateCollection(Ctx, "Chat")
	// if collection != nil {
	// 	fmt.Println("This is the result from creating the 'Chat' collection", collection)
	// }

	return collection, nil
}

func GetMessages() ([]*Messages, error) {
	var msg *Messages

	conn := InitiateMongoClient()
	chats := conn.Collection("chats")
	cursor, err := chats.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		fmt.Errorf("Error finding all messages: %v", err)
	}

	messages := make([]*Messages, 0)
	for cursor.Next(Ctx) {
		err := cursor.Decode(&msg)
		if err != nil {
			return nil, fmt.Errorf("Error that shows up when using 'Decode' : %v", err)
		}
		messages = append(messages, msg)
	}
	return messages, nil

}

func Create(msg *Messages) (*Messages, error) {
	conn := InitiateMongoClient()
	chats := conn.Collection("chats")
	message, err := chats.InsertOne(Ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("Error creating a messages")
	}
	fmt.Printf("The result of insertedId %v : ", message.InsertedID)
	return msg, err
}
func Update(msg *Messages, newMessage string) error {
	conn := InitiateMongoClient()
	chats := conn.Collection("chats")
	filter := bson.D{{"_id", msg.Id}}
	update := bson.D{{"Text", newMessage}}
	_, err := chats.UpdateOne(
		Ctx,
		filter,
		update,
	)
	return err
}

func Delete(msg *Messages) error {
	conn := InitiateMongoClient()
	chats := conn.Collection("chats")
	_, err := chats.DeleteOne(Ctx, bson.D{{"_id", msg.Id}})
	if err != nil {
		return err
	}
	return nil
}
