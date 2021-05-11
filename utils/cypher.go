package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type Cypher []byte

var cypher *Cypher

func SetupCypher() {
	Cypher := Cypher("1m4$up3R$3cR37p4ssw0Rd!@#qwertyu")

	setCypher(&Cypher)
}

func setCypher(Cypher *Cypher) {
	cypher = Cypher
}

func GetCypher() Cypher {
	return *cypher
}

func (cy *Cypher) Encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher(*cy)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func (cy *Cypher) Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher(*cy)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
