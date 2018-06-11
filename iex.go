package iex

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	baseEndpoint = "https://api.iextrading.com/1.0"
)

// IEX struct
type IEX struct {
	client *http.Client
}

// NewIEX returns a new IEX struct
func NewIEX(client *http.Client) *IEX {
	i := &IEX{
		client: client,
	}
	return i
}

func (i *IEX) getJSON(url string, result interface{}) error {
	if url == "" {
		return errors.New("url needs to be defined")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to retrive data from URL: " + url + "with a status code " + resp.Status)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &result)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEX) endpoint(route string) string {
	return baseEndpoint + route
}
