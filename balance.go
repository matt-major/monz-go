package monzgo

// Balance response definition
type Balance struct {
	Balance      int64  `json:"balance"`
	TotalBalance int64  `json:"total_balance"`
	Currency     string `json:"currency"`
	SpendToday   int64  `json:"spend_today"`
}

// Balance - get the balance for a given accountID
func (m *Monzgo) Balance(accountID string) (*Balance, error) {
	rspHolder := &Balance{}
	requestParams := map[string]string{
		"account_id": accountID,
	}

	if err := m.request("GET", "balance", rspHolder, requestParams); err != nil {
		return nil, err
	}

	return rspHolder, nil
}
