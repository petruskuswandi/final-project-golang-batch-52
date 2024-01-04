package config

import "os"

var DB_CONNECTION = ""
var DB_HOST = ""
var DB_PORT = ""
var DB_DATABASE = ""
var DB_USERNAME = ""
var DB_PASSWORD = ""

func InitDatabaseConfig() {
	connectionDriverEnv := os.Getenv("DB_CONNECTION")
	if connectionDriverEnv != "" {
		DB_CONNECTION = connectionDriverEnv
	}

	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")
	if portEnv != "" {
		DB_PORT = portEnv
	}

	databaseNameEnv := os.Getenv("DB_DATABASE")
	if databaseNameEnv != "" {
		DB_DATABASE = databaseNameEnv
	}

	usernameEnv := os.Getenv("DB_USERNAME")
	if usernameEnv != "" {
		DB_USERNAME = usernameEnv
	}

	passwordEnv := os.Getenv("DB_PASSWORD")
	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}
}
