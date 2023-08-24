package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"nathapon/src/models"
)

func embed(info models.ClipInfo, clipAuthor string) *bytes.Buffer {

	channel := info.Data[0]

	var duration int
	if d, ok := channel.Duration.(float64); ok {
		duration = int(d)
	} else {
		duration = int(d)
	}

	payload, _ := json.Marshal(map[string]string{
		"username":   "Nathapon",
		"avatar_url": "https://i.imgur.com/JEUHO92.png",
		"content": fmt.Sprintf(
			"Clipe do canal: **[%s](https://twitch.tv/%s)**\n"+
				"Clipe criado por: **%s**\n"+
				"Título: **%s** | Duração: **%d segundos**\n"+
				"**%s**",
			channel.BroadcasterName, channel.BroadcasterName,
			clipAuthor, channel.Title,
			duration, channel.URL),
	})

	return bytes.NewBuffer(payload)

}
