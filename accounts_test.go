package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestAccounts(t *testing.T) {
	t.Run("GetAll", testGetAllAccounts())
}

func testGetAllAccounts() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		accounts, err := m.Accounts()

		if err != nil {
			t.Errorf("Failed to fetch all accounts; %s", err)
		}

		if len(accounts) != 1 {
			t.Errorf("Received more accounts than expected; %s", err)
		}
	}
}
