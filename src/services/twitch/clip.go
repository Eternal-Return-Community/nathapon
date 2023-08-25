package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"nathapon/src/models"
	"nathapon/src/services"
	"nathapon/src/services/discord"
	"net"
	"time"
)

const endpoint = "https://api.twitch.tv/helix"

var send net.Conn

func Clip(conn net.Conn, channel models.Irc) {

	send = conn
	clipID, err := createClip(channel)
	if err != nil {
		fmt.Println(err)
		return
	}

	info, err := getClip(clipID, channel)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = discord.Webhook(info, channel.MessageAuthor)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func createClip(channel models.Irc) (string, error) {

	resp, err := services.Instance("POST", fmt.Sprintf("%s%s", endpoint+"/clips?broadcaster_id=", channel.ChannelRoom), nil)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer resp.Close()

	var response models.Clip
	err = json.NewDecoder(resp).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	if response.Status == 404 {
		message := fmt.Sprintf("-> @%s Esse canal está offline.", channel.MessageAuthor)
		fmt.Fprintf(send, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
		return "", errors.New("")
	}

	message := fmt.Sprintf("-> @%s Clipe criado com sucesso! https://clips.twitch.tv/%s", channel.MessageAuthor, response.Data[0].Id)
	fmt.Fprintf(send, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)

	return response.Data[0].Id, nil

}

func getClip(clipID string, channel models.Irc) (models.ClipInfo, error) {

	time.Sleep(20 * time.Second)
	resp, err := services.Instance("GET", fmt.Sprintf("%s%s", endpoint+"/clips?id=", clipID), nil)
	if err != nil {
		return models.ClipInfo{}, err
	}
	defer resp.Close()

	var response models.ClipInfo
	err = json.NewDecoder(resp).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return models.ClipInfo{}, nil
	}

	if len(response.Data) == 0 {
		message := fmt.Sprintf("-> @%s Ocorreu um erro durante a criação do clipe. Tente novamente.", channel.MessageAuthor)
		fmt.Fprintf(send, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
		return models.ClipInfo{}, errors.New("")
	}

	return response, nil
}
