package auth

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/Eldius/message-server-go/config"
	"github.com/Eldius/message-server-go/logger"
	"github.com/Eldius/message-server-go/repository"
	"github.com/Eldius/message-server-go/user"
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
	if err == nil {
		t.Error("Should return an error in login process")
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

func TestToJWT(t *testing.T) {
	u, err := user.NewCredentials("myUser", "myPass")
	if err != nil {
		t.Errorf("Failed to create the test user\n%s", err.Error())
	}
	token, err := ToJWT(u)
	if err != nil {
		t.Errorf("Failed to create token\n%s", err.Error())
	}
	t.Logf("token: %s", token)

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Errorf("Should return 3 parts separated by dot (.), but returned %d", len(parts))
	}
}
