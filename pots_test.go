package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestPots(t *testing.T) {
	t.Run("Get all Pots", testGetAllPots())
}

func testGetAllPots() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		pots, err := m.Pots(true)
		if err != nil {
			t.Errorf("Failed to fetch all pots; %s", err)
		}

		if len(pots) == 0 {
			t.Errorf("Received fewer pots than expected; %s", err)
		}
	}
}
