package lazyxchacha_test

import (
	"testing"

	"github.com/prongbang/lazyxchacha"
)

func TestKeyPair(t *testing.T) {
	// Generate KeyPair
	clientKp := lazyxchacha.NewKeyPair()
	serverKp := lazyxchacha.NewKeyPair()

	// KEy Exchange
	serverKx := serverKp.Exchange(clientKp.Pk)
	clientKx := clientKp.Exchange(serverKp.Pk)

	// Shared Key
	serverSharedKey, _ := serverKx.Secret()
	clientSharedKey, _ := clientKx.Secret()

	if serverSharedKey != clientSharedKey {
		t.Errorf("Error %s != %s", serverSharedKey, clientSharedKey)
	}
}
