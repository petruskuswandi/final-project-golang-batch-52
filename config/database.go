package config

import "os"

var DB_CONNECTION = "mysql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"
var DB_DATABASE = "final_project_golang_batch_52"
var DB_USERNAME = "root"
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
