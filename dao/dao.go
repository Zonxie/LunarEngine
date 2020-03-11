package dao

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao struct {
	Database string
}

type Client struct {
	client     *mongo.Client
	context    context.Context
	db         *mongo.Database
	collection *mongo.Collection
}

func NewDao() (*Dao, error) {
	return &Dao{
		Database: "mongodb+srv://gilbert:<databaseSecret>@lunarassignmentcluster0-5kxg3.mongodb.net/test?retryWrites=true&w=majority",
	}, nil
}

func (dao *Dao) client(collection string) (*Client, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://gilbertus:kode12345@lunarassignmentcluster0-5kxg3.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)

	lunarDatabase := client.Database("bank")
	resourceCollection := lunarDatabase.Collection("accounts")

	return &Client{
		client:     client,
		context:    ctx,
		collection: resourceCollection,
		db:         lunarDatabase,
	}, nil
}

func (c *Client) Disconnect(ctx context.Context) {
	c.client.Disconnect(ctx)
}
