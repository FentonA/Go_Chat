package chat

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx = context.TODO()
)

type Messages struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Sender   string
	SentTime time.Time
	Text     string
}

type Database struct {
	mongo *mongo.Database
}

func NewChat() (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gochat:gochat@cluster0.usyps.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{mongo: client.Database("Chat")}, nil
}

func (d *Database) Item(ID primitive.ObjectID) (*Messages, error) {
	var msg *Messages

	chat := d.mongo.Collection("Chat").FindOne(Ctx, ID)
	err := chat.Decode(&msg)
	if err != nil {
		return nil, fmt.Errorf("Error found when finding chat by ID, %v", err)
	}

	return msg, nil
}
func (msg *Messages) ID() primitive.ObjectID {
	return msg.Id
}

func (d *Database) All() ([]*Messages, error) {
	var msg *Messages

	chats := d.mongo.Collection("Chat")
	cursor, err := chats.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		fmt.Printf("Error finding all messages: %v", err)
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

func (d *Database) Create(msg *Messages) (*Messages, error) {

	chats := d.mongo.Collection("chats")
	message, err := chats.InsertOne(Ctx, msg)
	if err != nil {
		return nil, fmt.Errorf("Error creating a messages %v", err)
	}
	//This fmt.Printf is for debug purposes
	fmt.Printf("The result of insertedId %v : ", message.InsertedID)
	return msg, err
}

func (d *Database) Update(msg *Messages) error {
	chats := d.mongo.Collection("chats")
	filter := bson.D{{"_id", msg.Id}}
	update := bson.D{{"Text", msg}}
	_, err := chats.UpdateOne(
		Ctx,
		filter,
		update,
	)
	return err
}

func (d *Database) Delete(ID primitive.ObjectID) error {
	chats := d.mongo.Collection("chats")
	_, err := chats.DeleteOne(Ctx, bson.D{{"_id", ID}})
	if err != nil {
		return err
	}
	return nil
}
