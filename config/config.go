package config

import (
	"github.com/joho/godotenv"
	"strconv"
	"os"
)

type Config struct {
	Server  	string
	Database 	string
	PublicKey 	string
	PrivateKey 	string
	Port		int
}

func (config *Config) Read() (err error) {
	// load configuration
	err = godotenv.Load()
	config.Port, err = strconv.Atoi(os.Getenv("PORT"))
	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.PublicKey = os.Getenv("PUBLIC_KEY")
	config.Server = os.Getenv("SERVER")
	config.Database = os.Getenv("DATABASE")
	return
}