package database

import (
	"context"
	"fmt"
	"nathapon/src/models"
	"net"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func Part(conn net.Conn, channel models.Irc) error {

	if channel.ChannelName != "centraldonathpon" {
		return nil
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"name": strings.ToLower(channel.MessageAuthor)})
	if err != nil {
		panic(err)
	}

	if result.DeletedCount == 0 {
		message := fmt.Sprintf("-> @%s Central do Nathapon não está no seu canal. Caso queira adicionar digite !join.", channel.MessageAuthor)
		fmt.Fprintf(conn, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
		return nil
	}

	conn.Write([]byte(fmt.Sprintf("PART #%s\r\n", channel.MessageAuthor)))
	message := fmt.Sprintf("-> @%s Central do Nathapon foi removido do seu canal.", channel.MessageAuthor)
	fmt.Fprintf(conn, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
	return nil
}
