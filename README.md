# lazyxchacha

Lazy XChaCha20-Poly1305 in golang base on [golang.org/x/crypto](golang.org/x/crypto).

[![Go Report Card](https://goreportcard.com/badge/github.com/prongbang/lazyxchacha)](https://goreportcard.com/report/github.com/prongbang/lazyxchacha)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/prongbang)

### Algorithm details

- Key exchange: X25519
- Encryption: XChaCha20
- Authentication: Poly1305

### Install

```
go get github.com/prongbang/lazyxchacha
```

### Benchmark

```shell
BenchmarkEncrypt-10    	 1220955	       984.5 ns/op	     944 B/op	       5 allocs/op
BenchmarkDecrypt-10    	 1706475	       695.0 ns/op	     576 B/op	       4 allocs/op
```

### How to use

- Generate KeyPair

```go
keyPair := lazyxchacha.NewKeyPair()
```

- Key Exchange

```go
// Generate KeyPair
clientKp := lazyxchacha.NewKeyPair()
serverKp := lazyxchacha.NewKeyPair()

serverKx := serverKp.Exchange(clientKp.Pk)
clientKx := clientKp.Exchange(serverKp.Pk)
```

- Shared Key

```go
serverSharedKey, _ := serverKx.Secret()
clientSharedKey, _ := clientKx.Secret()
```

- Encrypt

```go
lazyXchacha := lazyxchacha.New()
key, _ := lazyXchacha.RandomKey()
plaintext := "text"
ciphertext, err := lazyXchacha.Encrypt(plaintext, key)
```

- Decrypt

```go
lazyXchacha := lazyxchacha.New()
key := "e7de22e8"
ciphertext := "f6a1bd8"
plaintext, err := lazyXchacha.Decrypt(ciphertext, key)
```
