package lazyxchacha

import (
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/chacha20poly1305"

	cryptorand "crypto/rand"
)

type LazyXChaCha interface {
	RandomKey() ([]byte, error)
	Encrypt(plaintext string, key []byte) (string, error)
	Decrypt(ciphertext string, key []byte) (string, error)
	EncryptBytes(plaintext []byte, key []byte) ([]byte, error)
	DecryptBytes(ciphertext []byte, key []byte) ([]byte, error)
}

type lazyXChaCha struct {
}

func (c *lazyXChaCha) EncryptBytes(plaintext []byte, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, err
	}

	// Select a random nonce, and leave capacity for the ciphertext.
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(plaintext)+aead.Overhead())
	if _, err := cryptorand.Read(nonce); err != nil {
		return nil, err
	}

	// Encrypt the message and append the ciphertext to the nonce.
	encrypted := aead.Seal(nonce, nonce, plaintext, nil)

	return encrypted, nil
}

func (c *lazyXChaCha) DecryptBytes(ciphertext []byte, key []byte) ([]byte, error) {
	aead, err := chacha20poly1305.NewX(key[:])
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aead.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}

	// Split nonce and ciphertext.
	nonce, encrypted := ciphertext[:aead.NonceSize()], ciphertext[aead.NonceSize():]

	// Decrypt the message and check it wasn't tampered with.
	plaintext, err := aead.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func (c *lazyXChaCha) RandomKey() ([]byte, error) {
	key := make([]byte, chacha20poly1305.KeySize)
	if _, err := cryptorand.Read(key); err != nil {
		return []byte{}, err
	}
	return key, nil
}

func (c *lazyXChaCha) Encrypt(plaintext string, key []byte) (string, error) {
	encrypted, err := c.EncryptBytes([]byte(plaintext), key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encrypted), nil
}

func (c *lazyXChaCha) Decrypt(ciphertext string, key []byte) (string, error) {
	encrypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := c.DecryptBytes(encrypted, key)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func New() LazyXChaCha {
	return &lazyXChaCha{}
}
