package auth

import (
	"github.com/Eldius/auth-server-go/repository"
	"github.com/Eldius/auth-server-go/user"
	u "github.com/Eldius/auth-server-go/user"
)

// ValidatePass validates user credentials
func ValidatePass(user string, pass string) *user.CredentialInfo {
	var usr u.CredentialInfo
	repository.GetDB().Where("User = ?", user).First(&usr)

	return &usr
}
