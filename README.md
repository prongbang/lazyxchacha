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
goos: darwin
goarch: arm64
cpu: Apple M4 Pro
BenchmarkEncrypt-12         	 1882269	       621.3 ns/op
BenchmarkDecrypt-12         	 2717778	       432.8 ns/op
BenchmarkEncryptBytes-12    	 2503562	       471.8 ns/op
BenchmarkDecryptBytes-12    	 3929430	       301.2 ns/op
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
sharedKey, _ := clientKx.Secret()
key, _ := hex.DecodeString(sharedKey)
plaintext := "text"
ciphertext, err := lazyXchacha.Encrypt(plaintext, key)
```

- Decrypt

```go
lazyXchacha := lazyxchacha.New()
sharedKey, _ := serverKx.Secret()
key, _ := hex.DecodeString(sharedKey)
ciphertext := "f6a1bd8"
plaintext, err := lazyXchacha.Decrypt(ciphertext, key)
```
