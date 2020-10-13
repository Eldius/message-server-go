package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Eldius/message-server-go/config"
	"github.com/Eldius/message-server-go/hashtools"
	"github.com/Eldius/message-server-go/repository"
	"github.com/Eldius/message-server-go/user"
)

const (
	invalidJwtFormat = "auth.jwt.validation.format.invalid"
	invalidJwtSign   = "auth.jwt.validation.sign.invalid"
)

// ValidatePass validates user credentials
func ValidatePass(username string, pass string) (u *user.CredentialInfo, err error) {
	var usr = repository.FindUser(username)
	if usr.Hash == nil {
		err = fmt.Errorf("Failed to authenticate user")
		return
	}

	var ph []byte
	ph, err = hashtools.Hash(pass, usr.Salt)
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

func ToJWT(u user.CredentialInfo) (jwt string, err error) {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	header, err := generateHeader(u)
	if err != nil {
		return
	}

	payload, err := generatePayload(u)
	if err != nil {
		return
	}

	jwtWOSign := fmt.Sprintf("%s.%s", header, payload)
	sign, err := signContent(jwtWOSign)
	if err != nil {
		return
	}
	jwt = fmt.Sprintf("%s.%s", jwtWOSign, sign)
	return
}

func FromJWT(jwt string) (u *user.CredentialInfo, err error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		err = fmt.Errorf(invalidJwtFormat)
		return
	}
	sign, err := signContent(fmt.Sprintf("%s.%s", parts[0], parts[1]))
	if err != nil {
		return
	}
	if sign != parts[2] {
		err = fmt.Errorf(invalidJwtSign)
		return
	}

	b, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return
	}

	var tmpData map[string]string
	err = json.Unmarshal([]byte(b), &tmpData)
	if err != nil {
		return
	}

	u = repository.FindUser(tmpData["user"])

	return
}

func generateHeader(u user.CredentialInfo) (headerStr string, err error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	headerByte, err := json.Marshal(header)
	if err != nil {
		return
	}

	headerStr = base64.StdEncoding.EncodeToString([]byte(headerByte))
	return
}

func generatePayload(u user.CredentialInfo) (payloadStr string, err error) {
	payload := map[string]string{
		"user": u.User,
		"name": u.Name,
	}
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return
	}

	payloadStr = base64.StdEncoding.EncodeToString([]byte(payloadByte))
	return
}

func signContent(content string) (sign string, err error) {
	h := hmac.New(sha256.New, []byte(config.GetJWTSecret()))

	// Write Data to it
	_, err = h.Write([]byte(content))
	if err != nil {
		return
	}
	sign = hex.EncodeToString(h.Sum(nil))

	return
}
