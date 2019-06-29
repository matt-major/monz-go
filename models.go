package monzgo

// PingResult - the result of the Ping / Verify request
type PingResult struct {
	Authentication bool
	ClientID       string `json:"client_id"`
	UserID         string `json:"user_id"`
}

// TransactionListEntry definition of a Transaction list entry
type TransactionListEntry struct {
	AccountBalance int64             `json:"account_balance"`
	Amount         int64             `json:"amount"`
	Created        string            `json:"created"`
	Currency       string            `json:"currency"`
	Description    string            `json:"description"`
	ID             string            `json:"id"`
	Metadata       map[string]string `json:"metadata"`
	Notes          string            `json:"notes"`
	IsLoad         bool              `json:"is_load"`
	Settled        string            `json:"settled"`
	Category       string            `json:"category"`
}

// TransactionMerchantAddress definition of a Merchant Address
type TransactionMerchantAddress struct {
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Postcode  string  `json:"postcode"`
	Region    string  `json:"region"`
}

// TransactionMerchant definition of a merchant
type TransactionMerchant struct {
	Address  TransactionMerchantAddress `json:"address"`
	Created  string                     `json:"created"`
	GroupID  string                     `json:"group_id"`
	ID       string                     `json:"id"`
	Logo     string                     `json:"logo"`
	Emoji    string                     `json:"emoji"`
	Name     string                     `json:"name"`
	Category string                     `json:"category"`
}

// TransactionDetails definition of an expanded transaction
type TransactionDetails struct {
	AccountBalance int64               `json:"account_balance"`
	Amount         int64               `json:"amount"`
	Created        string              `json:"created"`
	Currency       string              `json:"currency"`
	Description    string              `json:"description"`
	ID             string              `json:"id"`
	Metadata       map[string]string   `json:"metadata"`
	Notes          string              `json:"notes"`
	IsLoad         bool                `json:"is_load"`
	Settled        string              `json:"settled"`
	Merchant       TransactionMerchant `json:"merchant"`
}

// Pot definition
type Pot struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Style    string `json:"style"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Deleted  bool   `json:"deleted"`
}

// Balance response definition
type Balance struct {
	Balance      int64  `json:"balance"`
	TotalBalance int64  `json:"total_balance"`
	Currency     string `json:"currency"`
	SpendToday   int64  `json:"spend_today"`
}

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
	ID             string                          `json:"id"`
	Closed         bool                            `json:"closed"`
	Created        string                          `json:"created"`
	Description    string                          `json:"description"`
	Type           string                          `json:"types"`
	Currency       string                          `json:"currency"`
	CountryCode    string                          `json:"country_code"`
	Owners         []Owner                         `json:"owners"`
	AccountNumber  string                          `json:"account_number"`
	SortCode       string                          `json:"sort_code"`
	PaymentDetails map[string]PaymentDetailsLocale `json:"paymentdetails"`
}

// FeedItemOptions - Feed Item request options definition
type FeedItemOptions struct {
	Type            string
	URL             string
	Title           string
	ImageURL        string
	BackgroundColor string
	BodyColor       string
	TitleColor      string
	Body            string
}
