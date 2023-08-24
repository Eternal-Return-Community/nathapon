package models

type ClipInfo struct {
	Data []Info `json:"data"`
}

type Info struct {
	URL             string      `json:"url"`
	BroadcasterName string      `json:"broadcaster_name"`
	CreatorName     string      `json:"creator_name"`
	GameID          string      `json:"game_id"`
	Title           string      `json:"title"`
	Duration        interface{} `json:"duration"`
	Thumbnail_url   string      `json:"thumbnail_url"`
}
