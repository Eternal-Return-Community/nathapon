package database

import (
	"context"
	"fmt"
	"nathapon/src/models"
	"net"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func Join(conn net.Conn, channel models.Irc) error {

	if channel.ChannelName != "centraldonathpon" {
		return nil
	}

	checkUser, err := collection.CountDocuments(context.Background(), bson.M{"name": strings.ToLower(channel.MessageAuthor)})
	if err != nil {
		return err
	}

	if checkUser > 0 {
		message := fmt.Sprintf("-> @%s Central do Nathapon já está no seu canal.", channel.MessageAuthor)
		fmt.Fprintf(conn, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
		return err
	}

	_, err = collection.InsertOne(context.Background(), models.Join{Name: strings.ToLower(channel.MessageAuthor), ID: channel.MessageAuthorID})
	if err != nil {
		panic(err)
	}

	conn.Write([]byte(fmt.Sprintf("JOIN #%s\r\n", channel.MessageAuthor)))
	message := fmt.Sprintf("-> @%s Central do Nathapon foi adicionado com sucesso ao seu canal! Não esqueça de adicionar o Bot como moderador no seu canal /mod centraldonathpon.", channel.MessageAuthor)
	fmt.Fprintf(conn, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
	return nil
}
