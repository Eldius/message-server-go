package auth

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/Eldius/auth-server-go/config"
	"github.com/Eldius/auth-server-go/repository"
	"github.com/Eldius/auth-server-go/user"
	"github.com/spf13/viper"
)

var (
	tmpDir string
)

func init()  {
	//viper.SetConfigFile("../config/samples/auth-server-sqlite3.yml")
	var err error
	tmpDir, err = ioutil.TempDir("", "auth-server")
	if err != nil {
		log.Println("Failed to setup temp database")
		log.Fatal(err.Error())
	}
	os.RemoveAll("/tmp/auth-server-test")
	os.MkdirAll("/tmp/auth-server-test", os.ModePerm)
	viper.SetDefault("app.database.url", fmt.Sprintf("%s/test.db", tmpDir))
	viper.SetDefault("app.database.engine", "sqlite3")
	log.Println("db file:", config.GetDBURL())

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestValidatePass(t *testing.T)  {
	username := "user"
	passwd := "pass"
	u, err := user.NewCredentials(username, passwd)
	if err != nil {
		t.Errorf("Failed to prepare test user:\n%s", err.Error())
	}

	repository.SaveUser(&u)

	//au := repository.ListUSers()
	//log.Println("au", au)


	c, err := ValidatePass(username, passwd)
	if err != nil {
		t.Error(err)
	}

	if c == nil {
		t.Errorf("Failed to validate user (returned nil value)")
	}

}

func TestValidatePassInvalidCredentials(t *testing.T)  {
	username := "user1"
	passwd := "pass1"
	u, err := user.NewCredentials(username, passwd)
	if err != nil {
		t.Errorf("Failed to prepare test user:\n%s", err.Error())
	}

	repository.SaveUser(&u)

	//au := repository.ListUSers()
	//log.Println("au", au)


	c, err := ValidatePass(username, "pass")
	if err != nil {
		t.Error(err)
	}

	if c != nil {
		t.Errorf("Failed to validate user (returned nil value)")
	}

}
