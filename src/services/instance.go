package services

import (
	"fmt"
	"io"
	"nathapon/src/utils"
	"net/http"
)

func Instance(method string, url string, payload io.Reader) (io.ReadCloser, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", utils.Env.Token)
	req.Header.Add("Client-ID", utils.Env.Client)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	return resp.Body, nil

}
