package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func generateKey() []byte {
	key := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	return key
}

func encrypt(data []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	encrypted := make([]byte, len(data))
	stream.XORKeyStream(encrypted, data)

	return append(iv, encrypted...)
}
