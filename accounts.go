package monzgo

// Accounts - get all accounts with an optional type filter
func (m *Monzgo) Accounts(typeFilter string) ([]*Account, error) {
	requestParams := make(map[string]string)
	if typeFilter != "" {
		requestParams["account_type"] = typeFilter
	}

	response := &struct {
		Accounts []*Account `json:"accounts"`
	}{}

	if err := m.request("GET", "accounts", response, requestParams); err != nil {
		return nil, err
	}

	return response.Accounts, nil
}
