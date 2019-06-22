package monzgo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const BASE_URL string = "https://api.monzo.com"

type Monzgo struct {
	APIKey string
	http.Client
}

func Setup(apiKey string) *Monzgo {
	if apiKey == "" {
		log.Fatal("No Monzo API key provided")
	}

	return &Monzgo{
		APIKey: apiKey,
	}
}

func (m *Monzgo) Verify() error {
	request, err := m.CreateRequest(http.MethodGet, "ping/whoami", nil)
	if err != nil {
		return err
	}

	response, _ := m.Do(request)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Failure")
	}

	return nil
}

func (m *Monzgo) CreateRequest(method string, path string, requestBody io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, BASE_URL+"/"+path, requestBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+m.APIKey)
	return request, nil
}

func (m *Monzgo) Get(path string) (*http.Request, error) {
	request, err := m.CreateRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func ParseResponse(data []byte, key string, t interface{}) error {
	var keyMap map[string]*json.RawMessage
	if err := json.Unmarshal(data, &keyMap); err != nil {
		return err
	}

	if err := json.Unmarshal(*keyMap[key], t); err != nil {
		return err
	}

	return nil
}
