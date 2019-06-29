package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestFeedItems(t *testing.T) {
	t.Run("Add Feed item to account", testAddFeedItemToAccount())
}

func testAddFeedItemToAccount() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		accounts, err := m.Accounts("")

		if err != nil {
			t.Errorf("Failed to fetch all accounts; %s", err)
		}

		opts := monzgo.FeedItemOptions{
			Type:  "basic",
			Title: "Test Feed Item",
			Body:  "A simple test",
			ImageURL: "https://i.kym-cdn.com/entries/icons/mobile/000/005/608/nyan-cat-01-625x450.jpg",
		}

		err = m.AddFeedItem(accounts[0].ID, opts)
		if err != nil {
			t.Errorf("Failed to add Feed Item; %s", err)
		}
	}
}
