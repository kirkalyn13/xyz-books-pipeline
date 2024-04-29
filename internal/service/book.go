package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/kirkalyn13/xyz-books-pipeline/pkg/model"
)

const (
	url = "https://jsonplaceholder.typicode.com/todos/1"
)

// Get sends a GET request to the url
func Get() (model.Response, error) {
	var response model.Response

	resp, err := http.Get(url)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Put sends a PUT request to the url
func Put(data string) error {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Println("PUT:", string(body))
	return nil
}
