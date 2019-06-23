package monzgo

// Pot definition
type Pot struct {
	ID       string
	Name     string
	Style    string
	Balance  int64
	Currency string
	Created  string
	Updated  string
	Deleted  bool
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
