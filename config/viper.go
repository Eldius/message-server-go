package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func SetDefaults() {
	viper.SetDefault("app.database.url", "test.db")
	viper.SetDefault("app.database.engine", "sqlite3")
	viper.SetDefault("app.log.format", "json")
	viper.SetDefault("app.auth.user.pattern", "^[a-zA-Z0-9\\._-]*$")
	viper.SetDefault("app.auth.jwt.secret", uuid.New().String())
}

func BindEnvVars() {
	bindEnv("app.log.format", "APP_LOG_FORMAT")
	bindEnv("app.database.url", "APP_DATABASE_URL")
	bindEnv("app.database.engine", "APP_DATABASE_ENGINE")
	bindEnv("app.database.log", "APP_DATABASE_LOG")
	bindEnv("app.auth.pass.pattern", "APP_AUTH_PASS_PATTERN")
	bindEnv("app.auth.jwt.secret", "APP_AUTH_JWT_SECRET")
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
	BindEnvVars()
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
