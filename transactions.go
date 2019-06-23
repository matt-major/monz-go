package monzgo

import "strconv"

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

// Transactions - retrieve paginated results of Transactions
func (m *Monzgo) Transactions(accountID string, limit int, since string) ([]*TransactionListEntry, error) {
	requestParams := map[string]string{
		"account_id": accountID,
		"limit":      strconv.Itoa(limit),
		"since":      since,
	}

	rspHolder := &struct {
		Transactions []*TransactionListEntry `json:"transactions"`
	}{}

	if err := m.request("GET", "transactions", rspHolder, requestParams); err != nil {
		return nil, err
	}

	return rspHolder.Transactions, nil
}

// TransactionDetails - get the full details for a single transaction
func (m *Monzgo) TransactionDetails(transactionID string) (*TransactionDetails, error) {
	requestParams := map[string]string{
		"expand[]": "merchant",
	}

	rspHolder := &struct {
		Transaction *TransactionDetails `json:"transaction"`
	}{}

	if err := m.request("GET", "transactions/"+transactionID, rspHolder, requestParams); err != nil {
		return nil, err
	}

	return rspHolder.Transaction, nil
}
