package monzgo

// Owner - the detail of the owner of the Account
type Owner struct {
	UserID             string `json:"user_id"`
	PreferredName      string `json:"preferred_name"`
	PreferredFirstName string `json:"preferred_first_name"`
}

// PaymentDetailsLocale - details for the account in a given locale
type PaymentDetailsLocale struct {
	AccountNumber string `json:"account_number"`
	SortCode      string `json:"sort_code"`
}

// Account - top level Account definition
type Account struct {
	ID             string
	Closed         bool
	Created        string
	Description    string
	Type           string
	Currency       string
	CountryCode    string `json:"country_code"`
	Owners         []Owner
	AccountNumber  string `json:"account_number"`
	SortCode       string `json:"sort_code"`
	PaymentDetails map[string]PaymentDetailsLocale
}

// Accounts - get all accounts with an optional type filter
func (m *Monzgo) Accounts(typeFilter string) ([]*Account, error) {
	requestParams := map[string]string{}
	if typeFilter != "" {
		requestParams["account_type"] = typeFilter
	}

	rspHolder := &struct {
		Accounts []*Account `json:"accounts"`
	}{}

	if err := m.request("GET", "accounts", rspHolder, requestParams); err != nil {
		return nil, err
	}

	return rspHolder.Accounts, nil
}
