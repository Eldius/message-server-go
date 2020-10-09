package repository

import (
	"github.com/Eldius/auth-server-go/clients"
	"github.com/Eldius/auth-server-go/logger"
	"github.com/jinzhu/gorm"
)

// SaveUser saves the new user credential
func SaveClient(c *clients.ClientInfo) {
	if c == nil {
		return
	}
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Save(c).Error; err != nil {
			// return any error will rollback
			logger.Logger().Println("Error saving credentials")
			logger.Logger().Println(err.Error())
			return err
		}
		// return nil will commit
		return nil
	})
	if err != nil {
		logger.Logger().Panicln("Failed to insert data\n", err.Error())
	}
}

// FindUser finds the user
func FindClient(name string) *clients.ClientInfo {

	u := clients.ClientInfo{}
	GetDB().Where("Name = ?", name).First(&u)
	return &u
}

// ListUSers returns all users
func ListClients() (r []clients.ClientInfo) {
	GetDB().Find(&r, "")
	return
}
