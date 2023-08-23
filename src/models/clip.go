package models

type Clip struct {
	Error   string `json:"error"`
	Status  int16  `json:"status"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}

type Data struct {
	Id       string `json:"id"`
	Edit_url string `json:"edit_url"`
}

