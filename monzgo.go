package monzgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Monzgo client
type Monzgo struct {
	APIKey  string
	BaseURL string
}

// MonzoError - API error definition
type MonzoError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e *MonzoError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (m *Monzgo) request(requestMethod string, path string, response interface{}, queryParams map[string]string) error {
	var request *http.Request
	var err error

	request, err = http.NewRequest(requestMethod, m.BaseURL+path, nil)
	if err != nil {
		return err
	}

	if len(queryParams) > 0 {
		q := request.URL.Query()
		for k, v := range queryParams {
			q.Add(k, v)
		}
		request.URL.RawQuery = q.Encode()
	}

	return m.makeHTTPrequest(request, response)
}

func (m *Monzgo) makeHTTPrequest(request *http.Request, response interface{}) error {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.APIKey))

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		monzoErr := &MonzoError{
			StatusCode: res.StatusCode,
		}

		if err := json.Unmarshal(bytes, monzoErr); err != nil {
			return fmt.Errorf("Error during JSON umarshalling %d error: %v", res.StatusCode, err)
		}

		return monzoErr
	}

	if err := json.Unmarshal(bytes, response); err != nil {
		return err
	}

	return nil
}
