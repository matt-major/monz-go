package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestPing(t *testing.T) {
	m := monzgo.Monzgo{
		APIKey:  os.Getenv("MONZO_API_KEY"),
		BaseURL: "https://api.monzo.com/",
	}

	err := m.Ping()
	if err != nil {
		t.Errorf("Ping failed %s", err)
	}
}
