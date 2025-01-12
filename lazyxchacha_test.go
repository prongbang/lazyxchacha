package lazyxchacha_test

import (
	"encoding/hex"
	"fmt"
	"strings"
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

func TestEncryptBytes(t *testing.T) {
	lazyXchacha := lazyxchacha.New()
	key, _ := lazyXchacha.RandomKey()
	plaintext := []byte(`{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"}`)
	ciphertext, err := lazyXchacha.EncryptBytes(plaintext, key)

	keys := []string{}
	for _, b := range key {
		keys = append(keys, fmt.Sprintf("%v", b))
	}

	cipher := []string{}
	for _, b := range ciphertext {
		cipher = append(cipher, fmt.Sprintf("%v", b))
	}

	fmt.Println("key:", strings.Join(keys, ","))
	fmt.Println("ciphertext:", strings.Join(cipher, ","), err)
}

func TestDecryptBytes(t *testing.T) {
	key := []byte{185, 137, 144, 29, 162, 230, 64, 16, 229, 30, 85, 35, 181, 246, 92, 209, 127, 242, 67, 4, 205, 62, 253, 233, 148, 2, 143, 190, 241, 2, 154, 52}
	ciphertext := []byte{209, 9, 224, 73, 76, 138, 115, 100, 170, 10, 225, 80, 63, 76, 89, 108, 60, 53, 11, 160, 201, 125, 82, 0, 7, 49, 52, 164, 244, 138, 177, 76, 15, 219, 10, 61, 203, 16, 145, 213, 223, 63, 226, 162, 146, 224, 119, 53, 180, 205, 82, 167, 60, 252, 222, 61, 29, 146, 48, 223, 170, 183, 213, 214, 152, 202, 23, 207, 23, 89, 4, 72, 199, 128, 120, 129, 197, 137, 229, 136, 221, 98, 185, 42, 116, 72, 174, 90, 193, 40, 246, 71, 38, 107, 42, 119, 192, 1, 168, 81, 98, 252, 88, 168, 153, 116, 175, 50, 255, 152, 24, 21, 156, 155, 1, 69, 46, 205, 157, 242, 137, 236, 94, 53, 226, 235, 150, 211, 55, 85, 87, 154, 210, 205, 146, 185, 42, 180, 118, 39, 188, 32, 137, 228, 100, 26, 19, 112, 60}
	lazyXchacha := lazyxchacha.New()
	plaintext, err := lazyXchacha.DecryptBytes(ciphertext, key)
	fmt.Println(string(plaintext), err)
}

func BenchmarkEncrypt(b *testing.B) {
	lazyXchacha := lazyxchacha.New()
	key, _ := lazyXchacha.RandomKey()
	plaintext := `{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"}`
	for i := 0; i < b.N; i++ {
		_, err := lazyXchacha.Encrypt(plaintext, key)
		if err != nil {
			b.Errorf("Error %s", err)
		}
	}
}

func BenchmarkEncryptBytes(b *testing.B) {
	lazyXchacha := lazyxchacha.New()
	key, _ := lazyXchacha.RandomKey()
	plaintext := []byte(`{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"}`)
	for i := 0; i < b.N; i++ {
		_, err := lazyXchacha.EncryptBytes(plaintext, key)
		if err != nil {
			b.Errorf("Error %s", err)
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	lazyXchacha := lazyxchacha.New()
	key, _ := hex.DecodeString("e7de22e898b35cf5ed2c2339f702210429da35909c6e070af48e1c01e8a34d55")
	ciphertext := "f6a1bd8e40d1ac130a53c85f8eb3ce5c8f524dd16b6c844eba81f40430dc3f43a20f19ce0ff9cac0fb552e945c9c9eb03eef3d3ec120b0ff17e8181e2e3d949b2eb44180b494d72a33d79d30a4de4130488aabfb2922f7265c7010ddf649a231856f1a8dccd57284a53230c79d16a732d38a48f9a1fab78e3dca7eff3d48bb848a3f04169cefbf021523dc6e62def880ffefcd1e4d"
	for i := 0; i < b.N; i++ {
		_, err := lazyXchacha.Decrypt(ciphertext, key)
		if err != nil {
			b.Errorf("Error %s", err)
		}
	}
}

func BenchmarkDecryptBytes(b *testing.B) {
	lazyXchacha := lazyxchacha.New()
	key := []byte{185, 137, 144, 29, 162, 230, 64, 16, 229, 30, 85, 35, 181, 246, 92, 209, 127, 242, 67, 4, 205, 62, 253, 233, 148, 2, 143, 190, 241, 2, 154, 52}
	ciphertext := []byte{209, 9, 224, 73, 76, 138, 115, 100, 170, 10, 225, 80, 63, 76, 89, 108, 60, 53, 11, 160, 201, 125, 82, 0, 7, 49, 52, 164, 244, 138, 177, 76, 15, 219, 10, 61, 203, 16, 145, 213, 223, 63, 226, 162, 146, 224, 119, 53, 180, 205, 82, 167, 60, 252, 222, 61, 29, 146, 48, 223, 170, 183, 213, 214, 152, 202, 23, 207, 23, 89, 4, 72, 199, 128, 120, 129, 197, 137, 229, 136, 221, 98, 185, 42, 116, 72, 174, 90, 193, 40, 246, 71, 38, 107, 42, 119, 192, 1, 168, 81, 98, 252, 88, 168, 153, 116, 175, 50, 255, 152, 24, 21, 156, 155, 1, 69, 46, 205, 157, 242, 137, 236, 94, 53, 226, 235, 150, 211, 55, 85, 87, 154, 210, 205, 146, 185, 42, 180, 118, 39, 188, 32, 137, 228, 100, 26, 19, 112, 60}
	for i := 0; i < b.N; i++ {
		_, err := lazyXchacha.DecryptBytes(ciphertext, key)
		if err != nil {
			b.Errorf("Error %s", err)
		}
	}
}
