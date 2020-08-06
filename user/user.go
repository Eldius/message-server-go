package user

import (
	"crypto/rand"
	"crypto/sha512"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	_pwSaltBytes = 32
	_pwHashBytes = 64
)

/*
CredentialInfo represents the user credentials
*/
type CredentialInfo struct {
	ID     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	User   string `gorm:"unique;not null;UNIQUE_INDEX"`
	Hash   []byte `gorm:"not null"`
	Salt   []byte `gorm:"not null"`
	Active bool
}

/*
Profile is the user profile
*/
type Profile struct {
	ID     int
	Name   string
	Active bool
}

/*
NewCredentials  creates a new CredentialInfo
*/
func NewCredentials(user string, pass string) (cred CredentialInfo, err error) {

	h := sha512.New()
	h.Write([]byte(pass))
	salt := salt()
	hash, err := hash(pass, salt)
	if err != nil {
		return
	}
	cred = CredentialInfo{
		User:   user,
		Salt:   salt,
		Hash:   hash,
		Active: true,
	}

	return
}

func salt() []byte {
	salt := make([]byte, _pwSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	return salt
}

func hash(pass string, salt []byte) (hash []byte, err error) {
	hash, err = scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, _pwHashBytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%x\n", hash)
	return
}
