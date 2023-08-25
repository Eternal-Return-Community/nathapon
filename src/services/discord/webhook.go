package discord

import (
	"fmt"
	"nathapon/src/models"
	"nathapon/src/services"
	"nathapon/src/utils"
)

func Webhook(channel models.ClipInfo, clipAuthor string) error {

	resp, err := services.Instance("POST", utils.Env.Webhook, embed(channel, clipAuthor))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Close()

	return nil

}
