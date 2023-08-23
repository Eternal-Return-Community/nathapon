package services

import (
	"fmt"
	"nathapon/src/database"
	"nathapon/src/utils"
	"net"
	"os"
	"strings"
)

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	channels, err := database.ChannelList()
	if err != nil {
		fmt.Println(err)
	}

	conn.Write([]byte(fmt.Sprintf("PASS %s\r\n", utils.Env.Auth)))
	conn.Write([]byte("NICK CentralDoNathapon\r\n"))
	conn.Write([]byte(channels))
	conn.Write([]byte("CAP REQ :twitch.tv/tags twitch.tv/commands twitch.tv/membership\r\n"))

	return conn
}

// Test
func joinChannels() string {
	channels := []string{"nicaa", "anniesemtrema", "labnica"}
	return fmt.Sprintf("JOIN #%s\r\n", strings.Join(channels, ",#"))
}
