package monzgo_test

import (
	"os"
	"testing"

	"github.com/matt-major/monzgo"
)

var apiKey string

func setup() {
	apiKey, _ = os.LookupEnv("MONZO_API_KEY")
}

func TestSetup(t *testing.T) {
	setup()
	if apiKey == "" {
		t.Errorf("Tests require MONZO_API_KEY environment variable to be set")
	}
}

func TestMonzgoCore(t *testing.T) {
	setup()
	m := monzgo.Setup(apiKey)

	err := m.Verify()
	if err != nil {
		t.Errorf("Verify failed: %s", err)
	}
}
