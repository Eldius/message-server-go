package user

import (
	"crypto/sha512"
	"errors"
	"regexp"

	"github.com/Eldius/auth-server-go/config"
	"github.com/Eldius/auth-server-go/hashtools"
	"github.com/Eldius/auth-server-go/logger"
)

const (
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
	salt := hashtools.Salt()
	hash, err := hashtools.Hash(pass, salt)
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
