package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

func TestPing(t *testing.T) {
	t.Run("Success", testPingSuccess())
	t.Run("Failure", testPingError())
}

func testPingSuccess() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/",
		}

		err := m.Ping()
		if err != nil {
			t.Errorf("Ping failed %s", err)
		}
	}
}

func testPingError() func(*testing.T) {
	return func(t *testing.T) {
		m := monzgo.Monzgo{
			APIKey:  os.Getenv("MONZO_API_KEY"),
			BaseURL: "https://api.monzo.com/FAIL_ME_ALWAYS",
		}

		err := m.Ping()
		if err == nil {
			t.Error("Ping succeeded but expected failure")
		}
	}
}
