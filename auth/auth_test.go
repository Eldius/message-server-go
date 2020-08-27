package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Eldius/auth-server-go/config"
	"github.com/Eldius/auth-server-go/logger"
	"github.com/Eldius/auth-server-go/repository"
	"github.com/Eldius/auth-server-go/user"
	"github.com/spf13/viper"
)

var (
	tmpDir string
)

func init() {
	//viper.SetConfigFile("../config/samples/auth-server-sqlite3.yml")
	var err error
	tmpDir, err = ioutil.TempDir("", "auth-server")
	if err != nil {
		logger.Logger().Println("Failed to setup temp database")
		logger.Logger().Fatal(err.Error())
	}
	os.RemoveAll("/tmp/auth-server-test")
	if err := os.MkdirAll("/tmp/auth-server-test", os.ModePerm); err != nil {
		logger.Logger().Panic("Failed to create temp dir for tests.")
	}
	viper.SetDefault("app.database.url", fmt.Sprintf("%s/test.db", tmpDir))
	viper.SetDefault("app.database.engine", "sqlite3")
	logger.Logger().Println("db file:", config.GetDBURL())

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Logger().Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestValidatePass(t *testing.T) {
	username := "user"
	passwd := "pass"
	u, err := user.NewCredentials(username, passwd)
	if err != nil {
		t.Errorf("Failed to prepare test user:\n%s", err.Error())
	}

	repository.SaveUser(&u)

	c, err := ValidatePass(username, passwd)
	if err != nil {
		t.Error(err)
	}

	if c == nil {
		t.Errorf("Failed to validate user (returned nil value)")
	}

}

func TestValidatePassInvalidCredentials(t *testing.T) {
	username := "user1"
	passwd := "pass1"
	u, err := user.NewCredentials(username, passwd)
	if err != nil {
		t.Errorf("Failed to prepare test user:\n%s", err.Error())
	}

	repository.SaveUser(&u)

	c, err := ValidatePass(username, "pass")
	if err != nil {
		t.Error(err)
	}

	if c != nil {
		t.Errorf("Failed to validate user (returned nil value)")
	}

}

func TestValidatePassUserNotFound(t *testing.T) {
	username := "user2"
	passwd := "pass1"

	c, err := ValidatePass(username, passwd)
	if err == nil {
		t.Error("Should return an error")
	}

	if c != nil {
		t.Errorf("Failed to validate user (returned nil value)")
	}

}
