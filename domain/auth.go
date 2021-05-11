package domain

import (
	"encoding/hex"
	"gmauth/utils"

	uuid "github.com/satori/go.uuid"
)

type Auth struct {
	SID         uuid.UUID `json:"-"`
	State       uuid.UUID `json:"state"`
	SIDCrypt    string    `json:"sid"`
	RedirectURI string    `json:"redirect_uri"`
}

func NewAuth() (Auth, error) {
	sid := uuid.NewV4()
	state := uuid.NewV4()

	cypher := utils.GetCypher()

	ciphertext, err := cypher.Encrypt([]byte(sid.String()))
	if err != nil {
		return Auth{}, err
	}

	hex.EncodeToString(ciphertext)
	return Auth{SIDCrypt: hex.EncodeToString(ciphertext), State: state, RedirectURI: "http://localhost:8080"}, nil
}

/*
test2, _ := hex.DecodeString(test)
plaintext, err := cypher.Decrypt(test2)
if err != nil {
	// TODO: Properly handle error
	log.Fatal(err)
}
fmt.Println(string(plaintext))
*/
