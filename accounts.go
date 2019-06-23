package monzgo

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
	CountryCode    string `json:"country_code"`
	Owners         []Owner
	AccountNumber  string `json:"account_number"`
	SortCode       string `json:"sort_code"`
	PaymentDetails map[string]PaymentDetailsLocale
}

func (m *Monzgo) Accounts() ([]*Account, error) {
	rspHolder := &struct {
		Accounts []*Account `json:"accounts"`
	}{}

	if err := m.request("GET", "accounts", rspHolder); err != nil {
		return nil, err
	}

	return rspHolder.Accounts, nil
}
