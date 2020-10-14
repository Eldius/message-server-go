package repository

import (
	"github.com/Eldius/message-server-go/logger"
	"github.com/Eldius/message-server-go/messenger"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// SaveMessage saves the new user credential
func SaveMessage(c *messenger.Message) {
	if c == nil {
		return
	}
	err := GetDB().Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Save(c).Error; err != nil {
			// return any error will rollback
			logger.Logger().Println("Error saving message")
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

// FindMessageTo message based on the receiver
func FindMessageTo(to int) []messenger.Message {

	m := []messenger.Message{}
	GetDB().Where("Destination = ?", to).Order("Sent asc").Find(&m)
	return m
}

// FindMessageTo message based on the receiver
func FindMessageFrom(from int) []messenger.Message {

	m := []messenger.Message{}
	GetDB().Where("To = ?", from).First(&m)
	return m
}

// FindMessageTo message based on the receiver
func FindMessageByID(msgId uuid.UUID) []messenger.Message {

	m := []messenger.Message{}
	GetDB().Where("To = ?", msgId).First(&m)
	return m
}

// ListMessages returns all users
func ListMessages() (r []messenger.Message) {
	GetDB().Find(&r, "")
	return
}
