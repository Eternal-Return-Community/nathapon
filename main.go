package main

import (
	"bufio"
	"fmt"
	"nathapon/src/database"
	"nathapon/src/models"
	client "nathapon/src/services/twitch"
	"nathapon/src/utils"
	"net"
	"strings"
	"time"
)

var (
	channel models.Irc
	send net.Conn
)

func main() {
	// Env
	utils.Load()
	database.Connect()

	for {
		conn := reconnect()
		send = conn
		send.(*net.TCPConn).SetNoDelay(true)

		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				conn.Close()
				break
			}

			messageIRC(message)
			pong(message)
		}
		fmt.Printf("%s Nathapon foi reconectado com sucesso!\n", utils.Date())
	}
}

func pong(message string) {

	if strings.HasPrefix(message, "PING") {
		pong := strings.Replace(message, "PING", "PONG", 1)
		send.Write([]byte(fmt.Sprintf("%s\r\n", pong)))
		fmt.Printf("%s Nathapon pingou o servidor da Twitch.\n", utils.Date())
	}
	fmt.Println(message)
}

func reconnect() net.Conn {
	for {
		conn := client.Connect()
		if conn != nil {
			fmt.Printf("%s Conexão estabelecida com sucesso.\n", utils.Date())
			return conn
		}
		fmt.Printf("%s Falha na conexão. Tentando novamente em 5 segundos...\n", utils.Date())
		time.Sleep(5 * time.Second)
	}
}

func messageIRC(message string) {
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
		commands(result)
	}

}

func commands(channel models.Irc) {

	message := strings.TrimSpace(strings.ToLower(channel.Message))

	switch message {
	case "!clip":
		client.Clip(send, channel)
		break
	case "!join":
		database.Join(send, channel)
		break
	case "!part":
		database.Part(send, channel)
		break
	case "!ping":
		fmt.Fprintf(send, "PRIVMSG #%s :%s\r\n", channel.ChannelName, "🏓 Pong!")
		break
	}
}
