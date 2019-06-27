package monzgo

import "strconv"

// Transactions - retrieve paginated results of Transactions
func (m *Monzgo) Transactions(accountID string, limit int, since string) ([]*TransactionListEntry, error) {
	requestParams := make(map[string]string)
	requestParams["account_id"] = accountID
	requestParams["limit"] = strconv.Itoa(limit)
	requestParams["since"] = since

	response := &struct {
		Transactions []*TransactionListEntry `json:"transactions"`
	}{}

	if err := m.request("GET", "transactions", response, requestParams); err != nil {
		return nil, err
	}

	return response.Transactions, nil
}

// TransactionDetails - get the full details for a single transaction
func (m *Monzgo) TransactionDetails(transactionID string) (*TransactionDetails, error) {
	requestParams := make(map[string]string)
	requestParams["expand[]"] = "merchant"

	response := &struct {
		Transaction *TransactionDetails `json:"transaction"`
	}{}

	if err := m.request("GET", "transactions/"+transactionID, response, requestParams); err != nil {
		return nil, err
	}

	return response.Transaction, nil
}
