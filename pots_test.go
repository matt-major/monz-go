package monzgo_test

import (
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestPots(t *testing.T) {
	t.Run("Get all Pots", testGetAllPots())
	t.Run("Add funds to Pots", testAddFundsToPot())
	t.Run("Withdraw funds from Pots", testWithdrawFundsFromPot())
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

func testAddFundsToPot() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		accs, err := m.Accounts("")
		if err != nil {
			t.Errorf("%s", err)
		}

		pots, err := m.Pots(true)
		if err != nil {
			t.Errorf("%s", err)
		}

		_, err = m.AddToPot(pots[0].ID, accs[0].ID, 1, strconv.Itoa(rand.Intn(100)))
		if err != nil {
			t.Errorf("%s", err)
		}
	}
}

func testWithdrawFundsFromPot() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		accs, err := m.Accounts("")
		if err != nil {
			t.Errorf("%s", err)
		}

		pots, err := m.Pots(true)
		if err != nil {
			t.Errorf("%s", err)
		}

		_, err = m.WithdrawFromPot(pots[0].ID, accs[0].ID, 1, strconv.Itoa(rand.Intn(100)))
		if err != nil {
			t.Errorf("%s", err)
		}
	}
}
