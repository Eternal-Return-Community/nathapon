package discord

import (
	"fmt"
	"nathapon/src/models"
	"nathapon/src/utils"
	"net/http"
)

func Webhook(client *http.Client, channel models.ClipInfo, clipAuthor string) error {

	req, err := http.NewRequest("POST", utils.Env.Webhook, embed(channel, clipAuthor))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	return nil

}
