package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("app.database.url", "app.db")
	viper.SetDefault("app.database.engine", "sqlite3")
	viper.SetDefault("app.log.format", "json")
}

func SetupViper(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".message-server-go" (without extension).
		viper.AddConfigPath(filepath.Join(home, ".message-server-go"))
		viper.SetConfigName("auth-server")
		viper.SetConfigType("yml")
	}
	SetDefaults()
	viper.SetEnvPrefix("messenger")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func bindEnv(key string, envVar string) {
	if err := viper.BindEnv(key, envVar); err != nil {
		log.Panic(fmt.Sprintf("Failed to bind config key '%s' to environment variable '%s'", key, envVar))
	}
}
