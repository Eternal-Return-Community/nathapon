package database

import (
	"context"
	"fmt"
	"log"
	"nathapon/src/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func Connect() (client *mongo.Client, ctx context.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.Env.Database))
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("channels").Collection("info")

	fmt.Printf("%s Database online.\n", utils.Date())
	return
}
