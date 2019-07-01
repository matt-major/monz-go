package monzgo

import (
	"strconv"
)

// Pots - retrieve all pots for the user, optionally excluding deleted ones
func (m *Monzgo) Pots(ignoreDeleted bool) ([]*Pot, error) {
	response := &struct {
		Pots []*Pot `json:"pots"`
	}{}

	if err := m.request("GET", "pots", response, nil); err != nil {
		return nil, err
	}

	pots := response.Pots
	if ignoreDeleted {
		pots = filterPots(pots)
	}

	return pots, nil
}

// AddToPot - add the specified amount to a Pot from the provided Source Account
func (m *Monzgo) AddToPot(potID string, sourceAccountID string, amount int64, dedupeID string) (*Pot, error) {
	requestData := make(map[string]string)
	requestData["source_account_id"] = sourceAccountID
	requestData["amount"] = strconv.FormatInt(amount, 10)
	requestData["dedupe_id"] = dedupeID

	pot := &Pot{}
	if err := m.request("PUT", "pots/"+potID+"/deposit", pot, requestData); err != nil {
		return nil, err
	}

	return pot, nil
}

// WithdrawFromPot - withdraw the specified amount from a Pot in to the provided Destination Account
func (m *Monzgo) WithdrawFromPot(potID string, destinationAccountID string, amount int64, dedupeID string) (*Pot, error) {
	requestData := make(map[string]string)
	requestData["destination_account_id"] = destinationAccountID
	requestData["amount"] = strconv.FormatInt(amount, 10)
	requestData["dedupe_id"] = dedupeID

	pot := &Pot{}
	if err := m.request("PUT", "pots/"+potID+"/withdraw", pot, requestData); err != nil {
		return nil, err
	}

	return pot, nil
}

func filterPots(pots []*Pot) []*Pot {
	filtered := make([]*Pot, 0)

	for _, pot := range pots {
		if !pot.Deleted {
			filtered = append(filtered, pot)
		}
	}

	return filtered
}
