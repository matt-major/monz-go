package monzgo

// PingResult - the result of the Ping / Verify request
type PingResult struct {
	Authentication bool
	ClientID       string `json:"client_id"`
	UserID         string `json:"user_id"`
}

// Ping the Monzo API to verify the conenction
func (m *Monzgo) Ping() error {
	var result PingResult
	if err := m.request("GET", "ping/whoami", &result); err != nil {
		return err
	}

	return nil
}
