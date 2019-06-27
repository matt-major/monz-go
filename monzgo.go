package monzgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func (m *Monzgo) request(requestMethod string, path string, responseHolder interface{}, requestData map[string]string) error {
	var request *http.Request
	var err error

	if requestMethod == "GET" {
		request, err = m.createRequestWithQuery(path, requestData)
	} else {
		request, err = m.createRequestWithFormData(requestMethod, path, requestData)
	}

	if err != nil {
		return err
	}

	return m.performRequest(request, responseHolder)
}

func (m *Monzgo) createRequestWithQuery(path string, queryParams map[string]string) (*http.Request, error) {
	request, err := http.NewRequest("GET", m.BaseURL+path, nil)
	if err != nil {
		return nil, err
	}

	if len(queryParams) > 0 {
		q := request.URL.Query()
		for key, val := range queryParams {
			q.Add(key, val)
		}
		request.URL.RawQuery = q.Encode()
	}

	return request, err
}

func (m *Monzgo) createRequestWithFormData(requestMethod string, path string, parameters map[string]string) (*http.Request, error) {
	formData := url.Values{}
	for key, val := range parameters {
		formData.Set(key, val)
	}

	request, err := http.NewRequest(requestMethod, m.BaseURL+path, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}

func (m *Monzgo) performRequest(request *http.Request, responseHolder interface{}) error {
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

	if err := json.Unmarshal(bytes, responseHolder); err != nil {
		return err
	}

	return nil
}
