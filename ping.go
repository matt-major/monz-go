package monzgo

// Ping the Monzo API to verify the conenction
func (m *Monzgo) Ping() error {
	return m.request("GET", "ping/whoami", &PingResult{}, nil)
}
