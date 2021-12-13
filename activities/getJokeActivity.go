package activities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type apiresp struct {
	Result []Joke `json:"result"`
}

type Joke struct {
	Value      string   `json:"value"`
	Categories []string `json:"categories"`
}

func GetJoke(searchval string) ([]Joke, error) {

	resp, err := http.Get("https://api.chucknorris.io/jokes/search?query=" + searchval)

	var response apiresp
	if err == nil {

		body, readErr := ioutil.ReadAll(resp.Body)

		if readErr == nil {

			json.Unmarshal(body, &response)

		}

	}

	return response.Result, nil

}
