package helper

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(method, url string, body io.Reader) (*http.Response, error) {
	fmt.Println("Inside makeRequest:")
	fmt.Println("method:", method)
	fmt.Println("url:", url)
	fmt.Println("body:", body)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
