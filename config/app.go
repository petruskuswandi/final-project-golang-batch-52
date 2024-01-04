package config

import "os"

var PORT = ""
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"
var JWT_KEY = "JWT_KEY"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}

	staticRouteEnv := os.Getenv("STATIC_ROUTE")
	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}

	staticDirEnv := os.Getenv("STATIC_DIR")
	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}

	jwtKeyEnv := os.Getenv("JWT_SECRET")
	if jwtKeyEnv != "" {
		JWT_KEY = jwtKeyEnv
	}
}
