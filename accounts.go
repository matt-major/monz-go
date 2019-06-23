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

type Balance struct {
	Balance     int
	Total       int `json:"total_balance"`
	WithSavings int `json:"balance_including_flexible_savings"`
	Currency    string
}
