package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestGetAccounts(t *testing.T) {
	key, _ := os.LookupEnv("MONZO_API_KEY")
	m := monzgo.Setup(key)

	accounts, err := m.Accounts()
	if err != nil {
		t.Errorf("Get Accounts failed: %s", err)
	}

	want := 1
	got := len(accounts)
	if got != want {
		t.Errorf("Got %d Accounts, wanted %d", got, want)
	}
}
