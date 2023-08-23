package main

import (
	"bufio"
	"fmt"
	"nathapon/src/database"
	"nathapon/src/models"
	client "nathapon/src/services/twitch"
	"nathapon/src/utils"
	"net"
	"os"
	"strings"
)

var channel models.Irc

func main() {

	//Env
	utils.Load()
	database.Connect()

	conn := client.Connect()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		message = strings.TrimSuffix(message, "\r\n")
		if strings.Contains(message, "PRIVMSG") {
			parts := strings.Split(message, " ")
			room := strings.Split(message, ";")
			result := models.Irc{
				ChannelRoom:     strings.Replace(room[len(room)-6], "room-id=", "", 1),
				ChannelName:     strings.Replace(parts[3], "#", "", 1),
				MessageAuthor:   strings.Replace(room[4], "display-name=", "", 1),
				MessageAuthorID: strings.Replace(room[len(room)-2], "user-id=", "", 1),
				Message:         strings.Replace(parts[4], ":", "", 1),
			}
			messageContent(conn, result)
		}
	}

}

/*
Falta criar um sistema de handler.
Fiquei com preguiça, então vai ficar assim por um tempo.
*/
func messageContent(conn net.Conn, channel models.Irc) {

	if len(channel.ChannelName) == 0 {
		return
	}

	if strings.ToLower(channel.Message) == "!clip" {
		client.Clip(conn, channel)
	}

	if strings.ToLower(channel.Message) == "!join" {
		database.Join(conn, channel)
	}

	if strings.ToLower(channel.Message) == "!part" {
		database.Part(conn, channel)
	}
}
