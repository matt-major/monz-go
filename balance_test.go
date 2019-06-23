package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestBalance(t *testing.T) {
	t.Run("Get balance", testGetBalance())
}

func testGetBalance() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		accounts, err := m.Accounts("")

		_, err = m.Balance(accounts[0].ID)
		if err != nil {
			t.Errorf("Failed to fetch balance; %s", err)
		}
	}
}
