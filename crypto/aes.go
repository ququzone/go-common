package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// AES ...
type AES struct {
	Key []byte
	IV  []byte
}

// NewAES ...
func NewAES(key string, iv string) *AES {
	kb := make([]byte, 32)
	ib := make([]byte, 16)
	copy(kb, []byte(key))
	copy(ib, []byte(iv))
	return &AES{
		Key: kb,
		IV:  ib,
	}
}

// Encrypt ...
func (a *AES) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	stream := cipher.NewCTR(block, a.IV)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// Decrypt ...
func (a *AES) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	stream := cipher.NewCTR(block, a.IV)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
	return plaintext, nil
}
