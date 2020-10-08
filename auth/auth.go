package auth

import (
	"fmt"

	"github.com/Eldius/auth-server-go/repository"
	"github.com/Eldius/auth-server-go/user"
)

// ValidatePass validates user credentials
func ValidatePass(username string, pass string) (u *user.CredentialInfo, err error) {
	var usr = repository.FindUser(username)
	if usr.Hash == nil {
		err = fmt.Errorf("Failed to authenticate user")
		return
	}

	var ph []byte
	ph, err = user.Hash(pass, usr.Salt)
	if err != nil {
		return
	}

	if string(ph) == string(usr.Hash) {
		u = usr
	} else {
		err = fmt.Errorf("Failed to authenticate user")
	}

	return
}
