package googleplay

import (
	"os"
	"testing"
)

func TestToken(t *testing.T) {
	tok, err := NewToken(email, password)
	if err != nil {
		t.Fatal(err)
	}
	config, err := os.UserConfigDir()
	if err != nil {
		t.Fatal(err)
	}
	if err := tok.Create(config, "googleplay/token.json"); err != nil {
		t.Fatal(err)
	}
}
