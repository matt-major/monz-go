package monzgo

// Balance - get the balance for a given accountID
func (m *Monzgo) Balance(accountID string) (*Balance, error) {
	requestParams := make(map[string]string)
	requestParams["account_id"] = accountID

	balance := &Balance{}
	if err := m.request("GET", "balance", balance, requestParams); err != nil {
		return nil, err
	}

	return balance, nil
}
