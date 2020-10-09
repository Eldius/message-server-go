package clients

import (
	"time"

	"github.com/Eldius/auth-server-go/hashtools"
	"github.com/google/uuid"
)

type ClientInfo struct {
	ID        int
	CreatedAt time.Time
	Name      string
	Active    bool
	HashedKey []byte
	Salt      []byte
}

func NewClientInfo(name string) (*ClientInfo, string, error) {
	key := uuid.New().String()
	salt := hashtools.Salt()
	hashedKey, err := hashtools.HashKey(key, salt)
	if err == nil {
		return &ClientInfo{
			CreatedAt: time.Now(),
			Active:    true,
			Name:      name,
			HashedKey: hashedKey,
			Salt:      salt,
		}, key, nil
	} else {
		return nil, "", err
	}
}
