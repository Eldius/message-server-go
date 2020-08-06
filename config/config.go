package config

import "github.com/spf13/viper"

// GetDBURL returns the database url
func GetDBURL() string {
	return viper.GetString("app.database.url")
}

// GetDBEngine returns the database engine name
func GetDBEngine() string {
	return viper.GetString("app.database.engine")
}
