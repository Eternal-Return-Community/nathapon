package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"nathapon/src/models"
	"nathapon/src/services/discord"
	"nathapon/src/utils"
	"net"
	"net/http"
	"time"
)

const endpoint = "https://api.twitch.tv/helix"

var send net.Conn

func Clip(conn net.Conn, channel models.Irc) {

	send = conn
	client := &http.Client{}
	clipID, err := createClip(client, channel)
	if err != nil {
		fmt.Println(err)
		return
	}

	info, err := getClip(client, clipID, channel)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = discord.Webhook(client, info, channel.MessageAuthor)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func createClip(client *http.Client, channel models.Irc) (string, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", endpoint+"/clips?broadcaster_id=", channel.ChannelRoom), nil)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

	req.Header.Add("Authorization", utils.Env.Token)
	req.Header.Add("Client-ID", utils.Env.Client)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	defer resp.Body.Close()

	var response models.Clip
	err = json.NewDecoder(resp.Body).Decode(&response)
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

func getClip(client *http.Client, clipID string, channel models.Irc) (models.ClipInfo, error) {

	time.Sleep(1 * time.Second)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", endpoint+"/clips?id=", clipID), nil)
	if err != nil {
		return models.ClipInfo{}, err
	}

	req.Header.Add("Authorization", utils.Env.Token)
	req.Header.Add("Client-ID", utils.Env.Client)

	resp, err := client.Do(req)
	if err != nil {
		return models.ClipInfo{}, err
	}
	defer resp.Body.Close()

	var response models.ClipInfo
	err = json.NewDecoder(resp.Body).Decode(&response)
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
