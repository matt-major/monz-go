package monzgo

import (
	"bytes"
	"fmt"
	"net/http"
)

type Owner struct {
	UserID             string `json:"user_id"`
	PreferredName      string `json:"preferred_name"`
	PreferredFirstName string `json:"preferred_first_name"`
}

type PaymentDetailsLocale struct {
	AccountNumber string `json:"account_number"`
	SortCode      string `json:"sort_code"`
}

type Account struct {
	ID             string
	Closed         bool
	Created        string
	Description    string
	Type           string
	Currency       string
	CountryCode    string `json:"country_code`
	Owners         []Owner
	AccountNumber  string `json:"account_number"`
	SortCode       string `json:"sort_code"`
	PaymentDetails map[string]PaymentDetailsLocale
}

func (m *Monzgo) Accounts() ([]Account, error) {
	request, err := m.Get("accounts")
	if err != nil {
		return nil, err
	}

	response, _ := m.Do(request)
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(response.Body)
	responseString := buffer.String()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to retrieve Accounts: %s", responseString)
	}

	responseBytes := buffer.Bytes()
	var accounts []Account
	if err := ParseResponse(responseBytes, "accounts", &accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
