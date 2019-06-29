package monzgo

// Ping the Monzo API to verify the conenction
func (m *Monzgo) Ping() error {
	result := &PingResult{}
	return m.request("GET", "ping/whoami", result, nil)
}
