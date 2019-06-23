package monzgo

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

// Pots - retrieve all pots for the user, optionally excluding deleted ones
func (m *Monzgo) Pots(ignoreDeleted bool) ([]*Pot, error) {
	rspHolder := &struct {
		Pots []*Pot `json:"pots"`
	}{}

	if err := m.request("GET", "pots", rspHolder, nil); err != nil {
		return nil, err
	}

	pots := rspHolder.Pots
	if ignoreDeleted {
		pots = filterPots(pots)
	}

	return pots, nil
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
