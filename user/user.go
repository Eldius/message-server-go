package user

import (
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
	"regexp"

	"github.com/Eldius/auth-server-go/config"
	"github.com/Eldius/auth-server-go/logger"
	"golang.org/x/crypto/scrypt"
)

const (
	_pwSaltBytes    = 32
	_pwHashBytes    = 64
	emptyUsername   = "credentials.username.must.not.be.empty"
	invalidUsername = "credentials.username.must.match.pattern"
)

/*
CredentialInfo represents the user credentials
*/
type CredentialInfo struct {
	ID     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	User   string `gorm:"unique;not null;UNIQUE_INDEX"`
	Hash   []byte `gorm:"not null"`
	Salt   []byte `gorm:"not null"`
	Name   string
	Active bool
	Admin  bool
}

/*
Profile is the user profile
*/
type Profile struct {
	ID          int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	Name        string `gorm:"unique;not null;UNIQUE_INDEX"`
	Description string
	Active      bool
}

/*
NewCredentials  creates a new CredentialInfo
*/
func NewCredentials(user string, pass string) (cred CredentialInfo, err error) {

	if err = validateUsername(user); err != nil {
		return
	}

	h := sha512.New()
	_, err = h.Write([]byte(pass))
	if err != nil {
		logger.Logger().Println("Failed to handle pass")
		return
	}
	salt := salt()
	hash, err := Hash(pass, salt)
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

func validateUsername(username string) error {
	if username == "" {
		return errors.New(emptyUsername)
	}

	r := regexp.MustCompile(config.GetUsernamePattern())
	if !r.MatchString(username) {
		return errors.New(invalidUsername)
	}
	return nil
}

func salt() []byte {
	salt := make([]byte, _pwSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		logger.Logger().Fatal(err)
	}

	return salt
}

// Hash returns the user pass' hash
func Hash(pass string, salt []byte) (hash []byte, err error) {
	hash, err = scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, _pwHashBytes)
	if err != nil {
		logger.Logger().Fatal(err)
	}
	return
}
