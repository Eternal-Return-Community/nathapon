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

const (
	endpoint = "https://api.twitch.tv/helix"
)

func Clip(conn net.Conn, channel models.Irc) {

	client := &http.Client{}
	clipID, err := createClip(client, conn, channel)
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

func createClip(client *http.Client, conn net.Conn, channel models.Irc) (string, error) {

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
		fmt.Fprintf(conn, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)
		return "", errors.New("")
	}

	message := fmt.Sprintf("-> @%s Clipe criado com sucesso! https://clips.twitch.tv/%s", channel.MessageAuthor, response.Data[0].Id)
	fmt.Fprintf(conn, "PRIVMSG #%s :%s\r\n", channel.ChannelName, message)

	return response.Data[0].Id, nil

}

func getClip(client *http.Client, clipID string, channel models.Irc) (models.ClipInfo, error) {

	time.Sleep(5 * time.Second)
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

	return response, nil
}
