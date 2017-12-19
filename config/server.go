package config

import "os"

// GetListenAddress bind address and port
func GetListenAddress() string {
	listenAddr := os.Getenv("BACKEND_CERT2ADDR_LISTEN")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	return listenAddr
}

func GetHttpAuthUsername() string {
	return os.Getenv("BACKEND_CERT2ADDR_HTTPAUTH_USERNAME")
}

func GetHttpAuthPassword() string {
	return os.Getenv("BACKEND_CERT2ADDR_HTTPAUTH_PASSWORD")
}
