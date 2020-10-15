package config

import "github.com/spf13/viper"

var (
	version    string
	buildDate  string
	branchName string
)

// GetVersion returns app version
func GetVersion() string {
	return version
}

// GetBuildDate returns app build date
func GetBuildDate() string {
	return buildDate
}

// GetBuildDate returns app build date
func GetBranchName() string {
	return branchName
}

// GetDBURL returns the database url
func GetDBURL() string {
	return viper.GetString("app.database.url")
}

// GetDBEngine returns the database engine name
func GetDBEngine() string {
	return viper.GetString("app.database.engine")
}

// GetDBLogQueries enable query log info
func GetDBLogQueries() bool {
	return viper.GetBool("app.database.log")
}

/*
GetLoggerFormat returns the type of log
*/
func GetLoggerFormat() string {
	return viper.GetString("app.log.format")
}
