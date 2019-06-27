package monzgo

// Ping the Monzo API to verify the conenction
func (m *Monzgo) Ping() error {
	result := &PingResult{}
	if err := m.request("GET", "ping/whoami", result, nil); err != nil {
		return err
	}

	return nil
}
