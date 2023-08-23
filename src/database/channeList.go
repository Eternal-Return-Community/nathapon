package database

import (
	"context"
	"fmt"
	"nathapon/src/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ChannelList() (string, error) {

	var channels []string
	channel, err := collection.Find(context.Background(), bson.M{}, options.Find())
	if err != nil {
		return "", nil
	}

	defer channel.Close(context.Background())

	for channel.Next(context.Background()) {
		var result models.Join
		err := channel.Decode(&result)
		if err != nil {
			return "", nil
		}

		channels = append(channels, result.Name)
	}

	return fmt.Sprintf("JOIN #%s\r\n", strings.Join(channels, ",#")), err

}
