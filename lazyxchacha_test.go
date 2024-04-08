package lazyxchacha_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/prongbang/lazyxchacha"
)

func TestEncrypt(t *testing.T) {
	lazyXchacha := lazyxchacha.New()
	key, _ := lazyXchacha.RandomKey()
	plaintext := `{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"}`
	ciphertext, err := lazyXchacha.Encrypt(plaintext, key)
	fmt.Println(key)
	fmt.Println(ciphertext, err)
}

func TestDecrypt(t *testing.T) {
	key, _ := hex.DecodeString("e7de22e898b35cf5ed2c2339f702210429da35909c6e070af48e1c01e8a34d55")
	ciphertext := "f6a1bd8e40d1ac130a53c85f8eb3ce5c8f524dd16b6c844eba81f40430dc3f43a20f19ce0ff9cac0fb552e945c9c9eb03eef3d3ec120b0ff17e8181e2e3d949b2eb44180b494d72a33d79d30a4de4130488aabfb2922f7265c7010ddf649a231856f1a8dccd57284a53230c79d16a732d38a48f9a1fab78e3dca7eff3d48bb848a3f04169cefbf021523dc6e62def880ffefcd1e4d"
	lazyXchacha := lazyxchacha.New()
	plaintext, err := lazyXchacha.Decrypt(ciphertext, key)
	fmt.Println(plaintext, err)
}

// BenchmarkEncrypt-10    	 1220955	       984.5 ns/op	     944 B/op	       5 allocs/op
func BenchmarkEncrypt(b *testing.B) {
	lazyXchacha := lazyxchacha.New()
	key, _ := lazyXchacha.RandomKey()
	plaintext := `{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"}`
	for i := 0; i < b.N; i++ {
		lazyXchacha.Encrypt(plaintext, key)
	}
}

// BenchmarkDecrypt-10    	 1689169	       689.0 ns/op	     576 B/op	       4 allocs/op
func BenchmarkDecrypt(b *testing.B) {
	lazyXchacha := lazyxchacha.New()
	key, _ := hex.DecodeString("e7de22e898b35cf5ed2c2339f702210429da35909c6e070af48e1c01e8a34d55")
	ciphertext := "f6a1bd8e40d1ac130a53c85f8eb3ce5c8f524dd16b6c844eba81f40430dc3f43a20f19ce0ff9cac0fb552e945c9c9eb03eef3d3ec120b0ff17e8181e2e3d949b2eb44180b494d72a33d79d30a4de4130488aabfb2922f7265c7010ddf649a231856f1a8dccd57284a53230c79d16a732d38a48f9a1fab78e3dca7eff3d48bb848a3f04169cefbf021523dc6e62def880ffefcd1e4d"
	for i := 0; i < b.N; i++ {
		lazyXchacha.Decrypt(ciphertext, key)
	}
}
